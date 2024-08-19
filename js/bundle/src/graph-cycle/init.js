import initZoom from './zoom';
import initKeys from './keys';
import {initSwipe, blockSwipe, unblockSwipe} from './swipe';
//import stretchBody from "./body";
//import addHead from "./add-head";
//import initInfoBarScroll from "./info-bar";
//import initPreload from "./preload";

let cachedNext;
let prevUrl, nextUrl, parentUrl;

let blockNav = false;

export default function initCycle(page) { 
    nextUrl    = page.nav.next;
    prevUrl    = page.nav.prev;
    parentUrl  = page.nav.parent;

    initZoom();
    initSwipe();
    initKeys();

    initHideNav();

    window.addEventListener("popstate", (event) => {
        console.log(event.state);
        render(event.state);
    });
};

function cacheNext(url) {
    fetchPage(url).then(json => {
        cachedNext = json
        document.getElementById("preload").innerHTML = json.html;
    });
}

function fetchPage(url) {
    return fetch("/part" + url).then(res => res.json())
}

function block() {
    blockNav = true;
    //blockSwipe();
}

function unblock() {
    //unblockSwipe();
    blockNav = false;
}

function blocked() {
    return blockNav;
}

export function goPrev() {
    if (blocked()) {
        return
    }
    if (prevUrl == "") {
        //goParent();
        return;
    }
    block();
    fetchPage(prevUrl).then(page => {
        renderPush(page);
        unblock();
    });
}

export function goNext() {
    if (blocked()) {
        return
    }
    hideNavTip();
    if (nextUrl == "") {
        //goParent();
        return;
    }
    if (cachedNext != undefined) {
        if (cachedNext.url == nextUrl) {
            renderPush(cachedNext);
            return
        }
    }
    block();
    fetchPage(nextUrl).then(page => {
        renderPush(page);
        unblock();
    });
}

export function goParent() {
    location = parentUrl;
}

function renderPush(page) {
    render(page)
    history.pushState(page, page.title, page.url); 
    cacheNext(page.next);
}

function render(page) {
    document.title = page.title;
    document.getElementById("part").innerHTML = page.html;
    nextUrl = page.next;
    prevUrl = page.prev;
    parentUrl = page.parent;

    const l = document.getElementsByClassName("lnavl");
    l[0].href = page.langs.de;
    l[1].href = page.langs.en;

    initZoom();
    hideNavTip();
}

function hideNavTip() {
    let el = document.getElementById("nav-tip");
    if (el !== undefined) {
        el.style.display = "none";
        saveHideNavTip();
    }
}

function initHideNav() {
    if (window.localStorage.getItem("hideNavTip") !== null) {
        hideNavTip();
    }
}

function saveHideNavTip() {
    window.localStorage.setItem("hideNavTip", "true");
}

const handleKeyNav = evt => {
    switch (evt.keyCode) {
    case 37:
        evt.preventDefault();
        goPrev();
        break;
    case 39:
        evt.preventDefault();
        goNext();
        break;
    case 27:
        evt.preventDefault();
        goParent();
        return;
    }
}
