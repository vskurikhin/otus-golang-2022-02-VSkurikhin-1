package hw05parallelexecution

import (
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 23
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})

	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})
}

// Значение m <= 0 трактуется на усмотрение программиста.
// Считать это как "максимум 0 ошибок", значит функция всегда будет возвращать ErrErrorsLimitExceeded.
func TestMEqualOrLessZeroAlwaysErrErrorsLimitExceeded(t *testing.T) {
	setTests := []struct {
		fun      func() error
		expected error
	}{
		{fun: func() error {
			fmt.Println("Test 0")
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 1")
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 2")
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 3")
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 4")
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 5")
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 6")
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 7")
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 8")
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 9")
			return nil
		}},
	}
	tasks := make([]Task, 0, len(setTests))
	for _, t := range setTests {
		tasks = append(tasks, t.fun)
	}
	err := Run(tasks, 4, 0)
	require.Equal(t, err, ErrErrorsLimitExceeded)
	err1 := Run(tasks, 4, -1)
	require.Equal(t, err1, ErrErrorsLimitExceeded)
}

// -------ok-----ok-----ok-----ok  (1 воркер выполнил 4 задачи).
// -----------ok-------------ok    (2 воркер выполнил 2 задачи).
// -----ok---------ok---------ok   (3 воркер выполнил 3 задачи).
// --------------------ok          (4 воркер выполнил 1 задачу).
// Выполнится 10 задач (10 успешно): задач не осталось, воркеры остановились.

func TestRun10of10(t *testing.T) {
	var runTasksCount int32
	setTests := []struct {
		fun      func() error
		expected error
	}{
		{fun: func() error {
			time.Sleep(time.Millisecond * 7)
			atomic.AddInt32(&runTasksCount, 1)
			return nil
		}},
		{fun: func() error {
			time.Sleep(time.Millisecond * 11)
			atomic.AddInt32(&runTasksCount, 1)
			return nil
		}},
		{fun: func() error {
			time.Sleep(time.Millisecond * 5)
			atomic.AddInt32(&runTasksCount, 1)
			return nil
		}},
		{fun: func() error {
			time.Sleep(time.Millisecond * 20)
			atomic.AddInt32(&runTasksCount, 1)
			return nil
		}},
		{fun: func() error {
			time.Sleep(time.Millisecond * 5)
			atomic.AddInt32(&runTasksCount, 1)
			return nil
		}},
		{fun: func() error {
			time.Sleep(time.Millisecond * 13)
			atomic.AddInt32(&runTasksCount, 1)
			return nil
		}},
		{fun: func() error {
			time.Sleep(time.Millisecond * 9)
			atomic.AddInt32(&runTasksCount, 1)
			return nil
		}},
		{fun: func() error {
			time.Sleep(time.Millisecond * 9)
			atomic.AddInt32(&runTasksCount, 1)
			return nil
		}},
		{fun: func() error {
			time.Sleep(time.Millisecond * 5)
			atomic.AddInt32(&runTasksCount, 1)
			return nil
		}},
		{fun: func() error {
			time.Sleep(time.Millisecond * 5)
			atomic.AddInt32(&runTasksCount, 1)
			return nil
		}},
	}
	tasks := make([]Task, 0, len(setTests))
	for _, t := range setTests {
		tasks = append(tasks, t.fun)
	}
	err := Run(tasks, 4, 1)
	require.Nil(t, err)
	require.Equal(t, runTasksCount, int32(10), "not all tasks were completed")
}

// ------ok--------ok (узнал, что лимит превышен и остановился).
// -----------err.
// ---err.
// --------ok-------ok.

func TestRun6of10(t *testing.T) {
	setTests := []struct {
		fun      func() error
		expected error
	}{
		{fun: func() error {
			time.Sleep(time.Millisecond * 6)
			fmt.Println("Test 0")
			return nil
		}},
		{
			fun: func() error {
				fmt.Println("Test 1")
				time.Sleep(time.Millisecond * 11)
				return fmt.Errorf("Error 0 in Test 1")
			},
			expected: fmt.Errorf("Error 0 in Test 1"),
		},
		{
			fun: func() error {
				fmt.Println("Test 2")
				time.Sleep(time.Millisecond * 3)
				return fmt.Errorf("Error 1 in Test 2")
			},
			expected: fmt.Errorf("Error 1 in Test 2"),
		},
		{fun: func() error {
			fmt.Println("Test 3")
			time.Sleep(time.Millisecond * 8)
			return nil
		}},
		{fun: func() error {
			time.Sleep(time.Millisecond * 8)
			fmt.Println("Test 4")
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 5")
			time.Sleep(time.Millisecond * 7)
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 6")
			time.Sleep(time.Millisecond * 20)
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 7")
			time.Sleep(time.Millisecond * 20)
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 8")
			time.Sleep(time.Millisecond * 20)
			return nil
		}},
		{fun: func() error {
			fmt.Println("Test 9")
			time.Sleep(time.Millisecond * 20)
			return nil
		}},
	}
	tasks := make([]Task, 0, len(setTests))
	for _, t := range setTests {
		tasks = append(tasks, t.fun)
	}
	err := Run(tasks, 4, 2)
	require.Equal(t, err, ErrErrorsLimitExceeded)
}
