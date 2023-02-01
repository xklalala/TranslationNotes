package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func fyFile(source, dest string) {
	f, err := os.Open(source)
	if err != nil {
		panic(err)
	}

	var lineRes [][]byte
	r := bufio.NewReader(f)
	var newStr string = "// "
	for {
		line, prefix, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		lineRes = append(lineRes, append([]byte{}, line...))

		if prefix {
			break
		}

		if !strings.HasPrefix(string(line), "//") {
			if newStr != "// " {
				last := lineRes[len(lineRes)-1]
				res := fy(newStr)
				lineRes[len(lineRes)-1] = []byte(res)
				lineRes = append(lineRes, last)
				newStr = "// "
			}
			continue
		}

		str := string(line)
		newStr += str[2:]
	}

	f2, err := os.Create(dest)
	if err != nil {
		panic(err)
	}

	b2 := bufio.NewWriter(f2)
	for _, v := range lineRes {
		b2.Write(v)
		b2.WriteString("\n")
	}

	if err = b2.Flush(); err != nil {
		panic(err)
	}
}
