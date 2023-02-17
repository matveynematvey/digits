package main

import (
	"fmt"
	"strconv"
)

const versionCount = 39365 //1222222222 ternary notation

var n = [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

func toTernary(n int) (s [10]int) { //from decimal to ternary
	cnt := 9
	for n > 0 {
		s[cnt] = n % 3
		n /= 3
		cnt--
	}
	return
}

func calcPrevSum(str string) (int, bool) {
	sum, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error in conversion: %s, \n In str: %s", err, str)
	}

	return sum, err == nil

}

func eval(s [10]int) {
	str := ""
	sum := 0
	localStr := ""

	for ind, op := range s {
		switch op {
		case 0: //0
			{
				localStr += strconv.Itoa(n[ind])
				str += strconv.Itoa(n[ind])
			}
		case 1: //-
			{
				str += "-" + strconv.Itoa(n[ind])

				if ind != 0 {
					locSum, ok := calcPrevSum(localStr)
					if ok {
						sum += locSum
					}
				}

				localStr = "-" + strconv.Itoa(n[ind])
			}
		case 2: //+
			{
				str += "+" + strconv.Itoa(n[ind])

				locSum, ok := calcPrevSum(localStr)
				if ok {
					sum += locSum
				}

				localStr = "+" + strconv.Itoa(n[ind])
			}
		}
	}

	locSum, ok := calcPrevSum(localStr)
	if ok {
		sum += locSum
	}

	if sum == 200 {
		fmt.Println(str)
	}
}

func main() {

	for ver := 0; ver < versionCount; ver++ {
		eval(toTernary(ver))
	}

}
