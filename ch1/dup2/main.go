package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileNameCnts := make(map[string]string)
	pwd,err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNameCnts)
	} else {
		for _, arg := range files {
			f, err := os.Open(pwd+"/"+arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNameCnts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s [filename]: %s\n", n, line,fileNameCnts[line])
			//fmt.Printf("%d\n", counts[line])
			//fmt.Printf("%d\t%s\n", n, line)
		}
	}


}

func countLines(f *os.File, counts map[string]int, fileNameCnts map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileNameCnts[input.Text()] = f.Name()

		//cntMap[input.Text()][f.Name()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
