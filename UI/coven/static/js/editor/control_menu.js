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

    overrideCard() {
        const overrideForm = document.getElementById('override-result-form')
        if (!overrideForm) {
            console.error('failed to get override form')
            return
        }
        const cardTypeInput = document.getElementById('cardType')
        if (!cardTypeInput) {
            console.error('failed to override data input cardType not found')
            return
        }
        const cardNameInput = document.getElementById('cardName')
        if (!cardNameInput) {
            console.error('failed to override data cardName input not found')
            return
        }
        const cardNameInputOverride = document.getElementById('cardNameOverride')
        if (!cardNameInput) {
            console.error('failed to override data cardName input not found')
            return
        }
        cardNameInputOverride.value = cardNameInput.value
        const cardTypeInputOverride = document.getElementById('cardTypeOverride')
        if (!cardTypeInput) {
            console.error('failed to override data input cardType not found')
            return
        }
        cardTypeInputOverride.value = cardTypeInput.value
        const dataTypeInputOverride = document.getElementById('dataTypeOverride')
        if (!dataTypeInputOverride) {
            console.error('failed to override data input cardTypeInput not found')
            return
        }
        const dataIntputOverride = document.getElementById('dataOverride')
        if (!dataIntputOverride) {
            console.error('failed to override data input dataIntputOverride not found')
            return
        }

        dataTypeInputOverride.value = 'png'
        const dataAsURL = this.canvas.toDataURL('png')
        const base64Data = dataAsURL.split(',')[1]
        dataIntputOverride.value = base64Data
        overrideForm.dispatchEvent(new Event('overrideCardEvent'))

        dataTypeInputOverride.value = 'json'
        const json = JSON.stringify(this.canvas.toJSON(['id']))
        dataIntputOverride.value = json
        overrideForm.dispatchEvent(new Event('overrideCardEvent'))
    }

    savingCardTypeChanged(radio) {
        const cardType = document.getElementById('cardType')
        if (!cardType){
            console.error('failed to get card type')
            return
        }
        cardType.value = radio.value
        const overrideBtn = document.getElementById('overrideCardBtn')
        if (!overrideBtn){
            return
        }
        overrideBtn.remove()
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
        const json = JSON.stringify(this.canvas.toJSON(['id']))
        dataIntput.value = json
        form.dispatchEvent(new Event('saveCanvasEvent'))
    }
}