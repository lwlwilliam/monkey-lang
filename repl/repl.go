package repl

import (
	"bufio"
	"fmt"
	"github.com/lwlwilliam/monkey-lang/lexer"
	"github.com/lwlwilliam/monkey-lang/parser"
	"io"
)

const MONKEY_FACE = `           __,__
  .--.  .-"     "-.  .--.
 / .. \/  .-. .-.  \/ .. \
| |  '|  /   Y   \  |'  | |
| \   \  \ 0 | 0 /  /   / |
 \ '- ,\.-"""""""-./, -' /
  ''-' /_   ^ ^   _\ '-''
      |  \._   _./  |
      \   \ '~' /   /
       '._ '-=-' _.'
          '-----'
`

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
		/*
			p := parser.New(l)
			program := p.ParseProgram()
			if len(p.Errors()) > 0 {
				for _, e := range p.Errors() {
					fmt.Println(e)
				}
			} else {
				fmt.Println(program)
			}
		*/

		// demo3
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
