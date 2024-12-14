package main

import (
	"TLanguage/compile"
	"os"
)

// func main() {
// 	name, err := user.Current()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Hello %s! This is the T programming language!\n", name.Username)
// 	fmt.Println("Feel free to type in command")
// 	repl.Start(os.Stdin, os.Stdout)
// }

func main() {
	compile.Start(os.Stdin, os.Stdout, os.Args[1])
}
