package main

import (
	"fmt"
	"os/exec"
)

func main() {
	remote, err := exec.Command("git", "-C", "/personal/repos/daniele821/golang/", "rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}").Output()
	fmt.Println(string(remote), err)
	// git rev-parse --abbrev-ref --symbolic-full-name @{u}
}
