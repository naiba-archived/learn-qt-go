import QtQuick 2.9
import QtQuick.Window 2.2
import QtQuick.Controls 2.3

Window {
    visible: true
    width: 400
    height: 300
    title: qsTr("图片预览")

    BorderImage {
        id: borderImage
        x: 25
        y: 22
        width: 350
        height: 163
        source: "qrc:///assets/dog.jpg"
        visible: true
        clip: false
    }

    TextField {
        id: textField
        x: 25
        y: 219
        width: 233
        height: 40
        text: qsTr("")
        placeholderText: "图片地址 http(s)://"
    }

    Button {
        id: button
        x: 275
        y: 219
        text: qsTr("预览")
    }
}
