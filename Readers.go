// Exercise: Readers
// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

package main

import "code.google.com/p/go-tour/tour"
import "google.com/p/go/reader"


type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (reader MyReader) Read(bs []byte) (int, error){
	for b := range bs{
		bs[b] = 'A'
	}
	return len(bs), nil
}

func main() {
	reader.Validate(MyReader{})
}
