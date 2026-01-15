export class Positioning {
    static IsEditing = false

    constructor(canvas) {
        this.canvas = canvas
        this.updateTimeout = null
        canvas.on('selection:created', (e) => { this._onSelection(e) })
        canvas.on('selection:updated', (e) => { this._onSelection(e) })
        canvas.on('selection:cleared', (e) => { this._onSelectionCleared(e) })
        canvas.on('object:moving', (e) => { this._onMoving(e) })
        canvas.on('object:rotating', (e) => {
            this._onRotating(e)
        })
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

    bindRotation(rotationInput) {
        if (!rotationInput) {
            console.log('failed to bind rotation')
            return
        }
        rotationInput.addEventListener('blur', () => { this._setEditing(false) })
        rotationInput.addEventListener('focus', () => { this._setEditing(true) })
        rotationInput.addEventListener('change', () => {
            const active = this.canvas.getActiveObject()
            if (!active) { return }
            if (active.type === 'activeSelection') { return }
            else {
                const parsedRotation = parseFloat(rotationInput.value)
                active.set('angle', parsedRotation)
                this.canvas.renderAll()
            }
        })
        rotationInput.addEventListener('input', () => {
            const value = parseFloat(rotationInput.value)
            const min = parseFloat(rotationInput.min)
            const max = parseFloat(rotationInput.max)

            if (isNaN(value)) { return }

            if (value < min) rotationInput.value = min
            if (value > max) rotationInput.value = max
        })
        rotationInput.disabled = false
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
            objectRotation.disabled = true
            return
        }
        const singleSelected = e.selected[0]
        this._updateCoords(singleSelected)
        this._updateRotation(singleSelected)
    }

    _onSelectionCleared(e) {
        objectPosX.disabled = false
        objectPosY.disabled = false
        objectRotation.disabled = false

        objectPosX.value = 0
        objectPosY.value = 0
        objectRotation.value = 0
    }

    _onMoving(e) {
        const obj = e.target
        if (!obj) { return }

        const about60FPS = 16
        this.updateTimeout = setTimeout(() => {
            this._updateCoords(obj)
        }, about60FPS)
    }

    _onRotating(e) {
        const obj = e.target
        if (!obj) { return }
        const about60FPS = 16
        this.updateTimeout = setTimeout(() => {
            this._updateRotation(obj)
        }, about60FPS)
    }

    _updateRotation(obj) {
        if (!obj) { return }
        const rotation = obj.get('angle').toFixed(0)
        objectRotation.value = rotation
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