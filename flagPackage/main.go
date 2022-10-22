package main

import (
	"flag"
	"fmt"
	"net/url"
	"reflect"
)

type URLValue struct {
	URL *url.URL
}

// var endpoint = flag.String("endpoint", "myserver.aws.com", "The server this app will contact")
// var timeD = flag.Duration("duration", time.Duration(3), "")
var help = flag.Bool("help", false, "print help")
var name = flag.String("name", "", "your name")

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

var u = &url.URL{}

func init() {
	flag.Var(&URLValue{u}, "url", "URL to parse")
}

func main() {
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	if *name != "" {
		fmt.Println("Hello", *name)
	}

	if !reflect.ValueOf(*u).IsZero() {
		fmt.Printf("scheme: %q, host: %q, path: %q\n", u.Scheme, u.Host, u.Path)
	}
}
