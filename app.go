package main

import (
	"demo/message"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("commit.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	commitNum := strings.TrimSuffix(string(file), "\n")
	message, err := message.Convert(commitNum)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf(message)
}
