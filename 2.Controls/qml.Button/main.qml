import QtQuick 2.9
import QtQuick.Window 2.2
import QtQuick.Controls 2.3

Window {
    visible: true
    width: 400
    height: 200
    title: "Hello World"

    Button {
        id: button
        x: 132
        y: 79
        width: 136
        height: 43
        text: "Hello"
        font.pointSize: 31
        focusPolicy: Qt.NoFocus
        onClicked: {
            if (button.text == "Hello"){
                button.text = "world"
            }else{
                button.visible = false
            }
        }
    }
}
