package gen

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func loggingGen(s string) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	abs,err := filepath.Abs(before(dir,"src/"))
	var b []byte
	if abs == dir {
		b, err = ioutil.ReadFile(abs+"/github.com/AliasYermukanov/gkgen/templates/logging.gotxt") // just pass the file name
	}else{
		b, err = ioutil.ReadFile(abs+"/src/github.com/AliasYermukanov/gkgen/templates/logging.gotxt") // just pass the file name
	}
	if err != nil {
		fmt.Print(err)
	}

	f, _ := os.Create(s+"-api/src/"+s+"/logging.go")
	_, _ = f.Write([]byte(string(b)))
	_ = f.Close()

	err = parseEndpoint(s+"-api/src/"+s+"/logging.go",s,dir)
	if err != nil {
		return err
	}

	f, _ = os.Open(s+"-api/src/"+s+"/logging.go")
	_, _ = io.Copy(os.Stdout, f)
	_ = f.Close()

	return nil
}