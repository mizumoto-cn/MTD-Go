package core

import (
	"fmt"
	"strings"
)

func (d *Downloader) getPartDir(filename string) string {
	return strings.SplitN(filename, ".", 2)[0]
}

func (d *Downloader) getPartFilename(filename string, partID int) string {
	partdir := d.getPartDir(filename)
	return fmt.Sprintf("%s/%s-%d", partdir, filename, partID)
}
