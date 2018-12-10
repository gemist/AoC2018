package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	c, err := ioutil.ReadFile("input.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	c = c[:len(c)-1] // removing last character - end of file -not part of string

	//	fmt.Println(string(append(c[0:3], c[4:len(c)]...)))
	min_length := len(c)
	var find_ascii uint8
	var ascii uint8
	var i int
	for ascii = 65; ascii <= 90; ascii++ {
		b := append([]byte(nil), c...)

		//remove assci

		for i = 0; i < len(b); i++ {

			if b[i] == ascii || b[i] == ascii+32 {
				if i == 0 {
					b = b[1:]
				} else if i == len(b)-1 {
					b = b[:len(b)-1]
				} else {
					b = append(b[0:i], b[i+1:len(b)]...)
				}
				i--
			}

		}

		//repeat again like in part1
		i = 0
		for i < len(b)-1 {
			i++

			if (b[i] == b[i-1]+32) || (b[i] == b[i-1]-32) {
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

		if len(b) < min_length {
			min_length = len(b)
			find_ascii = ascii
		}

	}
	fmt.Println(min_length, string(find_ascii)) // print the content as 'bytes'

}
