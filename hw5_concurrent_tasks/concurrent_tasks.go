package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

// Run function runs functions from `[]func() error` with N concurrency.
// Returns an error if more than M functions return errors
func Run(task []func() error, N int, M int) error {
	queueChan := make(chan bool, N)
	shutdownChan := make(chan bool)
	errorsChan := make(chan bool, len(task))
	finishChan := make(chan bool, len(task))
	resultChan := make(chan bool)

	if M < 1 || N < 1 {
		return errors.New("wrong arguments")
	}

	go func(shutdown chan bool, errors <-chan bool, finish <-chan bool, result chan bool) {
		errorsCounter := 0
		finishCounter := 0
		for {
			select {
			case <-errors:
				errorsCounter++
				if errorsCounter == M {
					close(shutdown)
				}
			case <-finish:
				finishCounter++
				if finishCounter == len(task) {
					select {
					case <-shutdown:
						for {
							if len(queueChan) == 0 {
								break
							}
						}
						result <- false
						return
					default:
						result <- true
						return
					}
				}
			}
		}
	}(shutdownChan, errorsChan, finishChan, resultChan)

	for _, t := range task {
		queueChan <- true
		go wrapper(t, queueChan, errorsChan, finishChan, shutdownChan)
	}

	result := <-resultChan
	if result {
		return nil
	}
	return errors.New("too many errors")

}

func wrapper(f func() error, queueChan <-chan bool, errorChan chan<- bool, finishChan chan<- bool, shutdown <-chan bool) {
	select {
	default:
		err := f()
		if err != nil {
			errorChan <- true
			time.Sleep(100 * time.Millisecond) // time to prevent race condition
			finishChan <- true
		} else {
			finishChan <- true
		}
	case <-shutdown:
		finishChan <- true
	}

	<-queueChan
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
	tasks := []func() error{errorFunc, longNilFunc}
	fmt.Println(runtime.NumGoroutine())
	err := Run(tasks, 20, 12)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(err)
}
