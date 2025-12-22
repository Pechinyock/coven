export class ObjectsOreder {
    static IsEditingObjId = false

    constructor(canvas) {
        const objectsOrdering = document.getElementById('objects-ordering')

        if (!objectsOrdering) {
            console.error('failed to init layer objects-ordering')
            return
        }

        canvas.on('object:added', (e) => {
            const id = e.target.get('id')
            if (!id) {
                console.error('trying to add element with no id')
                return
            }
            this._appendNewElement(id)
        })

        canvas.on('object:removed', (e) => {
            const id = e.target.get('id')
            if (!id) {
                console.error('trying to remove element with no id')
                return
            }
            const viewObj = document.getElementById(id)
            if (!viewObj) {
                console.error(`failed to to sync deleted element ${id}`)
                return
            }
            viewObj.remove()
        })

        this.canvas = canvas
        this.elementsContainer = objectsOrdering
        this.randomNameIdx = 0
    }

    _getIco(icoClassName) {
        if (!icoClassName) {
            console.error('failed to get ico class name is not specified')
            return
        }
        const icoFontSize = 'font-size: 12px;'
        const result = document.createElement('i')
        result.classList.add('bi', icoClassName)
        result.style = icoFontSize
        return result
    }

    _selectObject(e) {
        if (!e) {
            console.error('failed to handle select event')
            return
        }
        const targetElem = e.currentTarget
        if (!targetElem) {
            console.error('failed to handle select event')
            return
        }

        const elementRootDiv = targetElem.parentElement
        if (!elementRootDiv) {
            console.error('failed to get div root element on delete clicked button')
            return
        }

        const elementId = elementRootDiv.id
        if (!elementId) {
            console.error('failed to get target element id')
            return
        }
        const elems = this.canvas.getObjects()
        if (!elems || elems.length === 0) {
            console.error('trying to select item on empty canvas')
            return
        }
        const selected = elems.find(obj => obj.id === elementId)
        if (!selected) {
            console.error(`failed to select object with id: ${elementId}`)
            return
        }

        this.canvas.setActiveObject(selected)
        this.canvas.renderAll()
    }

    _removeElement(e) {
        if (!e) {
            console.error('delete button is not specified')
            return
        }
        const delBtn = e.currentTarget
        if (!delBtn) {
            console.error('failed to get current target')
        }
        const elementRootDiv = delBtn.parentElement.parentElement
        if (!elementRootDiv) {
            console.error('failed to get div root element on delete clicked button')
            return
        }
        const targetId = elementRootDiv.id
        if (!targetId) {
            console.error('failed to get target element id')
            return
        }

        const obj = this.canvas.getObjects().find(o => o.id === targetId)
        if (!obj) {
            return console.error(`failed to get object with id: ${targetId}`)
        }
        this.canvas.remove(obj)
        elementRootDiv.remove()
    }

    _appendNewElement(elementId) {
        const delIco = this._getIco('bi-trash3')
        const arrowUp = this._getIco('bi-arrow-bar-up')
        const arrowDown = this._getIco('bi-arrow-bar-down')

        const delBtn = document.createElement('button')
        delBtn.classList.add('btn', 'btn-danger', 'btn-sm')
        delBtn.addEventListener('click', (e) => this._removeElement(e))
        delBtn.appendChild(delIco)

        const moveUpBtn = document.createElement('button')
        moveUpBtn.classList.add('btn', 'btn-primary', 'btn-sm', 'me-1')
        moveUpBtn.appendChild(arrowUp)
        moveUpBtn.addEventListener('click', (e) => this._moveElementHandler(e, 'up'))

        const moveDownBtn = document.createElement('button')
        moveDownBtn.classList.add('btn', 'btn-primary', 'btn-sm', 'me-1')
        moveDownBtn.appendChild(arrowDown)
        moveDownBtn.addEventListener('click', (e) => this._moveElementHandler(e, 'down'))

        const controlsWrapper = document.createElement('div')
        controlsWrapper.classList.add('d-flex', 'justify-content-end')
        controlsWrapper.appendChild(moveUpBtn)
        controlsWrapper.appendChild(moveDownBtn)
        controlsWrapper.appendChild(delBtn)

        const labelDiv = document.createElement('div')
        labelDiv.textContent = elementId
        labelDiv.id = `lb-${elementId}`
        labelDiv.classList.add('object-label')
        labelDiv.addEventListener('click', (e) => this._selectObject(e))
        labelDiv.addEventListener('dblclick', (e) => this._editObjectId(e))
        labelDiv.addEventListener('blur', (e) => this._saveAfterEdit(e))

        const elementDiv = document.createElement('div')
        elementDiv.classList.add('d-flex', 'justify-content-between', 'align-items-center', 'mt-1', 'object-view-element')
        elementDiv.id = elementId
        elementDiv.appendChild(labelDiv)
        elementDiv.appendChild(controlsWrapper)

        this.elementsContainer.prepend(elementDiv)
    }

    _editObjectId(e) {
        if (!e) {
            console.error('failed to handle edit object id')
            return
        }
        const targetDiv = e.currentTarget
        if (!targetDiv) {
            console.error('failed to handle edit object id')
            return
        }
        targetDiv.contentEditable = true
        targetDiv.focus()
        targetDiv.classList.remove('object-label')
        targetDiv.classList.add('form-control', 'form-control-sm')
        ObjectsOreder.IsEditingObjId = true
    }

    _saveAfterEdit(e) {
        if (!e) {
            console.error('failed to handle edit object id')
            return
        }
        const targetDiv = e.currentTarget
        if (!targetDiv) {
            console.error('failed to handle edit object id')
            return
        }
        const targetElement = targetDiv.parentElement
        if (!targetElement) {
            console.error('failed to get parent of label')
            return
        }
        let newValue = targetDiv.innerText.replace(/\n/g, '').trim()
        if (!newValue || newValue === '') {
            newValue = `rand_name_${this.randomNameIdx}`
            this.randomNameIdx++
        }

        const allObjs = this.canvas.getObjects()
        const target = allObjs.find(x => x.id === targetElement.id)
        target.set('id', newValue)
        targetElement.id = newValue
        targetDiv.id = `lb-${newValue}`
        if (targetDiv.innerText.replace(/\n/g, '').trim() === '') {
            targetDiv.innerText = newValue
        } else {
            targetDiv.innerText = targetDiv.innerText
                .replace(/\n/g, '')
                .trim()
                .replace(/[^\w\s]/g, '')
                .replace(/\s+/g, '_')
        }
        targetDiv.contentEditable = false
        targetDiv.classList.remove('form-control', 'form-control-sm')
        targetDiv.classList.add('object-label')
        ObjectsOreder.IsEditingObjId = false
    }

    _moveElementHandler(e, direction) {
        if (direction !== 'up' && direction !== 'down') {
            console.error('unknown direction')
            return
        }
        const theBtn = e.currentTarget
        const elementRootDiv = theBtn.parentElement.parentElement
        if (!elementRootDiv) {
            console.error('failed to get control wrapper')
            return
        }
        const elem = elementRootDiv.id
        if (!elem) {
            console.error('failed to get element id')
            return
        }
        const all = this.canvas.getObjects()
        const obj = all.find(x => x.id === elem)
        if (!obj) {
            console.error(`failed to find element ${elem}`)
            return
        }
        const currentIndx = all.indexOf(obj)

        if (currentIndx === -1) {
            console.error(`failed to find object ${elementId}`)
            return
        }
        const isMoveUp = direction === 'up'
        if (isMoveUp && currentIndx >= all.length - 1) {
            console.log(`can't move it up it is already on the top ${currentIndx}`);
            return
        }
        else if (!isMoveUp && currentIndx === 0) {
            console.log(`can't move it down it is already on the buttom ${currentIndx}`);
            return
        }

        let tmp = currentIndx
        const targetIndex = isMoveUp
            ? ++tmp
            : --tmp

        const viewElements = Array.from(this.elementsContainer.children).reverse()
        const movingItem = viewElements[currentIndx]
        if (!movingItem) {
            console.error('something wierd happened...')
            return
        }
        const targetPlaceItem = viewElements[targetIndex]
        if (!targetPlaceItem) {
            console.error('something wierd happened again...')
            return
        }
        this.canvas.moveTo(obj, targetIndex)
        this._swap(movingItem, targetPlaceItem)
    }

    _swap(elementOne, elementTwo) {
        if (!elementOne) {
            console.error('failed swap html elements: element one is now specified')
            return
        }
        if (!elementTwo) {
            console.error('failed swap html elements: element one is now specified')
            return
        }

        const labelOneId = `lb-${elementOne.id}`
        const labelOne = document.getElementById(labelOneId)
        if (!labelOne) {
            console.error(`failed to swap elements coulnd't get label ${labelOneId}`)
            return
        }

        const labelTwoId = `lb-${elementTwo.id}`
        const labelTwo = document.getElementById(labelTwoId)
        if (!labelTwo) {
            console.error(`failed to swap elements coulnd't get label ${labelOneId}`)
            return
        }

        const labelTempId = labelTwo.id
        labelTwo.id = labelOne.id
        labelOne.id = labelTempId

        const tempLabelText = labelTwo.innerText
        labelTwo.innerText = labelOne.innerText
        labelOne.innerText = tempLabelText

        const temp = elementTwo.id
        elementTwo.id = elementOne.id
        elementOne.id = temp
    }
}
