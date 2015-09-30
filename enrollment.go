package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type StringList []string

// add a set function for our StringList
func (l *StringList) Set(val string) error {
	fmt.Printf("Appending %s to files\n", val)
	*l = append(*l, val)
	return nil
}

// make it easy to print
func (l *StringList) String() string {
	return strings.Join(*l, ", ")
}

// create our Strings flag variable
func Strings(name string, def []string, desc string) *StringList {
	var s StringList = StringList(def)
	flag.Var(&s, name, desc)
	return &s
}

func main() {
	var envFiles = Strings("e", nil, "Pass any config files. Multiple allowed.")

	flag.Parse()

	for _, filename := range *envFiles {
		UpdateEnv(filename)
	}

	tail := flag.Args()

	if len(tail) > 0 {
		fmt.Println(tail)
		cmd := exec.Command(tail[0], tail[1:]...)

		cmd.Env = os.Environ() // this is the default...
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}
