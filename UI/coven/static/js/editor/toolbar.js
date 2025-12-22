import { FromHexToRGBA, IsHexFormat, GetOpacityFromRGBA, ChangeRGBAOpacity } from "./color.js"

export class Toolbar {
    constructor(canvas) {
        if (!canvas) {
            console.error('failed to initialze canvas toolbar: provided canvas is null')
            return
        }
        this.canvas = canvas
    }

    bindColorPicker(colorPickerId, textOpacityId) {
        const picker = document.getElementById(colorPickerId)
        if (!picker) {
            console.error('failed to bind color picker')
            return
        }
        const textOpacity = document.getElementById(textOpacityId)
        if (!textOpacity) {
            console.error('failed to bind color transparancy')
            return
        }

        textOpacity.addEventListener('input', (e) => {
            const active = this.canvas.getActiveObject()
            if (!active) { return }
            const opacity = parseInt(e.target.value) / 100
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (obj.type.includes('text')) {
                        let currentColor = obj.get('fill')
                        if (IsHexFormat(currentColor)) {
                            currentColor = FromHexToRGBA(currentColor, opacity)
                        }
                        const completeColor = ChangeRGBAOpacity(currentColor, opacity)
                        obj.set('fill', completeColor)
                    }
                });
            } else if (active.type.includes('text')) {
                let currentColor = active.get('fill')
                if (IsHexFormat(currentColor)) {
                    currentColor = FromHexToRGBA(currentColor, opacity)
                }
                const completeColor = ChangeRGBAOpacity(currentColor, opacity)
                active.set('fill', completeColor)
            }

            this.canvas.renderAll()
        })

        picker.addEventListener('input', (e) => {
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
                    if (obj.type.includes('text')) {
                        obj.set('fill', targetColor)
                    }
                });
            } else if (active.type.includes('text')) {
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

        this.textFillColor = picker
        this.textTransperancy = textOpacity

        picker.disabled = false
        textOpacity.disabled = false
    }

    bindFontSize(id) {
        const fontSizeInput = document.getElementById(id)
        if (!fontSizeInput) {
            console.error('failed to bind font size element')
            return
        }
        fontSizeInput.addEventListener('input', (e) => {
            const active = this.canvas.getActiveObject()
            if (!active) { return }

            const size = parseInt(e.target.value)

            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (obj.type.includes('text')) {
                        obj.set('fontSize', size)
                    }
                });
            } else if (active.type.includes('text')) {
                active.set('fontSize', size)
            }

            this.canvas.renderAll()
        })
        fontSizeInput.disabled = false
        this.textFontSize = fontSizeInput
    }

    bindAddText(id) {
        const btn = document.getElementById(id)
        if (!btn) {
            console.error(`failed to bind add text button, the button with id ${id} is not found`)
            return
        }
        btn.addEventListener('click', () => {
            this.addText('new text')
        })
        btn.disabled = false
    }

    bindTextStroke(strokeColorId, strokeOpacityId, strokeOpacityWidthId) {
        const colorPicker = document.getElementById(strokeColorId)
        if (!colorPicker) {
            console.error('failed to bind text stroke color')
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
                    if (obj.type.includes('text')) {
                        obj.set('stroke', targetColor)
                    }
                });
            } else if (active.type.includes('text')) {
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
        colorPicker.disabled = false

        const strokeOpacity = document.getElementById(strokeOpacityId)
        if (!strokeOpacity) {
            console.error('failed to bind text stroke opacity')
            return
        }

        strokeOpacity.addEventListener('input', (e) => {
            const active = this.canvas.getActiveObject()
            if (!active) { return }
            const opacity = parseInt(e.target.value) / 100
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (obj.type.includes('text')) {
                        let currentColor = obj.get('stroke')
                        if (IsHexFormat(currentColor)) {
                            currentColor = FromHexToRGBA(currentColor, opacity)
                        }
                        const completeColor = ChangeRGBAOpacity(currentColor, opacity)
                        obj.set('stroke', completeColor)
                    }
                });
            } else if (active.type.includes('text')) {
                let currentColor = active.get('stroke')
                if (IsHexFormat(currentColor)) {
                    currentColor = FromHexToRGBA(currentColor, opacity)
                }
                const completeColor = ChangeRGBAOpacity(currentColor, opacity)
                active.set('stroke', completeColor)
            }

            this.canvas.renderAll()
        })
        strokeOpacity.disabled = false

        const strokeWidth = document.getElementById(strokeOpacityWidthId)
        if (!strokeWidth) {
            console.error('failed to bind text stroke width')
            return
        }

        strokeWidth.addEventListener('input', (e) => {
            const active = this.canvas.getActiveObject()
            if (!active) { return }
            const size = parseFloat(e.target.value)
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (obj.type.includes('text')) {
                        obj.set('strokeWidth', size)
                    }
                });
            } else if (active.type.includes('text')) {
                active.set('strokeWidth', size)
            }

            this.canvas.renderAll()
        })
        strokeWidth.disabled = false

        this.textStrokeColor = colorPicker
        this.textStrokeOpacity = strokeOpacity
        this.textStrokeWidth = strokeWidth
    }

    addText(text) {
        const txtFillColor = FromHexToRGBA(this.textFillColor.value)
        const strokeWidth = parseInt(this.textStrokeWidth.value)
        const strokeOpacity = parseInt(this.textStrokeOpacity.value) / 100
        const strokeColor = FromHexToRGBA(this.textStrokeColor.value, strokeOpacity)
        const txt = new fabric.Textbox(text, {
            left: 50,
            top: 50,
            width: 90,
            fontSize: this.textFontSize.value,
            fill: txtFillColor,
            stroke: strokeColor,
            strokeWidth: strokeWidth
        });
        this.canvas.add(txt)
    }
}