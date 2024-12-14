package executor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var basic string = `package main 
import "fmt"

func main() {
	fmt.Println("%s")
}
`

func Exec(out []string, path string) {
	//创建临时文件
	wd, _ := os.Getwd()
	filename := filepath.Base(path)
	ext := filepath.Ext(path)
	name := filename[:len(filename)-len(ext)]
	c := "\\"
	if runtime.GOOS == "linux" {
		c = "/"
	}
	file, err := os.Create(wd + c + name + ".go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer os.Remove(file.Name())
	_, err = file.WriteString(fmt.Sprintf(basic, strings.Join(out, "\\n")))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
	cmd := exec.Command("go", "build", file.Name())
	fmt.Println(file.Name())
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error building executable:", err)
		os.Exit(1)
	}
}
