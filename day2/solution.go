package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var availableCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	gamesChan := make(chan int, 0)

	go findPossibleGameFromInput(gamesChan)

	sum := 0
	for game := range gamesChan {
		sum += game
	}

	fmt.Printf("result = %d", sum)
}

func findPossibleGameFromInput(gamesChan chan int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to read file. err = %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		findPossibleGame(scanner.Text(), gamesChan)
	}

	close(gamesChan)

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan file. err = %v", err)
	}
}

func findPossibleGame(game string, gamesChan chan int) {
	gameNumber := findGameNumber(game)
	if gameNumber == 0 {
		return
	}

	rounds := findGameRounds(game)
	if len(rounds) < 1 {
		return
	}

	for _, round := range rounds {
		if !isPossibleGame(round) {
			return
		}
	}

	gamesChan <- gameNumber
}

func findGameNumber(game string) int {
	exp := regexp.MustCompile(`Game (\d+):`)
	matches := exp.FindStringSubmatch(game)
	if len(matches) < 2 {
		return 0
	}

	value, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0
	}

	return value
}

func findGameRounds(game string) []string {
	separatedRounds := strings.Split(game, ":")

	return strings.Split(separatedRounds[1], ";")
}

func isPossibleGame(round string) bool {
	cubes := strings.Split(round, ",")

	for _, cube := range cubes {
		number, color := findCubeNumberAndColor(cube)
		if number == 0 || color == "" {
			return false
		}

		foundCube, ok := availableCubes[color]
		if !ok {
			return false
		}

		rest := foundCube - number

		if rest < 0 {
			return false
		}
	}

	return true
}

func findCubeNumberAndColor(cube string) (int, string) {
	cubeInfo := strings.Split(cube, " ")

	number, err := strconv.Atoi(cubeInfo[1])
	if err != nil {
		return 0, ""
	}

	return number, cubeInfo[2]
}
