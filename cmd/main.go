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
		if (shell.CheckInput() == true) { break } 
	}
	fmt.Println("Leave go-shell !");
}