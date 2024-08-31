package tasks

import (
	"fmt"
	"testing"
	testing_utils "todo-console/cmd/utils/testing"
)

func TestTodoDelete(t *testing.T) {
	FILE_PATH = "../../data/test_db.csv"
	t.Run("Complete task | not exist", func(t *testing.T) {
		err := Delete(0)
		if err == nil {
			t.Errorf("Delete function not working: %v", err)
		}
		testing_utils.DeleteFile(FILE_PATH, t)
	})

	t.Run("Complete task", func(t *testing.T) {
		for index := range 3 {
			err := Add(fmt.Sprintf("Test task #%v", index))
			if err != nil {
				t.Errorf("Add function not working: %v", err)
			}
		}

		err := Delete(0)
		if err != nil {
			t.Errorf("Delete function not working: %v", err)
		}

		testing_utils.DeleteFile(FILE_PATH, t)
	})
}
