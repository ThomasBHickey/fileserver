package fileserver

import (
	"fmt"
	"io/ioutil"
	//"path/filepath"
	"strings"
)

type Page struct {
	FileName string
	Contents []byte
}

const (
	baseDir = "C:/Thom/go/src/"
)

var whereToLook = map[string][]string{
	".go": {"./",
		baseDir,
		baseDir + "github.com/",
		baseDir + "github.com/ThomasBHickey/"},
	".xsl": {"./xsl/"},
}

func readFile(fname string) (*Page, error) {
	fmt.Println("readFile looking for ", fname)
	for suffix, list := range whereToLook {
		fmt.Println("checking for suffix ", suffix)
		if strings.HasSuffix(fname, suffix) {
			for _, prefix := range list {
				xfname := prefix + fname
				fmt.Println("trying file name ", xfname)
				contents, err := ioutil.ReadFile(xfname)
				if err == nil {
					fmt.Println("Found it!")
					return &Page{FileName: xfname, Contents: contents}, nil
				}
			}
			break
		}
	}
	contents, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("didn't find ", fname, "!")
		return nil, err
	}
	return &Page{FileName: fname, Contents: contents}, nil
}
