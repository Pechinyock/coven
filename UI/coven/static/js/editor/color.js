export function FromHexToRGBA(hexCode, opacity = 1) {
    if (!IsHexFormat(hexCode)) {
        console.warn(`trying to convert hex to rgba with hex wrong format ${hexCode}`)
        return
    }
    const hex = hexCode.replace('#', '')
    if (hex.lenght === 3) {
        hex = `${hex[0]}${hex[0]}${hex[1]}${hex[1]}${hex[2]}${hex[2]}`
    }
    const r = parseInt(hex.substring(0, 2), 16);
    const g = parseInt(hex.substring(2, 4), 16);
    const b = parseInt(hex.substring(4, 6), 16);

    return `rgba(${r}, ${g}, ${b}, ${opacity})`;
}

export function IsHexFormat(input) { return input.startsWith('#') }

export function GetOpacityFromRGBA(rgbaString) {
    const match = rgbaString.match(/rgba?\([^)]*,\s*([^)]+)\)/);
    return match ? parseFloat(match[1]) : 1
}

export function ChangeRGBAOpacity(rgbaString, newOpacity) {
    return rgbaString.replace(
        /rgba?\((\d+,\s*\d+,\s*\d+),?\s*[\d\.]+\)/,
        `rgba($1, ${newOpacity})`
    );
}