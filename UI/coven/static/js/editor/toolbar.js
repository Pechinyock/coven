export class Toolbar {
    constructor(canvas) {
        if (!canvas) {
            console.error('failed to initialze canvas toolbar: provided canvas is null')
            return
        }
        this.canvas = canvas
    }

    bindColorPicker(colorPickerId, trasparancyId) {
        const picker = document.getElementById(colorPickerId)
        if (!picker) {
            console.error('failed to bind color picker')
            return
        }
        const trasparancy = document.getElementById(trasparancyId)
        if (!trasparancy) {
            console.error('failed to bind color transparancy')
            return
        }

        this.colorPicker = picker
        this.textTransperancy = trasparancy

        picker.disabled = false
        trasparancy.disabled = false
    }

    bindFontSize(id) {
        const fontSizeInput = document.getElementById(id)
        if (!fontSizeInput) {
            console.error('failed to bind font size element')
            return
        }
        fontSizeInput.addEventListener('input', (e) => {
            const active = this.canvas.getActiveObject();
            if (!active) return;

            const size = parseInt(e.target.value) || 20;

            if (active.type === 'activeSelection') {
                const objs = active.getObjects();
                objs.forEach(obj => {
                    if (obj.type.includes('text')) { 
                        obj.set('fontSize', size);
                    }
                });
            } else if (active.type.includes('text')) {
                active.set('fontSize', size);
            }

            this.canvas.renderAll();
        });
        fontSizeInput.disabled = false
        this.fontSizeInput = fontSizeInput
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

    addText(text, fabricObj) {
        const opacity = parseInt(this.textTransperancy.value) / 100
        const txt = new fabric.Textbox(text, {
            left: 50,
            top: 50,
            width: 200,
            fontSize: this.fontSizeInput.value,
            fill: this.colorPicker.value,
            opacity: opacity
        });
        this.canvas.add(txt)
    }
}