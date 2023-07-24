package main

import (
	"bufio"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func wordCount(str string) map[string]int {
	strTrim := strings.Replace(str, ".", "", -1)
	wordList := strings.Fields(strTrim)
	counts := make(map[string]int)

	for _, word := range wordList {
		if _, ok := counts[word]; ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func craw(urls string) {
	resp, err := http.Get(urls)

	if err != nil {
		_, netErrors := http.Get("https://www.google.com")

		if netErrors != nil {
			fmt.Fprintf(os.Stderr, "no internet\n")
			os.Exit(1)
		}

		return
	}

	defer resp.Body.Close()

	page := html.NewTokenizer(resp.Body)

	for {
		token_type := page.Next()

		switch {
		case token_type == html.ErrorToken:
			return
		case token_type == html.StartTagToken:
			token := page.Token()

			isAnchor := token.Data == "a"
			if !isAnchor {
				continue
			}
			ok, url := getHref(token)
			if !ok {
				continue
			}
			// Make sure the url begines in http**
			isUrl := strings.Index(url, "http") == 0
			if isUrl {
				fmt.Println(url)
			}

		}
	}
}

func getHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	return
}

func lab1() {
	// str := "This is a sample text. This text is just an example."
	fmt.Print("Lab1: Enter text: ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	for word, count := range wordCount(str) {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func lab2() {
	var n int
	fmt.Print("Lab 2: N = ")
	fmt.Scanln(&n)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")

}

func lab3() {
	var urls string
	fmt.Print("nhap URL: ")
	fmt.Scanln(&urls)
	_, err := url.ParseRequestURI(urls)
	if err != nil {
		panic(err)
	}
	craw(urls)
}

func main() {
	lab1()
	lab2()
	lab3()
}
