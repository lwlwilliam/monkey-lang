package repl

import (
	"bufio"
	"fmt"
	"github.com/lwlwilliam/monkey-lang/lexer"
	"github.com/lwlwilliam/monkey-lang/parser"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// demo1
		//for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		//	fmt.Printf("%+v\n", tok)
		//}

		// demo2
		p := parser.New(l)
		program := p.ParseProgram()
		fmt.Println(program)
	}
}
