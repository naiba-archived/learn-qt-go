package main

import (
	"os"

	"log"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

type viewer struct {
	core.QObject

	_ func() bool `slot:"downloadImage"`
	_ string      `property:"imageURL"`
	_ string      `property:"remoteURL"`
	_ string      `property:"errorMsg"`
}

func (v *viewer) downloadImage() bool {
	remoteURL := v.RemoteURL()
	if !strings.HasPrefix(remoteURL, "http") {
		v.SetErrorMsg("输入地址不正确")
		return false
	}
	file, err := downloadFromURL(remoteURL)
	if err != nil {
		v.SetErrorMsg("下载失败:" + err.Error())
		return false
	}
	v.SetImageURL(file)
	return true
}

var v *viewer

var downloadDir string

func init() {
	currPath := getCurPath() + "/"
	if strings.HasSuffix(currPath, "/qml.ImageViewer.app/Contents/MacOS/") {
		currPath = strings.Replace(currPath, "/qml.ImageViewer.app/Contents/MacOS", "", -1)
	}

	f, _ := os.Create(currPath + "err.log")
	log.SetOutput(f)

	downloadDir = currPath + "demo-downloads/"
	_, err := os.Stat(downloadDir)
	if err != nil {
		err = os.Mkdir(downloadDir, os.ModePerm)
		if err != nil {
			log.Println(err)
			os.Exit(-1)
		}
	}
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	gui.NewQGuiApplication(len(os.Args), os.Args)

	initViewer()

	var engine = qml.NewQQmlApplicationEngine(nil)
	engine.RootContext().SetContextProperty("v", v)
	engine.Load(core.NewQUrl3("qrc:/main.qml", 0))
	gui.QGuiApplication_Exec()
}

func initViewer() {
	v = NewViewer(nil)
	v.SetImageURL("qrc:///assets/dog.jpg")
	v.SetRemoteURL(downloadDir)
	v.ConnectDownloadImage(v.downloadImage)
}
