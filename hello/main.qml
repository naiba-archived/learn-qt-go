import QtQuick 2.9
import QtQuick.Window 2.2

Window {
    visible: true
    width: 400
    height: 200
    title: qsTr("Hello World")

    Text {
        id: text1
        x: 26
        y: 62
        width: 349
        height: 77
        text: "Hello World"
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignHCenter
        font.pixelSize: 66
    }
}
