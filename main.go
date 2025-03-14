package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func commandTest(args []string, all bool, tests *[]Test, log_reader *LogReader) (bool, error) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	match := Match(cwd, tests)

	if match == nil {
		return false, nil
	}

	if match.PreTestCommand != "" {
		cmdErr := ExecuteCommandInteractive(match.PreTestCommand, []string{})
		ExitIfNonZero(cmdErr)
	}

	var testError error

	if all {
		if match.TestAllCommand == "" {
			fmt.Println("no all command")
			os.Exit(1)
		}

		testError = ExecuteCommandInteractive(match.TestAllCommand, []string{})
	} else {
		if match.TestCommand == "" {
			fmt.Println("no test command")
			os.Exit(1)
		}

		testError = ExecuteCommandInteractive(match.TestCommand, args)
	}
	if testError == nil {
		return false, nil
	}

	if match.FailedLogListCommand != "" {
		logs, cmdErr := ExecuteCommandCaptureStdout(match.FailedLogListCommand, []string{})
		if cmdErr != nil {
			fmt.Println("failed to list logs")
			os.Exit(1)
		}

		logFiles := strings.Split(strings.TrimSpace(logs), "\n")

		if len(logFiles) > 0 && log_reader.Command != "" {
			logCmdError := ExecuteCommandInteractive(log_reader.Command, logFiles)
			if logCmdError != nil {
				fmt.Println("failed to read logs")
				return false, logCmdError
			}

			return true, testError
		}

	}

	return false, testError

}

func commandLoop(args []string, all bool, tests *[]Test, log_reader *LogReader) {
	for {
		// Clear screen
		fmt.Print("\033[H\033[2J")

		implicitContinue, err := commandTest(args, all, tests, log_reader)
		if err == nil {
			return
		}

		if !implicitContinue {
			fmt.Print("\n\n[continue]")
			fmt.Scanln()
		}
	}
}

func main() {
	configFilename := flag.String("config", os.Getenv("TESTER_CONFIG"), "config file (yaml), or set TESTER_CONFIG")
	flag.Parse()

	config := ReadConfig(*configFilename)

	if flag.Arg(0) != "test" {
		fmt.Println("test is the only supported command")
		os.Exit(1)
	}

	fs := flag.NewFlagSet("test", flag.ExitOnError)
	all := fs.Bool("a", false, "run all tests")
	loop := fs.Bool("l", false, "run the test, show the logs, loop")
	fs.Parse(flag.Args()[1:])

	var args = fs.Args()
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}

	if *loop {
		commandLoop(args, *all, &config.Tester.Tests, &config.Tester.LogReader)
	} else {
		_, testError := commandTest(args, *all, &config.Tester.Tests, &config.Tester.LogReader)
		ExitIfNonZero(testError)
	}
}
