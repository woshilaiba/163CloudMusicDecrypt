package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

//read a file for input， or STDIN
func ConvertCacheFile(in, out string) {
	var fin, fout *os.File
	var er error
	if in == "" {
		fin = os.Stdin
	} else {
		fin, er = os.Open(in)
		if er != nil {
			fmt.Println(er.Error())
			panic(er)
		}
	}
	if out == "" {
		fout = os.Stdout
	} else {
		fout, er = os.OpenFile(out, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if er != nil {
			fmt.Println(er.Error())
			panic(er)
		}
	}

	r := bufio.NewReader(fin)
	w := bufio.NewWriter(fout)
	bf := make([]byte, 1024)
	for {
		tbf := bf
		n, err := r.Read(tbf)
		if err != nil {
			if err != io.EOF {
				//abnormal
				fmt.Println(err.Error())
				panic(err)
			} else if n == 0 {
				//finished reading
				return
			}
		}
		tbf = tbf[:n]
		for i := 0; i < n; i++ {
			tbf[i] ^= 0xa3
		}
		_, er := w.Write(tbf)
		if er != nil {
			//abnormal
			fmt.Println(err.Error())
			panic(err)
		}
	}
}

func main() {
	in := flag.String("in", "noinput", "path of the input cache file")
	out := flag.String("out", "nooutput", "path of the output file")
	flag.Parse()
	ConvertCacheFile(*in, *out)
}
