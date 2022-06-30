package upload

import (
	"path"
)

type FileType int

const TypeImage FileType = iota + 1

/*func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeM
}*/

func GetFileExt(name string) string {
	return path.Ext(name)
}

