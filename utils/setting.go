package utils

import (
	"bytes"
	"io/ioutil"

	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load(
		[]byte("raw data"), // 原始数据
		"filename",         // 文件路径
		ioutil.NopCloser(bytes.NewReader([]byte("some other data"))),
	)
}
