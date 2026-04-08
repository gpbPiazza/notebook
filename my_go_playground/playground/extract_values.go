package playground

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var moneyRegex = regexp.MustCompile(`\d+,\d+`)

func ExtractAndSum() map[string]float64 {
	sums := map[string]float64{}
	for _, line := range inputs {
		key := "BMG"
		if strings.Contains(line, "LECCA") {
			key = "LECCA"
		}

		matches := moneyRegex.FindAllString(line, -1)
		for _, m := range matches {
			fmt.Printf("extracted money -> %s\n", m)
			val, err := strconv.ParseFloat(strings.Replace(m, ",", ".", 1), 64)
			if err != nil {
				fmt.Printf("line %s - error -> %s", line, err)
				continue
			}
			sums[key] += val
		}

	}

	for k, v := range sums {
		fmt.Printf("%s: %.2f\n", k, v)
	}

	fmt.Printf("Total: %.2f\n", sums["BMG"]+sums["LECCA"])

	return sums
}

var inputs = []string{"|AMORT CARTAO CREDITO - BMG | | 1 | | | | 69,13| 70,78| 70,78",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 117,34| 118,23| 118,23| 117,26| 464,50| 464,50",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 464,50| 464,50| 464,50| 464,50| 468,12| 202,65",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 202,65| 468,13| 468,13| 468,13| 468,13| 140,86",
	"|AMORT CARTAO CREDITO - BMG | D | 1 | 140,86| 468,13| 468,13| 468,13| 468,13| 202,65",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 202,65| 468,13| 468,13| 468,13| 202,65| 202,65",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 468,13| 468,13| 468,13| 468,13| 438,57|",
	"|AMORT CARTAO CREDITO - BMG | | 1 | | 311,10| 311,10| 406,75| 468,13| 468,13",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 468,13| 404,30| 406,75| 401,48| 401,48|",
	"|AMORT CARTAO CREDITO - BMG | | 1 | | 406,75| 406,75| 406,75| 406,75| 406,75",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 396,60| 396,60| 39,66| 406,75| 406,75|",
	"|AMORT CARTAO CREDITO - BMG | | 1 | | 274,88| 274,88| 274,88| 406,75| 274,88",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 274,88| 274,88| 274,88| 274,88| 274,88|",
	"|AMORT CARTAO CREDITO - BMG | | 1 | | | 274,88| 200,44| 200,44| 274,88",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 274,88| 451,56| 451,56| 427,97| 427,97|",
	"|AMORT CARTAO CREDITO - BMG | | 1 | | 490,61| 490,61| 490,61| 490,61| 490,61",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 490,61| 490,61| 171,38| 171,38| 164,97| 171,38",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 171,38| 164,97| 171,38| 171,38| 174,71| 174,71",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 174,71| 174,71| 174,71| 174,71| 174,71| 174,71",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 174,71|",
	"|AMORT CARTAO CREDITO - BMG | | 1 | | 406,75| 406,75| 406,75| 406,75| 406,75",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 396,60| 396,60| 39,66| 406,75| 406,75|",
	"|AMORT CARTAO CREDITO - BMG | | 1 | | 274,88| 274,88| 274,88| 406,75| 274,88",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 274,88| 274,88| 274,88| 274,88| 274,88|",
	"|AMORT CARTAO CREDITO - BMG | | 1 | | | 274,88| 200,44| 200,44| 274,88",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 274,88| 451,56| 451,56| 427,97| 427,97|",
	"|AMORT CARTAO CREDITO - BMG | | 1 | | 490,61| 490,61| 490,61| 490,61| 490,61",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 490,61| 490,61| 171,38| 171,38| 164,97| 171,38",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 171,38| 164,97| 171,38| 171,38| 174,71| 174,71",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 174,71| 174,71| 174,71| 174,71| 174,71| 174,71",
	"|AMORT CARTAO CREDITO - BMG | | 1 | 174,71|",
	"|AMORT CARTAO BENEFICIO - LECCA | D | 1 | | | 439,68| 433,27| 439,68| 439,68",
	"|AMORT CARTAO BENEFICIO - LECCA | D | 1 | 439,68| 439,68| 439,68| 439,68| 439,68| 439,68",
	"|AMORT CARTAO BENEFICIO - LECCA | D | 1 | 439,68| 439,68| 439,68| 439,68| 439,68| 439,68",
	"|AMORT CARTAO BENEFICIO - LECCA | | 1 | 439,68|",
	"|AMORT CARTAO BENEFICIO - LECCA | D | 1 | | | 439,68| 433,27| 439,68| 439,68",
	"|AMORT CARTAO BENEFICIO - LECCA | D | 1 | 439,68| 439,68| 439,68| 439,68| 439,68| 439,68",
	"|AMORT CARTAO BENEFICIO - LECCA | D | 1 | 439,68| 439,68| 439,68| 439,68| 439,68| 439,68",
	"|AMORT CARTAO BENEFICIO - LECCA | | 1 | 439,68|",
}
