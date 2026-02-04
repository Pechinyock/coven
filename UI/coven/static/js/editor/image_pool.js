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
                button.disabled = true
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
    }

    onNewImageUploaded() {
        if (!cardTypeSelector) {
            console.error('failed to find card type selector')
            return
        }
        const currentSelectedType = cardTypeSelector.value
        if (!currentSelectedType) {
            console.error('failed to get selecet card type')
            return
        }
        const url = `/ui/image-pool/${currentSelectedType}`
        setTimeout(() => {
            htmx.ajax('GET', url, {
                target: '#images',
                swap: 'innerHTML'
            })
        }, 2000)
    }

    onCardTypeChanged(selected) {
        const url = `/ui/image-pool/${selected}`
        htmx.ajax('GET', url, {
            target: '#images',
            swap: 'innerHTML'
        })
        this.selectedLocalImage = null
        if (addFromLocal) {
            addFromLocal.disabled = true
        }
    }

    onImageClicked(img) {
        if (addFromLocal) {
            addFromLocal.disabled = false
        }
        const selectedBorderStyle = '2px solid rgb(3, 252, 78)'

        if (!this.selectedLocalImage) {
            img.style.border = selectedBorderStyle
            this.selectedLocalImage = img
            return
        }
        const currentFileName = img.dataset.fileName
        const selectedFileName = this.selectedLocalImage.dataset.fileName
        if (currentFileName === selectedFileName) {
            return
        }
        img.style.border = selectedBorderStyle
        this.selectedLocalImage.style.border = '2px solid transparent'
        this.selectedLocalImage = img
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