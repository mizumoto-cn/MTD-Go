package core

import (
	"net/http"
	"path"
)

type Downloader struct {
	concurrency_num int
}

func NewDownloader(cc int) *Downloader {
	return &Downloader{concurrency_num: cc}
}

func (d *Downloader) mDownload(url, filename string, content_size int) error {
	return nil
}

func (d *Downloader) sDownload(url, filename string) error {
	return nil
}

func (d *Downloader) Download(url string, filename string) error {
	if filename == "" {
		filename = path.Base(url) //Base returns the last element of path.
	}

	// "if err != nil oriented" programming
	res, err := http.Head(url)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusOK && res.Header.Get("Accept-Ranges") == "bytes" {
		return d.mDownload(url, filename, int(res.ContentLength))
	}

	return d.sDownload(url, filename)
}

func (d *Downloader) PartDownload(url, filename string, range_begin, range_end, i int) {
	if range_begin >= range_end {
		return
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {

	}
}
