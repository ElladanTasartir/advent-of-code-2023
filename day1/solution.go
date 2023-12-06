package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	calibrationChan := make(chan int, 0)

	go findCalibrationValue(calibrationChan)

	sum := 0
	for calibration := range calibrationChan {
		sum += calibration
	}

	fmt.Printf("result = %d", sum)
}

func findCalibrationValue(calibrationChan chan int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to read file. err = %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		getCalibrationValueFromText(scanner.Text(), calibrationChan)
	}

	close(calibrationChan)

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan file. err = %v", err)
	}
}

func getCalibrationValueFromText(text string, calibrationChan chan int) {
	foundFirst, foundSecond := false, false
	i, j := 0, len(text)-1

	calibration := 0

	for !foundFirst || !foundSecond {
		if !foundFirst {
			value, err := strconv.Atoi(string(text[i]))
			if err == nil {
				calibration += value * 10
				foundFirst = true
			}

			i++
		}

		if !foundSecond {
			value, err := strconv.Atoi(string(text[j]))
			if err == nil {
				calibration += value
				foundSecond = true
			}

			j--
		}
	}

	fmt.Println(calibration)
	fmt.Println(text)

	calibrationChan <- calibration
}
