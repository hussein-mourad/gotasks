/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hussein-mourad/gotasks/internal/tasks"
	"github.com/hussein-mourad/gotasks/utils"
)

func main() {
	// GetFile()
	// cmd.Execute()
	taskMap := make(map[int]tasks.Task)
	taskMap[1] = tasks.Task{ID: 1, Task: "Task 1", Completed: false, Created: time.Now().UTC()}
	taskMap[2] = tasks.Task{ID: 2, Task: "Task 2", Completed: false, Created: time.Now().UTC()}
	taskMap[3] = tasks.Task{ID: 3, Task: "Task 3", Completed: false, Created: time.Now().UTC()}
	taskMap[4] = tasks.Task{ID: 4, Task: "Task 4", Completed: false, Created: time.Now().UTC()}
	tasks.WriteTasks(taskMap)
	fmt.Println(tasks.ReadTasks())
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
