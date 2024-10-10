/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"os"

	"github.com/hussein-mourad/gotasks/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add \"Your Task\"",
	Short: "add Tasks to the list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile("data/tasks.csv", os.O_APPEND|os.O_WRONLY, 0o644)
		utils.HandleErr(err)
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		Task{}

		row := []string{"David", "30", "Male"}
		err = writer.Write(row)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
