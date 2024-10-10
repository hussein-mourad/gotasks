package tasks

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/hussein-mourad/gotasks/utils"
	"github.com/mergestat/timediff"
)

var w = tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)

type Store struct {
	Tasks []Task
}

func NewStore() *Store {
	s := &Store{Tasks: make([]Task, 0)}
	s.ReadTasks()
	return s
}

type Task struct {
	ID        int       `csv:"id"`
	Task      string    `csv:"task"`
	Completed bool      `csv:"completed"`
	Created   time.Time `csv:"created"`
}

var (
	ErrTaskNotFound = errors.New("Task not found.")
	TasksFile       = "data/tasks.csv"
)

func NewTask(id int, task string) *Task {
	return &Task{ID: id, Task: task, Completed: false, Created: time.Now().UTC()}
}

func (s *Store) CreateTask(Task string) (Task, error) {
	t := *NewTask(s.GetInsertID(), Task)
	s.Tasks = append(s.Tasks, t)
	return t, s.WriteTasks()
}

func (s *Store) GetInsertID() int {
	insertID := 1
	if len(s.Tasks) > 0 {
		insertID = s.Tasks[len(s.Tasks)-1].ID + 1
	}
	return insertID
}

func (s *Store) MarkCompleted(id int) (Task, error) {
	index, err := s.findByID(id)
	if err != nil {
		return Task{}, err
	}
	task := &s.Tasks[index]
	task.Completed = true
	return *task, s.WriteTasks()
}

func (s *Store) DeleteTask(id int) error {
	index, err := s.findByID(id)
	if err != nil {
		return ErrTaskNotFound
	}
	s.Tasks = append(s.Tasks[:index], s.Tasks[index+1:]...)
	return s.WriteTasks()
}

func (s *Store) findByID(id int) (int, error) {
	index := -1
	for i, t := range s.Tasks {
		if t.ID == id {
			index = i
			break
		}
	}
	if index != -1 {
		return index, nil
	}
	return index, ErrTaskNotFound
}

func PrintTask(t Task) {
	fmt.Fprintln(w, "ID\tTask\tCreated")
	fmt.Fprintf(w, "%v\t%v\t%v\n", t.ID, t.Task, formatTime(t.Created))
	w.Flush()
}

func PrintTaskCompleted(t Task) {
	fmt.Fprintln(w, "ID\tTask\tCreated\tCompleted")
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", t.ID, t.Task, formatTime(t.Created), t.Completed)
	w.Flush()
}

func (s *Store) ListTasks() {
	fmt.Fprintln(w, "ID\tTask\tCreated")
	for _, t := range s.Tasks {
		if !t.Completed {
			fmt.Fprintf(w, "%v\t%v\t%v\n", t.ID, t.Task, formatTime(t.Created))
		}
	}
	w.Flush()
}

func (s *Store) ListTasksCompleted() {
	fmt.Printf("\n\n")
	fmt.Fprintln(w, "ID\tTask\tCreated\tCompleted")
	for _, t := range s.Tasks {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", t.ID, t.Task, formatTime(t.Created), t.Completed)
	}
	w.Flush()
}

func formatTime(t time.Time) string {
	return timediff.TimeDiff(t)
}

func (s *Store) ReadTasks() {
	file, err := os.OpenFile(TasksFile, os.O_CREATE|os.O_RDWR, 0o644)
	utils.HandleErr(err)
	defer file.Close()
	r := csv.NewReader(file)

	i := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading CSV data:", err)
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
		s.Tasks = append(s.Tasks, t)
		i++
	}
}

func (s *Store) WriteTasks() error {
	file, err := os.Create(TasksFile)
	if err != nil {
		return err
	}
	defer file.Close()
	w := csv.NewWriter(file)
	header := []string{"id", "task", "completed", "created"}
	records := [][]string{
		header,
	}
	for _, task := range s.Tasks {
		record := []string{
			strconv.Itoa(task.ID),
			task.Task,
			strconv.FormatBool(task.Completed),
			task.Created.Format(time.DateTime),
		}
		records = append(records, record)
	}

	return w.WriteAll(records)
}
