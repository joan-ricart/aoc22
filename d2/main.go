package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type ShapeName string

type Shape struct {
	Name   ShapeName
	Points int
	Beats  []ShapeName
}

const (
	Rock     ShapeName = "Rock"
	Paper    ShapeName = "Paper"
	Scissors ShapeName = "Scissors"
)

const (
	LossBonus int = 0
	DrawBonus     = 3
	WinBonus      = 6
)

func main() {

	file, err := os.Open("input-example.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		round_choices := strings.Split(scanner.Text(), " ")
		p1_choice := round_choices[0]
		p2_choice := round_choices[1]

		// fmt.Println(s)
		fmt.Println(getShape(p1_choice))
		fmt.Println(getShape(p2_choice))
	}

}

// func getRoundScore(s1 Shape, s2 Shape) int {
// 	if s1.Beats
// }

func getShape(choice string) Shape {

	s := Shape{}

	switch choice {
	case "A", "X":
		s = Shape{
			Name:   Rock,
			Points: 1,
		}
	case "B", "Y":
		s = Shape{
			Name:   "Paper",
			Points: 2,
		}
	case "C", "Z":
		s = Shape{
			Name:   "Scissors",
			Points: 3,
		}
	}

	return s
}
