package wikimedia

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const api = "/api/rest_v1/"

// Page consts and vars
const cPage = "page"
const cSummary = "summary"
const cRandom = "random"
const cTitle = "title"

var cPageSummaryEndpoint = path.Join(cPage, cSummary)
var cPageTitleEndpoint = path.Join(cPage, cTitle)
var cPageSummaryRandomEndpoint = path.Join(cPage, cRandom, cSummary)
var cPageTitleRandomEndpoint = path.Join(cPage, cRandom, cTitle)

// Feed consts and vars
const cFeed = "feed"
const cOnThisDay = "onthisday"

var cFeedOnThisDayEndpoint = path.Join(cFeed, cOnThisDay)

// Wikimedia represents an instance of a wikimedia REST API.
// Specs: https://en.wikipedia.org/api/rest_v1/?spec
type Wikimedia struct {
	BaseURL *url.URL
}

// NewWikimedia creates a new Wikimedia object
// @param wikimediaHost represents the baseurl of the desired Wikimedia project. E.g. https://en.wikipedia.org, https://nl.wikipedia.org, etc.
func NewWikimedia(wikimediaHost string) (*Wikimedia, error) {
	u, err := url.Parse(wikimediaHost)
	if err != nil {
		return nil, err
	}
	u.Path = api
	return &Wikimedia{BaseURL: u}, nil
}

// getQueryURL appends the Wikimedia base url with the provided query path
// and returns the resulting url.
func (w *Wikimedia) getQueryURL(path string) (string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	return w.BaseURL.ResolveReference(u).String(), nil
}

// Query queries the provided url and tries to unmarshal the response body
// to the provided interface v.
func (w *Wikimedia) Query(queryURL string, v interface{}) error {
	query, err := w.getQueryURL(queryURL)
	if err != nil {
		return err
	}
	respBody, err := get(query)
	if err != nil {
		return err
	}
	err = json.Unmarshal(respBody, v)
	if err != nil {
		return err
	}
	return nil
}

// get does a get request for the provided url and
// returns the response as a byte array
func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//////////////////
//     Page     //
//////////////////

// GetPageSummary represents the /page/summary/{title} endpoint
func (w *Wikimedia) GetPageSummary(title string) (*Summary, error) {
	var summary Summary
	err := w.Query(path.Join(cPageSummaryEndpoint, title), &summary)
	if err != nil {
		return nil, err
	}
	return &summary, nil
}

// GetPageTitle represents the /page/title/{title} endpoint
func (w *Wikimedia) GetPageTitle(title string) (*TitleRevisions, error) {
	var titleRevisions TitleRevisions
	err := w.Query(path.Join(cPageTitleEndpoint, title), &titleRevisions)
	if err != nil {
		return nil, err
	}
	return &titleRevisions, nil
}

// GetPageSummaryRandom represents the /page/random/summary endpoint
func (w *Wikimedia) GetPageSummaryRandom() (*Summary, error) {
	var summary Summary
	err := w.Query(cPageSummaryRandomEndpoint, &summary)
	if err != nil {
		return nil, err
	}
	return &summary, nil
}

// GetPageTitleRandom represents the /page/random/title endpoint
func (w *Wikimedia) GetPageTitleRandom() (*TitleRevisions, error) {
	var revs TitleRevisions
	err := w.Query(cPageTitleRandomEndpoint, &revs)
	if err != nil {
		return nil, err
	}
	return &revs, nil
}

//////////////////
//     Feed     //
//////////////////

// GetFeedOnThisDay represents the /feed/onthisday/{eventType}/{mm}/{dd} endpoint
func (w *Wikimedia) GetFeedOnThisDay(month int, day int, eventType EventType) (*OnThisDayFeed, error) {
	var feed OnThisDayFeed
	err := w.Query(path.Join(cFeedOnThisDayEndpoint, eventType.String(), fmt.Sprintf("%02d", month), fmt.Sprintf("%02d", day)), &feed)
	if err != nil {
		return nil, err
	}
	return &feed, nil
}
