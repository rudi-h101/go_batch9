package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score int
}

func main() {
	var students []Student

	for i := 0; i < 5; i++ {
		var name string
		var score int

		fmt.Printf("Input %d Student's Name: ", i+1)
		fmt.Scan(&name)
		fmt.Printf("Input %d Student's Score: ", i+1)
		fmt.Scan(&score)

		student := Student{Name: name, Score: score}
		students = append(students, student)
	}

	totalScore := 0
	minScore := math.MaxInt32
	maxScore := math.MinInt32
	var minStudent, maxStudent Student

	for _, student := range students {
		totalScore += student.Score

		if student.Score < minScore {
			minScore = student.Score
			minStudent = student
		}

		if student.Score > maxScore {
			maxScore = student.Score
			maxStudent = student
		}
	}

	averageScore := totalScore / len(students)

	fmt.Println("Average Score:", averageScore)
	
	fmt.Printf("Min Score of Students: %s (%d)\n", minStudent.Name, minStudent.Score)
	fmt.Printf("Max Score of Students: %s (%d)\n", maxStudent.Name, maxStudent.Score)
}
