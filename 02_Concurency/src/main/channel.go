package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var character string

func read_input(file_name string, channel chan []string) _, error {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		character = scanner.Text()
	}
	if nil == channel {
		fmt.Println("The content is null\n")
		err = err.new("The content is null")
		return err
	} else {
		split_string := strings.Split(character, " ")
		channel <- split_string
	}
	return
}

func count_occurences(data []string, counts map[string]int) {
	for _, i := range data {
		counts[i]++
	}
}

func main() {
	//Create list of potential file name
	var Input_file_name []string
	//Path to the Input files
	Input_file_path := "D:/04_Go_course_pratice/02_Concurency/src/main/Input"
	//Output
	counts := make(map[string]int)
	//channel between routines
	channel := make(chan []string)
	//Create files from the directory
	files, err := ioutil.ReadDir(Input_file_path)
	//If no file exist will return error
	if err != nil {
		log.Fatal(err)
	}
	//Save file's name to list
	for _, file := range files {
		Input_file_name = append(Input_file_name, file.Name())
	}
	//Read content of files
	//Number of routines depend on how many files exist
	for _, input := range Input_file_name {
		go read_input(Input_file_path+"/"+input, channel)
	}
	//Count the occurency of each character
	go func() {
		for {
			count_occurences(<-channel, counts)
		}
	}()
	//Delay for all routines
	time.Sleep(10 * time.Second)
	//Print the final result
	for key, value := range counts {
		fmt.Printf("The occurency of %s is : %d \n", key, value)
	}

}
