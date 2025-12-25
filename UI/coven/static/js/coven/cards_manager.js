export class CardsManager {
    constructor() {
        document.addEventListener('visibilitychange', () => {
            if (document.visibilityState === 'visible') {
                this._updateCardsView()
            }
        })
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
        const cardViewElem = delBtn.parentElement
        if (!cardViewElem) {
            console.error('failed to delete card delBtn parent not found')
            return
        }
        const cardName = cardViewElem.dataset.name
        if (!cardName || cardName === '') {
            console.error('failed to delete card, card name is empty')
            return
        }
        const cardType = cardViewElem.dataset.cardType
        if (!cardType || cardType === '') {
            console.error('failed to delete card, card type is empty')
            return
        }

        delCardNameInput.value = cardName
        delCardTypeInput.value = cardType
        
        delForm.dispatchEvent(new Event('deleteCardEvent'))
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