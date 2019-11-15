# go-consoletime
consoletime is a golang small package for log any execute time of each code stagement

## Install
```bash
go get github.com/spksoft/go-consoletime
```

## Using
```go
// main.go
package main

import (
	consoletime "github.com/spksoft/go-consoletime"
)

func main() {
	consoletime.Log("Loop#1") // Log time for Loop#1
	for i := 0; i < 10000000000; i++ {
	}
	consoletime.Log("Loop#1") // End log time of Loop#1 and display it
	consoletime.Log("Loop#2") // Log time for Loop#2
	for i := 0; i < 10000000000; i++ {
	} 
    consoletime.Log("Loop#2") // End log time of Loop#1 and display it

}
```

Display on console
```
$ go run main.go
2019/11/16 02:52:08 Loop#1 took 2.68317789s
2019/11/16 02:52:10 Loop#2 took 2.684548133s
```