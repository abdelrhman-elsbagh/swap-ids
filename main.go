package main

import (
	"fmt"
)

func swapSeats(students map[int]string) map[int]string {
	result := make(map[int]string)

	keys := make([]int, 0, len(students))
	for key := range students {
		keys = append(keys, key)
	}

	for i := 0; i < len(keys); i += 2 {
		if i+1 < len(keys) {
			result[keys[i]] = students[keys[i+1]]
			result[keys[i+1]] = students[keys[i]]
		} else {
			result[keys[i]] = students[keys[i]]
		}
	}

	return result
}

func main() {
	students := map[int]string{
		1: "Abbot",
		2: "Doris",
		3: "Emerson",
		4: "Green",
		5: "Jeames",
		6: "Abdel",
	}

	swappedStudents := swapSeats(students)

	fmt.Printf("| id | student |\n")
	fmt.Printf("+----+---------+\n")

	for id, student := range swappedStudents {
		fmt.Printf("| %2d | %-7s |\n", id, student)
	}
}
