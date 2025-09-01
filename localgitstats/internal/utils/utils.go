package utils

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/user"
)

func GetDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.localgitstats"
	return dotFile
}

func ClearDotFile() bool {
	dotFile := GetDotFilePath()
	err := os.Truncate(dotFile, 0)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func ParseFileLinesToSlice(filePath string) []string {
	f := openFile(filePath)
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	return lines
}

func openFile(filePath string) *os.File {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(filePath)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	return f
}
