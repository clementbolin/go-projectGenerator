package shell

import (
	"os"
	"fmt"
	"bufio"
	"strings"

	"github.com/ClementBolin/go-shell/pkg/makefile"
)

// Shell Structure
type Shell struct {
	input string;
	cmd []string;
}

// Init your Shell Structure
func (shell *Shell) Init() {
	shell.input = "";
	shell.cmd = append(shell.cmd, "init");
	shell.cmd = append(shell.cmd, "clean");
	shell.cmd = append(shell.cmd, "Quit");
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
func (shell Shell) CheckInput() {
	cmdArray := strings.Split(shell.input, " ");
	for _, cmd := range shell.cmd {
		if (cmdArray[0] == cmd) {
			shell.leaveShell();
			shell.initMakefile(cmdArray);
			return
		}
	}
	fmt.Println(cmdArray[0], ": invalid command");
}

func (shell Shell) leaveShell() {
	if (shell.input == "Quit") {
		fmt.Println("Leave go-shell Bye !")
		os.Exit(0);
	}
}

func (shell Shell) initMakefile(cmd []string) {
	var make makefile.Makefile;

	if (len(cmd) != 2) {
		fmt.Print("init Usage:\n\n<name> this is name of your project\n\nExample:\ninit go-shell\n");
		return
	}
	fmt.Println("Start init Makefile");
	make.Init(cmd[1])
	make.DebugEnv();
}
