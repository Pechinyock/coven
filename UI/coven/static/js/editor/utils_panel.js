export class UtilsPanel {
    constructor(canvas) {
        this.canvas = canvas
        this.cardEdgeLineList = [];
        this._isLoaded = false
        this.canvas.on('object:added', (e) => {
            if (!this._isLoaded) { return }
            this.cardEdgeLineList.forEach(edge => {
                edge.bringToFront()
            })
        })
    }

    bindCardEdges(toggleButton) {
        if (!toggleButton) {
            console.error('failed to bind card edges')
            return
        }
        toggleButton.addEventListener('click', () => {
            if (this.cardEdgeLineList.length !== 4) {
                console.error('failed to toggle card edge visibility')
                return
            }
            let descition = false
            this.cardEdgeLineList.forEach(edge => {
                const currentValue = edge.visible
                descition = currentValue
                edge.set('visible', !currentValue)
            })
            const hideText = 'Скрыть границу карточки'
            const showText = 'Показать границу карточки'
            if (descition) {
                // toggleButton.textContent = showText
                toggleButton.classList.remove('btn-primary')
                toggleButton.classList.add('btn-outline-primary')
            } else {
                // toggleButton.textContent = hideText
                toggleButton.classList.add('btn-primary')
                toggleButton.classList.remove('btn-outline-primary')
            }
            this.canvas.renderAll()
        })
    }

    initCardEdges() {
        if (this._isLoaded) { return }
        const strokeWidth = 2
        const skipSavePrefix = 'unsave'
        this._addCardEdge(100
            , 100 - strokeWidth
            , this.canvas.width - 100 + strokeWidth
            , 100 - strokeWidth
            , strokeWidth
            , `${skipSavePrefix}_top_horizontal`
        )
        this._addCardEdge(100 - strokeWidth
            , 100
            , 100 - strokeWidth
            , this.canvas.height - 100
            , strokeWidth
            , `${skipSavePrefix}_left_vertical`
        )
        this._addCardEdge(this.canvas.width - 100 + strokeWidth
            , 100
            , this.canvas.width - 100 + strokeWidth
            , this.canvas.height - 100
            , strokeWidth
            , `${skipSavePrefix}_right_vertical`
        )
        this._addCardEdge(100 + strokeWidth
            , this.canvas.height - 100
            , this.canvas.width - 100 + strokeWidth
            , this.canvas.height - 100
            , strokeWidth
            , `${skipSavePrefix}_bottom_horizontal`
        )
        this._isLoaded = true
    }

    _addCardEdge(x1, y1, x2, y2, edgeWidth, id) {
        const edge = new fabric.Line([x1, y1, x2, y2], {
            stroke: 'red',
            strokeWidth: edgeWidth,
            strokeDashArray: [8, 2],
            selectable: false,
            evented: false,
            id: id,
        })
        this.canvas.add(edge)
        this.canvas.renderAll()
        this.cardEdgeLineList.push(edge)
    }
}