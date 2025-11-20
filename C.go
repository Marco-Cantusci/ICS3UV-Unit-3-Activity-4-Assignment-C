/*
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-11-20
 * @fileoverview If statement problem.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// declare variables
	var name string
	var ageString string
	var ageNumber int
	var student string

	reader := bufio.NewReader(os.Stdin)

	// input

	// ask for name
	fmt.Print("What is your name? \n")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// ask for age
	fmt.Print("How old are you? \n")
	ageString, _ = reader.ReadString('\n')
	ageString = strings.TrimSpace(ageString)
	ageNumber, _ = strconv.Atoi(ageString)

	// ask if they are a student
	fmt.Print("Are you a student? \n")
	student, _ = reader.ReadString('\n')
	student = strings.TrimSpace(student)

	// processing
if strings.ToLower(student) == "yes" {
	if ageNumber >= 13 && ageNumber <= 19 {
		fmt.Printf("Student %s is a teenager. \n", name)
	} else if ageNumber >= 5 && ageNumber <= 12 {
		fmt.Printf("Student %s is a child. \n", name)
	} else {
		fmt.Printf("%s is in a different life stage. \n", name)
	}
} else if strings.ToLower(student) == "no" {
	if ageNumber >= 20 && ageNumber <= 30 {
		fmt.Printf("%s is a young adult. \n", name)
	} else {
	fmt.Printf("%s is in a different life stage. \n", name)
	}
}

fmt.Println("\nDone.")
}
