package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

var character string

func RemoveExpression(raw_text string) string {
	var line string
	doublespace := regexp.MustCompile(`\s+`)
	breakLine := regexp.MustCompile(`\n`)
	spaceEOL := regexp.MustCompile(`\s\n`)
	line = doublespace.ReplaceAllString(raw_text, " ")
	line = breakLine.ReplaceAllString(line, " ")
	line = spaceEOL.ReplaceAllString(line, " ")
	return line
}

//Read the content of files
func ReadInput(file_name string, channel chan []string) {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(reader)
	validstring := RemoveExpression(string(content))
	split_string := strings.Split(validstring, " ")

	fmt.Printf("split_string %s\n", split_string)
	channel <- split_string
}

//Count ocurrency of data
func CountOccurences(data []string, counts map[string]int) {
	for _, i := range data {
		counts[i]++
	}
}

func main() {
	//Create list of potential file name
	var InputFileName []string
	//Path to the Input files
	InputFilePath := "D:/05_Go_Repo_boot_Camp/go_bootcamp/02_Concurency/src/main/Input"
	//Output
	counts := make(map[string]int)
	//channel between routines
	channel := make(chan []string)
	//Create files from the directory
	files, err := ioutil.ReadDir(InputFilePath)
	//If no file exist will return error
	if err != nil {
		log.Fatal(err)
	}
	//Save file's name to list
	for _, file := range files {
		InputFileName = append(InputFileName, file.Name())
	}
	//Read content of files
	//Number of routines depend on how many files exist
	for _, input := range InputFileName {
		go ReadInput(InputFilePath+"/"+input, channel)
	}
	//Count the occurency of each character
	go func() {
		for {
			CountOccurences(<-channel, counts)
		}
	}()
	//Delay for all routines
	time.Sleep(5 * time.Second)
	//Print the final result
	for key, value := range counts {
		fmt.Printf("The occurency of %s is : %d \n", key, value)
	}

}
