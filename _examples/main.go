package main

import (
	"fmt"
	"os"

	"github.com/ngc224/wcfile"
)

var patterns = []*wcfile.Pattern{
	{
		FileName: "logs/access/my.log",
		Word:     "192.168.1.14",
	},
	{
		FileName: "logs/access/private.log",
		Word:     "192.168.",
	},
	{
		FileName: "logs/access/global.log",
	},
}

func main() {
	w, err := wcfile.NewContains(patterns)

	if err != nil {
		os.Exit(1)
	}

	fmt.Fprintln(w, "192.168.1.14 ...")
	fmt.Fprintln(w, "192.168.1.128 ...")
	fmt.Fprintln(w, "100.100.100.100 ...")
}
