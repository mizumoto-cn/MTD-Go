package core

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type Downloader struct {
	concurrency_num int
}

func NewDownloader(cc int) *Downloader {
	return &Downloader{concurrency_num: cc}
}

// func (d *Downloader) mDownload(url, filename string, content_size int) error {
// 	return nil
// }

// func (d *Downloader) sDownload(url, filename string) error {
// 	return nil
// }

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

func (d *Downloader) getPartDir(filename string) string {
	return strings.SplitN(filename, ".", 2)[0]
}

func (d *Downloader) getPartFilename(filename string, partID int) string {
	partdir := d.getPartDir(filename)
	return fmt.Sprintf("%s/%s-%d", partdir, filename, partID)
}

func (d *Downloader) PartDownload(url, filename string, range_begin, range_end, id int) {
	if range_begin >= range_end {
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", range_begin, range_end))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	f_flag := os.O_CREATE | os.O_WRONLY // creates the path. if exist overwrite
	var filepart *os.File
	filepart, err = os.OpenFile(d.getPartFilename(filename, id), f_flag, 0666) // Unix permission bits
	if err != nil {
		log.Fatal(err)
	}
	defer filepart.Close()

	buf := make([]byte, 32*1024)
	_, err = io.CopyBuffer(filepart, res.Body, buf)
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
}
