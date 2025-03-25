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
