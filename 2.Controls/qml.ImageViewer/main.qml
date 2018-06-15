import QtQuick 2.9
import QtQuick.Window 2.2
import QtQuick.Controls 2.3

Window {
    visible: true
    width: 400
    maximumWidth: 400
    minimumWidth: 400
    height: 300
    maximumHeight: 300
    minimumHeight: 300
    title: "线上图像预览工具"

    BorderImage {
        id: borderImage
        x: 25
        y: 22
        width: 350
        height: 163
        source: "qrc:///assets/dog.jpg"
        visible: true
    }

    TextField {
        id: textField
        x: 25
        y: 219
        width: 233
        height: 40
        text: v.imageURL
        placeholderText: "图片地址 http(s)://"
    }

    Button {
        id: button
        x: 275
        y: 219
        text: "预览"
        onClicked: {
            popup.open()
        }
    }

    Popup {
        id: popup
        x: 100
        y: 108
        width: 200
        height: 85
        modal: true
        focus: true
        closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutside

        Rectangle
        {
            //width: 400
            //height: 300
            anchors.fill: popup
            Text {
                id: mytext
                font.pixelSize: 13
                text: qsTr("Popup 内容显示模块")
            }
        }
    }
}
