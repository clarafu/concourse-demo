package demo

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("commit.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	commitNum := strings.TrimSuffix(string(file), "\n")
	message, err := ConvertToMessage(commitNum)
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

	return fmt.Sprintf("This app is using commit number %d.\n", commitNumber), nil
}
