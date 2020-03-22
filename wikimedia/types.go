package wikimedia

import (
	"time"

	"github.com/sirupsen/logrus"
)

// TitleRevisions
type TitleRevisions struct {
	Items []Title `json:"items"`
	Count uint32  `json:"count,omitempty"`
}

// Title
type Title struct {
	Title        string    `json:"title,omitempty"`
	PageID       uint32    `json:"page_id,omitempty"`
	Revision     uint32    `json:"rev,omitempty"`
	Comment      string    `json:"comment,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
	PageLanguage string    `json:"page_language,omitempty"`
}

// Summary
type Summary struct {
	Title         string      `json:"title"`
	DisplayTitle  string      `json:"displaytitle,omitempty"`
	Extract       string      `json:"extract"`
	ExtractHTML   string      `json:"extract_html,omitempty"`
	Thumbnail     Image       `json:"thumbnail,omitempty"`
	OriginalImage Image       `json:"originalimage,omitempty"`
	Description   string      `json:"description,omitempty"`
	ContentURLs   ContentURLs `json:"content_urls,omitempty"`
	ApiURLs       ApiURLs     `json:"api_urls,omitempty"`
}

// Image
type Image struct {
	Source string `json:"source"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// ContentURLs
type ContentURLs struct {
	Desktop ContentURL `json:"desktop,omitempty"`
	Mobile  ContentURL `json:"mobile,omitempty"`
}

// ContentURL
type ContentURL struct {
	Page      string `json:"page,omitempty"`
	Revisions string `json:"revisions,omitempty"`
	Edit      string `json:"edit,omitempty"`
	Talk      string `json:"talk,omitempty"`
}

// ApiURLs
type ApiURLs struct {
	Summary      string `json:"summary,omitempty"`
	Metadata     string `json:"metadata,omitempty"`
	References   string `json:"references,omitempty"`
	Media        string `json:"media,omitempty"`
	EditHTML     string `json:"edit_html,omitempty"`
	TalkPageHTML string `json:"talk_page_html,omitempty"`
}

// OnThisDayFeed
type OnThisDayFeed struct {
	Births   []OnThisDay `json:"births,omitempty"`
	Deaths   []OnThisDay `json:"deaths,omitempty"`
	Events   []OnThisDay `json:"events,omitempty"`
	Holidays []OnThisDay `json:"holidays,omitempty"`
	Selected []OnThisDay `json:"selected,omitempty"`
}

// OnThisDay
type OnThisDay struct {
	Text  string    `json:"text,omitempty"`
	Pages []Summary `json:"pages,omitempty"`
}

// EventType
type EventType uint8

const (
	EventAll EventType = iota
	EventSelected
	EventBirths
	EventDeaths
	EventHolidays
	EventEvents
)

var eventTypeStrings = [...]string{
	"all",
	"selected",
	"births",
	"deaths",
	"holidays",
	"events",
}

func (e EventType) String() string {
	if EventAll <= e && e <= EventEvents {
		return eventTypeStrings[e]
	}
	logrus.Errorf("Unsupported EventType: %v", e)
	return ""
}

// PageFormat
type PageFormat uint8

const (
	TitleFormat PageFormat = iota
	SummaryFormat
	HTMLFormat
	Related
)
