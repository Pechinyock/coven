export class ObjectsView {
    static _cardObjects = new Map()

    constructor(containerId, canvas) {
        if (!canvas) {
            console.error('failed to initialize ObjectsView provided canvas is null')
            return
        }

        const contaier = document.getElementById(containerId)
        if (!contaier) {
            console.error(`failed to initialize ObjectsView container with id ${containerId} is not found`)
            return
        }

        this.container = contaier
        this.canvas = canvas

        this.elementsIncrementer = canvas.size()

        canvas.on('object:added', (e) => {
            const newElementId = `${e.target.type}_${this.elementsIncrementer++}`
            e.target.set('id', newElementId)
            ObjectsView._cardObjects.set(newElementId, e)

            const delIco = document.createElement('i')
            delIco.classList.add('bi', 'bi-trash3')
            delIco.style = "font-size: 12px;"

            const delBtn = document.createElement('button')
            delBtn.classList.add('btn', 'btn-danger', 'btn-sm')
            const newElement = document.createElement('div')
            newElement.classList.add('d-flex', 'justify-content-between', 'align-items-center', 'mt-1')
            delBtn.appendChild(delIco)
            delBtn.addEventListener('click', () => {
                const delTargetId = newElementId
                const obj = canvas.getObjects().find(o => o.id === delTargetId);
                if (obj) {
                    canvas.remove(obj)
                }
            })

            const nameDiv = document.createElement('div')
            nameDiv.textContent = newElementId
            nameDiv.style = "cursor: pointer;"
            nameDiv.addEventListener('click', () => {
                const elems = canvas.getObjects()
                if (elems) {
                    const selected = elems.find(obj => obj.id === newElementId)
                    if (selected) {
                        canvas.setActiveObject(selected)
                        canvas.renderAll()
                    }
                }
            })
            newElement.id = newElementId
            contaier.appendChild(newElement)
            newElement.appendChild(delBtn)
            newElement.appendChild(nameDiv)
        });

        canvas.on('object:removed', (e) => {
            const targetId = e.target.id
            const viewElement = document.getElementById(targetId)
            if (viewElement) {
                viewElement.remove()
            }
            ObjectsView._cardObjects.delete(targetId);
        });
    }
}