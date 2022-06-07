package core

import (
	"io"
	"log"
	"net/http"
	"os"
)

func (d *Downloader) singleDownload(url, filename string) error {
	return d.fake_single_down(url, filename)
}

func (d *Downloader) fake_single_down(url, filename string) error {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := make([]byte, 32*1024)
	_, err = io.CopyBuffer(io.MultiWriter(f), res.Body, buf)

	return err
}
