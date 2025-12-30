import { FromHexToRGBA, IsHexFormat, GetOpacityFromRGBA, ChangeRGBAOpacity } from "./color.js"

export class Geometry {
    constructor(canvas) {
        if (!canvas) {
            console.error('failed to create geometry canvas is not found')
            return
        }
        this.canvas = canvas
        this._newObjSpan = {
            x: 50,
            y: 50,
        }
    }

    bindCreateRect(btn) {
        if (!btn) {
            console.error('failed to bind create rectangle button, passed btn is undefined or null')
            return
        }
        btn.addEventListener('click', () => {
            const fillColor = FromHexToRGBA(this.fillColor.value)
            const strokeColor = FromHexToRGBA(this.strokeColor.value)
            const strokeWidth = parseFloat(this.strokeWidth.value)
            const roundScale = parseFloat(this.roundScale.value)
            const roundX = parseFloat(this.roundRectX.value) * roundScale
            const roundY = parseFloat(this.roundRectY.value) * roundScale
            const newRect = new fabric.Rect({
                left: this._newObjSpan.x,
                top: this._newObjSpan.y,
                width: 100,
                height: 100,
                fill: fillColor,
                stroke: strokeColor,
                rx: roundX || 0,
                ry: roundY || 0,
                strokeWidth: strokeWidth
            })
            this.canvas.add(newRect)
            this._updateNewObjCreatedPosition()
        })
        btn.disabled = false
    }

    bindCreateTriangle(btn) {
        if (!btn) {
            console.error('failed to bind create triangle button, passed btn is undefined or null')
            return
        }
        btn.addEventListener('click', () => {
            const fillColor = FromHexToRGBA(this.fillColor.value)
            const strokeColor = FromHexToRGBA(this.strokeColor.value)
            const strokeWidth = parseFloat(this.strokeWidth.value)
            const newTriangle = new fabric.Triangle({
                left: this._newObjSpan.x,
                top: this._newObjSpan.y,
                width: 100,
                height: 100,
                stroke: strokeColor,
                strokeWidth: strokeWidth,
                fill: fillColor
            })
            this.canvas.add(newTriangle)
            this._updateNewObjCreatedPosition()
        })
        btn.disabled = false
    }

    bindCreateCircle(btn) {
        if (!btn) {
            console.error('failed to bind create circle button, passed btn is undefined or null')
            return
        }
        btn.addEventListener('click', () => {
            const fillColor = FromHexToRGBA(this.fillColor.value)
            const strokeColor = FromHexToRGBA(this.strokeColor.value)
            const strokeWidth = parseFloat(this.strokeWidth.value)
            const newTriangle = new fabric.Circle({
                left: this._newObjSpan.x,
                top: this._newObjSpan.y,
                radius: 50,
                stroke: strokeColor,
                strokeWidth: strokeWidth,
                fill: fillColor
            })
            this.canvas.add(newTriangle)
            this._updateNewObjCreatedPosition()
        })
        btn.disabled = false
    }

    bindFillColor(colorPicker) {
        if (!colorPicker) {
            console.error('failed to bind fill color picker')
            return
        }
        this.fillColor = colorPicker
        colorPicker.addEventListener('input', (e) => {
            const active = this.canvas.getActiveObject()
            if (!active) { return }
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    let currentColor = obj.get('fill')
                    if (IsHexFormat(currentColor)) {
                        currentColor = FromHexToRGBA(currentColor)
                    }
                    const currentOpacity = GetOpacityFromRGBA(currentColor)
                    const targetColor = FromHexToRGBA(e.target.value, currentOpacity)
                    if (this._isGeometry(obj)) {
                        obj.set('fill', targetColor)
                    }
                });
            } else if (this._isGeometry(active)) {
                let currentColor = active.get('fill')
                if (IsHexFormat(currentColor)) {
                    currentColor = FromHexToRGBA(currentColor)
                }
                const currentOpacity = GetOpacityFromRGBA(currentColor)
                const targetColor = FromHexToRGBA(e.target.value, currentOpacity)
                active.set('fill', targetColor)
            }

            this.canvas.renderAll()
        })
        colorPicker.disabled = false
    }

    bindFillTransparency(input) {
        if (!input) {
            console.error('failed to bind transparency')
            return
        }
        input.addEventListener('input', (e) => {
            const active = this.canvas.getActiveObject()
            if (!active) { return }
            const opacity = parseInt(e.target.value) / 100
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (this._isGeometry(obj)) {
                        let currentColor = obj.get('fill')
                        if (IsHexFormat(currentColor)) {
                            currentColor = FromHexToRGBA(currentColor, opacity)
                        }
                        const completeColor = ChangeRGBAOpacity(currentColor, opacity)
                        obj.set('fill', completeColor)
                    }
                });
            } else if (this._isGeometry(active)) {
                let currentColor = active.get('fill')
                if (IsHexFormat(currentColor)) {
                    currentColor = FromHexToRGBA(currentColor, opacity)
                }
                const completeColor = ChangeRGBAOpacity(currentColor, opacity)
                active.set('fill', completeColor)
            }

            this.canvas.renderAll()
        })
        input.disabled = false
    }

    bindStrokeColor(colorPicker) {
        if (!colorPicker) {
            console.error('failed to bind geometry stroke color')
            return
        }

        colorPicker.addEventListener('input', (e) => {
            const active = this.canvas.getActiveObject()
            if (!active) { return }
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    let currentColor = obj.get('stroke')
                    if (IsHexFormat(currentColor)) {
                        currentColor = FromHexToRGBA(currentColor)
                    }
                    const currentOpacity = GetOpacityFromRGBA(currentColor)
                    const targetColor = FromHexToRGBA(e.target.value, currentOpacity)
                    if (this._isGeometry(obj)) {
                        obj.set('stroke', targetColor)
                    }
                });
            } else if (this._isGeometry(active)) {
                let currentColor = active.get('stroke')
                if (IsHexFormat(currentColor)) {
                    currentColor = FromHexToRGBA(currentColor)
                }
                const currentOpacity = GetOpacityFromRGBA(currentColor)
                const targetColor = FromHexToRGBA(e.target.value, currentOpacity)
                active.set('stroke', targetColor)
            }

            this.canvas.renderAll()
        })
        this.strokeColor = colorPicker
        colorPicker.disabled = false
    }

    bindStrokeWidth(input) {
        if (!input) {
            console.error('failed to bind geometry stroke width')
            return
        }
        input.addEventListener('input', (e) => {
            const active = this.canvas.getActiveObject()
            if (!active) { return }
            const size = parseFloat(e.target.value)
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (this._isGeometry(obj)) {
                        obj.set('strokeWidth', size)
                    }
                });
            } else if (this._isGeometry(active)) {
                active.set('strokeWidth', size)
            }

            this.canvas.renderAll()
        })
        this.strokeWidth = input
        input.disabled = false
    }

    bindRectRound(inputX, inputY, inputScale) {
        if (!inputX) {
            console.error('failed to bind rect round')
            return
        }
        if (!inputY) {
            console.error('failed to bind rect round')
            return
        }
        if (!inputScale) {
            console.error('failed to bind rect round')
            return
        }

        inputX.addEventListener('input', (e) => {
            const active = this.canvas.getActiveObject()

            if (!active) { return }
            const round = parseFloat(e.target.value) || 0
            const roundScale = parseFloat(this.roundScale.value) || 1
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (obj.type === 'rect') {
                        obj.set('rx', round * roundScale)
                    }
                });
            } else if (active.type === 'rect') {
                active.set('rx', round * roundScale)
            }

            this.canvas.renderAll()
        })

        inputY.addEventListener('input', (e) => {
            const active = this.canvas.getActiveObject()

            if (!active) { return }
            const round = parseFloat(e.target.value) || 0
            const roundScale = parseFloat(this.roundScale.value) || 1
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (obj.type === 'rect') {
                        obj.set('ry', round * roundScale)
                    }
                });
            } else if (active.type === 'rect') {
                active.set('ry', round * roundScale)
            }

            this.canvas.renderAll()
        })

        inputScale.addEventListener('input', (e) => {
            const active = this.canvas.getActiveObject()

            if (!active) { return }
            const roundScale = parseFloat(e.target.value) || 1
            const roundX = parseFloat(this.roundRectX.value) || 0
            const roundY = parseFloat(this.roundRectY.value) || 0
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (obj.type === 'rect') {
                        obj.set('ry', roundY * roundScale)
                        obj.set('rx', roundX * roundScale)
                    }
                });
            } else if (active.type === 'rect') {
                active.set('ry', roundY * roundScale)
                active.set('rx', roundX * roundScale)
            }

            this.canvas.renderAll()
        })

        this.roundRectX = inputX
        this.roundRectY = inputY
        this.roundScale = inputScale

        inputX.disabled = false
        inputY.disabled = false
        inputScale.disabled = false
    }

    _isGeometry(obj) {
        if (!obj) {
            console.error('failed to define obj type the type is null')
            return false
        }
        return obj.type === 'rect' || obj.type === 'triangle' || obj.type === 'circle'
    }

    _updateNewObjCreatedPosition() {
        const step = 20;
        const limit = 270;
        if (this._newObjSpan.x >= limit) {
            this._newObjSpan.x = 30
        }
        if (this._newObjSpan.y >= limit) {
            this._newObjSpan.y = 30
        }
        this._newObjSpan.x += step
        this._newObjSpan.y += step
    }
}