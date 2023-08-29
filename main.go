package main

import (
	"fmt"
	"log"
	"os"
	"reslang/lexer"
	"runtime"
)

func main() {
	raw, err := os.ReadFile("index.rsq")
	if err != nil {
		log.Fatal(err)
	}

	var l *lexer.Lexer = lexer.New(string(raw))

	// create a file
	file, err := os.Create("parsed.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	var mems = make([]int64, 0)

	for {
		var pointA runtime.MemStats
		runtime.ReadMemStats(&pointA)
		token := l.Parse()
		var pointB runtime.MemStats
		runtime.ReadMemStats(&pointB)

		_, err := file.Write([]byte(fmt.Sprintf("{%v %v}\n", token.Type, token.Literal)))
		if err != nil {
			log.Fatal(err)
		}

		if token.Type == "EOF" {
			break
		}

		mems = append(mems, int64(pointB.TotalAlloc-pointA.TotalAlloc))
	}

	var sum int64
	for _, v := range mems {
		sum += v
	}

	log.Println(sum, "bytes total allocated memory")
	log.Println(sum/int64(len(mems)), "bytes avg allocated memory per iteration")
}
