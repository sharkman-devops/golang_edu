package main

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Run function runs functions from `[]func() error` with N concurrency.
// Returns an error if more than M functions return errors
func Run(task []func() error, N int, M int) error {
	taskChan := make(chan func() error, len(task))
	shutdownChan := make(chan bool)
	errorsChan := make(chan bool, len(task))
	waitGroup := sync.WaitGroup{}
	workersPool := N
	if len(task) < N {
		workersPool = len(task)
	}

	if M < 1 || N < 1 {
		return errors.New("wrong arguments")
	}

	for _, t := range task {
		taskChan <- t
	}
	close(taskChan)

	for i := 0; i < workersPool; i++ {
		waitGroup.Add(1)
		go worker(taskChan, &waitGroup, errorsChan, shutdownChan, M)
	}

	waitGroup.Wait()
	if len(errorsChan) >= M {
		return errors.New("too many errors")
	}
	return nil
}

func worker(tasks <-chan func() error, wg *sync.WaitGroup, errors chan<- bool, shutdown <-chan bool, maxErrors int) {
	for {
		select {
		case <-shutdown:
			wg.Done()
			return
		default:
		}

		select {
		case f, ok := <-tasks:
			if !ok {
				wg.Done()
				return
			}

			if len(errors) >= maxErrors {
				wg.Done()
				return
			}

			err := f()
			if err != nil {
				errors <- true
			}

		case <-shutdown:
			wg.Done()
			return
		}
	}
}

func nilFunc() error {
	time.Sleep(1 * time.Second)
	fmt.Println("nil func")
	return nil
}

func longNilFunc() error {
	fmt.Println("long nil func started")
	time.Sleep(5 * time.Second)
	fmt.Println("long nil func finished")
	return nil
}

func errorFunc() error {
	time.Sleep(1 * time.Second)
	fmt.Println("err func")
	return errors.New("some error")
}

func main() {
	tasks := []func() error{nilFunc, errorFunc, errorFunc}
	fmt.Println(runtime.NumGoroutine())
	err := Run(tasks, 1, 1)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(err)
}
