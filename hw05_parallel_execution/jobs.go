package hw05parallelexecution

import "log"

// Структура для задачь.
// Инкапсулирует канал с задачами и предоставляет методы доступа.
type Jobs struct {
	debug bool
	jobs  chan Task
}

func NewJobs() *Jobs {
	return newJobs(false)
}

func newJobs(debug bool) *Jobs {
	return &Jobs{
		debug: debug,
		jobs:  make(chan Task),
	}
}

func (j *Jobs) AddTasks(tasks []Task) {
	go func() {
		for _, task := range tasks {
			j.jobs <- task
		}
		close(j.jobs)
	}()
}

func (j *Jobs) Get() *chan Task {
	return &j.jobs
}

func (j *Jobs) Clear() {
	count := 0
	for range j.jobs {
		count++
	}
	if j.debug {
		log.Printf("skip %d tasks", count)
	}
}
