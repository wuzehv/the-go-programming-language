package q1_2

import (
	"fmt"
	"os"
)

func Echo() {
	for k, v := range os.Args {
		fmt.Println(k, v)
	}
}
