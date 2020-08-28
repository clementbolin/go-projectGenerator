# Go-projectGenerator

## How to use

The goal to ```go-projectGenerator``` is to setup your golang project.

![](Assets/example.png)

##### Tree structure

```
.
├── Makefile
├── cmd
│   └── main.go
├── go.mod
└── pkg

2 directories, 3 files
```

for run your project generate you can run ```make run```
for display all rules of your Makefile, you can run ```make```

```
Makefile Rules
 Choose a command run in go-projectGenerator: 
   > make setup (setup project)
   > make build (build project)
   > make clean (clean Project)
   > make clean-cache-mod (clean cahe in /Users/clementbolin/Epitech/Tek2/Go/go-projectGenerator/vendor:/Users/clementbolin/Epitech/Tek2/Go/go-projectGenerator/pkg/mod)
```

## How to install

    git clone https://github.com/ClementBolin/go-projectGenerator
    cd go-projectGenerator
    chmod +x install.sh
    ./install.sh

if ```install.sh``` failed you follow next instruction :

#### MacOs

    make build
    sudo cp ./bin/go-projectGenerator /usr/local/bin

#### Linux

    make build
    sudo cp ./bin/go-projectGenerator /usr/bin

