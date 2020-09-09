package main

import (
	"fmt"
	"regexp"
)

func main() {
	var re = regexp.MustCompile(`(?m)--cluster-name[=|' ']([^\s]*)`)
	var str = `gett --cluster-name minikube  dgdsgsdg`
	fmt.Println(re.FindStringSubmatch(str)[1])
}
