package tasks

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
	file_utils "todo-console/cmd/utils/file"

	"github.com/mergestat/timediff"
)

const filePath = "./data/db.csv"

type Task struct {
	ID          int64
	description string
	createdAt   time.Time
	isComplete  bool
}

type Tasks = []Task

var COLUMNS_NAMES = []string{
	"ID", "Description", "CreatedAt", "IsComplete",
}

func Add(description string) error {
	var data file_utils.TCsvData

	csvList := getList()

	listLength := int64(len(csvList))

	if listLength == 0 {
		data = append(data, COLUMNS_NAMES)
	} else {
		data = append(data, csvList...)
	}

	lastRowId := getLastRowId()
	taskRow := createTaskRow(lastRowId, description, false, time.Now())

	data = append(data, taskRow)

	err := file_utils.WriteToCSV(filePath, data)
	if err != nil {
		return err
	}

	return nil
}

func List(showAll bool) {
	w := tabwriter.NewWriter(os.Stdout, 1, 2, 3, ' ', tabwriter.Debug)
	defer w.Flush()

	if showAll {
		fmt.Fprintln(w, " ID \t Task \t Created  \t Done")
	} else {
		fmt.Fprintln(w, " ID \t Task \t Created ")
	}

	for _, task := range getListTasks() {
		duration := time.Since(task.createdAt)
		timeDiff := timediff.TimeDiff(time.Now().Add(duration))
		idStr := strconv.FormatInt(task.ID, 10)
		isCompleteStr := strconv.FormatBool(task.isComplete)
		if showAll {
			fmt.Fprintf(w, " %v \t %v \t %v \t %v \n", idStr, task.description, timeDiff, isCompleteStr)
		} else {
			fmt.Fprintf(w, " %v \t %v \t %v \n", idStr, task.description, timeDiff)
		}

	}

}

func Complete(id int64) error {
	index, exist := findTaskById(id)

	if !exist {
		return errors.New("can't complete task as can't find index")
	}

	list := getListTasks()
	row := list[index]
	row.isComplete = true
	list[index] = row

	var data file_utils.TCsvData
	data = append(data, COLUMNS_NAMES)

	for _, value := range list {
		taskRow := createTaskRow(value.ID, value.description, value.isComplete, value.createdAt)
		data = append(data, taskRow)
	}

	err := file_utils.WriteToCSV(filePath, data)
	if err != nil {
		return err
	}

	return nil
}

func Delete(id int64) error {
	index, exist := findTaskById(id)

	if !exist {
		return errors.New("can't delete task as can't find index")
	}

	list := getListTasks()

	var data file_utils.TCsvData
	data = append(data, COLUMNS_NAMES)

	for listIndex, value := range list {
		if int64(listIndex) == index {
			continue
		}
		taskRow := createTaskRow(value.ID, value.description, value.isComplete, value.createdAt)
		data = append(data, taskRow)
	}

	err := file_utils.WriteToCSV(filePath, data)
	if err != nil {
		return err
	}

	return nil
}
