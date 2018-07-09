package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	_type, _it, _container, _gen, _filter, _map, _reduce, _foreach, _key, _value string
)

type A struct {
	a int
	b string
}

func init() {
	flag.StringVar(&_type, "t", "int", "type name")
	flag.StringVar(&_container, "c", "List", "container name, support List, PriorityQueue, Queue, Stack, Deque now")
	flag.StringVar(&_gen, "g", _type+_container, "generated name")
	flag.StringVar(&_it, "i", "ListIt", "iterator name")
	flag.StringVar(&_filter, "ft", _type+"FilterFunc", "filter function name")
	flag.StringVar(&_map, "mp", _type+"MapFunc", "map function name")
	flag.StringVar(&_reduce, "rd", _type+"ReduceFunc", "reduce func name")
	flag.StringVar(&_foreach, "fe", _type+"ForEachFunc", "foreach function name")
	flag.StringVar(&_key, "k", "string", "key type name")
	flag.StringVar(&_value, "v", "string", "value type name")
}

type vector []int

func main() {
	flag.Parse()
	f, err := os.OpenFile("template_"+_container+".txt", os.O_RDONLY, 0666)
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
	str = strings.Replace(str, "{{container}}", _gen, -1)
	str = strings.Replace(str, "{{key}}", _key, -1)
	str = strings.Replace(str, "{{value}}", _value, -1)
	str = strings.Replace(str, "{{mapfunc}}", _map, -1)
	str = strings.Replace(str, "{{filterfunc}}", _filter, -1)
	str = strings.Replace(str, "{{foreachfunc}}", _foreach, -1)
	str = strings.Replace(str, "{{reducefunc}}", _reduce, -1)
	fmt.Println(str)
}
