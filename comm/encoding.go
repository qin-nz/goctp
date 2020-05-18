package comm

import (
	"bytes"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// DecodeGBK convert GBK to UTF-8
func DecodeGBK(s string) (string, error) {
	I := bytes.NewReader([]byte(s))
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, err := ioutil.ReadAll(O)
	return string(d), err
}

func ShouldDecodeGBK(s string) string {
	s, err := DecodeGBK(s)
	if err != nil {
		logrus.Fatal("Decode GBK err %v", err)
		return ""
	}
	return s
}
