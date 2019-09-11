
export default function initText() {
    let links = document.getElementsByClassName("notel");
    for (let i = 0; i < links.length; i++) {
        links[i].addEventListener("click", showNote);
    }
}

function showNote(evt) {
    evt.preventDefault();
    let article = findParent(evt.target, "ARTICLE");
    if (article) {
        let notes = article.getElementsByClassName("footnotes");
        if (notes.length > 0) {
            notes[0].classList.toggle("show");
        }
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
