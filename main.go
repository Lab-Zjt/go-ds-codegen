package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	_type, _it, _container, _name string
)

func init() {
	flag.StringVar(&_type, "t", "int", "type name")
	flag.StringVar(&_container, "c", "List", "container name, support List, PriorityQueue, Queue, Stack, Deque now")
	flag.StringVar(&_name, "g", _type+_container, "generated name")
	flag.StringVar(&_it, "i", "ListIt", "iterator name")
}


func main() {
	flag.Parse()
	f, err := os.OpenFile("/home/zjt/code-repo/go/codegen/template_"+_container+".txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	stat, _ := f.Stat()
	size := stat.Size()
	buf := make([]byte, size)
	f.Read(buf)
	str := string(buf)
	str = strings.Replace(str, "{{type}}", _type, -1)
	str = strings.Replace(str, "{{it}}", _it, -1)
	str = strings.Replace(str, "{{container}}", _name, -1)
	fmt.Println(str)
}
