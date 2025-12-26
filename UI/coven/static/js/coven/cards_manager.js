export class CardsManager {
    constructor() {
        document.addEventListener('visibilitychange', () => {
            if (document.visibilityState === 'visible') {
                this._updateCardsView()
            }
        })

        this.deleteTimeGap = null
    }

    openEditor(btn) {
        const cardType = btn.dataset.cardType
        if (!cardType) {
            console.error('failed to open editor for specific card, card type not found')
            return
        }
        const cardName = btn.dataset.name
        if (!cardName) {
            console.error('failed to open editor for specific card, card name not found')
            return
        }
        const url = `/editor?type=${encodeURIComponent(cardType)}&name=${encodeURIComponent(cardName)}`
        window.open(url, `${cardType} ${cardName}`)
    }

    startDeleteHold(btn) {
        if (this.deleteTimeGap) {
            console.warn('how does it even possible???')
            return
        }
        btn.classList.add('holding')
        const waitUntillDeleteMs = 1000
        this.deleteTimeGap = setTimeout(() => {
            this.deleteCard(btn)
        }, waitUntillDeleteMs)
    }

    cancelDeleteHold(btn) {
        if (this.deleteTimeGap) {
            clearTimeout(this.deleteTimeGap)
            this.deleteTimeGap = null
        }
        btn.classList.remove('holding')
    }

    deleteCard(delBtn) {
        const delForm = document.getElementById('delete-card-form')
        if (!delForm) {
            console.error('failed to delete card del form not found')
            return
        }
        const delCardTypeInput = document.getElementById('deleteCardType')
        if (!delCardTypeInput) {
            console.error('failed to delete card del card type input not found')
            return
        }
        const delCardNameInput = document.getElementById('deleteCardName')
        if (!delCardNameInput) {
            console.error('failed to delete card del card name input not found')
            return
        }
        const cardName = delBtn.dataset.name
        if (!cardName || cardName === '') {
            console.error('failed to delete card, card name is empty')
            return
        }
        const cardType = delBtn.dataset.cardType
        if (!cardType || cardType === '') {
            console.error('failed to delete card, card type is empty')
            return
        }

        const cardViewElem = document.getElementById(`card-${cardType}-${cardName}`)
        if (!cardViewElem) {
            console.error('failed to delete card delBtn parent not found')
            return
        }

        delCardNameInput.value = cardName
        delCardTypeInput.value = cardType

        delForm.dispatchEvent(new Event('deleteCardEvent'))
        if (this.deleteTimeGap) {
            clearTimeout(this.deleteTimeGap)
            this.deleteTimeGap = null
        }
        cardViewElem.remove()
    }

    _updateCardsView() {
        const viewRoot = document.getElementById('cards-view-root')
        if (!viewRoot) {
            console.error('failed to update cards view view root not found')
            return
        }
        viewRoot.dispatchEvent(new Event('updateCards'))
    }
}

const cardsManager = new CardsManager()
window.cardsManager = cardsManager