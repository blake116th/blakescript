package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Appleby43/blakescript/lexer"
)

const prompt = "BLAKESCRIPT: "

func Start(in io.Reader, out io.Writer) {
	for {
		fmt.Print(prompt)
		scanner := bufio.NewScanner(in)

		stillScanning := scanner.Scan()
		if !stillScanning {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for !l.AtEnd() {
			next := l.NextToken()
			fmt.Printf("{%s, %s} \n", next.Type.String(), next.Literal)
		}
	}
}
