package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"imooc.com/ccmouse/u2pppw/crawler/config"
)

var (
	rateLimiter    = time.Tick(time.Second / config.Qps)
	verboseLogging = false
)

func SetVerboseLogging() {
	verboseLogging = true
}

func Fetch(url string) ([]byte, error) {
	<-rateLimiter

	if verboseLogging {
		log.Printf("Fetching url %s", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return nil,
			fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	r := bufio.NewReader(resp.Body)
	e := determineEncoding(r)
	utf8Reader := transform.NewReader(r, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
