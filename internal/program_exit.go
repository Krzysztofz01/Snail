package internal

func RegisterExitProgram() *Program {
	return &Program{
		Name: "exit",
		Execute: func(args []string) *EmitterResults {
			return &EmitterResults{
				Executed:                   true,
				ProgramExitCode:            0,
				PassProgramExitCodeToShell: true,
				Errors:                     make([]error, 0),
				ExternalProgram:            false,
			}
		},
	}
}
