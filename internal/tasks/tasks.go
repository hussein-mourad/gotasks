package tasks

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/hussein-mourad/gotasks/utils"
)

type Task struct {
	ID        int       `csv:"id"`
	Task      string    `csv:"task"`
	Completed bool      `csv:"completed"`
	Created   time.Time `csv:"created"`
}

func NewTask(id int, task string) *Task {
	return &Task{Task: task, Completed: false, Created: time.Now()}
}

func ReadTasks() map[int]Task {
	file, err := os.OpenFile("data/tasks.csv", os.O_CREATE|os.O_RDWR, 0o644)
	utils.HandleErr(err)
	defer file.Close()
	r := csv.NewReader(file)
	tasks := make(map[int]Task, 2)

	i := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV data:", err)
			break
		}
		if i == 0 {
			i++
			continue // Skip Reading the header
		}

		t := Task{}
		for j, field := range record {
			if j == 0 {
				t.ID, err = strconv.Atoi(field)
				utils.HandleErr(err)
			}
			if j == 1 {
				t.Task = field
			}
			if j == 2 {
				t.Completed, err = strconv.ParseBool(field)
				utils.HandleErr(err)
			}
			if j == 3 {
				t.Created, err = time.Parse(time.DateTime, field)
				utils.HandleErr(err)
			}
		}
		tasks[i] = t
		i++
	}
	return tasks
}

func WriteTasks(tasks map[int]Task) {
	file, err := os.OpenFile("data/tasks.csv", os.O_CREATE|os.O_RDWR, 0o644)
	utils.HandleErr(err)
	defer file.Close()
	w := csv.NewWriter(file)
	header := []string{"id", "task", "completed", "created"}
	records := [][]string{
		header,
	}
	for _, task := range tasks {
		record := []string{
			strconv.Itoa(task.ID),
			task.Task,
			strconv.FormatBool(task.Completed),
			task.Created.Format(time.DateTime),
		}
		records = append(records, record)
	}

	w.WriteAll(records)
}
