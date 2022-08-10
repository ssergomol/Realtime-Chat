function inputAutoGrow() {
    // Dealing with Textarea Height
    function calcHeight(value) {
        let numberOfLineBreaks = (value.match(/\n/g) || []).length;

        // min-height + lines x line-height + padding + border
        let newHeight = 20 + numberOfLineBreaks * 20 + 12 + 2;
        return newHeight;
    }

    let textarea = document.querySelector(".resize");
    textarea.addEventListener("keyup", () => {

        // max-height 
        const maxHeight = 200;
        let calculatedHeight = calcHeight(textarea.value);
        if (calculatedHeight < maxHeight) {
            textarea.style.height = calculatedHeight + "px";
        } 
       
    });
}

export default inputAutoGrow;