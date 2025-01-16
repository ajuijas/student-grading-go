package main

import (
	"os"
	"strconv"
	"strings"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func stringToInt(s string) int {
	i, error := strconv.Atoi(s)
	if error!=nil {
		panic(error)
	}
	return i
}

func parseCSV(filePath string) []student {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	students := []student{}
	for _, line := range lines[1:] {
		studentData := strings.Split(line, ",")
		var stud student
		stud.firstName = studentData[0]
		stud.lastName = studentData[1]
		stud.university = studentData[2]
		stud.test1Score = stringToInt(studentData[3])
		stud.test2Score = stringToInt(studentData[4])
		stud.test3Score = stringToInt(studentData[5])
		stud.test4Score = stringToInt(studentData[6])
		students = append(students, stud)
	}
	return students
}

func setGrade(score float32) Grade {
	var grade Grade
	if score >= 70 {
		grade = A
	}else if score >=50 {
		grade = B
	}else if score >= 35 {
		grade = C
	}else {
		grade = F
	}
	return grade
}

func setGradeAndScore(stud student) studentStat {
	var gradedStud studentStat
	gradedStud.student = stud
	gradedStud.finalScore = (float32(stud.test1Score) + float32(stud.test2Score) + float32(stud.test3Score) + float32(stud.test4Score))/4
	gradedStud.grade = setGrade(gradedStud.finalScore)
	return gradedStud
}

func calculateGrade(students []student) []studentStat {
	var gradedStudents []studentStat
	for _, stud := range students {
		gradedStud := setGradeAndScore(stud)
		gradedStudents = append(gradedStudents, gradedStud)
	}
	return gradedStudents
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	var toper studentStat
	for _, stud := range gradedStudents {
		if stud.finalScore > toper.finalScore{
			toper = stud
		}
	}
	return toper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	var topperPerUniversity map[string]studentStat = map[string]studentStat{}
	for _, stud := range gs{
		if topperPerUniversity[stud.student.university].finalScore < stud.finalScore {
			topperPerUniversity[stud.student.university]=stud
		}
	}
	return topperPerUniversity
}
