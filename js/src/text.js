
export default function initText() {
    let links = document.getElementsByClassName("ref");
    for (let i = 0; i < links.length; i++) {
        links[i].addEventListener("click", showNote);
    }
}

function showNote(evt) {
    evt.preventDefault();
    let p = findParent(evt.target, "P");
    if (p) {
        let notes = p.getElementsByClassName("footnotes");
        if (notes.length > 0) {
            notes[0].classList.toggle("show");
        }
        console.log(notes)
    }
}

function findParent(el, tag) {
    let parentEl = el.parentElement;
    if (!parentEl) {
        return;
    }
    if (parentEl.tagName == tag) {
        return parentEl;
    }
    return findParent(parentEl, tag)
}
