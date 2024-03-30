package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func compareStrings() {

	january := "2023-01"
	febuary := "2023-02"

	if january < febuary {
		fmt.Println(january)
		return
	}
	fmt.Println("num da galo")

}

func zapper() {

	ss := []string{"12/12", "2/12", "3/12", "8/12", "4/12", "5/12", "6/12", "11/12", "7/12", "1/12", "9/12", "10/12", "3/12"}
	sort.Slice(ss, func(i, j int) bool {
		// next := strings.Split(ss[i], "/")
		// current := strings.Split(ss[j], "/")

		// nextDenominador, _ := strconv.ParseFloat(next[0], 32)
		// nextDivisor, _ := strconv.ParseFloat(next[1], 32)
		// nextVal := nextDenominador / nextDivisor

		// currentDenominador, _ := strconv.ParseFloat(current[0], 32)
		// currentDivisor, _ := strconv.ParseFloat(current[1], 32)
		// currentVal := currentDenominador / currentDivisor

		// return nextVal < currentVal

		next := strings.Split(ss[i], "/")
		current := strings.Split(ss[j], "/")

		nextDenominator, _ := strconv.ParseFloat(next[1], 32)
		nextNumerator, _ := strconv.ParseFloat(next[0], 32)
		nextVal := nextNumerator / nextDenominator

		currentDenominator, _ := strconv.ParseFloat(current[1], 32)
		currentNumerator, _ := strconv.ParseFloat(current[0], 32)
		currentVal := currentNumerator / currentDenominator

		return nextVal < currentVal
	})

	fmt.Println(ss)
}
