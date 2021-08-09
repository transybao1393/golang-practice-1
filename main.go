package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	keyword   string = "vaccine"
	file1Path string = "sample1.txt"
	file2Path string = "sample2.txt"
)

func main() {
	//- sample
	// numberChan := make(chan int)
	// stringChan := make(chan string)

	// go test1(numberChan)
	// go test2(stringChan)
	// fmt.Println("number channel result", <-numberChan)
	// fmt.Println("string channel result", <-stringChan)

	//- file count
	file1Channel := make(chan int)
	file2Channel := make(chan int)

	go file1Count(file1Channel)
	go file2Count(file2Channel)

	fmt.Printf("total count: %d", <-file1Channel+<-file2Channel)

}

func fileCount(wordCount chan int, filePath string) {
	var count int
	fileContent, err := ioutil.ReadFile(file1Path)
	//- if error
	if err != nil {
		//- put the file content to the channel
		wordCount <- 0
		return
	}
	//- if not error
	count = strings.Count(string(fileContent), keyword)
	fmt.Printf("count %d of file name %s", count, filePath)
	wordCount <- count
}

func file1Count(wordCount chan int) {
	var count int
	fileContent, err := ioutil.ReadFile(file1Path)
	//- if error
	if err != nil {
		//- put the file content to the channel
		wordCount <- 0
		return
	}
	//- if not error
	count = strings.Count(string(fileContent), keyword)
	wordCount <- count
}

func file2Count(wordCount chan int) {
	var count int
	fileContent, err := ioutil.ReadFile(file2Path)
	//- if error
	if err != nil {
		//- put the file content to the channel
		wordCount <- 0
		return
	}
	//- if not error
	count = strings.Count(string(fileContent), keyword)
	wordCount <- count
}

func test1(numberChan chan int) {
	var result int
	for i := 0; i < 100; i++ {
		result += i
	}
	numberChan <- result
}

func test2(stringChan chan string) {
	var result string
	for i := 'A'; i < 'A'+26; i++ {
		result += string(i)
	}
	stringChan <- result
}
