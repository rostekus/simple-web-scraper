package controller

import (
	"log/slog"
	"runtime"
	"sync"

	"github.com/rostekus/simple-web-scraper/internal/utils/logger/sl"
	"github.com/rostekus/simple-web-scraper/internal/words"
)

type Scraper interface {
	CalcFreqWords(string) (map[string]uint, error)
}

type Cache interface {
	Get(string) (map[string]uint, bool)
	Set(string, map[string]uint)
}

type UrlIter interface {
	Next() (string, error)
}

type ControllerOpts struct {
	Cache       Cache
	ResultsChan chan words.WordFreq
	UrlIter     UrlIter
	Log         *slog.Logger
	MaxGo       int
	Scraper     Scraper
}

type Controller struct {
	*ControllerOpts
	urlChan   chan string
	boundChan chan struct{}
	wg        sync.WaitGroup
}

func NewController(opts *ControllerOpts) *Controller {
	if opts.MaxGo == 0 {
		opts.MaxGo = runtime.NumCPU()
	}
	ctrl := &Controller{
		ControllerOpts: opts,
		urlChan:        make(chan string),
		wg:             sync.WaitGroup{},
		boundChan:      make(chan struct{}, opts.MaxGo),
	}

	return ctrl
}

func (c *Controller) Serve() {
	go c.readUrls()
	go c.processUrls()
}

func (c *Controller) readUrls() {
	// write urls to channel
	defer close(c.urlChan)
	for {
		url, err := c.UrlIter.Next()
		if err != nil {
			return
		}
		c.wg.Add(1)
		c.urlChan <- url
	}

}
func (c *Controller) processUrls() {
	// read urls from chan -> send resault
	defer close(c.ResultsChan)
	for url := range c.urlChan {
		url := url
		// check if we already processed url
		if _, ok := c.Cache.Get(url); ok {
			c.Log.Info("skipping url", slog.String("url", url))
			// We can do something with data in cache here
			// maybe read data from cashe and send to results chan
			c.wg.Done()
			continue
		}
		c.Cache.Set(url, nil)
		c.Log.Info("processing url", slog.String("url", url))
		// process url in saperate goroutine
		// use chan to block starting of new goroutine
		// used empty chan for saving memeory
		c.boundChan <- struct{}{}
		go c.processUrl(url)
	}

	c.wg.Wait()
}

func (c *Controller) processUrl(url string) {
	defer c.wg.Done()
	freqMap, err := c.Scraper.CalcFreqWords(url)
	if err != nil {
		c.Log.Error("cannot process url", slog.String("url", url), sl.Err(err))
	}
	c.Cache.Set(url, freqMap)
	for word, freq := range freqMap {
		wF := words.WordFreq{
			Word: word,
			Freq: freq,
		}
		c.ResultsChan <- wF
	}

}
