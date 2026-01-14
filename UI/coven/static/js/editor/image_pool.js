export class ImagePool {
    constructor(canvas) {
        this.canvas = canvas
    }

    bindTestButton(button) {
        if (!button) {
            console.error('failed to bind test button')
            return
        }
        button.addEventListener('click', () => {
            const imgUri = '/image-pool/characters/vanessa.png'
            fabric.Image.fromURL(imgUri, (img) => {
                this.canvas.add(img)
                this.canvas.centerObject(img)
                this.canvas.renderAll()
            })
        })
    }
}