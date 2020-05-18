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
	for i := int(8); i < 0xffffffffffff; i <<= 1 {
		find2(&pn, i)
	}
}

func main2() {
	for i := int(8); i < 0xffffffffffff; i <<= 1 {
		find(int64(i))
	}
}

func find2(pn *primenum.PrimeIntList, n int) {
	st := time.Now()
	pn.MultiAppendFindTo(n)
	// pn.AppendFindTo(n)
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

/* test data
1.69µs 4 7
880ns 6 13
1.07µs 11 31
3.07µs 18 61
5.5µs 31 127
11.18µs 54 251
10.72µs 97 509
24.84µs 172 1021
53.52µs 309 2039
124.25µs 564 4093
296.99µs 1028 8191
684.909µs 1900 16381
1.62772ms 3512 32749
3.868778ms 6542 65521
6.936947ms 12251 131071
11.301715ms 23000 262139
27.615097ms 43390 524287
68.094249ms 82025 1048573
168.868742ms 155611 2097143
428.764485ms 295947 4194301
1.068120231s 564163 8388593
2.71166386s 1077871 16777213
6.908527526s 2063689 33554393

*/
