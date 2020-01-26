package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	var tasks []func() error
	assert.Error(t, Run(tasks, 0, 0))
	assert.Error(t, Run(tasks, 1, 0))
	assert.Error(t, Run(tasks, 0, 1))

	tasks = []func() error{errorFunc}
	assert.Error(t, Run(tasks, 0, 0))
	assert.Error(t, Run(tasks, 1, 0))
	assert.Error(t, Run(tasks, 0, 1))
	assert.Error(t, Run(tasks, 1, 1))
	assert.Nil(t, Run(tasks, 1, 2))
	assert.Error(t, Run(tasks, 2, 1))
	assert.Nil(t, Run(tasks, 2, 2))

	tasks = []func() error{nilFunc}
	assert.Nil(t, Run(tasks, 1, 1))
	assert.Nil(t, Run(tasks, 1, 2))
	assert.Nil(t, Run(tasks, 2, 1))
	assert.Nil(t, Run(tasks, 2, 2))

	tasks = []func() error{nilFunc, errorFunc, errorFunc}
	assert.Error(t, Run(tasks, 1, 1))
	assert.Error(t, Run(tasks, 1, 2))
	assert.Nil(t, Run(tasks, 1, 3))
	assert.Error(t, Run(tasks, 2, 1))
	assert.Error(t, Run(tasks, 2, 2))
	assert.Nil(t, Run(tasks, 2, 3))
	assert.Error(t, Run(tasks, 3, 1))
	assert.Error(t, Run(tasks, 3, 2))
	assert.Nil(t, Run(tasks, 3, 3))
	assert.Error(t, Run(tasks, 4, 1))
	assert.Error(t, Run(tasks, 4, 2))
	assert.Nil(t, Run(tasks, 4, 3))
	assert.Nil(t, Run(tasks, 4, 4))

	countGoroutinesBefore := runtime.NumGoroutine()
	assert.Nil(t, Run(tasks, 4, 4))
	countGoroutinesAfter := runtime.NumGoroutine()
	assert.Equal(t, countGoroutinesBefore, countGoroutinesAfter)

	tasks = []func() error{longNilFunc, errorFunc, nilFunc}
	countGoroutinesBefore = runtime.NumGoroutine()
	assert.Error(t, Run(tasks, 3, 1))
	countGoroutinesAfter = runtime.NumGoroutine()
	assert.Equal(t, countGoroutinesBefore, countGoroutinesAfter)

	countGoroutinesBefore = runtime.NumGoroutine()
	assert.Nil(t, Run(tasks, 3, 2))
	countGoroutinesAfter = runtime.NumGoroutine()
	assert.Equal(t, countGoroutinesBefore, countGoroutinesAfter)
}
