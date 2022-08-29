package internal

import (
	"os"
	"os/user"
	"strings"
)

type Prompt struct {
}

func (prompt *Prompt) Init() error {
	return nil
}

func (prompt *Prompt) GetPrompt() *string {
	username := "Unknown user"
	workingDirectory := "Unknown directory"

	if user, err := user.Current(); err == nil && user != nil {
		username = user.Username
	}

	if dir, err := os.Getwd(); err == nil {
		workingDirectory = dir
	}

	promptBuilder := strings.Builder{}

	promptBuilder.WriteString(username)
	promptBuilder.WriteString(" λ ")
	promptBuilder.WriteString(workingDirectory)
	promptBuilder.WriteString(" > ")

	formatedPrompt := promptBuilder.String()
	return &formatedPrompt
}
