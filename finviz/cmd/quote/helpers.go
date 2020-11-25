package quote

import (
	"fmt"
	"os"
)

func er(msg interface{}) {
	fmt.Println("Error: ", msg)
	os.Exit(1)
}
