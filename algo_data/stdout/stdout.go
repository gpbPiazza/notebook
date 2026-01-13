package stdout

import (
	"io"
	"log"
	"os"
)

func String(f func()) string {
	original := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	f()
	os.Stdout = original

	err := writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	result, _ := io.ReadAll(reader)
	return string(result)
}
