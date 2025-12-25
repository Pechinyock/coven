export class CardsManager {
    constructor() {
        document.addEventListener('visibilitychange', () =>{ 
            if(document.visibilityState === 'visible'){
                this._updateCardsView()
            }
        })
    }

    _updateCardsView(){
        const viewRoot = document.getElementById('cards-view-root')
        if (!viewRoot){
            console.error('failed to update cards view view root not found')
            return
        }
        viewRoot.dispatchEvent(new CustomEvent('updateCards'))
    }
}

const cardsManager = new CardsManager()
window.cardsManager = cardsManager