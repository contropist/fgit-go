package main

import (
	"fmt"
	"os"
	"strings"
)

type JsdFunc struct {
}

func (j *JsdFunc) Run(args []string) {
	switch len(args) {
	case 0:
		jsdget("", "")
	case 1:
		jsdget(args[0], "")
	case 2:
		jsdget(args[0], args[1])
	}
	os.Exit(0)
}

func jsdget(url string, path string) {
	if url == "" || url == "--help" || url == "-h" {
		fmt.Println(jsdHelpMsg)
	} else {
		downloadFile(parseToJsdUrl(url), path)
	}
}

func parseToJsdUrl(url string) string {
	url = removeHttpAndHttps(url)
	if !strings.HasPrefix(url, "raw.githubusercontent.com") {
		fmt.Print("Url is not supported!")
		os.Exit(1)
	}
	i := strings.Index(url, "/")
	url = url[i+1:]
	// <OWN>/<Repo>/<Branch>/<Path>
	// <OWN>/<Repo>@<Branch>/<Path>
	url = "https://cdn.jsdelivr.net/gh/" + replaceNth(url, "/", "@", 2)
	fmt.Print("Url -> ")
	fmt.Println(url)
	return url
}
