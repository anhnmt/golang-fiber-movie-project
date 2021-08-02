package util

import (
	"fmt"
	"path"
	"time"
)

func StoragePoster(filename string) string {
	return fmt.Sprintf("/storage/poster/%s", StorageFilename(filename))
}

func StorageBanner(filename string) string {
	return fmt.Sprintf("/storage/banner/%s", StorageFilename(filename))
}

func StorageFilename(filename string) string {
	return fmt.Sprintf("%d-%s%s", time.Now().Unix(), MD5(filename), path.Ext(filename))
}

func Storage(url string) string {
	return fmt.Sprintf(".%s", url)
}
