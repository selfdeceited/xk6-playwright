package playwright

import (
	"io/fs"
	"io/ioutil"
	"log"
	"time"

	"github.com/mxschmitt/playwright-go"
	"go.k6.io/k6/js/modules"
)

//Register the extension on module initilization, avaialable to
//import from JS as "k6/x/playwright".
func init() {
	modules.Register("k6/x/playwright", new(Playwright))
}

//Playwright is the k6 extension for a playwright-go client.
type Playwright struct {
	Self    *playwright.Playwright
	Browser playwright.Browser
	Page    playwright.Page
}

//Launch starts the playwright client and launches a browser
func (p *Playwright) Launch(args []string) {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Args:     args,
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	p.Self = pw
	p.Browser = browser
}

func (p *Playwright) LaunchCustomBrowser(args []string, browserPath string, driver string) {
	runOptions := playwright.RunOptions{
		DriverDirectory: driver,
		SkipInstallBrowsers: true,
		Browsers: []string {browserPath},
	}
	pw, err := playwright.Run(&runOptions)
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Args:     args,
		ExecutablePath: &browserPath,
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	p.Self = pw
	p.Browser = browser
}

//NewPage opens a new page within the browser
func (p *Playwright) NewPage() {
	page, err := p.Browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	p.Page = page
}

//Kill closes browser instance and stops puppeteer client
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

//Goto wrapper around playwright goto page function that takes in a url and a set of options
func (p *Playwright) Goto(url string, opts playwright.PageGotoOptions) {
	if _, err := p.Page.Goto(url, opts); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
}

//WaitForSelector wrapper around playwright waitForSelector page function that takes in a selector and a set of options
func (p *Playwright) WaitForSelector(selector string, opts playwright.PageWaitForSelectorOptions) {
	if _, err := p.Page.WaitForSelector(selector, opts); err != nil {
		log.Fatalf("error with waiting for the selector: %v", err)
	}
}

//Click wrapper around playwright click page function that takes in a selector and a set of options
func (p *Playwright) Click(selector string, opts playwright.PageClickOptions) {
	if err := p.Page.Click(selector, opts); err != nil {
		log.Fatalf("error with clicking: %v", err)
	}
}

//Type wrapper around playwright type page function that takes in a selector, string, and a set of options
func (p *Playwright) Type(selector string, typedString string, opts playwright.PageTypeOptions) {
	if err := p.Page.Type(selector, typedString, opts); err != nil {
		log.Fatalf("error with typing: %v", err)
	}
}

//PressKey wrapper around playwright Press page function that takes in a selector, key, and a set of options
func (p *Playwright) PressKey(selector string, key string, opts playwright.PagePressOptions) {
	if err := p.Page.Press(selector, key, opts); err != nil {
		log.Fatalf("error with pressing the key: %v", err)
	}
}

//Sleep wrapper around playwright waitForTimeout page function that sleeps for the given `timeout` in milliseconds
func (p *Playwright) Sleep(time float64) {
	p.Page.WaitForTimeout(time)
}

//Screenshot wrapper around playwright screenshot page function that attempts to take and save a png image of the current screen.
func (p *Playwright) Screenshot(filename string, perm fs.FileMode, opts playwright.PageScreenshotOptions) {
	image, err := p.Page.Screenshot(opts)
	if err != nil {
		log.Fatalf("error with taking a screenshot: %v", err)
	}
	err = ioutil.WriteFile("Screenshot_" + time.Now().Format("2017-09-07 17:06:06") + ".png", image, perm)
	if err != nil {
		log.Fatalf("error with writing the screenshot to the file system: %v", err)
	}
}

//Focus wrapper around playwright focus page function that takes in a selector and a set of options
func (p *Playwright) Focus(selector string, opts playwright.PageFocusOptions) {
	if err := p.Page.Focus(selector); err != nil {
		log.Fatalf("error focusing on the element: %v", err)
	}
}

//Fill wrapper around playwright fill page function that takes in a selector, text, and a set of options
func (p *Playwright) Fill(selector string, filledString string, opts playwright.FrameFillOptions) {
	if err := p.Page.Fill(selector, filledString, opts); err != nil {
		log.Fatalf("error with fill: %v", err)
	}
}

//DragAndDrop wrapper around playwright draganddrop page function that takes in two selectors(source and target) and a set of options
func (p *Playwright) DragAndDrop(sourceSelector string, targetSelector string, opts playwright.FrameDragAndDropOptions) {
	if err := p.Page.DragAndDrop(sourceSelector, targetSelector, opts); err != nil {
		log.Fatalf("error with drag and drop: %v", err)
	}
}

//Evaluate wrapper around playwright evaluate page function that takes in an expresion and a set of options and evaluates the expression/function returning the resulting information.
func (p *Playwright) Evaluate(expression string, opts playwright.PageEvaluateOptions) interface{} {
	returnedValue, err := p.Page.Evaluate(expression, opts);
	if err != nil {
		log.Fatalf("error evaluating the expression: %v", err)
	}
	return returnedValue;
}