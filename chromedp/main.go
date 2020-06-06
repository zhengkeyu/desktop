package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
)

func main() {
	parent, cancel := chromedp.NewExecAllocator(context.Background(),
		chromedp.ExecPath(`C:\Users\Administrator\AppData\Local\Google\Chrome\Application\chrome.exe`))
	defer cancel()

	ctx, cancel := chromedp.NewContext(parent)
	defer cancel()

	html := ""
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.zhipin.com/c100010000-p100123/?page=1&ka=page-1"),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.OuterHTML("body", &html, chromedp.ByQuery),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(html)
}
