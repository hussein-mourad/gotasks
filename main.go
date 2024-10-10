/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/hussein-mourad/gotasks/internal/tasks"
	"github.com/hussein-mourad/gotasks/utils"
	"github.com/mergestat/timediff"
)

func main() {
	// GetFile()
	// cmd.Execute()
	taskStore := tasks.NewStore()
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)

	t1, _ := taskStore.CreateTask("Task")
	PrintTask(tw, t1)

	taskStore.MarkCompleted(t1.ID)

	fmt.Printf("\n\n")
	fmt.Fprintln(tw, "ID\tTask\tCreated")
	for _, t := range taskStore.Tasks {
		fmt.Fprintf(tw, "%v\t%v\t%v\n", t.ID, t.Task, formatTime(t.Created))
	}
	tw.Flush()

	fmt.Printf("\n\n")
	fmt.Fprintln(tw, "ID\tTask\tCreated\tCompleted")
	for _, t := range taskStore.Tasks {
		fmt.Fprintf(tw, "%v\t%v\t%v\t%v\n", t.ID, t.Task, formatTime(t.Created), t.Completed)
	}
	tw.Flush()
}

func formatTime(t time.Time) string {
	return timediff.TimeDiff(t)
}

func PrintTask(w *tabwriter.Writer, t tasks.Task) {
	fmt.Fprintln(w, "ID\tTask\tCreated")
	fmt.Fprintf(w, "%v\t%v\t%v\n", t.ID, t.Task, t.Created)
	w.Flush()
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
