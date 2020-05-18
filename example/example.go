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

package main

import (
	"fmt"
	"time"

	"github.com/kasworld/primenum"
)

func main() {
	bench1()
	// bench2()
	// loadsave()
}

func loadsave() {
	filename := "pnff.gob"
	pn := primenum.NewPrimeIntList(0xff)
	fmt.Println(pn)
	err := pn.Save(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	pn2, err := primenum.LoadPrimeIntList(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(pn2)
}

func bench2() {
	pn := primenum.NewPrimeIntList(8)
	for i := int(8); i < 0xffffffffffff; i <<= 1 {
		st := time.Now()
		pn = pn.MultiAppendFindTo(i)
		fmt.Printf("multi %v %v %v\n",
			time.Now().Sub(st),
			len(pn),
			(pn)[len(pn)-1],
		)
	}
}

func bench1() {
	// for i := int(8); i < 0xff; i <<= 1 {
	for i := int(8); i < 0xffffffffffff; i <<= 1 {
		// find1(int64(i))
		find2(i)
		find3(i)
		find4(i)
		fmt.Println()
	}
}

func find1(n int64) {
	st := time.Now()
	pn := primenum.MakePrimes(n)
	fmt.Printf("simple %v %v %v\n",
		time.Now().Sub(st),
		len(pn),
		pn[len(pn)-1],
	)
	// fmt.Println(pn)
}

func find2(n int) {
	st := time.Now()
	pn := primenum.NewPrimeIntList(8)
	pn = pn.AppendFindTo(n)
	fmt.Printf("single %v %v %v\n",
		time.Now().Sub(st),
		len(pn),
		(pn)[len(pn)-1],
	)
	// fmt.Println(pn)
}

func find3(n int) {
	st := time.Now()
	pn := primenum.NewPrimeIntList(8)
	pn = pn.MultiAppendFindTo(n)
	fmt.Printf("multi %v %v %v\n",
		time.Now().Sub(st),
		len(pn),
		(pn)[len(pn)-1],
	)
	// fmt.Println(pn)
}

func find4(n int) {
	st := time.Now()
	pn := primenum.NewPrimeIntList(8)
	pn = pn.MultiAppendFindTo2(n)
	fmt.Printf("multi2 %v %v %v\n",
		time.Now().Sub(st),
		len(pn),
		(pn)[len(pn)-1],
	)
	// fmt.Println(pn)
}
