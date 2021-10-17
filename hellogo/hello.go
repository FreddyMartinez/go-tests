package hello

import (
	"fmt"

	"rsc.io/quote"
)

func Greet(name string) string {
	msg := quote.Go()
	return fmt.Sprintf("Hello %v, welcome. %v", name, msg)
}
