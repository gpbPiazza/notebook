package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func getRFCLetters(rfcID int) (string, error) {
	httpClient := &http.Client{}
	url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", rfcID)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	return string(body), nil
}

func countLetters(rfcID int, frequency *[26]int32) {
	letters, _ := getRFCLetters(rfcID)

	for _, b := range letters {
		c := strings.ToLower(string(b))
		index := strings.Index(allLetters, c)
		if index < 0 {
			continue
		}
		frequency[index] += 1
	}
}

func main() {
	var frequency [26]int32
	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		countLetters(i, &frequency)
	}

	elapsed := time.Since(start)
	fmt.Println("Done")
	fmt.Printf("Processing took %s\n", elapsed)
	for i, f := range frequency {
		fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	}
}
