import {goNext, goPrev, goParent} from './init';

const handleKeyNav = evt => {
    if (evt.keyCode == 37) {
        evt.preventDefault();
        goPrev();
        return;
    }
    if (evt.keyCode == 39) {
        evt.preventDefault();
        goNext();
        return;
    }
    if (evt.keyCode == 27) {
        evt.preventDefault();
        goParent();
        return;
    }
}

export default function initKeys(pageNav) {
    document.body.addEventListener("keydown", handleKeyNav);
}
