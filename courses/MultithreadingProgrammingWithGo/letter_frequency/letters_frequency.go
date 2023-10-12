package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

type RFCGateway struct {
	URL        string
	httpClient *http.Client
}

const Timeout = time.Millisecond * 1500

func NewRFCGateway() *RFCGateway {
	return &RFCGateway{
		httpClient: &http.Client{
			Timeout: Timeout,
		},
		URL: "https://www.rfc-editor.org",
	}
}

type LettersGetter interface {
	getRFCByID(rfcID int) (string, error)
}

func (rfc *RFCGateway) getRFCByID(rfcID int) (string, error) {
	entryPoint := fmt.Sprintf("/rfc/rfc%d.txt", rfcID)
	url := rfc.URL + entryPoint

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	resp, err := rfc.httpClient.Do(request)
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return "", err
	}

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

func countLetters(rfcID int, frequency *[26]int32, lettersGetter LettersGetter) {
	letters, _ := lettersGetter.getRFCByID(rfcID)

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
	rFCGateway := NewRFCGateway()

	var frequency [26]int32
	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		countLetters(i, &frequency, rFCGateway)
	}

	elapsed := time.Since(start)
	fmt.Println("Done")
	fmt.Printf("Processing took %s\n", elapsed)
	for i, f := range frequency {
		fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	}
}
