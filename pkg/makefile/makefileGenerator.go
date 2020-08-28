package makefile

import (
	"os"
	"fmt"
)

// Makefile : Generation Makefile structure
type Makefile struct {
	// env var
	GOPATH string;
	GOBIN string;
	GOBASE string;
	// other varibles
	projectName string
}

// Init : inti Makefile Structure
func (make *Makefile) Init(projectName string) {
	make.GOPATH = os.Getenv("GOPATH");
	make.GOBIN = os.Getenv("GOBIN");
	make.GOBASE = os.Getenv("PWD");
	make.projectName = projectName;
	// Conctat vendor to GOPATH
	make.GOPATH = make.GOPATH + "/vendor:" + make.GOBASE;
}

func (make Makefile) DebugEnv() {
	fmt.Println("GOPATH: ", make.GOPATH);
	fmt.Println("GOBIN: ", make.GOBIN);
	fmt.Println("GOBASE: ", make.GOBASE);
}