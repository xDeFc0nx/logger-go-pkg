package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"
	"testing"
)

func TestInfoFunctionUsage(t *testing.T) {
	// Create a buffer to capture the log output
	var buf bytes.Buffer

	// Redirect os.Stdout to the buffer
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Use a WaitGroup to wait for the Info function to complete
	var wg sync.WaitGroup
	wg.Add(1)

	// Ensure the original os.Stdout is restored at the end of the test
	defer func() {
		// Restore original stdout and close the write end
		os.Stdout = originalStdout
		w.Close()

		// Wait for Info function to complete
		wg.Wait()

		// Read from the pipe and write to the buffer
		out, err := io.ReadAll(r)
		if err != nil {
			t.Errorf("Error reading from pipe: %v", err)
		}
		buf.Write(out)
	}()

	// Define a message to be logged
	message := "This is an info message."

	// Call the Info function with the message
	go func() {
		defer wg.Done()
		Info(message)
	}()

	// Check if the log contains the expected message with the correct prefix
	logOutput := buf.String()
    matched, err := regexp.MatchString(`INFO:.*`, logOutput)
    	if err != nil {
			t.Errorf("Error reading from pipe: %v", err)
		}
	fmt.Println(matched, err)
    

}
