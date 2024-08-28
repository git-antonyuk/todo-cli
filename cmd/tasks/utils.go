package tasks

import (
	"strconv"
	"time"
	file_utils "todo-console/cmd/utils/file"
)

func getList() file_utils.TCsvData {
	csvList, err := file_utils.ReadFromCSV("./data/db.csv")
	if err != nil {
		csvList = [][]string{}
	}
	return csvList
}

func convertCsvRowToTaskStruct(columnIndex int64, value string, task *Task) error {
	var _err error = nil
	switch columnIndex {
	case 0:
		{
			id, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				_err = err
			}
			task.ID = id
		}
	case 1:
		{
			task.description = value
		}
	case 2:
		{
			createdAt, err := time.Parse(time.RFC3339, value)
			if err != nil {
				_err = err
			}
			task.createdAt = createdAt
		}
	case 3:
		isComplete, err := strconv.ParseBool(value)
		if err != nil {
			_err = err
		}
		task.isComplete = isComplete
	}
	return _err
}

func getListTasks() Tasks {
	csvList := getList()
	listLength := int64(len(csvList))

	if listLength == 0 {
		return []Task{}
	}

	var tasks []Task

	for rowIndex, row := range csvList {
		if rowIndex == 0 {
			continue
		}
		var task Task

		// TODO: Add separate function for this logic
		for columnIndex, value := range row {
			err := convertCsvRowToTaskStruct(int64(columnIndex), value, &task)

			if err != nil {
				continue
			}
		}
		tasks = append(tasks, task)
	}

	return tasks
}

func createTaskRow(id int64, description string, isComplete bool, createdAt time.Time) []string {
	task := Task{
		ID:          id,
		description: description,
		isComplete:  isComplete,
		createdAt:   createdAt,
	}
	createdAtStr := task.createdAt.Format(time.RFC3339)
	isCompleteStr := strconv.FormatBool(task.isComplete)
	idStr := strconv.FormatInt(task.ID, 10)
	taskRow := []string{idStr, task.description, createdAtStr, isCompleteStr}

	return taskRow
}

func findTaskById(id int64) (int64, bool) {
	for index, value := range getListTasks() {
		if value.ID == id {
			return int64(index), true
		}
	}

	return -1, false
}

func getLastRowId() int64 {
	tasks := getListTasks()
	length := len(tasks)
	if length == 0 {
		return 0
	}
	lastRow := tasks[length-1]
	return lastRow.ID + 1
}
