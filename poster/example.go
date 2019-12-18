package poster_go

import (
	"github.com/gogf/gf/util/gconv"
	"github.com/skip2/go-qrcode"
	"image"
	"image/png"
	"os"
	"time"
)

func TestPoster()(bool,error){
	var bgPath = "/d.png"
	var typePath = "/ziti.ttc"
	p := "./public"
	url := "https://www.baidu.com"
	//获取二维码
	qrImage,errn := GetQRImage(url,qrcode.Medium,164)
	if errn != nil{
		return false ,errn
	}
	//获取背景
	bg,errb := os.Open(p+bgPath)
	if errb != nil{
		return false ,errb
	}
	bgImage,errb := png.Decode(bg)
	if errb != nil{
		return false ,errb
	}
	//配置二维码坐标参数
	QR_Pt := Pt{288,800}
	//新建载体图片
	png := NewPNG(0,0,640,1040)
	//设置图片切片
	MergeImage(png, bgImage, bgImage.Bounds().Min)
	MergeImage(png, qrImage, qrImage.Bounds().Min.Sub(image.Pt(QR_Pt.X,QR_Pt.Y)))

	//新建文件载体
	merged,err := NewMerged(p+"/poster"+gconv.String(time.Now().Unix())+".png")
	if err != nil{
		return false,err
	}

	//设置字体切片
	trueTypeFont,err := LoadTextType(p+typePath)
	if err != nil{
		return false,err
	}

	d1 := NewDrawText(png)
	err1 := d1.MergeText("你好",18,trueTypeFont,80,45)
	if err1 != nil{
		return false,err1
	}

	d2 := NewDrawText(png)
	//设置颜色
	d2.SetColor(173,173,173)
	title := "邀请你免费使用XXX"
	err2 := d2.MergeText(title,18,trueTypeFont,80,80)
	//合并
	err = Merge(png,merged)
	if err2 != nil{
		return false,err2
	}

	defer merged.Close()
	return true,nil
}