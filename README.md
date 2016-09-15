# wcfile
Word contains files in Go


### Pattern
Pattern applies only one in order.  

Contains word "chrome"
```
*wcfile.Pattern{
    FileName: "logs/access/chrome.log",
    Word:     "chrome",
}
```

Not contains word "chrome"
```
*wcfile.Pattern{
    FileName: "logs/access/not_chrome.log",
    Word:     "chrome",
    Not:      true,
}
```

All contains words
```
*wcfile.Pattern{
    FileName: "logs/access/all.log",
}
```





### Usage

```
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
		Word:     "192.168.",
		Not:      true,
	},
}

func main() {
	w, err := wcfile.NewContains(patterns)

	if err != nil {
		os.Exit(1)
	}

	fmt.Fprintln(w, "192.168.1.14 ...") //=> logs/access/my.log
	fmt.Fprintln(w, "192.168.1.128 ...") //=> logs/access/private.log
	fmt.Fprintln(w, "100.100.100.100 ...") //=> logs/access/global.log
}
```