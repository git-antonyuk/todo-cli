package tasks

import (
	"fmt"
	"testing"
	testing_utils "todo-console/cmd/utils/testing"
)

func TestTodoAdd(t *testing.T) {
	FILE_PATH = "../../data/test_db.csv"
	t.Run("Add one task | create new file", func(t *testing.T) {
		err := Add("Test task #1")
		if err != nil {
			t.Errorf("Add function not working: %v", err)
		}
		testing_utils.DeleteFile(FILE_PATH, t)
	})

	t.Run("Add several taks | create new file", func(t *testing.T) {
		for index := range 3 {
			err := Add(fmt.Sprintf("Test task #%v", index))
			if err != nil {
				t.Errorf("Add function not working: %v", err)
			}
		}

		testing_utils.DeleteFile(FILE_PATH, t)
	})
}
