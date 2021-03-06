package utils

import "os"

func Check(e error) {
	if e != nil {
		panic("Checked error fails: " + e.Error())
	}
}

func Filesize(f *os.File) uint64 {
	fileInfo, err := f.Stat()
	Check(err)
	return uint64(fileInfo.Size())
}
