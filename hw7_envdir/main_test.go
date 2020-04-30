package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvVars(t *testing.T) {
	type TD struct {
		fileNames    []string
		fileContents [][]byte
		result       []string
		err          bool
	}

	testData := []TD{
		{[]string{"testVar1"}, [][]byte{[]byte("1")}, []string{"testVar1=1"}, false},
		{[]string{"testVar1", "PATH"}, [][]byte{[]byte("1"), []byte{}}, []string{"testVar1=1", "PATH="}, false},
		{[]string{"Var"}, [][]byte{[]byte("\"VARIABLE\"")}, []string{"Var=\"VARIABLE\""}, false},
		{[]string{"URL"}, [][]byte{[]byte("https://www.google.com/search?q=gopher&source=lmns&bih=910&biw=1920&hl=ru&ved=2ahUKEwj9l_PDp5DpAhVBxSoKHXfID7wQ_AUoAHoECAEQAA")}, []string{"URL=https://www.google.com/search?q=gopher&source=lmns&bih=910&biw=1920&hl=ru&ved=2ahUKEwj9l_PDp5DpAhVBxSoKHXfID7wQ_AUoAHoECAEQAA"}, false},
		{[]string{}, [][]byte{}, []string{}, false},
	}

	_ = os.Mkdir("testEnv", 0755)

	for _, testCase := range testData {
		for index, fileName := range testCase.fileNames {
			_ = ioutil.WriteFile("testEnv/"+fileName, testCase.fileContents[index], 0644)
		}
		result, err := getEnvVars("testEnv")

		assert.ElementsMatch(t, testCase.result, result)
		if testCase.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

		for _, fileName := range testCase.fileNames {
			_ = os.Remove("testEnv/" + fileName)
		}
	}

	_ = os.Remove("testEnv")
}

func TestRunCommand(t *testing.T) {
	type TD struct {
		cmd      []string
		env      []string
		stdout   string
		stderr   string
		exitCode int
		err      bool
	}

	testData := []TD{
		{[]string{"sh", "-c", "echo $TERM"}, []string{"TERM=123"}, "123\n", "", 0, false},
		{[]string{"sh", "-c", "exit 3"}, []string{}, "", "", 3, true},
		{[]string{"sh", "-c", "echo $V1$V2"}, []string{"V1=VIC", "V2=TORY"}, "VICTORY\n", "", 0, false},
		{[]string{"sh", "-c", "echo $OUT; echo $ERR 1>&2"}, []string{"ERR=stderr", "OUT=stdout"}, "stdout\n", "stderr\n", 0, false},
		{[]string{"sh", "-c", "echo $ERR 1>&2; exit 1"}, []string{"ERR=some error"}, "", "some error\n", 1, true},
	}

	for _, testCase := range testData {
		stdout, stderr, exitCode, err := runCommand(testCase.cmd, testCase.env)
		assert.Equal(t, testCase.stdout, stdout)
		assert.Equal(t, testCase.stderr, stderr)
		assert.Equal(t, testCase.exitCode, exitCode)
		if testCase.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
