package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	file, err := ioutil.ReadFile("commit.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	message, err := ConvertToMessage(string(file))
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf(message)
}

func ConvertToMessage(commit string) (string, error) {
	commitNumber, err := strconv.Atoi(commit)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("This is commit number %d", commitNumber), nil
}
