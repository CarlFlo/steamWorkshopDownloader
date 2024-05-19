package utils

import "os"

func FileSize(filepath string) int64 {
	info, err := os.Stat(filepath)
	if err != nil {
		return -1
	}
	return info.Size()
}
