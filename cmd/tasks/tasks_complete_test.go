package tasks

import (
	"testing"
	testing_utils "todo-console/cmd/utils/testing"
)

func TestTodoComplete(t *testing.T) {
	FILE_PATH = "../../data/test_db.csv"
	t.Run("Complete task | not exist", func(t *testing.T) {
		err := Complete(0)
		if err == nil {
			t.Errorf("Complete function not working: %v", err)
		}
		testing_utils.DeleteFile(FILE_PATH, t)
	})

	t.Run("Complete task", func(t *testing.T) {
		err := Add("Test task")
		if err != nil {
			t.Errorf("Add function not working: %v", err)
		}

		err = Complete(0)
		if err != nil {
			t.Errorf("Complete function not working: %v", err)
		}

		testing_utils.DeleteFile(FILE_PATH, t)
	})
}
