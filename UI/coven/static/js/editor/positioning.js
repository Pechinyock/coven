export class Positioning {
    static IsEditing = false

    constructor(canvas) {
        this.canvas = canvas
        this.updateTimeout = null
        canvas.on('selection:created', (e) => { this._onSelection(e) })
        canvas.on('selection:updated', (e) => { this._onSelection(e) })
        canvas.on('selection:cleared', (e) => { this._onSelectionCleared(e) })
        canvas.on('object:moving', (e) => { this._onMoving(e) })
        canvas.on('object:modified', (e) => {
            this._onMoving(e)
            clearTimeout(this.updateTimeout)
        })
    }

    bindCoords(posXInput, posYInput) {
        if (!posXInput || !posYInput) {
            console.error('failed to bind coors control')
            return
        }

        posXInput.addEventListener('focus', () => { this._setEditing(true) })
        posXInput.addEventListener('change', () => {
            const parsedX = parseFloat(posXInput.value)
            const active = this.canvas.getActiveObject()
            if (!active) { return }
            if (active.type === 'activeSelection') { return }
            else {
                active.set('left', parsedX)
                this.canvas.renderAll()
            }
        })
        posXInput.addEventListener('blur', () => { this._setEditing(false) })

        posYInput.addEventListener('focus', () => { this._setEditing(true) })
        posYInput.addEventListener('change', () => {
            const parsedY = parseFloat(posYInput.value)
            const active = this.canvas.getActiveObject()
            if (!active) { return }
            if (active.type === 'activeSelection') { return }
            else {
                const convertedY = this.convertCoords(0, parsedY)
                active.set('top', convertedY.y)
                this.canvas.renderAll()
            }
        })
        posYInput.addEventListener('blur', () => { this._setEditing(false) })

        posXInput.disabled = false
        posYInput.disabled = false

        this.posX = posXInput
        this.posY = posYInput
    }

    convertCoords(x, y) {
        const height = this.canvas.height
        return {
            x: x,
            y: height - y
        }
    }

    _onSelection(e) {
        if (!e || !e.selected) { return }
        if (e.selected.length > 1) {
            objectPosX.disabled = true
            objectPosY.disabled = true
            return
        }
        const singleSelected = e.selected[0]
        this._updateCoords(singleSelected)
    }

    _onSelectionCleared(e) {
        objectPosX.disabled = false
        objectPosY.disabled = false

        objectPosX.value = 0
        objectPosY.value = 0
    }

    _onMoving(e) {
        const obj = e.target;
        if (!obj) { return }

        const about60FPS = 16
        this.updateTimeout = setTimeout(() => {
            this._updateCoords(e.target)
        }, about60FPS)
    }

    _updateCoords(obj) {
        if (!obj) { return }
        const canvasX = obj.get('left')
        const canvasY = obj.get('top')
        const coords = this.convertCoords(canvasX, canvasY)
        objectPosX.value = coords.x.toFixed(2)
        objectPosY.value = coords.y.toFixed(2)
    }

    _setEditing(value) {
        Positioning.IsEditing = value
    }
}