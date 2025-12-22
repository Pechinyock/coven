import { Toolbar } from "./toolbar.js";
import { ObjectsOreder } from "./objects_ordering.js";

class Editor {
    constructor(canvasId) {
        this.canvas = new fabric.Canvas(canvasId)
        this.elementsIncrementer = this.canvas.size()

        this.setBgColor('#ffffffff')
        this.setSize(600, 800)
        this.initToolbar()
        this.initCanvasEvents()
        this.initKeyEvents()
        this.initObjectOrdering()
    }

    setBgColor(color) {
        this.canvas.set({
            backgroundColor: color
        });
        this.canvas.renderAll();
    }

    setSize(width, height) {
        this.canvas.setWidth(width)
        this.canvas.setHeight(height)
        this.canvas.renderAll();
    }

    initToolbar() {
        const toolbar = new Toolbar(this.canvas)
        toolbar.bindAddText('toolbar-add-txt')
        toolbar.bindColorPicker('toolbar-txt-color', 'toolbar-txt-opacity')
        toolbar.bindFontSize('toolbar-txt-font-size')
        toolbar.bindTextStroke('toolbar-txt-stroke-color', 'toolbar-txt-stroke-opacity', 'toolbar-txt-stroke-width')
        this.toolbar = toolbar
    }

    initObjectOrdering(ordering) {
        this.objectOrdering = new ObjectsOreder(this.canvas)
    }

    initCanvasEvents() {
        this.canvas.on('object:added', (e) => {
            const newElementId = `${e.target.type}_${this.elementsIncrementer++}`
            e.target.set('id', newElementId)
        })
    }

    initKeyEvents() {
        document.addEventListener('keyup', (e) => {
            const active = this.canvas.getActiveObject()
            if (e.key === 'Delete') {
                if (ObjectsOreder.IsEditingObjId){
                    return
                }
                if (active) {
                    if (active.type === 'activeSelection') {
                        const objs = active.getObjects()
                        if (objs) {
                            objs.forEach(obj => { this.canvas.remove(obj) })
                        }
                    } else {
                        if (active.isEditing) {
                            return
                        }
                        this.canvas.remove(active)
                        this.canvas.renderAll()
                    }
                }
            }
        })
    }
}

const editor = new Editor('card-canvas')