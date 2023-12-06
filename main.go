package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	comment, branch := argument()
	gitSend(comment, branch)

}

func argument() (string, string) {
	var branch, comment string
	fmt.Println("enter a comment")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	comment, _ = reader.ReadString('\n')
	fmt.Println("enter the name of the branch")
	fmt.Scan(&branch)
	comment = "\"" + comment + "\""
	return comment, branch
}

func gitSend(comment, branch string) {
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
