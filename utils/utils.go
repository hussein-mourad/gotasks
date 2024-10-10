package utils

import (
	"fmt"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
