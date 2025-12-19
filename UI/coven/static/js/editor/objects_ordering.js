export class ObjectsOreder {
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
            this.appendNewElement(id)
        })

        canvas.on('object:removed', (e) => {
            const id = e.target.get('id')
            if (!id) {
                console.error('trying to remove element with no id')
                return
            }
            this.removeElement(id)
        })

        this.canvas = canvas
        this.elementsContainer = objectsOrdering
    }

    appendNewElement(elementId) {
        const delIco = document.createElement('i')
        delIco.classList.add('bi', 'bi-trash3')
        delIco.style = 'font-size: 12px;'

        const arrowUp = document.createElement('i')
        arrowUp.classList.add('bi', 'bi-arrow-bar-up')
        arrowUp.style = 'font-size: 12px;'

        const arrowDown = document.createElement('i')
        arrowDown.classList.add('bi', 'bi-arrow-bar-down')
        arrowDown.style = 'font-size: 12px;'

        const delBtn = document.createElement('button')
        delBtn.classList.add('btn', 'btn-danger', 'btn-sm')
        delBtn.addEventListener('click', () => {
            const obj = this.canvas.getObjects().find(o => o.id === elementId);
            if (obj) {
                this.canvas.remove(obj)
            }
        })
        delBtn.appendChild(delIco)

        const moveUpBtn = document.createElement('button')
        moveUpBtn.classList.add('btn', 'btn-primary', 'btn-sm', 'me-1')
        moveUpBtn.appendChild(arrowUp)
        moveUpBtn.addEventListener('click', () => {
            /* [TODO] implement */
            const all = this.canvas.getObjects()
            const obj = all.find(x => x.id === elementId)
            const currentIndx = all.indexOf(obj)
            if (currentIndx === -1){
                console.error(`failed to find object ${elementId}`)
                return
            }
            if (currentIndx >= all.lenght){
                console.log(`can't move it up it is already on the top ${currentIndx}`);
                return
            }
            const targetIndex = currentIndx + 1
            const viewElements = Array.from(this.elementsContainer.children)
            if (targetIndex > viewElements.length){
                console.error('failed to move object the target index outside of elements')
                return
            }

            this.canvas.moveTo(obj, targetIndex)
        })

        const moveDownBtn = document.createElement('button')
        moveDownBtn.classList.add('btn', 'btn-primary', 'btn-sm', 'me-1')
        moveDownBtn.appendChild(arrowDown)

        const controlsWrapper = document.createElement('div')
        controlsWrapper.classList.add('d-flex', 'justify-content-end')
        controlsWrapper.appendChild(moveUpBtn)
        controlsWrapper.appendChild(moveDownBtn)
        controlsWrapper.appendChild(delBtn)

        const labelDiv = document.createElement('div')
        labelDiv.textContent = elementId
        labelDiv.addEventListener('click', () => {
            const elems = this.canvas.getObjects()
            if (elems) {
                const selected = elems.find(obj => obj.id === elementId)
                if (selected) {
                    this.canvas.setActiveObject(selected)
                    this.canvas.renderAll()
                }
            }
        })

        const elementDiv = document.createElement('div')
        elementDiv.classList.add('d-flex', 'justify-content-between', 'align-items-center', 'mt-1', 'object-view-element')
        elementDiv.id = elementId
        elementDiv.appendChild(labelDiv)
        elementDiv.appendChild(controlsWrapper)

        this.elementsContainer.prepend(elementDiv)
    }

    removeElement(id) {
        const element = document.getElementById(id)
        if (!element) { return }
        element.remove()
    }
}
