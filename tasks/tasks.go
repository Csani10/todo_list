package tasks

import (
	"fmt"
	"strings"
	"time"
)

type Task struct {
	task     string
	added    time.Time
	due      time.Time
	finished bool
}

func (t *Task) Finish() {
	t.finished = true
}

func (t *Task) TimeLeft() time.Duration {
	now := time.Now()
	if t.due.Before(now) {
		return 0
	}

	return t.due.Sub(now)
}

func (t *Task) GetTask() string {
	return t.task
}

func (t *Task) ModifyTask(task string) {
	t.task = task
}

func (t *Task) Serialize() string {
	str := t.task + ";&;" + t.due.String() + ";&;" + t.added.String() + ";&;"
	if t.finished {
		str += "1"
	} else {
		str += "0"
	}

	return str
}

func NewTask(task string, due time.Time) *Task {
	return &Task{
		task:     task,
		added:    time.Now(),
		due:      due,
		finished: false,
	}
}

func Deserialize(task_string string) *Task {
	task_split := strings.Split(task_string, ";&;")

	task := task_split[0]
	due, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", task_split[1])
	if err != nil {
		fmt.Println("Error while parsing task due date")
		return nil
	}

	added, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", task_split[2])
	if err != nil {
		fmt.Println("Error while parsing task added date")
		return nil
	}

	finished := false
	if task_split[3] == "1" {
		finished = true
	}

	return &Task{
		task:     task,
		added:    added,
		due:      due,
		finished: finished,
	}
}
