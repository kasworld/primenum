// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package primenum

import (
	"encoding/gob"
	"fmt"
	"math"
	"os"
	"runtime"
	"sort"
	"sync"
)

type Element int

type PrimeIntList []Element

func New() PrimeIntList {
	pn := PrimeIntList{2, 3}
	return pn
}

func NewWithCap(n int) PrimeIntList {
	pn := make(PrimeIntList, 0, n)
	pn = append(pn, PrimeIntList{2, 3}...)
	return pn
}

func (pn PrimeIntList) String() string {
	return fmt.Sprintf("%v %v", len(pn), pn.GetLast())
}

func (pn PrimeIntList) GetLast() Element {
	return pn[len(pn)-1]
}

func (pn PrimeIntList) MaxCanCheck() Element {
	last := pn.GetLast()
	return last * last
}

func (pn PrimeIntList) IsPrime(n Element) bool {
	to := Element(math.Sqrt(float64(n)))
	for _, v := range pn {
		if v > to {
			break
		}
		if n%v == 0 {
			return false
		}
	}
	return true
}

func (pn PrimeIntList) AppendFindTo(n Element) PrimeIntList {
	last := pn.GetLast()
	if last >= n {
		return pn
	}
	for i := last + 2; i <= n; i += 2 {
		if pn.IsPrime(i) {
			pn = append(pn, i)
		}
	}
	return pn
}

func (pn PrimeIntList) MultiAppendFindTo(n Element) PrimeIntList {
	oldLen := len(pn)
	last := pn.GetLast()
	if last >= n {
		return pn
	}

	if n >= pn.MaxCanCheck() {
		pn = pn.MultiAppendFindTo(n / 2)
		oldLen = len(pn)
		last = pn.GetLast()
	}

	bufl := runtime.NumCPU() * 1

	var wgWorker sync.WaitGroup
	var wgAppend sync.WaitGroup

	// recv result
	appendCh := make(chan Element, bufl*2)
	wgAppend.Add(1)
	go func() {
		for n := range appendCh {
			pn = append(pn, n)
		}
		wgAppend.Done()
	}()

	// prepare need check data
	argCh := make(chan Element, bufl*1000)
	go func() {
		for i := last + 2; i <= n; i += 2 {
			argCh <- i
		}
		close(argCh)
	}()

	// run worker
	for i := 0; i < bufl; i++ {
		wgWorker.Add(1)
		go func() {
			for n := range argCh {
				if pn.IsPrime(n) {
					appendCh <- n
				}
			}
			wgWorker.Done()
		}()
	}
	wgWorker.Wait()
	close(appendCh)
	wgAppend.Wait()

	sort.Slice(pn[oldLen:], func(i, j int) bool {
		return pn[i+oldLen] < pn[j+oldLen]
	})
	return pn
}

func (pn PrimeIntList) MultiAppendFindTo2(n Element) PrimeIntList {
	oldLen := len(pn)
	last := pn.GetLast()
	if last >= n {
		return pn
	}

	if n >= pn.MaxCanCheck() {
		pn = pn.MultiAppendFindTo2(n / 2)
		oldLen = len(pn)
		last = pn.GetLast()
	}

	workerCount := runtime.NumCPU() * 1

	var wgWorker sync.WaitGroup
	var wgAppend sync.WaitGroup

	// recv result
	appendCh := make(chan Element, workerCount*2)
	wgAppend.Add(1)
	go func() {
		for n := range appendCh {
			pn = append(pn, n)
		}
		wgAppend.Done()
	}()

	// run worker
	for workerid := 0; workerid < workerCount; workerid++ {
		wgWorker.Add(1)
		go func(workerid int) {
			for i := last + 2 + Element(workerid)*2; i <= n; i += Element(workerCount) * 2 {
				if pn.IsPrime(i) {
					appendCh <- i
				}
			}
			wgWorker.Done()
		}(workerid)
	}
	wgWorker.Wait()
	close(appendCh)
	wgAppend.Wait()

	sort.Slice(pn[oldLen:], func(i, j int) bool {
		return pn[i+oldLen] < pn[j+oldLen]
	})
	return pn
}

func (pn PrimeIntList) MultiAppendFindTo3(n Element) PrimeIntList {
	last := pn.GetLast()
	if last >= n {
		return pn
	}

	if n >= pn.MaxCanCheck() {
		pn = pn.MultiAppendFindTo3(n / 2)
		last = pn.GetLast()
	}

	workerCount := runtime.NumCPU() * 1

	var wgWorker sync.WaitGroup

	workResult := make([][]Element, workerCount)
	// run worker
	for workerID := 0; workerID < workerCount; workerID++ {
		wgWorker.Add(1)
		go func(workerID int) {
			var rtn []Element
			for i := last + 2 + Element(workerID)*2; i <= n; i += Element(workerCount) * 2 {
				if pn.IsPrime(i) {
					rtn = append(rtn, i)
				}
			}
			workResult[workerID] = rtn
			wgWorker.Done()
		}(workerID)
	}
	wgWorker.Wait()

	return pn.MergeSort(workResult)
}

func (pn PrimeIntList) MultiAppendFindTo4(n Element) PrimeIntList {
	last := pn.GetLast()
	if last >= n {
		return pn
	}

	if n >= pn.MaxCanCheck() {
		pn = pn.MultiAppendFindTo4(n / 2)
		last = pn.GetLast()
	}

	workerCount := runtime.NumCPU() * 1

	var wgWorker sync.WaitGroup

	workResult := make([][]Element, workerCount)
	workerBufferLen := int(n-last) / workerCount / 16
	// run worker
	for workerID := 0; workerID < workerCount; workerID++ {
		wgWorker.Add(1)
		go func(workerID int) {
			rtn := make([]Element, 0, workerBufferLen)
			for i := last + 2 + Element(workerID)*2; i <= n; i += Element(workerCount) * 2 {
				if pn.IsPrime(i) {
					rtn = append(rtn, i)
				}
			}
			workResult[workerID] = rtn
			wgWorker.Done()
		}(workerID)
	}
	wgWorker.Wait()

	return pn.MergeSort(workResult)
}

// merge sort
func (pn PrimeIntList) MergeSort(workResult [][]Element) PrimeIntList {
	workerCount := len(workResult)
	workerPos := make([]int, workerCount)
	for {
		minFound := Element(0)
		minWorkerID := 0
		for workerID := 0; workerID < workerCount; workerID++ {
			pos := workerPos[workerID]
			if pos >= len(workResult[workerID]) {
				continue
			}
			if minFound == 0 || workResult[workerID][pos] < minFound {
				minFound = workResult[workerID][pos]
				minWorkerID = workerID
			}
		}
		if minFound != 0 {
			pn = append(pn, minFound)
			workerPos[minWorkerID]++
		} else {
			break
		}
	}
	return pn
}

func (pn PrimeIntList) Save(filename string) error {
	fd, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("err in create %v", err)
	}
	defer fd.Close()
	enc := gob.NewEncoder(fd)
	err = enc.Encode(pn)
	if err != nil {
		return err
	}
	return nil
}

func LoadPrimeIntList(filename string) (PrimeIntList, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("fail to %v", err)
	}
	defer fd.Close()
	var rtn PrimeIntList
	dec := gob.NewDecoder(fd)
	err = dec.Decode(&rtn)
	if err != nil {
		return nil, err
	}
	return rtn, nil
}
