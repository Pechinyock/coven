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

export function FromRGBAToHex(rgba) {
    const match = rgba.match(/rgba?\((\d+),\s*(\d+),\s*(\d+)(?:,\s*(\d*\.?\d+))?\)/);

    if (match) {
        const r = match[1];
        const g = match[2];
        const b = match[3];

        return "#" + ((1 << 24) + (+r << 16) + (+g << 8) + +b).toString(16).slice(1);
    }
}

export function ExtractAlpha(rgba) {
    const match = rgba.match(/rgba?\((\d+),\s*(\d+),\s*(\d+)(?:,\s*(\d*\.?\d+))?\)/);
    if (match) {
        const alpha = match[4] || '1'
        return alpha
    }
}