package filereader

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	logFatal = log.Fatal
)

func ReadFile(path string) []string {
	//read file and return error in case of issues
	file, err := os.Open(path)
	if err != nil {
		logFatal("Error in reading prefix file:", err)
	}
	defer file.Close() //close file resources once functions exits
	//Init Variables
	lineRecord := make(chan string)
	outSlice := make([]string, 0)
	//this goroutines read the file contents line by line
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineRecord <- scanner.Text()
		}
		close(lineRecord)
	}()
	for val := range lineRecord {
		//trim leading and trailing space if found in line record
		if temp := strings.TrimSpace(val); temp != "" {
			outSlice = append(outSlice, temp)
		}
	}
	return outSlice
}
