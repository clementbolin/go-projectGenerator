package main

import (
	"os"
	"fmt"

	"github.com/ClementBolin/go-shell/cmd/shell"
)

func main() {
	var shell shell.Shell;

	if (len(os.Args) != 3) {
		fmt.Print("init Usage:\n\n<name> this is name of your project\n\nExample:\ninit go-shell\n")
		return
	}
	shell.CheckInput(os.Args);
}
