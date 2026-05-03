package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const resultFilename = "combined.txt"

func mergeDirTextFiles(directory []os.DirEntry, output *os.File) {
	for _, item := range directory {
		name := item.Name()
		fmt.Printf("Processing file %s\n", name)

		if name == resultFilename || item.IsDir() {
			continue
		}

		if extension := filepath.Ext(name); extension != ".txt" {
			fmt.Printf("Skipping file %s (not a text file)\n", name)
			continue
		}

		file, err := os.Open(name)
		if err != nil {
			fmt.Printf("Error opening file %s: %s\n", name, err)
			continue
		}

		data := make([]byte, 64)

		for {
			n, err := file.Read(data)
			if err == io.EOF {
				break
			}
			_, err = output.Write(data[:n])
			if err != nil {
				fmt.Printf("Failed to write data from file %s: %s\n", name, err)
				return
			}
		}

		if _, err = output.WriteString("\n"); err != nil {
			fmt.Printf("Failed to insert new line to output")
			return
		}

		fmt.Printf("File %s has been sucessfully processed\n", name)
		file.Close()
	}
}

func main() {
	directory, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.OpenFile(resultFilename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	mergeDirTextFiles(directory, file)
}
