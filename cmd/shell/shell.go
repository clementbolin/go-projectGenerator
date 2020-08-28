package shell

import (
	"os"
	"fmt"
	"bufio"
)

type Shell struct {
	input string;
	cmd []string;
}


func (shell *Shell) Init() {
	shell.input = "";
	shell.cmd = append(shell.cmd, "init");
	shell.cmd = append(shell.cmd, "clean");
	shell.cmd = append(shell.cmd, "Quit");
}

func (shell *Shell) ReadInput() {
	scan := bufio.NewScanner(os.Stdin);
	fmt.Print("> ");
	scan.Scan();
	shell.input = scan.Text();
}

func (shell Shell) CheckInput() bool {
	var status bool = false;

	for _, cmd := range shell.cmd {
		if (shell.input == cmd) {
			status = shell.leaveShell();
			return (status);
		}
	}
	return (status);
}

func (shell Shell) leaveShell() bool {
	if (shell.input == "Quit") {
		return (true)
	}
	return (true)
}