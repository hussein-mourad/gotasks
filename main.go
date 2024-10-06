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
	filepath := "data/data.csv"
	_, err := os.Open(filepath)
	if err != nil {
		_, err = os.Create(filepath)
	}
	if err != nil {
		log.Fatal(err)
	}
}
