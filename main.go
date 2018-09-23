package main

import (
	"github.com/gocolly/colly"
	"fmt"
	"time"
)

func main() {
	c := colly.NewCollector()
	c2 := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"
	c2.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("%+v\r\n%+v\r\n", *r, *(r.Headers))
	})

	c2.OnHTML("div>div.view-l>div.view-l-des", func(e *colly.HTMLElement) {
		fmt.Println("ok")
		ss := e.ChildText("div.t")
		//ss1 := e.ChildText("div.d:first-child")
		//ss2 := e.ChildText("div.d:nth-child(2)")
		//ss3 := e.ChildText("div.d:nth-child(3)")

		fmt.Println("chanpin=", ss)
	})

		// Find and visit all links
	c.OnHTML("table>tbody>tr[target]", func(e *colly.HTMLElement) {
		//time.Sleep(1* time.Second)
		ss := e.ChildText("td:first-child")
		ss2 := e.ChildText("td:nth-child(2)")
		ss3 := e.ChildText("td:nth-child(3)")
		ss4 := e.ChildText("td:nth-child(4)")
		ss5 := e.ChildText("td:nth-child(5)")
		ss6 := e.ChildText("td:nth-child(6)")
		ss7 := e.ChildText("td:nth-child(7)")
		ss8 := e.ChildText("td:nth-child(8)")
		ss9 := e.ChildText("td:nth-child(9)")
		fmt.Println(ss, ss2, ss3, ss4, ss5, ss6, ss7, ss8,ss9)

		chanpinLink := e.Attr("click-url")
		fmt.Println("--------------", chanpinLink)
		c2.Visit(chanpinLink)
	})

	last := ""
	c.OnHTML("div#page_section", func(e *colly.HTMLElement) {
		time.Sleep(2*time.Second)
		next := e.ChildAttr("a.next-page:last-of-type", "href")
		//next :=  e.Attr("href")
		fmt.Println("next=",next)

		if next < last {
			fmt.Println("end")
			return
		}
		last = next

		//c.Visit("https://www.rong360.com" +next)
	})


	c.Visit("https://www.rong360.com/licai-bank/list/")
}
