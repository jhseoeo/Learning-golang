package cleanup

import (
	"fmt"
	"os"
	"testing"
)

func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tempfile")
	if err != nil {
		return "", err
	}
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func TestFileProcessing(t *testing.T) {
	fName, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fName)
}
