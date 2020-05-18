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
	"math"
	"runtime"
	"sort"
	"sync"
)

type PrimeIntList []int

func (pn PrimeIntList) FindPos(n int) (int, bool) {
	i := sort.SearchInts(pn, n)
	if i < len(pn) && pn[i] == n {
		// x is present at pn[i]
		return i, true
	} else {
		// x is not present in pn,
		// but i is the index where it would be inserted.
		return i, false
	}
}

// func (pn PrimeIntList) Sort() {
// 	if !sort.IntsAreSorted(pn) {
// 		sort.Ints(pn)
// 	}
// }

func (pn PrimeIntList) MaxCanCheck() int {
	last := pn[len(pn)-1]
	return last * last
}

func (pn PrimeIntList) CanFindIn(n int) bool {
	return pn.MaxCanCheck() > n
}

func (pn PrimeIntList) CalcPrime(n int) bool {
	to := int(math.Sqrt(float64(n)))
	for _, v := range pn {
		if n%v == 0 {
			return false
		}
		if v > to {
			break
		}
	}
	return true
}

func (pn *PrimeIntList) AppendFindTo(n int) {
	last := (*pn)[len(*pn)-1]
	if last >= n {
		return
	}
	if pn.MaxCanCheck() < n {
		pn.AppendFindTo(n / 2)
	}
	for i := last + 2; i < n; i += 2 {
		if pn.CalcPrime(i) {
			*pn = append(*pn, i)
		}
	}
}

func NewPrimeIntList(n int) PrimeIntList {
	pn := make(PrimeIntList, 0, n/10)
	pn = append(pn, []int{2, 3}...)
	pn.AppendFindTo(n)
	return pn
}

func (pn *PrimeIntList) MultiAppendFindTo(n int) {
	lastIndex := len(*pn) - 1
	last := (*pn)[lastIndex]
	if last >= n {
		return
	}
	if pn.MaxCanCheck() < n {
		pn.AppendFindTo(n / 2)
	}

	bufl := runtime.NumCPU() * 1

	var wgWorker sync.WaitGroup
	var wgAppend sync.WaitGroup

	// recv result
	appendCh := make(chan int, bufl*2)
	wgAppend.Add(1)
	go func() {
		for n := range appendCh {
			*pn = append(*pn, n)
		}
		wgAppend.Done()
	}()

	// prepare need check data
	argCh := make(chan int, bufl*1000)
	go func() {
		for i := last + 2; i < n; i += 2 {
			argCh <- i
		}
		close(argCh)
	}()

	// run worker
	for i := 0; i < bufl; i++ {
		wgWorker.Add(1)
		go func() {
			for n := range argCh {
				if pn.CalcPrime(n) {
					appendCh <- n
				}
			}
			wgWorker.Done()
		}()
	}
	wgWorker.Wait()
	close(appendCh)
	wgAppend.Wait()

	sort.Ints((*pn)[lastIndex+1:])
}
