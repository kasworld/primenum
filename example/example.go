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
40.05µs 4 7
66.43µs 6 13
55.09µs 11 31
47.68µs 18 61
63.53µs 31 127
94.96µs 54 251
97.17µs 97 509
158.18µs 172 1021
342.47µs 309 2039
588.4µs 564 4093
1.00994ms 1028 8191
2.06324ms 1900 16381
3.93391ms 3512 32749
6.76159ms 6542 65521
10.4871ms 12251 131071
22.794339ms 23000 262139
46.186039ms 43390 524287
91.819998ms 82025 1048573
183.172376ms 155611 2097143
363.645092ms 295947 4194301
731.649163ms 564163 8388593
1.479949937s 1077871 16777213
2.930412283s 2063689 33554393
5.856430208s 3957809 67108859
11.573030381s 7603553 134217689
22.638488304s 14630843 268435399
42.723481076s 28192750 536870909
1m22.717911066s 54400028 1073741789
3m13.73109361s 105097565 2147483647
8m15.594820156s 203280221 4294967291

*/
