package main

import (
	"github.com/bertvanpoecke/wikimedia-rest-to-go-interface/wikimedia"
	"github.com/sirupsen/logrus"
)

func main() {
	w, err := wikimedia.NewWikimedia("https://en.wikipedia.org")
	if err != nil {
		logrus.Fatal(err)
	}

	sum, err := w.GetPageSummary("Belgium")
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("%v", sum)

	title, err := w.GetPageTitle("Belgium")
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("%v", title)

	sumRandom, err := w.GetPageSummaryRandom()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("%v", sumRandom)

	titleRandom, err := w.GetPageTitleRandom()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("%v", titleRandom)

	onthisday, err := w.GetFeedOnThisDay(3, 22, wikimedia.EventAll)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("%v", onthisday)
}
