package playground

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

func readsDataFromFile() {
	file, errOpen := os.OpenFile("messages.txt", os.O_RDONLY, fs.ModeType)
	if errOpen != nil {
		log.Fatalf("error on open file err: %s", errOpen)
	}

	var err error
	var parts []string
	for !errors.Is(err, io.EOF) {
		data := make([]byte, 8)

		_, err = file.Read(data)
		if errors.Is(err, io.EOF) {
			fmt.Printf("read: %s\n", "end")
			log.Print("finish reading from file")
			os.Exit(0)
		}

		if err != nil {
			log.Printf("err while reading from file err: %s", err)
		}

		nLine := "\n"
		dataStr := string(data)

		lineCut := strings.Split(dataStr, nLine)

		hasNewLineCut := len(lineCut) <= 1
		if hasNewLineCut {
			parts = append(parts, dataStr)
			continue
		}
		parts = append(parts, lineCut[0])
		line := strings.Join(parts, "")

		fmt.Printf("read: %s\n", line)

		parts = nil
		parts = append(parts, lineCut[1:]...)
	}
}

// does the same thing that reads data from File but with a channel
func lines(f io.ReadCloser) <-chan string {
	linesChan := make(chan string)

	go func() {
		defer func() {
			log.Print("close chann and conn")
			if err := f.Close(); err != nil {
				log.Printf("err while reading closing conn err: %s", err)
			}
			close(linesChan)
		}()

		var err error
		var parts []string
		for !errors.Is(err, io.EOF) {
			data := make([]byte, 8)
			_, err = f.Read(data)
			dataStr := string(data)

			if errors.Is(err, io.EOF) {
				sendLine(linesChan, append(parts, dataStr))
				break
			}

			if err != nil {
				log.Printf("err while reading from conn err: %s", err.Error())
			}

			nLine := "\n"
			nullByte := "\x00"

			if strings.Contains(dataStr, nullByte) {
				sendLine(linesChan, append(parts, dataStr))
				break
			}

			lineSegments := strings.Split(dataStr, nLine)

			hasLineCut := len(lineSegments) > 1
			if !hasLineCut {
				parts = append(parts, dataStr)
				continue
			}

			parts = append(parts, lineSegments[0])
			sendLine(linesChan, parts)

			parts = nil
			parts = append(parts, lineSegments[1:]...)
		}
	}()

	return linesChan
}

func sendLine(lines chan string, lineParts []string) {
	line := strings.Join(lineParts, "")
	lines <- line
}
