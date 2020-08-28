package shell

import (
	"os"
	"fmt"
	"bufio"

	"github.com/ClementBolin/go-shell/pkg/makefile"
	"github.com/ClementBolin/go-shell/pkg/env"
)

// Shell Structure
type Shell struct {
	input string;
	cmd []string;
}

// ReadInput os.Stdin Line and change value Shell.input
func (shell *Shell) ReadInput() {
	scan := bufio.NewScanner(os.Stdin);
	fmt.Print("> ");
	scan.Scan();
	shell.input = scan.Text();
}

// CheckInput : CheckInput Value
// Return : return false if input == Quit or false in other case
func (shell Shell) CheckInput(cmdArray []string) {
	if (cmdArray[1] == "init") {
		shell.initProject(cmdArray);
		return
	}
	fmt.Println(cmdArray[1], ": invalid command");
}

func (shell Shell) leaveShell() {
	if (shell.input == "Quit") {
		fmt.Println("Leave go-shell Bye !")
		os.Exit(0);
	}
}

func (shell Shell) initProject(cmd []string) {
	var make makefile.Makefile;

	if (len(cmd) != 3) {
		fmt.Print("init Usage:\n\n<name> this is name of your project\n\nExample:\ninit go-shell\n");
		return
	}
	fmt.Println("Start init Project:");
	var env env.Env;
	env.Init(cmd[2])
	if err := env.SetupWorkdir(); err != nil {
		fmt.Printf("Failed to Setup project ❌\nDetail Error: %s", err)
	}
	if err := make.Generation(env.ProjectName); err != nil {
		fmt.Println(err);
		return
	}
	fmt.Println("Project setup ✅")
}
