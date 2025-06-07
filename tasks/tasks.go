package tasks

import (
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

func NewTask(task string, due time.Time) *Task {
	return &Task{
		task:     task,
		added:    time.Now(),
		due:      due,
		finished: false,
	}
}
