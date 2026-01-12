export let IsEditingRange = false

export function BindDivToRange(divLabel, range, rangeValueType) {
    divLabel.addEventListener('dblclick', (e) => {
        IsEditingRange = true
        const targetDiv = e.currentTarget
        targetDiv.contentEditable = true
        targetDiv.focus()
        targetDiv.classList.add('form-control', 'form-control-sm')
    })

    divLabel.addEventListener('blur', (e) => {
        const targetDiv = e.currentTarget
        const value = rangeValueType === 'int' 
            ? parseInt(targetDiv.textContent)
            : parseFloat(targetDiv.textContent)

        if (!isNaN(value) && value >= range.min && value <= range.max) {
            range.value = value
            range.dispatchEvent(new Event('input'))
        } else {
            targetDiv.textContent = range.value
        }

        targetDiv.contentEditable = false
        targetDiv.classList.remove('form-control', 'form-control-sm')
        IsEditingRange = false
    })
}