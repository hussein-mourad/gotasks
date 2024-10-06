/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/hussein-mourad/gotasks/internal/tasks"
	"github.com/hussein-mourad/gotasks/utils"
)

func main() {
	// GetFile()
	// cmd.Execute()
	taskStore := tasks.NewStore()

	taskStore.CreateTask("Task")
	taskStore.CreateTask("Task")
	taskStore.CreateTask("Task")

	for _, t := range taskStore.Tasks {
		fmt.Println(t)
	}
}

func GetFile() {
	filepath := "data/tasks.csv"
	file, err := os.Open(filepath)
	if err != nil {
		file, err = os.Create(filepath)
	}
	utils.HandleErr(err)
	defer file.Close()
}
