package core

import (
	"os"
	"sync"
)

func (d *Downloader) multiDownload(url, filename string, content_size int) error {
	part_size := content_size / d.concurrency_num

	part_dir := d.getPartDir(filename)
	os.Mkdir(part_dir, 0777) //rwx
	defer os.RemoveAll(part_dir)

	var waitgroup sync.WaitGroup
	waitgroup.Add(d.concurrency_num)

	range_init := 0

	for i := 0; i < d.concurrency_num; i++ {
		// concurrency request, i for thread id
		go func(i, range_begin int) {
			defer waitgroup.Done() //Done decrements the WaitGroup counter by one.
			range_end := range_begin + part_size
			if i == d.concurrency_num-1 {
				range_end = content_size
			} // for the last data block
			d.PartDownload(url, filename, range_begin, range_end, i)
		}(i, range_init)

		range_init += part_size + 1
	}
	waitgroup.Wait() // Wait blocks until the WaitGroup counter is zero.

}
