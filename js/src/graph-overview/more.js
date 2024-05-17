import yall from './yall';
import initTopLink from './top-link';

export default function initGraphMore() {
    initTopLink();
    let link = document.getElementById("graph-main-morel");
    if (link) {
        link.addEventListener("click", showGraphMore);
    }
}

function showGraphMore(evt) {
    evt.preventDefault();
    evt.target.parentElement.classList.add("hide");
    let graphMore = document.getElementById("graph-main-more");
    if (graphMore) {
        graphMore.classList.add("show");
        yall();
    }
}
