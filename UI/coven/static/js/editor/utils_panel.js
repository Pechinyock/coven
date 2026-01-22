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
            if (descition) {
                toggleButton.classList.remove('btn-primary')
                toggleButton.classList.add('btn-outline-primary')
            } else {
                toggleButton.classList.add('btn-primary')
                toggleButton.classList.remove('btn-outline-primary')
            }
            this.canvas.renderAll()
        })
    }

    bindGroup(groupButton) {
        if (!groupButton) {
            return
        }
        groupButton.addEventListener('click', () => {
            const objs = this.canvas.getActiveObject()
            if (!objs) { return }
            const isActiveSelection = objs.type === 'activeSelection'
            const isGroup = objs.type === 'group'
            if (!isActiveSelection && !isGroup) { return }
            if (isActiveSelection) {
                objs.toGroup()
                this._toggleGroupButton(groupButton, true)
            }
            if (isGroup) {
                objs.toActiveSelection()
                this._toggleGroupButton(groupButton, false)
            }
            this.canvas.renderAll()
        })

        this.canvas.on('selection:created', (e) => {
            this._onObjectSelected(groupButton)
            if (this.openSavePartialBtn) {
                this.openSavePartialBtn.disabled = false
            }
        })
        this.canvas.on('selection:updated', (e) => {
            this._onObjectSelected(groupButton)
            if (this.openSavePartialBtn) {
                this.openSavePartialBtn.disabled = false
            }
        })
        this.canvas.on('selection:cleared', (e) => {
            this._toggleGroupButton(groupButton)
            if (this.openSavePartialBtn) {
                this.openSavePartialBtn.disabled = true
            }
        })
        groupButton.disabled = false
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

    bindSavePartial(openSavePartialBtn, savePartialBtn) {
        if (!savePartialBtn) {
            console.error('failed to bind save partial')
            return
        }
        if (!openSavePartialBtn) {
            console.error('failed to bund save partial')
            return
        }

        openSavePartialBtn.addEventListener('click', () => {
            this.previewCanvas.clear()
            this.previewCanvas.backgroundColor = '';
            this.previewCanvas.overlayImage = null;
            this.previewCanvas.renderAll();

            const obj = this.canvas.getActiveObject()
            if (!obj) {
                return
            }
            this._copySelectedToPartial(obj)
        })
        savePartialBtn.addEventListener('click', () => {
            const pngData = this._getPngData()

            this._savePartial(this._partialJson, pngData, false)
        })

        savePartialBtn.disabled = true
        this.previewCanvas = new fabric.Canvas(partialPreview)
        this.openSavePartialBtn = openSavePartialBtn
        this.savePartialBtn = savePartialBtn
    }

    overridePartial() {
        const pngData = this._getPngData()
        this._savePartial(this._partialJson, pngData, true)
    }

    _getPngData() {
        const prev = this.previewCanvas.get('backgroundColor')

        this.previewCanvas.set({ backgroundColor: 'transparent' })
        const canvasData = this.previewCanvas.toDataURL({ format: 'png' })
        const base64Data = canvasData.split(',')[1]
        this.previewCanvas.set({ backgroundColor: prev })
        return base64Data
    }

    _copySelectedToPartial(selected) {
        this.savePartialBtn.disabled = true
        const partialCanvas = this.previewCanvas
        const selectedBounds = selected.getBoundingRect()
        partialCanvas.setWidth(selectedBounds.width)
        partialCanvas.setHeight(selectedBounds.height)
        this._setChecker(20, partialCanvas)

        this.canvas.discardActiveObject();
        this.canvas.renderAll();

        let jsonData
        if (selected.type === 'activeSelection') {
            const selectedObjs = selected.getObjects()
            if (!selectedObjs || selectedObjs.length === 0) {
                return
            }
            const selectedLeft = selected.left
            const selectedTop = selected.top
            const toSave = selectedObjs.map(obj => {
                const clone = fabric.util.object.clone(obj)
                clone.left = clone.left - selectedLeft
                clone.top = clone.top - selectedTop
                clone.selectable = false
                clone.hoverCursor = 'default'
                partialCanvas.add(clone)
                return clone
            })
            jsonData = new fabric.Group(toSave).toJSON()
        } else {
            const clone = fabric.util.object.clone(selected)
            if (clone.type.includes('text')) {
                partialCanvas.setWidth(selectedBounds.width + 20)
                partialCanvas.setHeight(selectedBounds.height + 20)
                clone.left = 20
                clone.top = 20
            } else {
                clone.left = 0
                clone.top = 0
            }
            clone.selectable = false
            clone.angle = 0
            clone.hoverCursor = 'default'
            partialCanvas.add(clone)
            jsonData = clone.toJSON()
        }

        this._partialJson = jsonData

        partialCanvas.renderAll()
        this.savePartialBtn.disabled = false
    }

    _savePartial(jsonData, pngData, override = false) {
        if (!jsonData || !pngData) {
            console.error('failed to move data to form')
            return
        }
        const name = partialName.value

        const data = {
            name: name,
            jsonData: jsonData,
            pngData: pngData
        }

        const method = override ? "PATCH" : "POST"

        htmx.ajax(method, '/partial', {
            values: data,
            target: '#savePartialResponse',
            swap: 'innerHTML'
        })

        setTimeout(() => {
            htmx.ajax('GET', '/ui/partials', {
                target: '#parials-loader',
                swap: 'innerHTML'
            })
        }, 2000)
    }

    _setChecker(cellSize, canvas) {
        const patternCanvas = document.createElement('canvas')
        patternCanvas.width = cellSize * 2
        patternCanvas.height = cellSize * 2
        const ctx = patternCanvas.getContext('2d')

        ctx.fillStyle = '#e0e0e0'
        ctx.fillRect(0, 0, cellSize, cellSize)
        ctx.fillRect(cellSize, cellSize, cellSize, cellSize)
        const pattern = new fabric.Pattern({
            source: patternCanvas,
            repeat: 'repeat'
        })

        canvas.backgroundColor = pattern
        canvas.renderAll()
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

    _onObjectSelected(btn) {
        const objs = this.canvas.getActiveObject()
        if (!objs) { return }
        const isActiveSelection = objs.type === 'activeSelection'
        const isGroup = objs.type === 'group'
        if (isActiveSelection) {
            this._toggleGroupButton(btn, false)
        }
        if (isGroup) {
            this._toggleGroupButton(btn, true)
        }
    }

    _toggleGroupButton(btn, value) {
        const activeBtnClassName = 'btn-primary'
        const notActiveBtnClassName = 'btn-outline-primary'
        if (!value) {
            btn.classList.remove(activeBtnClassName)
            btn.classList.remove(notActiveBtnClassName)
            btn.classList.add(notActiveBtnClassName)
        } else {
            btn.classList.remove(activeBtnClassName)
            btn.classList.remove(notActiveBtnClassName)
            btn.classList.add(activeBtnClassName)
        }
    }
}