
let subnavStart = 0;
let subnavFixEl;
let isShown = false;

export default function initSubnavFix() {
    subnavStart = getSubnavPos();
    subnavFixEl = document.getElementById("subnav-fixed");
    window.addEventListener("scroll", handleScroll);

}

const handleScroll = evt => {
    if (subnavStart <= 0) {
        return;
    }
    if (window.scrollY >= subnavStart) {
        showSubnavFix();
    } else {
        hideSubnavFix();
    }
}

const getSubnavPos = () => {
    const subnav = document.getElementById("subnav");
    const l = subnav.getElementsByClassName("subnav");
    return l[l.length - 1].offsetTop - 6 ; // - 4
}

const showSubnavFix = () => {
    if (!isShown) {
        document.getElementById("subnav-fixed").classList.remove("hide");
        isShown = true;
    }
}

const hideSubnavFix = () => {
    if (isShown) {
        document.getElementById("subnav-fixed").classList.add("hide");
        isShown = false;
    }
}
