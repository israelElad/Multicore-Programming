//Exercise 5.13: Modify crawl to make local copies of the pages it finds, creating directories as necessary.
//Don’t make copies of pages that come from a different domain. For example, if the original page comes from golang.org,
//save all files from there, but exclude ones from vimeo.com.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	save_local_copy(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

var originalDomain string

func save_local_copy(urlAsStr string) {
	//Parse parses rawurl into a URL structure.
	urlStruct, err := url.Parse(urlAsStr)
	if err != nil {
		log.Print(err)
	}
	//first time- define the original domain as current url
	if len(originalDomain) == 0 {
		originalDomain = urlStruct.Host
	}
	//same domain as the original domain - save it
	if originalDomain == urlStruct.Host {
		var directory string
		var filename string

		fileCheck, _ := regexp.MatchString(".+\\..+", path.Base(urlStruct.Path))
		//file
		if fileCheck {
			//domain + path to file without the file name(only Dir)
			directory = path.Join(urlStruct.Host, path.Dir(urlStruct.Path))
			//Trim returns a slice of the string with all leading and trailing Unicode code points contained in cutset removed.
			filename = strings.Trim(path.Base(urlStruct.Path), "/")
		} else { //dir
			//domain + path to dir
			directory = path.Join(urlStruct.Host, urlStruct.Path)
			filename = "page.html"
		}

		//MkdirAll creates a directory named path, along with any necessary parents, and returns nil, or else returns an error.
		//The permission bits perm (before umask) are used for all directories that MkdirAll creates. If path is already a directory, MkdirAll does nothing and returns nil.
		err = os.MkdirAll(directory, 0777)
		if err != nil {
			log.Print(err)
		}
		resp, err := http.Get(urlAsStr)
		if err != nil {
			log.Print(err)
		}
		defer resp.Body.Close()
		file, err := os.Create(path.Join(directory, filename))
		if err != nil {
			log.Print(err)
		}
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Print(err)
		}
		err = file.Close()
		if err != nil {
			log.Print(err)
		}
	}
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
