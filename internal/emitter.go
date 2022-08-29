package internal

import (
	"os"
	"os/exec"
	"strings"
)

type Emitter struct {
	inputStream     *os.File
	outputStream    *os.File
	errorStream     *os.File
	builtInPrograms []Program
}

// Initialize the Emitter struct
func (emitter *Emitter) Init(inputStream *os.File, outputStream *os.File, errorStream *os.File) error {
	emitter.inputStream = inputStream
	emitter.outputStream = outputStream
	emitter.errorStream = errorStream

	// TODO: Implement ,,clear''
	emitter.builtInPrograms = []Program{
		*RegisterChangeDirectoryProgram(),
		*RegisterExitProgram(),
	}

	return nil
}

// Execute the command represented by the parsed tokens and return the result structure instance
func (emitter *Emitter) Execute(tokenizedArgs []string) *EmitterResults {
	if len(tokenizedArgs) < 1 {
		return &EmitterResults{
			Executed:                   false,
			ProgramExitCode:            0,
			PassProgramExitCodeToShell: false,
			Errors:                     nil,
			ExternalProgram:            false,
		}
	}

	programName := tokenizedArgs[0]
	programArgs := tokenizedArgs[1:]

	if wasBuiltIn, results := emitter.executeBuiltInProgram(programName, programArgs); wasBuiltIn {
		return results
	}

	return emitter.executeExternalProgram(programName, programArgs)

}

// Execute the command as build in program and return the result struct instance
func (emitter *Emitter) executeBuiltInProgram(programName string, args []string) (bool, *EmitterResults) {
	if len(emitter.builtInPrograms) == 0 {
		return false, nil
	}

	for _, program := range emitter.builtInPrograms {
		if strings.EqualFold(program.Name, programName) {
			return true, program.Execute(args)
		}
	}

	return false, nil
}

// Execute the command as external program and return the result struct instance
func (emitter *Emitter) executeExternalProgram(programName string, args []string) *EmitterResults {
	command := exec.Command(programName, args...)

	command.Stdin = emitter.inputStream
	command.Stdout = emitter.outputStream
	command.Stderr = emitter.errorStream

	errors := make([]error, 0)
	if err := command.Run(); err != nil {
		errors = append(errors, err)
	}

	return &EmitterResults{
		Executed: true,
		// TODO: Access exit code
		ProgramExitCode:            0,
		PassProgramExitCodeToShell: false,
		Errors:                     errors,
		ExternalProgram:            true,
	}
}

// Structure representing the result of a command/program execution
type EmitterResults struct {
	Executed                   bool
	ProgramExitCode            int
	PassProgramExitCodeToShell bool
	Errors                     []error
	ExternalProgram            bool
}
