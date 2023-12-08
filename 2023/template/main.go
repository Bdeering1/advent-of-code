package main

import (
    "flag"
    "fmt"
    "os"
)

const template =
`package solutions

import (

)

func (Solutions) Day%d_1(input []string) int {
    return -1
}

func (Solutions) Day%d_2(input []string) int {
    return -1
}`

func main() {
    day := flag.Int("day", 0, "day to create template for")    
    flag.Parse()

    var file *os.File
    path := fmt.Sprintf("solutions/day%d.go", *day)
    if *day == 0 {
	for i := 0; i < 25; i++ {
	    *day += 1
	    path = fmt.Sprintf("solutions/day%d.go", *day)
	    if f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666); err == nil {
		fmt.Printf("created %s\n", path)
		file = f
		break
	    }
	}
    } else {
	if f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666); err != nil {
	    panic(err)	    
	} else {
	    file = f
	}
    }
    defer file.Close()
    
    finfo, err := file.Stat()
    if err != nil {
	panic(err)
    }

    if finfo.Size() > 0 {
	fmt.Printf("%s exists and is non-empty, exiting", path)
	return
    }
    
    template_str := fmt.Sprintf(template, *day, *day)
    if n, err := file.WriteString(template_str); err != nil {
	panic(err)
    } else {
	fmt.Printf("wrote %d bytes to %s\n", n, path)
    }

    input_path := fmt.Sprintf("input/day%d.txt", *day)
    if f, err := os.OpenFile(input_path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666); err != nil {
	fmt.Printf("Unable to create %s", input_path)
    } else {
	fmt.Printf("Created %s\n", input_path)
	f.Close()
    }
}
