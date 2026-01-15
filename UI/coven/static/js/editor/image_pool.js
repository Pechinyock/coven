export class ImagePool {
    constructor(canvas) {
        this.canvas = canvas
        this.selectedLocalImage = null
    }

    bindAddImageFromLocal(button) {
        if (!button) {
            console.error('failed to bind test button')
            return
        }
        button.addEventListener('click', () => {
            if (!images) {
                console.warn('trying to access images while they are not loaded')
                return
            }
            const imgHtmlElement = this.selectedLocalImage
            if (!imgHtmlElement) {
                console.error('trying add nothing')
                return
            }
            const imgUri = this._getImgSrc(imgHtmlElement)
            this._addImage(imgUri)
            fabric.Image.fromURL(imgUri, (img) => {
                this.canvas.add(img)
                this.canvas.centerObject(img)
                this.canvas.renderAll()
            })
        })
        button.disabled = false
    }

    onImageClicked(img) {
        const selectedBorderStyle = '2px solid rgb(3, 252, 78)'
        if (!this.selectedLocalImage) {
            img.style.border = selectedBorderStyle
            this.selectedLocalImage = img
            return
        }
        if (this.selectedLocalImage.src === img.src) { return }
        else {
            img.style.border = selectedBorderStyle
            this.selectedLocalImage.style.border = '2px solid transparent'
            this.selectedLocalImage = img
        }
    }

    _getImgSrc(img) {
        const srcValue = img.getAttribute('src')
        if (!srcValue) {
            console.error('failed to get src value')
            return
        }
        return srcValue
    }

    _addImage(url) {
        if (!url || url === '') {
            console.error('trying to add empty src img')
            return
        }
    }
}