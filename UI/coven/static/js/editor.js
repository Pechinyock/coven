const canvas = new fabric.Canvas('card-canvas', {
    backgroundColor: '#ffffffff'
})

const text = new fabric.Textbox('some text', {
    left: 50,
    top: 50,
    width: 200,
    fontSize: 20
})
canvas.add(text)