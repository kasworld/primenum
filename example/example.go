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
	pn := primenum.NewPrimeIntList(8)
	for i := int(8); i < 0xffffffff; i <<= 1 {
		find2(&pn, i)
	}
}

func find2(pn *primenum.PrimeIntList, n int) {
	st := time.Now()
	pn.AppendFindTo(n)
	fmt.Printf("%v %v %v\n",
		time.Now().Sub(st),
		len(*pn),
		(*pn)[len(*pn)-1],
	)
}

func find(n int64) {
	st := time.Now()
	primes := primenum.MakePrimes(n)
	fmt.Printf("%v %v %v\n",
		time.Now().Sub(st),
		len(primes),
		primes[len(primes)-1],
	)
}
