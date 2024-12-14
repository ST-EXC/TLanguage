package compile

import (
	"TLanguage/evaluator"
	"TLanguage/executor"
	"TLanguage/lexer"
	"TLanguage/object"
	"TLanguage/parser"
	"io"
	"os"
)

func Start(in io.Reader, out io.Writer, path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	env := object.NewEnvironment()
	prog := string(content)
	l := lexer.NewLexer(prog)
	p := parser.NewParser(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParseErrors(out, p.Errors())
		return
	}
	evaluator.Eval(program, env)
	executor.Exec(evaluator.OUT, path)
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
