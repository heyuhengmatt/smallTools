// Package picscan 处理一些图片文件夹的逻辑
package picscan

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/schollz/progressbar/v3"
)

type list []string

// ScanAndRenamePics  扫描指定目录下的所有图片文件，并重命名为文件夹名+序号的文件名
//
//	@param baseDir
func ScanAndRenamePics(baseDir string) {
	fileList := new(list)
	filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		if d.Type().IsDir() {
			fileList.processFileList()
			return nil
		}
		ext := filepath.Ext(path)
		if strings.Contains(path, "图包") && (ext == ".jpg" || ext == ".png" || ext == ".mp4") {
			*fileList = append(*fileList, path)
		}
		if ext == ".html" || ext == "" {
			os.Remove(path)
		}
		return nil
	})
	fileList.processFileList()
}

func (fileList *list) processFileList() {
	if len(*fileList) > 0 {
		tmpBasePath := filepath.Dir(filepath.Dir((*fileList)[0]))
		realName := filepath.Base(tmpBasePath)
		bar := progressbar.Default(int64(len(*fileList)), "Processing "+tmpBasePath)
		for i, file := range *fileList {
			os.Rename(file, filepath.Join(tmpBasePath, fmt.Sprintf("%s-%d", realName, i)))
			bar.Add(1)
		}
		*fileList = make([]string, 0)
	}
}
