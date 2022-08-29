package internal

import (
	"errors"
	"strings"
)

type Parser struct {
}

// Initializer the Parser struct
func (parser *Parser) Init() error {
	return nil
}

// Split input string into tokens and return a pointer to the token array
func (parser *Parser) TokenizeInput(input string) (*[]string, *[]error) {
	tokens := make([]string, 0)
	tokenBuilder := strings.Builder{}

	parserError := new(parserError)
	if err := parserError.Init(); err != nil {
		return nil, &[]error{err}
	}

	parserState := new(parserState)
	if err := parserState.Init(); err != nil {
		parserError.PushError(err)
		return nil, parserError.GetErrors()
	}

	for _, char := range input {
		switch char {
		case ' ':
			parserError.PushError(parser.handleSpaceChar(char, &tokens, &tokenBuilder, parserState))
		case '"':
			parserError.PushError(parser.handleDoubleQuoteChar(char, &tokens, &tokenBuilder, parserState))
		default:
			parserError.PushError(parser.handleDefaultChar(char, &tokens, &tokenBuilder, parserState))
		}
	}

	if tokenBuilder.Len() > 0 {
		tokens = append(tokens, tokenBuilder.String())
	}

	if !parserState.IsStateValid() {
		parserError.PushError(errors.New("parser: tokenization finished witha invalid parser state"))
	}

	if parserError.ErrorsPresent() {
		return nil, parserError.GetErrors()
	}

	return &tokens, nil
}

// (Space) - Parser char handler function
func (parser *Parser) handleSpaceChar(char rune, tokens *[]string, tokenBuilder *strings.Builder, state *parserState) error {
	if state.DoubleQuoteOpen {
		tokenBuilder.WriteRune(char)
		return nil
	}

	*tokens = append(*tokens, tokenBuilder.String())
	tokenBuilder.Reset()

	return nil
}

// " (Double quote) - Parser char handler function
func (parser *Parser) handleDoubleQuoteChar(char rune, tokens *[]string, tokenBuilder *strings.Builder, state *parserState) error {
	state.DoubleQuoteOpen = !state.DoubleQuoteOpen
	return nil
}

// DEFAULT - Parser char handler function
func (parser *Parser) handleDefaultChar(char rune, tokens *[]string, tokenBuilder *strings.Builder, state *parserState) error {
	tokenBuilder.WriteRune(char)
	return nil
}

// Structure that represents the state of the parser state-machine
type parserState struct {
	DoubleQuoteOpen bool
	SingleQuoteOpen bool
}

// Initialzie the parserState struct with default values
func (parserState *parserState) Init() error {
	parserState.DoubleQuoteOpen = false
	parserState.SingleQuoteOpen = false

	return nil
}

// This function returns a boolean value indicating if the current state is valid
func (parserState *parserState) IsStateValid() bool {
	if parserState.SingleQuoteOpen {
		return false
	}

	if parserState.DoubleQuoteOpen {
		return false
	}

	return true
}

// Structure that holds all the errors raised during the parsing process
type parserError struct {
	parsingErrors []error
}

// Initialize the parserError struct with default values
func (parserError *parserError) Init() error {
	parserError.parsingErrors = make([]error, 0)
	return nil
}

// Push a error to the parserError struct, error is ignored if nil
func (parserError *parserError) PushError(err error) {
	if err != nil {
		parserError.parsingErrors = append(parserError.parsingErrors, err)
	}
}

// Retrieve error slice containing all errors from parserError
func (parserError *parserError) GetErrors() *[]error {
	return &parserError.parsingErrors
}

// Return a boolean value indication if any errors were collected
func (parserError *parserError) ErrorsPresent() bool {
	return len(parserError.parsingErrors) > 0
}
