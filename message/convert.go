package message

import (
	"fmt"
	"strconv"
)

func Convert(commit string) (string, error) {
	commitNumber, err := strconv.Atoi(commit)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("This app is using commit number %d.\n", commitNumber), nil
}
