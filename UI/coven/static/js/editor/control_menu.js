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
        const form = document.getElementById('save-result-form');
        if (!form) {
            console.error('failed to override form not found')
            return
        }
        const cardTypeInput = document.getElementById('cardType')
        if (!cardTypeInput) {
            console.error('failed to override data input cardType not found')
            return
        }
        const cardNameInput = document.getElementById('cardName')
        if (!cardNameInput) {
            console.error('failed to override data input cardNameInput not found')
            return
        }
        const jsonDataInput = document.getElementById('jsonData')
        if (!jsonDataInput) {
            console.error('failed to override data input jsonDataInput not found')
            return
        }
        const pngDataInput = document.getElementById('pngData')
        if (!pngDataInput) {
            console.error('failed to override data input pngDataInput not found')
            return
        }
        if (!pngDataInput.value) {
            console.error('failed to override data input pngDataInput has no value')
            return
        }
        if (!jsonDataInput.value) {
            console.error('failed to override data input jsonDataInput has no value')
            return
        }

        const data = {
            cardType: cardTypeInput.value,
            cardName: cardNameInput.value,
            jsonData: jsonDataInput.value,
            pngData: pngDataInput.value
        }

        htmx.ajax('PATCH', '/card', {
            values: data,
            target: '#save-response',
            swap: 'innerHTML'
        })
    }

    savingCardTypeChanged(radio) {
        const cardType = document.getElementById('cardType')
        if (!cardType) {
            console.error('failed to get card type')
            return
        }
        cardType.value = radio.value
        const overrideBtn = document.getElementById('overrideCardBtn')
        if (!overrideBtn) {
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
        const jsonDataInput = document.getElementById('jsonData')
        if (!jsonDataInput) {
            console.error('failed to save data input jsonDataInput not found')
            return
        }
        const pngDataInput = document.getElementById('pngData')
        if (!pngDataInput) {
            console.error('failed to save data input pngDataInput not found')
            return
        }

        const dataAsURL = this.canvas.toDataURL('png')
        const base64Data = dataAsURL.split(',')[1]
        const json = JSON.stringify(this.canvas.toJSON(['id']))

        jsonDataInput.value = json
        pngDataInput.value = base64Data

        form.dispatchEvent(new Event('saveCanvasEvent'))
    }
}