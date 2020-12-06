package main

import (
    "io/ioutil"
    "net/http"
	"net/url"
	"os"
	"log"
	"regexp"
	"html"
)


func main() {
    
	crawlUrl(os.Getenv("APP_BASE_URL"))  	
}

func crawlUrl(inputUrl string) ([]string, []string){
 var (
		err      error
		content  string
		activeContent string
		urlToGet *url.URL
		links    []string = make([]string, 0)
		activelinks []string = make([]string, 0)
	)
   	
	if urlToGet, err = url.Parse(inputUrl); err != nil {
		log.Println(err)
		return links, activelinks
	}
	
	log.Printf("%+v", urlToGet.String())
	
	// Retrieve content of URL
	if content, activeContent, err = getUrlContent(urlToGet.String()); err != nil {
		log.Println(err)
		return links, activelinks
	}
    
	
	// Clean up HTML entities
	content = html.UnescapeString(content)
	

	if links, err = parseLinks(urlToGet, content); err != nil {
		log.Println(err)
		return links, activelinks
	}
	
	if activelinks, err = parseLinks(urlToGet, activeContent); err != nil {
		log.Println(err)
		return links, activelinks
	}

	return links, activelinks	
}


func getUrlContent(urlToGet string) (string, string, error) {
	var (
		err     error
		content []byte
		resp    *http.Response
		activeContent string
		excludeComments = regexp.MustCompile("<!--(.*?)-->")
	)

	// GET content of URL
	if resp, err = http.Get(urlToGet); err != nil {
		return "","", err
	}
	defer resp.Body.Close()

	// Check if request was successful
	if resp.StatusCode != 200 {
		return "","", err
	}

	// Read the body of the HTTP response
	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		return "","", err
	}
	activeContent=excludeComments.ReplaceAllString(string(content), "")
	
	return string(content), activeContent, err
}


func parseLinks(urlToGet *url.URL, content string) ([]string, error) {
	var (
		err       error
		links     []string = make([]string, 0)
		matches   [][]string
		findLinks = regexp.MustCompile("<a.*?href=\"(.*?)\"")
		
	)
 

	// Retrieve all anchor tag URLs from string
	matches = findLinks.FindAllStringSubmatch(content, -1)
	

	for _, val := range matches {
		var linkUrl *url.URL

		// Parse the anchr tag URL
		if linkUrl, err = url.Parse(val[1]); err != nil {
			return links, err
		}

		// If the URL is absolute, add it to the slice
		// If the URL is relative, build an absolute URL
		if linkUrl.IsAbs() {
			links = append(links, linkUrl.String())
		} else {
			links = append(links, urlToGet.Scheme+"://"+urlToGet.Host+linkUrl.String())
		}
	}

	return links, err
}


