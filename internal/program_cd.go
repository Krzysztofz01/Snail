package internal

import "os"

func RegisterChangeDirectoryProgram() *Program {
	return &Program{
		Name: "cd",
		Execute: func(args []string) *EmitterResults {
			exitCode := 0
			targetPath := "~"

			if len(args) > 0 && len(args[0]) > 0 {
				targetPath = args[0]
			}

			errors := make([]error, 0)
			if err := os.Chdir(targetPath); err != nil {
				errors = append(errors, err)
				exitCode = 1
			}

			return &EmitterResults{
				Executed:                   true,
				ProgramExitCode:            exitCode,
				PassProgramExitCodeToShell: false,
				Errors:                     errors,
				ExternalProgram:            false,
			}
		},
	}
}
