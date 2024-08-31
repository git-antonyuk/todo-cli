package tasks

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
	testing_utils "todo-console/cmd/utils/testing"
)

func getExpectedBuf(extended bool, timeValue string) bytes.Buffer {
	var expectedBuf bytes.Buffer
	wr := tabwriter.NewWriter(&expectedBuf, 1, 2, 3, ' ', tabwriter.Debug)
	if extended {
		fmt.Fprintln(wr, " ID \t Task \t Created \t Done")
		for index := 0; index < 3; index++ {
			fmt.Fprintf(wr, " %v \t Test task #%v \t %v \t false \n", index, index, timeValue)
		}
	} else {
		fmt.Fprintln(wr, " ID \t Task \t Created ")
		for index := 0; index < 3; index++ {
			fmt.Fprintf(wr, " %v \t Test task #%v \t %v \n", index, index, timeValue)
		}
	}

	wr.Flush()

	return expectedBuf
}

func getOutputList(t *testing.T, extended bool) {
	// Create tasks
	for index := range 3 {
		err := Add(fmt.Sprintf("Test task #%v", index))
		if err != nil {
			t.Errorf("Add function not working: %v", err)
		}
	}

	// Create a pipe to capture stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	// Redirect stdout to the write end of the pipe
	oldStdout := os.Stdout
	os.Stdout = w

	// Call the List function
	List(extended)

	// Restore stdout and close the write end of the pipe
	os.Stdout = oldStdout
	w.Close()

	// Read the output from the read end of the pipe
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}

	// Close the read end of the pipe
	r.Close()

	expectedBuf := getExpectedBuf(extended, "in a few seconds")
	expectedBuf2 := getExpectedBuf(extended, "a few seconds ago")

	// Compare the captured output with the expected output
	got := buf.String()
	expected := expectedBuf.String()
	expected2 := expectedBuf2.String()

	if got != expected && got != expected2 {
		t.Errorf("List() = %v, want %v", got, expected)
	}

	testing_utils.DeleteFile(FILE_PATH, t)
}

func TestTodoList(t *testing.T) {
	t.Run("Get list output", func(t *testing.T) {
		getOutputList(t, true)
		getOutputList(t, false)

	})
}
