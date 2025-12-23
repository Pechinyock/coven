export class ControlMenu {
    constructor(canvas) {
        if (!canvas) {
            console.error('failed to initialize control menu, canvas is null')
            return
        }
        this.canvas = canvas
    }

    bindSaveButton(id) {
        const saveButton = document.getElementById(id)
        if (!saveButton) {
            console.error('failed to bind save button')
            return
        }
        saveButton.addEventListener('click', () => this._saveCanvas())
        saveButton.disabled = false
    }

    _saveCanvas() {
        const form = document.getElementById('save-result-form');
        if (!form) {
            console.error('failed to save form not found')
            return
        }
        const cardTypeInput = document.getElementById('cardType')
        if (!cardTypeInput) {
            console.error('failed to save data input cardType not found')
            return
        }
        const dataTypeInput = document.getElementById('dataType')
        if (!cardTypeInput) {
            console.error('failed to save data input cardType not found')
            return
        }
        const dataIntput = document.getElementById('data')
        if (!cardTypeInput) {
            console.error('failed to save data input cardType not found')
            return
        }
        dataTypeInput.value = 'png'
        const dataAsURL = this.canvas.toDataURL('png')
        const base64Data = dataAsURL.split(',')[1]
        dataIntput.value = base64Data
        form.dispatchEvent(new Event('saveCanvasEvent'))

        dataTypeInput.value = 'json'
        const json = JSON.stringify(this.canvas.toJSON())
        dataIntput.value = json
        form.dispatchEvent(new Event('saveCanvasEvent'))
    }
}