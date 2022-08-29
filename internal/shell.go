package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shell struct {
	inputStream  *os.File
	outputStream *os.File
	errorStream  *os.File
	prompt       *Prompt
	parser       *Parser
	emitter      *Emitter
}

func (shell *Shell) Init(inputStream *os.File, outputStream *os.File, errorStream *os.File) error {
	shell.inputStream = inputStream
	shell.outputStream = outputStream
	shell.errorStream = errorStream

	shell.prompt = new(Prompt)
	if err := shell.prompt.Init(); err != nil {
		return err
	}

	shell.parser = new(Parser)
	if err := shell.parser.Init(); err != nil {
		return err
	}

	shell.emitter = new(Emitter)
	if err := shell.emitter.Init(shell.inputStream, shell.outputStream, shell.errorStream); err != nil {
		return err
	}

	return nil
}

func (shell *Shell) StartShellLoop() *ShellResults {
	for {
		fmt.Print(*shell.prompt.GetPrompt())

		input, err := shell.getShellInput()
		if err != nil {
			// TODO: Error handle
		}

		tokens, errSlice := shell.parser.TokenizeInput(*input)
		if errSlice != nil {
			// TODO: Error handle
		}

		result := shell.emitter.Execute(*tokens)
		// TODO: Handle all result data

		if result.PassProgramExitCodeToShell {
			return &ShellResults{
				ExitCode: result.ProgramExitCode,
				Errors:   result.Errors,
			}
		}
	}
}

func (shell *Shell) getShellInput() (*string, error) {
	reader := bufio.NewReader(shell.inputStream)

	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	// Remove CR (Carriage Return) 0x0D
	inputNoCr := strings.Replace(input, "\r", "", -1)

	// Remove LF (Line Feed) 0x0A
	inputNoLf := strings.Replace(inputNoCr, "\n", "", -1)

	return &inputNoLf, nil
}

type ShellResults struct {
	ExitCode int
	Errors   []error
}
