package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	var branch, comment string
	fmt.Println("enter the name of the branch")
	fmt.Scan(&branch)
	fmt.Println("enter a comment")
	fmt.Scan(&comment)
	fmt.Println(comment, branch)
	comment = "\"" + comment + "\""

	out, err := exec.Command("git", "add", ".").Output()
	output(out, err)

	out, err = exec.Command("git", "commit", "-m", comment).Output()
	output(out, err)

	out, err = exec.Command("git", "push", "origin", branch).Output()
	output(out, err)

}

func output(out []byte, err error) {
	if err != nil {
		log.Fatal("Error")
	} else {
		fmt.Printf("\nOutput  is : %s \n", out)
	}

}
