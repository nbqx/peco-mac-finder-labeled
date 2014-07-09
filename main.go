package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"bufio"
	"strings"
	"sync"
)

var labelColors = map[string] string {
	"grey": "1",
	"gray": "1",
	"green": "2",
	"purple": "3",
	"blue": "4",
	"yellow": "5",
	"red": "6",
	"orange": "7",
}

func getkMDItemFSLabel(name string) (color string) {
	color = ""
	mdls := exec.Command("mdls","-name","kMDItemFSLabel",name)
	b,err := mdls.Output()
	if err != nil {
		os.Exit(1)
	}

	label := string(b)
	
	if label != "" {
		num := strings.TrimSpace(strings.SplitAfter(label,"=")[1])
		if num != "0" {
			color = num
		}
	}
	
	return
}

func main() {
	// os.Args -> $QUERY

	if runtime.GOOS != "darwin" {
		os.Exit(1)
	}

	// no query to exit
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	labeled := labelColors[strings.ToLower(os.Args[1])]

	var wg sync.WaitGroup

	// input from STDIN
	buf := bufio.NewReader(os.Stdin)
	for {
		b,_,err := buf.ReadLine()
		if err != nil {
			break
		}
		
		wg.Add(1)
		go func(line string){
			defer wg.Done()
			color := getkMDItemFSLabel(line)
			if labeled == color {
				fmt.Println(line)
			}
		}(string(b))
	}
	wg.Wait()
	
}
