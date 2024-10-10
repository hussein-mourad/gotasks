/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var showAll bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if showAll {
			store.ListTasksCompleted()
		} else {
			store.ListTasks()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Display all tasks, including completed")
}
