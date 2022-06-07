package core

import (
	"io"
	"os"
)

func (d *Downloader) merge(filename string) error {
	// return nil // to be implemented
	dest_file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer dest_file.Close()

	for i := 0; i < d.concurrency_num; i++ {
		partfile_name := d.getPartFilename(filename, i)
		part_file, err := os.Open(partfile_name)
		if err != nil {
			return err
		}
		io.Copy(dest_file, part_file)
		part_file.Close()
		os.Remove(partfile_name)
	}
	return nil
}
