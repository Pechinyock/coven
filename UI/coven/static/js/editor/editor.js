import { Toolbar } from "./toolbar.js";
import { ObjectsView } from "./objects_view.js";

const canvas = new fabric.Canvas('card-canvas', {
    backgroundColor: '#ffffffff'
})

canvas.setHeight(800);
canvas.setWidth(1000);

const toolbar = new Toolbar(canvas)
toolbar.bindAddText('toolbar-add-txt')
toolbar.bindColorPicker('toolbar-txt-color', 'toolbar-txt-transparancy')
toolbar.bindFontSize('toolbar-txt-font-size')

const objsView = new ObjectsView('objects-view', canvas)

document.addEventListener('keyup', (e) => {
    const active = canvas.getActiveObject()
    if (e.key === 'Delete') {
        if (active) {
            if (active.type === 'activeSelection') {
                const objs = active.getObjects()
                if (objs) {
                    objs.forEach(obj => {
                        canvas.remove(obj)
                    })
                }

            } else {
                e.preventDefault()
                canvas.remove(active)
                canvas.renderAll()
            }
        }
    }
})