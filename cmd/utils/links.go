package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLinksFromFile(path string) []string {

	file, err := os.Open(path)
	var links []string
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		links = append(links, fileScanner.Text())
	}

	file.Close()
	return links
}
