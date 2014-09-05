// Based loosely on the Go Wiki server at https://golang.org/doc/articles/wiki/

package fileserver

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Page struct {
	FileName string
	Contents []byte
}

const (
	//baseDir = "C:/Thom/go/src/"
	baseDir = "./" // run me in 'src'!
	fsPath  = baseDir + "github.com/ThomasBHickey/fileserver/"
)

var whereToLook = map[string][]string{
	".css": {"./css/"},
	".go": { //"./",
		baseDir,
		baseDir + "github.com/",
		baseDir + "github.com/ThomasBHickey/"},
	".gone": {"./"},
	".html": {fsPath + "./html/"},
	".js":   {fsPath + "./script/"},
	".xsl":  {fsPath + "./xsl/",
	    baseDir + "github.com/",
	    baseDir + "github.com/ThomasBHickey/"},
	".png":  {fsPath + "./image/"},
	".gif":  {fsPath + "./image/"},
	".ico":  {fsPath + "image/"},
}

func readFile(fname string) (*Page, error) {
	//fmt.Println("readFile looking for ", fname)
	for suffix, list := range whereToLook {
		//fmt.Println("checking for suffix ", suffix)
		if strings.HasSuffix(fname, suffix) {
			for _, prefix := range list {
				xfname := prefix + fname
				fmt.Println("trying file name ", xfname)
				contents, err := ioutil.ReadFile(xfname)
				if err == nil {
					//fmt.Println("Found it!")
					return &Page{FileName: xfname, Contents: contents}, nil
				}
			}
			break
		}
	}
	contents, err := ioutil.ReadFile(fname)
	if err != nil {
		//fmt.Println("didn't find ", fname, "!")
		return nil, err
	}
	return &Page{FileName: fname, Contents: contents}, nil
}

func (p *Page) writeFile() error {
	//fmt.Println("writing:", p.FileName)
	return ioutil.WriteFile(p.FileName, p.Contents, 0600)
}

func (p *Page) writeEditor(w http.ResponseWriter) error {
	fmt.Println("writeEditor '", p.FileName, "'")
	_, err := w.Write([]byte(`<?xml version='1.0' encoding='utf-8'?>
<?xml-stylesheet type='text/xsl' href='/editor.xsl'?>
`))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("<edit><fileName>"))
	if err != nil {
		return err
	}
	err = xml.EscapeText(w, []byte(p.FileName))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("</fileName>\n<contents>"))
	if err != nil {
		return err
	}
	err = xml.EscapeText(w, p.Contents)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("</contents>\n</edit>"))
	return err
}

func saveFile(r *http.Request) {
	fmt.Println("saveFile '", r.URL.Path, "'")
	body := new(bytes.Buffer)
	body.ReadFrom(r.Body)
	//fmt.Println("contents ", body)
	p := &Page{FileName: r.URL.Path[1:], Contents: body.Bytes()}
	err := p.writeFile()
	if err!=nil{
		fmt.Println("error writing file", p.FileName)
		}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler:", r.URL.Path, r.Method)
	if r.Method == "POST" {
		fmt.Println("File needs to be saved")
		saveFile(r)
	}
	page, err := readFile(r.URL.Path[1:])
	if err != nil {
		fmt.Println("Unable to read ", r.URL.Path[1:], err)
		return
	}
	u, err := url.Parse(r.URL.String())
	if err!=nil{
	    fmt.Println("Unable to parse URL", err)
	    return
	}
	fmt.Println("u.Query():", u.Query())
	_, doEdit := u.Query()["edit"]
	fmt.Println("doEdit", doEdit)
	if doEdit {
		page.writeEditor(w)
	} else {
		w.Write(page.Contents)
	}
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Server() {
	fmt.Println("Server called")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
