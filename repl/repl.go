// REPL stands for "Read Eval Print Loop"
// reads input, sends it to the interpreter for evaluation, prints the result/output of the interpreter and starts again
package repl

import (
	"bufio"
	"fmt"
	"github.com/lwlwilliam/monkey-lang/lexer"
	"github.com/lwlwilliam/monkey-lang/token"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
