package hw05parallelexecution

import (
	"errors"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type countHolder struct {
	ops int32
}

// Run запускает задачи в n горутинах и останавливает свою работу при получении m ошибок от задач.
func Run(tasks []Task, n, m int) error {
	// создавайте все задачи и отправить их в канал.
	jobs := NewJobs()
	jobs.AddTasks(tasks)

	counter := &countHolder{}

	// создавать и запланировать результаты закрытия, когда вся работа будет выполнена.
	result := NewErrorsSync(n)

	for i := 0; i < n; i++ {
		go worker(jobs, result, counter, uint(m))
	}
	result.LogErrors()

	if counter.get() >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

// Логика воркера.
func worker(jobs *Jobs, result *ErrorsSync, counter *countHolder, m uint) {
	defer result.Done()
	for task := range *jobs.Get() {
		if counter.get() >= int32(m) {
			jobs.Clear()
			return
		}
		err := task()
		result.Send(err)
		if err != nil {
			counter.inc()
		}
	}
}

func (c *countHolder) get() int32 {
	return atomic.LoadInt32(&c.ops)
}

func (c *countHolder) inc() {
	atomic.AddInt32(&c.ops, 1)
}
