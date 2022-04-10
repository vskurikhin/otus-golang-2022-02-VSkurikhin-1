package hw05parallelexecution

import (
	"log"
	"sync"
)

// Структура оперирует ошибками и синхронизацией.
// Инкапсулирует канал с ошибками и  и предоставляет методы доступа.
type ErrorsSync struct {
	debug  bool
	errors chan error
	wg     sync.WaitGroup
}

func NewErrorsSync(n int) *ErrorsSync {
	return newErrorsSync(n, false)
}

func newErrorsSync(n int, debug bool) *ErrorsSync {
	s := ErrorsSync{
		debug:  debug,
		errors: make(chan error),
	}
	s.wg.Add(n)

	go func() {
		s.wg.Wait()
		close(s.errors)
	}()
	return &s
}

func (s *ErrorsSync) Send(err error) {
	s.errors <- err
}

func (s *ErrorsSync) Done() {
	s.wg.Done()
}

func (s *ErrorsSync) Wait() {
	s.wg.Wait()
}

func (s *ErrorsSync) LogErrors() {
	for e := range s.errors {
		if s.debug {
			log.Printf("error: %v", e)
		}
	}
}
