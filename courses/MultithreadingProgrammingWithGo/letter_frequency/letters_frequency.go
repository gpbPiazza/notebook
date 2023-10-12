package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
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
	getRFC(rfcID int) (string, error)
}

func (rfc *RFCGateway) getRFC(rfcID int) (string, error) {
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

func countLetters(rfcID int, frequency *[26]int32, lettersGetter LettersGetter, wg *sync.WaitGroup, locker *sync.Mutex) {
	letters, _ := lettersGetter.getRFC(rfcID)

	for j := 0; j < 25; j++ {
		for _, b := range letters {
			c := strings.ToLower(string(b))
			index := strings.Index(allLetters, c)
			if index < 0 {
				continue
			}
			// locker.Lock()
			// frequency[index] += 1
			atomic.AddInt32(&frequency[index], 1)
			// locker.Unlock()
		}
	}

	wg.Done()
}

func main() {
	rFCGateway := NewRFCGateway()
	wg := sync.WaitGroup{}
	locker := sync.Mutex{}
	var frequency [26]int32
	fmt.Println("Start")
	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		wg.Add(1)
		go countLetters(i, &frequency, rFCGateway, &wg, &locker)
	}
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("Done")
	fmt.Printf("Processing took %s\n", elapsed)
	for i, f := range frequency {
		fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	}
}

// Using 1 times load
// Processing took 42.203930625s linear proccess
// Processing took 1.102896416s Using lock and unlock in multithread
// Processing took 506.188709ms using atomic variables in multithread

// Using more 25 times load
// atomic variables -> Processing took 6.7137705s
// lock and unlock -> Processing took 23.287942666s
// linear -> Processing took 9.31410875s

// this shows the cost of lock and unlcok a peace of code.
