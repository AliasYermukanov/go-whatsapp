package gen

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func transportGen(s string) error{
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	abs,err := filepath.Abs(before(dir,"src/"))
	var b []byte
	if abs == dir {
		b, err = ioutil.ReadFile(abs+"/github.com/AliasYermukanov/gkgen/templates/transport.gotxt") // just pass the file name
	}else{
		b, err = ioutil.ReadFile(abs+"/src/github.com/AliasYermukanov/gkgen/templates/transport.gotxt") // just pass the file name
	}
	if err != nil {
		fmt.Print(err)
	}

	f, _ := os.Create(s+"-api/src/"+s+"/transport.go")
	_, _ = f.Write([]byte(string(b)))
	_ = f.Close()

	err = parseTransport(s+"-api/src/"+s+"/transport.go",s,dir)
	if err != nil {
		return err
	}

	f, _ = os.Open(s+"-api/src/"+s+"/transport.go")
	_, _ = io.Copy(os.Stdout, f)
	_ = f.Close()

	return nil
}

func parseTransport(path ,s,src string)  error {
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Print(err)
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		log.Println("create file: ", err)
		return err
	}

	// A sample config
	config := map[string]string{
		"Sname": s,
		"Dir" : src,
	}

	err = t.Execute(f, config)
	if err != nil {
		log.Print("execute: ", err)
		return err
	}
	f.Close()
	return nil
}