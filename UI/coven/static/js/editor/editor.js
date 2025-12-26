import { Toolbar } from "./toolbar.js";
import { ObjectsOreder } from "./objects_ordering.js";
import { ControlMenu } from "./control_menu.js";

class Editor {
    constructor(canvasId) {
        this.canvas = new fabric.Canvas(canvasId)
        this.elementsIncrementer = this.canvas.size()

        this.setBgColor('#ffffffff')
        this.setSize(600, 800)
        this._initToolbar()
        this._initCanvasEvents()
        this._initKeyEvents()
        this._initObjectOrdering()
        this._initControlMenu()

        const urlParams = new URLSearchParams(window.location.search)
        if (!urlParams || urlParams.size === 0) {
            return
        }
        const edtCardType = urlParams.get('type')
        if (!edtCardType) {
            console.error('parameters provided, type not found')
            return
        }
        this.editingCardType = edtCardType
        const edtCardName = urlParams.get('name')
        if (!edtCardName) {
            console.error('parameters provided, name not found')
            return
        }
        this.editingCardName = edtCardName
        window.addEventListener('DOMContentLoaded', async () => {
            await this._loadCard(edtCardType, edtCardName)
            this._setCardTypeSelected(edtCardType)
            this._setCardName(edtCardName)
        })
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

    async _loadCard(type, name) {
        try {
            const response = await fetch(`/card?cardType=${type}&cardName=${name}`)
            const editingData = await response.json()

            this.canvas.loadFromJSON(editingData, () => {
                this.canvas.renderAll();
            }, (jsonObj, fabricObj) => {
                if (jsonObj.id) {
                    fabricObj.set('id', jsonObj.id)
                }
                return fabricObj;
            })
        } catch (error) {
            console.error(`failed to fetch data for ${type} ${name}`, error)
        }
    }

    _setCardName(name) {
        const saveForm = document.getElementById('save-result-form')
        if (!saveForm){
            console.error('failed to set editing card name, save result form not found')
            return
        }
        const cardNameInput = document.getElementById('cardName')
        if (!cardNameInput){
            console.error('failed to set editing card name cardName input not found')
            return
        }
        cardNameInput.value = name
        const cardNameProxyInput = document.getElementById('name-proxy-control')
        if (!cardNameInput){
            console.error('failed to set editing card name name-proxy-control input not found')
            return
        }
        cardNameProxyInput.value = name
    }

    _setCardTypeSelected(type) {
        const typeSelector = document.getElementById('cards-type-selector')
        if (!typeSelector) {
            console.error('failed to set card type selector div not found')
            return
        }
        const radios = typeSelector.querySelectorAll('input[type="radio"]')
        if (!radios) {
            console.error(`failed to set card type there's no any radio`)
            return
        }
        radios.forEach(x => { x.checked = (x.value === type) })
        const cardTypeInput = document.getElementById('cardType')
        if (!cardTypeInput){
            console.error('failed to set card type cardType not found')
            return
        }
        cardTypeInput.value = type
    }

    _initToolbar() {
        const toolbar = new Toolbar(this.canvas)
        toolbar.bindAddText('toolbar-add-txt')
        toolbar.bindColorPicker('toolbar-txt-color', 'toolbar-txt-opacity')
        toolbar.bindFontSize('toolbar-txt-font-size')
        toolbar.bindTextStroke('toolbar-txt-stroke-color', 'toolbar-txt-stroke-opacity', 'toolbar-txt-stroke-width')
        toolbar.bindFontSelector('font-selector')
        this.toolbar = toolbar
    }

    _initObjectOrdering() {
        this.objectOrdering = new ObjectsOreder(this.canvas)
    }

    _initCanvasEvents() {
        this.canvas.on('object:added', (e) => {
            const newElementId = `${e.target.type}_${this.elementsIncrementer++}`
            const currentId = e.target.get('id')
            if (!currentId) {
                e.target.set('id', newElementId)
            }
        })
    }

    _initControlMenu() {
        this.controlMenu = new ControlMenu(this.canvas)
        this.controlMenu.bindSaveButton('save-result-btn')
    }

    _initKeyEvents() {
        document.addEventListener('keyup', (e) => {
            const active = this.canvas.getActiveObject()
            if (e.key === 'Delete') {
                if (ObjectsOreder.IsEditingObjId) {
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
window.editor = editor