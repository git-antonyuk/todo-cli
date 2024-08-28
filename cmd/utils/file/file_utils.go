package file_utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"syscall"
)

// One major consideration is that the underlying data file should be locked by the process to prevent concurrent read/writes. This can
// be achieved using the flock system call in unix like systems to obtain an exclusive lock on the file.

// You can achieve this in go using the following code:
func loadFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	// Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

// Then to unlock the file, use the following:
func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

type TCsvData = [][]string

func WriteToCSV(filepath string, data TCsvData) error {
	loadedFile, err := loadFile(filepath)
	if err != nil {
		return err
	}
	defer closeFile(loadedFile)

	// Create or open the file
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush() // Ensure data is flushed to the file

	// Write all data at once
	err = writer.WriteAll(data)
	if err != nil {
		return fmt.Errorf("could not write to csv: %v", err)
	}

	return nil
}

func ReadFromCSV(filepath string) (TCsvData, error) {
	loadedFile, err := loadFile(filepath)
	if err != nil {
		return nil, err
	}
	defer closeFile(loadedFile)

	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all rows from the file
	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read from csv: %v", err)
	}

	return data, nil
}
