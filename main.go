package main

import (
	"fmt"
	"log"
	"os"
	"reslang/lexer"
)

func main() {
	code := `
	let one = 1;
	let two = 2;

	let sum = fn(a, b) {
		a + b;
	};

	let result = sum(one, two);
	`

	var l *lexer.Lexer = lexer.New(code)

	// create a file
	file, err := os.Create("parsed.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	for {
		token := l.Parse()

		_, err := file.Write([]byte(fmt.Sprintf("{%v %v}\n", token.Type, token.Literal)))
		if err != nil {
			log.Fatal(err)
		}

		if token.Type == "EOF" {
			break
		}
	}
}
