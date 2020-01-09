package qrcode

import (
	"FirstGin/pkg/file"
	"FirstGin/pkg/setting"
	"FirstGin/pkg/util"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/jpeg"
)

type QrCode struct{
	URL string
	Width int
	Height int
	Ext string
	Level qr.ErrorCorrectionLevel
	Mode qr.Encoding
}

const(
	EXT_JPG = ".jpg"
)

func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode{
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Ext:    EXT_JPG,
		Level:  level,
		Mode:   mode,
	}
}

func GetQrCodePath() string{
	return setting.AppSetting.QrCodeSavePath
}

func GetQrCodeFullPath() string{
	return setting.AppSetting.RuntimeRootPath + setting.AppSetting.QrCodeSavePath
}

func GetQrCodeFullUrl(name string) string{
	return setting.AppSetting.PrefixUrl + "/" + GetQrCodePath() + name
}

func GetQrCodeFileName(value string) string{
	return util.EncodeMD5(value)
}

func (q *QrCode) GetQrCodeExt() string{
	return q.Ext
}

func (q *QrCode) CheckEncode(path string) bool {
	src := path + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	if file.CheckExist(src) == true{
		return true
	}

	return false
}

func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name
	if file.CheckExist(src) == false {
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}

		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}

		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()

		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}

	return name, path, nil
}
