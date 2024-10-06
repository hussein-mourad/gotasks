/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"os"

	"github.com/hussein-mourad/gotasks/cmd"
)

func main() {
	createFile()
	cmd.Execute()
}

func createFile() {
	filename := "data.csv"
	_, err := os.Open(filename)
	if err != nil {
		_, err = os.Create(filename)
	}
	if err != nil {
		log.Fatal(err)
	}
}
