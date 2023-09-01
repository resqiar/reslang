package repl

import (
	"bufio"
	"fmt"
	"io"
	"reslang/lexer"
	"reslang/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print("--> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		for i := l.Parse(); i.Type != token.EOF; i = l.Parse() {
			fmt.Printf("%+v\n", i)
		}
	}
}
