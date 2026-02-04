export class Parials {
    constructor(canvas) {
        this.canvas = canvas
    }

    bindAddPartial(btn) {
        if (!btn) {
            console.error('failed to bind partial')
            return
        }
        btn.addEventListener('click', async () => {
            const name = this.partialTargetName
            if (!name || name === '') {
                return
            }
            const response = await fetch(`/partial?name=${encodeURIComponent(name)}`)
            const jsonData = await response.json()
            fabric.util.enlivenObjects([jsonData], (objects) => {
                const fabricObject = objects[0]
                fabricObject.center()
                this.canvas.add(fabricObject)
                this.canvas.setActiveObject(fabricObject)
                this.canvas.viewportCenterObject(fabricObject)
                this.canvas.renderAll()
            })
        })
        btn.disabled = false
    }

    onPartialSelected(element) {
        const fullName = element.dataset.fileName
        const name = fullName.lastIndexOf('.') > 0
            ? fullName.slice(0, fullName.lastIndexOf('.'))
            : fullName;
        this.partialTargetName = name

        const selectedBorderStyle = '2px solid rgb(3, 252, 78)'
        if (!this.selectedPartial) {
            element.style.border = selectedBorderStyle
            this.selectedPartial = element
            return
        }

        const currentTargetName = element.dataset.fileName
        const selecetedPartialName = this.selectedPartial.dataset.fileName
        if (currentTargetName === selecetedPartialName) {
            return
        }
        element.style.border = selectedBorderStyle
        this.selectedPartial.style.border = '2px solid transparent'
        this.selectedPartial = element
    }
}