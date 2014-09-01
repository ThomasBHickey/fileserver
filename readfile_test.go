package fileserver

import (
	//"errors"
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	wd, _ := os.Getwd()
	fmt.Println("TestReadFile ", wd)
	page, err := readFile("fileserver/readfile_test.go")
	if err != nil {
		t.FailNow()
	}
	if !bytes.HasPrefix(page.Contents, []byte("package fileserver")) {
		t.Fail()
	}
	page, err = readFile("test.xsl")
	if err != nil {
		t.Fail()
	}
}
