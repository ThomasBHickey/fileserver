package fileserver

import (
	//"errors"
	"bytes"
	"fmt"
	//"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	var tcontent = []byte("Now is the time")
	p := &Page{FileName: "testWriteFile.txt", Contents: tcontent}
	if p.writeFile() != nil {
		fmt.Println("TestWriteFile", p.writeFile())
		t.Fail()
	}
	p, err := readFile(p.FileName)
	if err == nil {
		//fmt.Println("read in ", p.FileName)
		if !bytes.Equal(p.Contents, tcontent) {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}
