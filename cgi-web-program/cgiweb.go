// Go CGI web program without any dependencies or external libraries.  Should
// run on any unix web server, or even windows, that has CGI (almost all unix
// has cgi enabled)
// TODO: split this into multiple examples, this example is a bit long

package main

import (
	"fmt"
	"os"
	"strings"
)

type NameVal struct {
	name string
	value string
}

var UrlVars []NameVal

func BR() {
	fmt.Println(`<br />`)
}

func outln(s ...interface{}) {
	fmt.Println(s...)
}

func outbr(s ...interface{}) {
	outln(s...)
	BR()
}

func showArgs() {
	fmt.Println(`How many arguments:`, len(os.Args)-1)
	BR()
	for i, arg := range os.Args[1:] {
		fmt.Println(`Arg`, i+1, `:`, arg)
		BR()
	}
}

func showEnv() {
	outbr(`Showing environment variables...`)
	for i, env := range os.Environ() {
		outbr(`Env pair`, i, `:`, env)
	}
}

func getEnv(name string) string {
	env := os.Getenv(name)
	return env
}

func showQueryStr() {
	outbr(`Showing query string...`)
	v := os.Getenv(`QUERY_STRING`)
	outbr(v)
}

// takes a slice and returns a name value struct, ok false if a problem
func makeNameVal(strSlice []string) (ok bool, nv NameVal) {
	ok = false
	length := len(strSlice)
	// only one name/value pair allowed (one equal = per pair)
	if length != 2 {
		fmt.Println("ERROR: problem making a name/value pair. Possible duplicate equal sign in string")
		return
	} else {
		nv.name = strSlice[0]
		nv.value = strSlice[1]
		ok = true
	}
	return
}

func parse(pairs []string) (result []NameVal) {
	for _, pair := range pairs {
		s := strings.Split(pair, `=`)
		ok, nv := makeNameVal(s)
		if ok {
			result = append(result, nv)
		}
	}
	return
}

func processUrlVars() {
	var pairs []string
	v := os.Getenv(`QUERY_STRING`)
	pairs = strings.Split(v, `&`)
	UrlVars = parse(pairs)
}

// returns url variable data by name, i.e ?p=page returns "page"
func getUrlVar(name string) (val string) {
	for _, v := range UrlVars {
		if v.name == name {
			val = v.value
			break
		}
	}
	return val
}

func showUrlVars () {
	outbr(`Global URL Variables:`)
	for _, v := range UrlVars {
		outbr(`name:`,v.name)
		outbr(`value:`,v.value)
	}
}

func header() {
	fmt.Println(`Content-type: text/html`)
	fmt.Println()
}

// path to server program such as /cgi-bin/someprog.cgi
func cgiPath() string {
    page := getUrlVar(`p`) // page name url var
	path := getEnv(`REQUEST_URI`)
	path += `?p=`+page
	return path
}

func form() {
	outbr(
`<form  action='`+ cgiPath() +`' method="post">
	<input type='textarea' name='txta'></input>
	<input type='text' name='txt'></input>
	<input type='password' name='pw'></input>
</form>`)

}

func htmlStart() {
	fmt.Println("<html><head></head><body>")
}

func htmlEnd() {
	fmt.Println("</body></html>")
}


func main() {
	header()
	htmlStart()
	processUrlVars()
	outbr(`Go cgi program, showing URL arguments:`)
	outbr(`URL "p" var:`, getUrlVar(`p`))
	outbr(`URL "other" var:`, getUrlVar(`other`))
	BR()
	showQueryStr()
	showUrlVars()
	showArgs()
	showEnv()
	htmlEnd()
}
