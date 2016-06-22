package main

import (
	"fmt"
  "os"
	"io/ioutil"
	"strings"
)

func writefile(path string, content string, number int){
	fmt.Printf("writefile %s\n", path)
	// create file
	f, err := os.Create(path)
	if err != nil {
		panic(err)
  }

  // close fi on exit and check for its returned error
	defer func() {
			if err := f.Close(); err != nil {
					panic(err)
			}
  }()
	for i := 0;  i<number; i++ {
		    // write x lines with content
				f.Write([]byte(fmt.Sprintf("%d %s\n", i, content)))
  }

}

func readfile(path string) []byte {
	fmt.Printf("readfile %s\n", path)
  // Open
	data, err := ioutil.ReadFile(path)
	if err != nil {
			panic(err)
	}
	// fmt.Println(string(data))
	return data
}

func changefile(file []byte, pattern string, change string){
  fmt.Println(string(file))
  // ioutil.WriteFile("filename", contents, 0644)
	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
					if strings.Contains(line, pattern) {
				      lines[i] = strings.Replace(line, pattern, change, -1)
					}
	}
	changedfile := strings.Join(lines, "\n")
	fmt.Println(changedfile)
}

func main() {

  writefile("/tmp/test", "bla bla\ntest test", 10)
	infile := readfile("/tmp/test")
	changefile(infile, "bla", "foobar")
}
