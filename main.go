package main

import (
	"Snail/internal"
	"os"
)

type ExitCode int32

const (
	Success      ExitCode = 0
	FailureMinor ExitCode = 1
	FailureMajor ExitCode = 2
)

func main() {
	shell := shellInitialization()
	if shell == nil {
		// TODO: Error panic
		os.Exit(int(FailureMajor))
	}

	results := shell.StartShellLoop()
	if len(results.Errors) > 0 {
		// TODO: Error panic
		os.Exit(int(FailureMinor))
	}

	if err := shellCleanup(shell); err != nil {
		// TODO: Error panic
		os.Exit(int(FailureMajor))
	}

	os.Exit(int(Success))
}

func shellInitialization() *internal.Shell {
	// TODO: Import config from file
	// TODO: Apply config to shell instance
	shell := new(internal.Shell)
	if err := shell.Init(os.Stdin, os.Stdout, os.Stderr); err != nil {
		// TODO: Error panic
		return nil
	}

	return shell
}

func shellCleanup(shell *internal.Shell) error {
	// TODO: Config save
	// TODO: Cleanup

	return nil
}
