/*
  MIT License

  Copyright (c) 2018 Star Brilliant

  Permission is hereby granted, free of charge, to any person obtaining a copy
  of this software and associated documentation files (the "Software"), to deal
  in the Software without restriction, including without limitation the rights
  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
  copies of the Software, and to permit persons to whom the Software is
  furnished to do so, subject to the following conditions:

  The above copyright notice and this permission notice shall be included in
  all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
  SOFTWARE.
*/

package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	midimark "../.."
)

func warningCallback(err error) {
	log.Println(err)
}

func main() {
	var input, output *os.File
	var err error
	switch len(os.Args) {
	case 0:
		fmt.Print("Usage: mark2midi INPUT.midml OUTPUT.mid\n\n")
		os.Exit(1)
	case 1, 2:
		fmt.Printf("Usage: %s INPUT.midml OUTPUT.mid\n\n", os.Args[0])
		os.Exit(1)
	case 3:
		input, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalln(err)
		}
		defer input.Close()
		output, err = os.Create(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}
		defer output.Close()
	}
	markup, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatalln(err)
	}
	var sequence midimark.MThd
	err = xml.Unmarshal(markup, &sequence)
	if err != nil {
		log.Fatalln(err)
	}
	w := bufio.NewWriter(output)
	defer w.Flush()
	err = sequence.Encode(w)
	if err != nil {
		log.Fatalln(err)
	}
}
