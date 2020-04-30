package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	name    string
	version string
	commit  string
)

func main() {
	dir, cmd, function, err := parseArgs()
	if err != nil {
		log.Println(err)
		printHelp()
		os.Exit(0)
	}
	if function != nil {
		function()
		os.Exit(0)
	}
	//fmt.Println(dir)
	//fmt.Println(cmd)

	envVars, err := getEnvVars(dir)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(envVars)
	stdOut, stdErr, exitCode, err := runCommand(cmd, envVars)
	if err != nil {
		if exitCode != 0 {
			fmt.Printf("%v", stdOut)
			fmt.Fprintf(os.Stderr, "%v", stdErr)
			os.Exit(exitCode)
		}
		log.Fatal(err)
	}
	fmt.Printf("%v", stdOut)
	fmt.Fprintf(os.Stderr, "%v", stdErr)
}

func parseArgs() (string, []string, func(), error) {
	args := os.Args[1:]
	if len(args) == 1 && (args[0] == "--version" || args[0] == "-v") {
		return "", []string{}, printVersion, nil
	} else if len(args) == 1 && (args[0] == "--help" || args[0] == "-h") {
		return "", []string{}, printHelp, nil
	} else if len(args) < 2 {
		err := errors.New("too few arguments")
		return "", []string{}, nil, err
	}
	return args[0], args[1:], nil, nil
}

func printHelp() {
	fmt.Printf("Usage: %v DIR COMMAND\n"+
		"%v - runs another program with environment modified according to files in a specified directory.\n"+
		"DIR is a single argument. COMMAND consists of one or more arguments.\n"+
		"%v sets various environment variables as specified by files in the directory named DIR. It then runs COMMAND.\n"+
		"If DIR contains a file named s whose first line is t, %s removes an environment variable\n"+
		"named s if one exists, and then adds an environment variable named s with value t.\n"+
		"The name s must not contain =.\n", name, name, name, name)
}

func printVersion() {
	fmt.Printf("%v %v %v\n", name, version, commit)
}

func getEnvVars(dir string) ([]string, error) {
	_, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	var envVars []string
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		//fmt.Println("path=" + path)
		if !info.IsDir() {
			dat, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			envVars = append(envVars, fmt.Sprintf("%s=%s", info.Name(), string(dat)))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return envVars, nil
}

func runCommand(cmd []string, env []string) (string, string, int, error) {
	exitCode := 0
	command := exec.Command(cmd[0])
	if len(cmd) > 1 {
		command = exec.Command(cmd[0], cmd[1:]...)
	}
	command.Env = append(os.Environ(), env...)
	var stdOut, stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr

	if isInputFromPipe() {
		r, w := io.Pipe()
		go func() {
			defer w.Close()
			_, err := io.Copy(w, os.Stdin)
			if err != nil {
				log.Fatal(err)
			}
		}()

		command.Stdin = r
	}

	err := command.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		}
		return string(stdOut.Bytes()), string(stdErr.Bytes()), exitCode, err
	}

	return string(stdOut.Bytes()), string(stdErr.Bytes()), exitCode, nil
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}
