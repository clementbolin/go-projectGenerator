package main

import (
	"fmt"

	"github.com/ClementBolin/go-shell/cmd/shell"
)

func main() {
	var shell shell.Shell;

	shell.Init();
	for {
		shell.ReadInput();
		shell.CheckInput()
	}
	fmt.Println("Leave go-shell !");
}
