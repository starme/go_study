package config

import (
	"io/ioutil"
	"path/filepath"
)

type FileInfo struct {
	Path string
	Name string
	Ext  string
}

func Scan(p string) []FileInfo {
	var fileList []FileInfo
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return []FileInfo{}
	}
	for _, file := range files {
		if file.IsDir() {
			fileList = append(fileList, Scan(p+"/"+file.Name())...)
		} else {
			name, ext := parseFileInfo(file.Name())
			fileList = append(fileList, FileInfo{
				Path: p, Name: name, Ext: ext[1:],
			})
		}
	}
	return fileList
}

func parseFileInfo(filename string) (name, ext string) {
	ext = filepath.Ext(filename)
	name = filename[:len(filename)-len(ext)]
	return
}
