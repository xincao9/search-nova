package chromium

import (
	"bytes"
	"github.com/playwright-community/playwright-go"
	"golang.org/x/net/html/charset"
	"io"
	"search-nova/internal/logger"
	"search-nova/internal/shutdown"
	"strings"
)

var (
	C *chromium
)

func init() {
	var err error
	C, err = new()
	if err != nil {
		logger.L.Fatalf("Fatal error chromium: %v\n", err)
	}
}

func new() (*chromium, error) {
	c := &chromium{}
	var err error
	c.pw, err = playwright.Run()
	if err != nil {
		return nil, err
	}
	c.browser, err = c.pw.Chromium.Launch()
	if err != nil {
		return nil, err
	}
	shutdown.S.Add(func() {
		err = c.browser.Close()
		if err != nil {
			logger.L.Errorf("browser.Close() err %v\n", err)
		}
		err = c.pw.Stop()
		if err != nil {
			logger.L.Errorf("playwright.Close() err %v\n", err)
		}
	})
	return c, nil
}

type chromium struct {
	pw      *playwright.Playwright
	browser playwright.Browser
}

func (c *chromium) Html(urlS string) (io.Reader, error) {
	newPage, err := c.browser.NewPage()
	if err != nil {
		return nil, err
	}
	defer newPage.Close()
	resp, err := newPage.Goto(urlS)
	if err != nil {
		return nil, err
	}
	ct := resp.Headers()["Content-Type"]
	if ct != "" && !strings.Contains(ct, "text/html") {
		logger.L.Infof("url: %s content-type is %s\n", urlS, ct)
		return nil, nil
	}
	body, err := resp.Body()
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	b.Write(body)
	reader, err := charset.NewReader(&b, ct)
	if err != nil {
		return nil, err
	}
	return reader, nil
}
