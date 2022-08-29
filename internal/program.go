package internal

// Strcture representig the built-in shell program
type Program struct {
	Name    string
	Execute func(args []string) *EmitterResults
}
