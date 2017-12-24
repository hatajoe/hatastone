package console

import (
	"bufio"
	"os"
)

type Reader struct{}

func (r Reader) Read() ([]byte, error) {
	var stdin string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stdin = scanner.Text()
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return []byte(stdin), nil
}
