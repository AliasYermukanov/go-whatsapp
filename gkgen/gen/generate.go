package gen

import (
	"fmt"
	"os"
	"strings"
)

func ServiceGen(args []string){
	var s string
	for _,v := range args {
		s += v
	}


	err := os.MkdirAll(s+"-api", 0777)
	if err != nil {
		fmt.Println("error:",err.Error())
	}
	err = os.MkdirAll(s+"-api/src/"+s, 0777)
	if err != nil {
		fmt.Println("error:",err.Error())
	}
	err = os.MkdirAll(s+"-api/src/config", 0777)
	if err != nil {
		fmt.Println("error:",err.Error())
	}

	err = maingen(s)
	if err != nil {
		panic(err)
	}
	configGen(s)
	err = transportGen(s)
	if err != nil {
		panic(err)
	}
	err = endpointGen(s)
	if err != nil {
		panic(err)
	}
	err = serviceGen(s)
	if err != nil {
		panic(err)
	}
	err = loggingGen(s)
	if err != nil {
		panic(err)
	}
}

func before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}