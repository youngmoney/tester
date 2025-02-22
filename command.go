package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func ExitIfNonZero(err interface{}) {
	if err != nil {
		if e, ok := err.(interface{ ExitCode() int }); ok {
			os.Exit(e.ExitCode())
		}
	}
}

func ExecuteCommand(command string, args []string) error {
	bashArgs := []string{"-s"}
	cmd := exec.Command("bash", append(bashArgs, args...)...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, command)
	}()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func ExecuteCommandCaptureStdout(command string, args []string) (string, error) {
	bashArgs := []string{"-c", command, "command"}
	cmd := exec.Command("bash", append(bashArgs, args...)...)

	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	cmd.Start()

	out, _ := io.ReadAll(stdout)

	return string(out), cmd.Wait()
}

func ExecuteCommandInteractive(command string, args []string) error {
	bashArgs := []string{"-c", command, "command"}
	cmd := exec.Command("bash", append(bashArgs, args...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
