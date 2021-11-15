package playwright

import (
	"log"

	"github.com/mxschmitt/playwright-go"
	"go.k6.io/k6/js/modules"
)

// Register the extension on module initilization, avaialable to
// import from JS as "k6/x/playwright".
func init() {
	modules.Register("k6/x/playwright", new(Playwright))
}

// Playwright is the k6 extension for a playwright-go client.
type Playwright struct {
	Self *playwright.Playwright
	Browser playwright.Browser
	Page playwright.Page
}

// Starts the playwright client and launches a browser
func (p *Playwright) Launch(args []string) {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
		Args: args,
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	p.Self = pw
	p.Browser= browser
}

// Opens a new page within the browser
func (p *Playwright) NewPage(){
	page, err := p.Browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	p.Page = page
}

// Closes browser instance and stops puppeteer client
func (p *Playwright) Kill() {
	if err := p.Browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err := p.Self.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}

//---------------------------------------------------------------------
//                         ACTIONS
//---------------------------------------------------------------------

func (p *Playwright) Goto(url string, opts playwright.PageGotoOptions) {
	if _, err := p.Page.Goto(url, opts); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
}
func (p *Playwright) WaitForSelector(selector string, opts playwright.PageWaitForSelectorOptions) {
	if _, err := p.Page.WaitForSelector(selector, opts); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
}