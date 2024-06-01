package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"ascii/arguments"
	"ascii/readfile"

	"github.com/stretchr/testify/assert" // Using testify for assertions
)

// MockExitHandler is a mock implementation of os.Exit to capture the exit status.
type MockExitHandler struct {
	status        int
	name          string
	args          []string
	bannerFile    string
	expectedLines []string
	expectedExit  int
}

var MockExitHandlers = []MockExitHandler{
	{
		name:          "Valid Arguments and Banner File",
		args:          []string{"programName", "bannerfile.txt"},
		bannerFile:    "bannerfile.txt",
		expectedLines: []string{"Line 1", "Line 2", "Line 3"},
		expectedExit:  0,
	},
	{
		name:         "Missing Arguments",
		args:         []string{"programName"},
		expectedExit: 0,
	},
	{
		name:         "Invalid Banner File",
		args:         []string{"programName", "invalidfile.txt"},
		bannerFile:   "invalidfile.txt",
		expectedExit: 0,
	},
}

func (m *MockExitHandler) Exit(code int) {
	m.status = code
}

func TestMain(t *testing.T) {
	for _, tc := range MockExitHandlers {
		t.Run(tc.name, func(t *testing.T) {
			// Capture os.Exit calls using MockExitHandler
			exitHandler := &MockExitHandler{}
			osExit = exitHandler.Exit
			defer func() { osExit = os.Exit }() // Restore osExit after the test

			// Mock readfile.ReadFile function
			readfileMock := func(filename string) ([]byte, error) {
				if filename == tc.bannerFile {
					return []byte(strings.Join(tc.expectedLines, "\n")), nil
				}
				return nil, fmt.Errorf("no such file")
			}

			// Mock arguments.HandleNewLines function
			argumentsMock := func(args []string) string {
				return strings.Join(args[1:], " ")
			}

			readfile.ReadFile = readfileMock
			arguments.HandleNewLines = argumentsMock

			// Execute the main function with test arguments
			os.Args = tc.args
			main()

			// Assert the expected behavior
			assert.Equal(t, tc.expectedExit, exitHandler.status)

			if tc.expectedExit == 0 {
				// Check the output (ASCII art)
				expectedOutput := strings.Join(tc.expectedLines, "\n") + "\n" + strings.Join(tc.args[1:], " ") + "\n"
				assert.Equal(t, expectedOutput, buf.String())
			}
		})
	}
}
