package lib

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func AppendSuffixNumberIfExistsFile(filePath string) string {
	suffix_number := 1
	ext := filepath.Ext(filePath)
	base_name := filepath.Base(strings.Replace(filePath, ext, "", -1))

	if FileExists(filePath) {
		for FileExists(addSuffixName(base_name, ext, suffix_number)) {
			suffix_number++
		}

		filePath = addSuffixName(base_name, ext, suffix_number)
	}

	return filePath
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	if pathError, ok := err.(*os.PathError); ok {
		if pathError.Err == syscall.ENOTDIR {
			return false
		}
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func addSuffixName(filename string, ext_name string, suffix_number int) string {
	return filename + "_" + strconv.Itoa(suffix_number) + ext_name
}
