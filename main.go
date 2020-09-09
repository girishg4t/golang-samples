package main

import (
	"flag"
	f "flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// p1 := p.PillType{Pill: 3}
	// fmt.Println(p1.String())
	// greet.Hello()
	test4()
}
func test5() {

}
func test4() {
	str := "get pods -n=test --cluster-name minikube "
	i := strings.Index(str, "--cluster-name") + 15
	var s string
	for j := i; j < len(str); j++ {
		if string(str[j]) == " " {
			break
		}
		s += string(str[j])
	}
	fmt.Println(s)
}
func test3() {
	var flags f.FlagSet

	flags.Init("test", f.ContinueOnError)

	//	var v flagVar

	str := flags.String("param1", "", "usage")
	flags.Parse([]string{"-param1", "3", "--param2=2"})
	fmt.Println(*str)
}
func test2() {
	var flags f.FlagSet

	flags.Init("test", f.ContinueOnError)

	var v flagVar
	//argument := "--param1=value -param2 other rest of the string"

	//fileSlice := strings.Split(argument, " ")
	//	sr := flags.String("param1", "", "a string")
	os.Args = []string{"cmd", "--param1=value", "-param2", "other", "rest", "of the string"}
	flags.Parse(os.Args)
	flags.Var(&v, "param1", "usage")

	fmt.Println(v.String())

}
func (f *flagVar) String() string {

	return fmt.Sprint([]string(*f))

}

func (f *flagVar) Set(value string) error {

	*f = append(*f, value)

	return nil

}

type flagVar []string

func test() {

	os.Args = []string{"cmd", "--param1=value", "-param2", "other", "rest", "of the string"}

	sr := flag.Lookup("param1")
	flag.Parse()

	fmt.Println(*sr)

}
