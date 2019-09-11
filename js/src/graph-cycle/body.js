/**
 * This sets the #body element to the viewport height on mobile devices. This is necessary because “100vh” doesn’t work on mobile devices as on desktop.
 **/


export default function stretchBody() {
    if (window.innerWidth >= 1024) {
        return;
    }
    const b = document.getElementById("body");
    const l = b.getElementsByTagName("img");

    if (l.length > 0 && l[0].height == 0) {
        setTimeout(stretchBody, 50);
        // l[0].addEventListener("load", stretchBody);
        return;
    }

    const viewportHeight = window.innerHeight;

    const titleHeight = document.getElementById("title").offsetHeight;
    const bodyHeight = b.offsetHeight;
    const prevNextHeight = document.getElementById("prev-next").offsetHeight;

    const bodyTooSmall = (titleHeight + bodyHeight + prevNextHeight) < viewportHeight;

    if (bodyTooSmall) {
        b.style.height = (viewportHeight - titleHeight) + "px";
    }
};
