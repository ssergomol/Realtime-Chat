function inputAutoGrow() {
    const element = document.querySelector('.textarea');
    const style = getComputedStyle(element);
    let minHeight = parseInt(style.minHeight.slice(0, -2));
    let maxHeight = parseInt(style.maxHeight.slice(0, -2));
    let lineHeight = parseInt(style.lineHeight.slice(0, -2));
    let paddingTop = parseInt(style.paddingTop.slice(0, -2));
    let borderTop = parseInt(style.borderTop.slice(0, -2));

    function calcHeight(value) {
        let numberOfLineBreaks = (value.match(/\n/g) || []).length;
        let newHeight = minHeight + numberOfLineBreaks * lineHeight + paddingTop + borderTop;
        return newHeight;
    }

    let textarea = document.querySelector(".resize");
    textarea.addEventListener("keyup", () => {

        let calculatedHeight = calcHeight(textarea.value);
        if (calculatedHeight < maxHeight) {
            textarea.style.height = calculatedHeight + "px";
        } 
       
    });
}

export default inputAutoGrow;