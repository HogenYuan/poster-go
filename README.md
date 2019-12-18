# poster-go(海报生成包)


## Introduction
- 开发目的：用go创建海报（包括生成二维码，拼接头像 字符）


## Installation

```bash
import (
    "fmt"
     "github.com/HogenYuan/poster-go/poster"
)
```

## Run

#### 设置路径
```golang
    //工作路径
    p := "./public/"
    //背景图文件
    var bgPath = "d.png"    
    //字体文件(可选)
    var typePath = "/ziti.ttc"  
	
    //新建图片载体png
    png := NewPNG(0,0,640,1040)
    //新建文件载体
    merged,err := NewMerged(p+"/poster"+gconv.String(time.Now().Unix())+".png")
    
    defer merged.Close()
```
#### 图像操作 [如：二维码]
```golang
    //生成二维码(可选)
	url := "https://www.baidu.com"
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
   
    //设置图片切片
    MergeImage(png, bgImage, bgImage.Bounds().Min)
    MergeImage(png, qrImage, qrImage.Bounds().Min.Sub(image.Pt(QR_Pt.X,QR_Pt.Y)))
 
```
#### 文字操作 [如：昵称]
```golang
    //设置字体切片
    trueTypeFont,err := LoadTextType(p+typePath)
    if err != nil{
        return false,err
    }

    //文本1
    d1 := NewDrawText(png)
    err1 := d1.MergeText("你好",18,trueTypeFont,80,45)   //文字，大小，trueTypeFont，X，Y
    if err1 != nil{
        return false,err1
    }

    //文本2
    d2 := NewDrawText(png)
    //设置颜色(R,G,B)
    d2.SetColor(173,173,173)    
    title := "邀请你免费使用XXX"
    err2 := d2.MergeText(title,18,trueTypeFont,80,80)
```
#### 合并生成
```
    err = Merge(png,merged)
```

#### Example
```
    git clone https://github.com/HogenYuan/poster-go
    go run main.go
```
