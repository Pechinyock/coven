import { FromHexToRGBA, IsHexFormat, GetOpacityFromRGBA, ChangeRGBAOpacity } from "./color.js"
import { BindDivToRange } from "./utils.js"

export class TextEdit {
    static IsEditingRangeValue = false

    constructor(canvas) {
        if (!canvas) {
            console.error('failed to initialze canvas toolbar: provided canvas is null')
            return
        }
        this.fonts = ['Montserrat-regular', 'Tataric2_0', 'Deutsch Gothic Regular', 'Helmswald']
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
        if (opacityValueLabel) {
            opacityValueLabel.textContent = textOpacity.value
        }
        textOpacity.addEventListener('input', (e) => {
            if (opacityValueLabel) {
                opacityValueLabel.textContent = e.currentTarget.value
            }
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

        BindDivToRange(opacityValueLabel, textOpacity, 'int')

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

    bindTextStyle() {
        const boldBtn = txtBoldBtn
        if (!boldBtn) {
            console.error('failed to get bold text btn')
            return
        }
        boldBtn.addEventListener('click', () => {
            const active = this.canvas.getActiveObject()
            if (!active) {
                this._usrTxtStyleSelet.bold = this._getIsAreaPressed(boldBtn)
                return
            }
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (!obj.type.includes('text')) { return }
                    const newValue = this._getIsAreaPressed(boldBtn) ? 'bold' : 'normal'
                    obj.set('fontWeight', newValue)
                });
                boldBtn.classList.remove('mixed-state')
            } else if (active.type.includes('text')) {
                const newValue = this._getIsAreaPressed(boldBtn) ? 'bold' : 'normal'
                active.set('fontWeight', newValue)
            }
            this.canvas.renderAll()
        })

        const italicBtn = txtItalicBtn
        if (!italicBtn) {
            console.error('failed to get italic text btn')
            return
        }
        italicBtn.addEventListener('click', () => {
            const active = this.canvas.getActiveObject()
            if (!active) {
                this._usrTxtStyleSelet.italic = this._getIsAreaPressed(italicBtn)
                return
            }
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (!obj.type.includes('text')) { return }
                    const newValue = this._getIsAreaPressed(italicBtn) ? 'italic' : 'normal'
                    obj.set('fontStyle', newValue)
                });
                italicBtn.classList.remove('mixed-state')
            } else if (active.type.includes('text')) {
                const newValue = this._getIsAreaPressed(italicBtn) ? 'italic' : 'normal'
                active.set('fontStyle', newValue)
            }
            this.canvas.renderAll()
        })

        const underlineBtn = txtUnderlineBtn
        if (!underlineBtn) {
            console.error('failed to get underline text btn')
            return
        }
        underlineBtn.addEventListener('click', () => {
            const active = this.canvas.getActiveObject()
            if (!active) {
                this._usrTxtStyleSelet.underline = this._getIsAreaPressed(underlineBtn)
                return
            }
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (!obj.type.includes('text')) { return }
                    const newValue = this._getIsAreaPressed(underlineBtn)
                    obj.set('underline', newValue)
                });
                underlineBtn.classList.remove('mixed-state')
            } else if (active.type.includes('text')) {
                const newValue = this._getIsAreaPressed(underlineBtn)
                active.set('underline', newValue)
            }
            this.canvas.renderAll()
        })

        const strikethroughBtn = txtStrikethroughBtn
        if (!strikethroughBtn) {
            console.error('failed to get strikethrough text btn')
            return
        }
        strikethroughBtn.addEventListener('click', () => {
            const active = this.canvas.getActiveObject()
            if (!active) {
                this._usrTxtStyleSelet.strikethrough = this._getIsAreaPressed(strikethroughBtn)
                return
            }
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (!obj.type.includes('text')) { return }
                    const newValue = this._getIsAreaPressed(strikethroughBtn)
                    obj.set('linethrough', newValue)
                });
                strikethroughBtn.classList.remove('mixed-state')
            } else if (active.type.includes('text')) {
                const newValue = this._getIsAreaPressed(strikethroughBtn)
                active.set('linethrough', newValue)
            }
            this.canvas.renderAll()
        })

        this._usrTxtStyleSelet = {
            bold: this._getIsAreaPressed(boldBtn),
            italic: this._getIsAreaPressed(italicBtn),
            underline: this._getIsAreaPressed(underlineBtn),
            strikethrough: this._getIsAreaPressed(strikethroughBtn)
        }

        const selectionHandler = async (objects) => {
            const txtObjs = objects.filter(obj => obj.type.includes('text'))
            if (!txtObjs || txtObjs.length === 0) {
                return
            }
            const isPropertyMixed = (propName) => {
                if (txtObjs.length === 1) { return false }
                const fstValue = txtObjs[0].get(propName)
                for (let i = 1; i < txtObjs.length; ++i) {
                    const propVal = txtObjs[i].get(propName)
                    if (propVal !== fstValue) {
                        return true
                    }
                }
                return false
            }

            if (isPropertyMixed('fontStyle')) {
                italicBtn.classList.add('mixed-state')
            } else {
                const isEnabledNow = this._getIsAreaPressed(italicBtn)
                const isEnabledOnTargets = txtObjs[0].get('fontStyle') === 'italic' ? true : false
                if (isEnabledNow !== isEnabledOnTargets) {
                    await this._toggleButton(italicBtn)
                }
            }

            if (isPropertyMixed('fontWeight')) {
                boldBtn.classList.add('mixed-state')
            } else {
                const isEnabledNow = this._getIsAreaPressed(boldBtn)
                const isEnabledOnTargets = txtObjs[0].get('fontWeight') === 'bold' ? true : false
                if (isEnabledNow !== isEnabledOnTargets) {
                    await this._toggleButton(boldBtn)
                }
            }

            if (isPropertyMixed('underline')) {
                underlineBtn.classList.add('mixed-state')
            } else {
                const isEnabledNow = this._getIsAreaPressed(underlineBtn)
                const isEnabledOnTargets = txtObjs[0].get('underline')
                if (isEnabledNow !== isEnabledOnTargets) {
                    await this._toggleButton(underlineBtn)
                }
            }

            if (isPropertyMixed('linethrough')) {
                strikethroughBtn.classList.add('mixed-state')
            } else {
                const isEnabledNow = this._getIsAreaPressed(strikethroughBtn)
                const isEnabledOnTargets = txtObjs[0].get('linethrough')
                if (isEnabledNow !== isEnabledOnTargets) {
                    await this._toggleButton(strikethroughBtn)
                }
            }
        }

        this.canvas.on('selection:created', async (e) => {
            const objects = e.selected
            if (!objects || e.lenght === 0) {
                return
            }
            await selectionHandler(objects)
        });

        this.canvas.on('selection:updated', async (e) => {
            const objects = e.selected
            if (!objects || e.lenght === 0) {
                return
            }
            await selectionHandler(objects)
        });

        this.canvas.on('selection:cleared', async (e) => {
            const allBts = [boldBtn, italicBtn, underlineBtn, strikethroughBtn]
            allBts.forEach(x => {
                x.classList.remove('mixed-state')
            })

            const currentStateBold = this._getIsAreaPressed(boldBtn)
            if (currentStateBold !== this._usrTxtStyleSelet.bold) {
                await this._toggleButton(boldBtn)
            }

            const currentStateItalic = this._getIsAreaPressed(italicBtn)
            if (currentStateItalic !== this._usrTxtStyleSelet.italic) {
                await this._toggleButton(italicBtn)
            }

            const currentStateUnderline = this._getIsAreaPressed(underlineBtn)
            if (currentStateUnderline !== this._usrTxtStyleSelet.underline) {
                await this._toggleButton(underlineBtn)
            }

            const currentStateStrikethrough = this._getIsAreaPressed(strikethroughBtn)
            if (currentStateStrikethrough !== this._usrTxtStyleSelet.strikethrough) {
                await this._toggleButton(strikethroughBtn)
            }
        });

        strikethroughBtn.disabled = false
        underlineBtn.disabled = false
        italicBtn.disabled = false
        boldBtn.disabled = false
    }

    bindFontSize(id) {
        const fontSizeInput = document.getElementById(id)
        if (!fontSizeInput) {
            console.error('failed to bind font size element')
            return
        }
        if (fontSizeValueLabel) {
            fontSizeValueLabel.textContent = fontSizeInput.value
        }
        fontSizeInput.addEventListener('input', (e) => {
            if (fontSizeValueLabel) {
                fontSizeValueLabel.textContent = fontSizeInput.value
            }
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
        BindDivToRange(fontSizeValueLabel, fontSizeInput, 'int')
        fontSizeInput.disabled = false
        this.textFontSize = fontSizeInput
    }

    bindFontSelector(id) {
        const selector = document.getElementById(id)
        if (!selector) {
            console.error('failed to get selector')
            return
        }
        const fonts = this.fonts
        fonts.forEach((font, index) => {
            const option = document.createElement('option')
            option.value = font
            option.textContent = font
            option.style.fontFamily = font
            if (index === 0) {
                option.selected = true
            }
            selector.appendChild(option)
        })
        selector.addEventListener('change', () => {
            const active = this.canvas.getActiveObject()
            if (!active) { return }
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                objs.forEach(obj => {
                    if (active.type.includes('text')) {
                        return
                    }
                    obj.set('fontFamily', selector.value)
                });
                this.canvas.renderAll()
            } else if (active.type.includes('text')) {
                active.set('fontFamily', selector.value)
                this.canvas.renderAll()
            }
        })
        selector.disabled = false
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
        if (txtStrokeOpacityValueLabel) {
            txtStrokeOpacityValueLabel.textContent = strokeOpacity.value
        }
        strokeOpacity.addEventListener('input', (e) => {
            if (txtStrokeOpacityValueLabel) {
                txtStrokeOpacityValueLabel.textContent = strokeOpacity.value
            }
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

        BindDivToRange(txtStrokeOpacityValueLabel, strokeOpacity, 'int')

        strokeOpacity.disabled = false

        const strokeWidth = document.getElementById(strokeOpacityWidthId)
        if (!strokeWidth) {
            console.error('failed to bind text stroke width')
            return
        }
        if (txtStrokeWidthValueLabel) {
            txtStrokeWidthValueLabel.textContent = strokeWidth.value
        }
        strokeWidth.addEventListener('input', (e) => {
            if (txtStrokeWidthValueLabel) {
                txtStrokeWidthValueLabel.textContent = strokeWidth.value
            }
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

        BindDivToRange(txtStrokeWidthValueLabel, strokeWidth, 'float')

        strokeWidth.disabled = false

        this.textStrokeColor = colorPicker
        this.textStrokeOpacity = strokeOpacity
        this.textStrokeWidth = strokeWidth
    }

    addText(text) {
        const txtFillColor = FromHexToRGBA(this.textFillColor.value)
        const txtOpacity = parseInt(this.textTransperancy.value) / 100
        const completeFillColor = ChangeRGBAOpacity(txtFillColor, txtOpacity)
        const strokeWidth = parseFloat(this.textStrokeWidth.value)
        const strokeOpacity = parseFloat(this.textStrokeOpacity.value) / 100
        const strokeColor = FromHexToRGBA(this.textStrokeColor.value, strokeOpacity)
        const selectedFont = document.getElementById('font-selector')
        const fontStyle = this._getIsAreaPressed(txtItalicBtn) ? 'italic' : 'normal'
        const fontWeight = this._getIsAreaPressed(txtBoldBtn) ? 'bold' : 'normal'

        let font = null
        if (!selectedFont) {
            font = 'Montserrat-regular'
            console.error('failed to find font selector')
        } else {
            font = selectedFont.value
        }
        const txt = new fabric.Textbox(text, {
            left: 50,
            top: 50,
            width: 90,
            fontSize: this.textFontSize.value,
            fill: completeFillColor,
            stroke: strokeColor,
            strokeWidth: strokeWidth,
            fontFamily: font,
            fontStyle: fontStyle,
            fontWeight: fontWeight,
            underline: this._getIsAreaPressed(txtUnderlineBtn),
            linethrough: this._getIsAreaPressed(txtStrikethroughBtn)
        });
        this.canvas.add(txt)
    }

    _bindRangeFloatValueLabel(divLabel, range) {
        divLabel.addEventListener('dblclick', (e) => {
            TextEdit.IsEditingRangeValue = true
            const targetDiv = e.currentTarget
            targetDiv.contentEditable = true
            targetDiv.focus()
            targetDiv.classList.add('form-control', 'form-control-sm')
        })

        divLabel.addEventListener('blur', (e) => {
            const targetDiv = e.currentTarget
            const value = parseFloat(targetDiv.textContent)

            if (!isNaN(value) && value >= range.min && value <= range.max) {
                range.value = value
                range.dispatchEvent(new Event('input'))
            } else {
                targetDiv.textContent = range.value
            }
            targetDiv.contentEditable = false
            targetDiv.classList.remove('form-control', 'form-control-sm')
            TextEdit.IsEditingRangeValue = false
        })
    }

    _getIsAreaPressed(element) {
        if (!element) { return false }
        const attrib = element.getAttribute('aria-pressed')
        if (!attrib) { return false }
        return attrib === 'true'
    }

    async _toggleButton(btn) {
        const button = await bootstrap.Button.getOrCreateInstance(btn)
        await button.toggle()
    }

}