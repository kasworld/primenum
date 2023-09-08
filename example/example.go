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
	"os"
	"strconv"
	"time"

	"github.com/kasworld/primenum"
)

func main() {
	num := 1000000
	cmd := "multi4"
	if len(os.Args) > 1 {
		v, err := strconv.Atoi(os.Args[1])
		if err != nil {
			println(err)
			help()
		} else {
			num = v
		}
	}
	if len(os.Args) > 2 {
		cmd = os.Args[2]
	}
	st := time.Now()
	fmt.Printf("%v %v\n", cmd, num)
	switch cmd {
	case "simple":
		pn := primenum.MakePrimes(int64(num))
		fmt.Printf("%v %v\n", time.Since(st), pn)
	case "single":
		pn := primenum.New().AppendFindTo(num)
		fmt.Printf("%v %v\n", time.Since(st), pn)
	case "single2":
		pn := primenum.NewWithCap(num / 16).AppendFindTo(num)
		fmt.Printf("%v %v\n", time.Since(st), pn)
	case "multi1":
		pn := primenum.New().MultiAppendFindTo(num)
		fmt.Printf("%v %v\n", time.Since(st), pn)
	case "multi2":
		pn := primenum.New().MultiAppendFindTo2(num)
		fmt.Printf("%v %v\n", time.Since(st), pn)
	case "multi3":
		pn := primenum.New().MultiAppendFindTo3(num)
		fmt.Printf("%v %v\n", time.Since(st), pn)
	case "multi4":
		pn := primenum.NewWithCap(num / 16).MultiAppendFindTo4(num)
		fmt.Printf("%v %v\n", time.Since(st), pn)
	default:
		help()
	}
}

func help() {
	fmt.Printf("%v num cmd\n", os.Args[0])
	println("num : prime calc upto num")
	println("cmd : multi1 ~ multi4 : multi thread calc")
	println("cmd : simple single single2 : single thread calc")
}

func loadsave() {
	filename := "pnff.gob"
	pn := primenum.New().AppendFindTo(0xff)
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
