package tools

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloat(file_name string) ([3]float64, error) {
	var numbers [3]float64

	// open file
	file, err := os.Open(file_name)

	//check error
	if err != nil {
		return numbers, err
	}
	// create scanner
	scanner := bufio.NewScanner(file)

	//check error

	//scan
	for i := 0; scanner.Scan(); i++ {

		value, err := strconv.ParseFloat(scanner.Text(), 64)

		//checking for error
		if err != nil {
			return numbers, err
		}

		numbers[i] = value
	}

	//checking scann errors
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	//close and check errors
	err = file.Close()
	if err != nil {
		return numbers, err
	}

	// in case every thing went good we just return the table, aka the content of the file in the form of a table
	return numbers, nil
}
