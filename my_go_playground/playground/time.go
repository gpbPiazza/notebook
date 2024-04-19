package playground

import (
	"fmt"
	"time"
)

func timeZaper() {
	ontem := time.Now().AddDate(0, 0, -1)
	hoje := time.Now()
	semanaPassasda := time.Now().AddDate(0, 0, -7)
	semanaQuevem := time.Now().AddDate(0, 0, 7)
	timeVazio := time.Time{}

	times := []time.Time{ontem, semanaPassasda, hoje, timeVazio, semanaQuevem}

	mostRecent := time.Time{}
	for _, t := range times {
		if mostRecent.Before(t) {
			mostRecent = t
			fmt.Println(t)
		}
	}

	fmt.Println("Result ->", mostRecent)
}
