
function inputAutoGrow() {
    const element = document.querySelector('.textarea');
    const style = getComputedStyle(element);
    const minHeight = parseInt(style.minHeight.slice(0, -2));
    const maxHeight = parseInt(style.maxHeight.slice(0, -2));
    const lineHeight = parseInt(style.lineHeight.slice(0, -2));
    const paddingTop = parseInt(style.paddingTop.slice(0, -2));
    const borderTop = parseInt(style.borderTop.slice(0, -2));

    const calcHeight = function (value) {
        let numberOfLineBreaks = (value.match(/\n/g) || []).length;
        let newHeight = minHeight + numberOfLineBreaks * lineHeight + paddingTop + borderTop;
        return newHeight;
    }

    let textarea = document.querySelector(".resize");
    textarea.addEventListener("keyup", (event) => {
        if (event.keyCode === 13 && !event.shiftKey) {
            let height = minHeight + paddingTop + borderTop;
            textarea.style.height = height + "px";
        } else {
            let calculatedHeight = calcHeight(textarea.value);
            if (calculatedHeight < maxHeight) {
                textarea.style.height = calculatedHeight + "px";
            }
        }

    });
}

function setDefaultHeight() {
    const element = document.querySelector('.textarea');
    const style = getComputedStyle(element);
    const minHeight = parseInt(style.minHeight.slice(0, -2));
    const paddingTop = parseInt(style.paddingTop.slice(0, -2));
    const borderTop = parseInt(style.borderTop.slice(0, -2));

    let height = minHeight + paddingTop + borderTop;
    let textarea = document.querySelector(".resize");
    textarea.style.height = height + "px";
}

function scrollToBottom() {
    window.scrollTo(0, document.querySelector(".body").scrollHeight);
}

export { setDefaultHeight, inputAutoGrow, scrollToBottom };