package main

import "os"

func process(data []byte) {
	// do something
}

func processFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// this pattern is good at reducing the garbage collector's workload
	data := make([]byte, 100) // using slice as a buffer
	for {
		count, err := file.Read(data)
		if err != nil {
			return err
		}
		if count == 0 {
			return nil
		}
		process(data[:count]) // it passes next block of bytes in to the slice (up to 100)
	}

}

func main() {
	processFile("asdasd")
}
