package env

import (
	"os"
	"os/exec"
	"errors"
	"fmt"
	"bufio"
)

// Env : Env structure 
type Env struct {
	// env var
	GOPATH string;
	GOBIN string;
	GOBASE string;
	PWD string;
	// other varibles
	ProjectName string
}

// Init : init workdir environment
func (env *Env) Init(ProjectName string) {
	env.GOPATH = os.Getenv("GOPATH");
	env.GOBIN = os.Getenv("GOBIN");
	env.GOBASE = os.Getenv("PWD");
	env.PWD = os.Getenv("PWD");
	env.ProjectName = ProjectName;
	// Conctat vendor to GOPATH
	env.GOPATH = env.GOPATH + "/vendor:" + env.GOBASE;
}

// SetupWorkdir : create folder and file of your project
func (env Env) SetupWorkdir() error {
	if _, err := os.Stat(env.ProjectName); err == nil { return errors.New("Error Folder " + env.ProjectName + " exist\nFailure to create the project") }
	if err := os.Mkdir(env.ProjectName, os.ModePerm); err != nil { return err }
	if err := os.Mkdir("./" + env.ProjectName + "/cmd", os.ModePerm); err != nil { return err }
	fmt.Printf("   > ./%s/cmd/ create ✅\n", env.ProjectName)
	if err := os.Mkdir("./" + env.ProjectName + "/pkg", os.ModePerm); err != nil { return err }
	fmt.Printf("   > ./%s/pkg/ create ✅\n", env.ProjectName)
	if err := env.createMain(); err != nil { return err }
	if err := env.createGoMod(); err != nil { return err }
	if err := env.gitInit(); err != nil { return nil }
	return nil
}

// Create main.go in ./(ProjectName)/cmd/main.go
func (env Env) createMain() error {
	fd, err := os.Create("./" + env.ProjectName + "/cmd/main.go");
	if (err != nil) { return err }
	var content string = `package main
import (
	"fmt"
)

func main() {
	fmt.Println("Hello World");
}
	`
	if _, err := fd.WriteString(content); err != nil { return err }
	fmt.Printf("   > ./%s/cmd/main.go create ✅\n", env.ProjectName);
	return nil
}

// Create go.mod in ./(ProjectName)/
func (env Env) createGoMod() error {
	scan := bufio.NewScanner(os.Stdin);
	for {
		fmt.Print("   > Init go.mod [y/n]: ")
		scan.Scan()
		if (scan.Text() == "y") { break }
		if (scan.Text() == "n") { return nil }
	}

	var name string;
	for {
		fmt.Print("   > Name of your module: ")
		scan.Scan()
		name = scan.Text()
		if (len(name) > 0) { break }
	}
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = "./" + env.ProjectName + "/"
	if err := cmd.Run(); err != nil { return err }
	fmt.Println("   > go.mod init ✅");
	return nil
}

// Init git
func (env Env) gitInit() error {
	git := exec.Command("git", "init");
	git.Dir = "./" + env.ProjectName + "/"
	if err := git.Run(); err != nil { return err }
	// Create .gitignore
	fd, err := os.Create("./" + env.ProjectName + ".gitignore")
	if (err != nil) { return err }
	if _, err := fd.WriteString("bin\nvendor"); err != nil { return err }
	// Add files
	addI := exec.Command("git", "add .*")
	addI.Dir = "./" + env.ProjectName + "/"
	if err := addI.Run(); err != nil { return err }
	add := exec.Command("git", "add *");
	add.Dir = "./" + env.ProjectName + "/"
	if err := add.Run(); err != nil { return nil }
	commit := exec.Command("git", "commit -m \"setup golang project\"");
	commit.Dir = "./" + env.ProjectName + "/"
	if err := commit.Run(); err != nil { return err }
	fmt.Println("   > git init ✅")
	return nil
}
