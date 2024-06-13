package stdout

import (
	"io"
	"os"
)

func String(f func()) string {
	original := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	f()
	os.Stdout = original
	writer.Close()

	result, _ := io.ReadAll(reader)
	return string(result)
}
