package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("input.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	b = b[:len(b)-1] // removing last character - end of file -not part of string

	i := 0
	for i < len(b)-1 {
		i++

		if b[i] == b[i-1]+32 || b[i] == b[i-1]-32 {
			if i == 1 && len(b) > 2 {
				b = b[2:len(b)]
			} else if i == len(b)-1 {
				b = b[0 : len(b)-2]
			} else {

				b = append(b[0:i-1], b[i+1:len(b)]...)
			}
			i = 0
		}

	}

	fmt.Println(len(b)) // print the content as 'bytes'

	//	str := string(b) // convert content to a 'string'

	//	fmt.Println(str[:len(str)-1]) // print the content as a 'string'
}
