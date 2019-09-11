import {goNext, goPrev} from "./init";
//import Hammer from "hammerjs";

export function initSwipe() {
    /*
    let el = document.getElementById("part");
    let mc = new Hammer.Manager(el);

    mc.add(new Hammer.Swipe({velocity: 0.1}));

    mc.on("swipeleft", (event) => {
        console.log(event);
        console.log("swipe");
        goNext()
    });
    mc.on("swiperight", (event) => {goPrev()});

    /*
    let stage = document.getElementById("part");
    let mc = new Hammer.Manager(stage);
    let swipe = new Hammer.Swipe;
    mc.add(swipe);

    mc.on("swipeleft", (event) => {
        console.log("swipe");
        goNext()});
    mc.on("swiperight", (event) => {goPrev()});
}
    */

    document.addEventListener("touchstart", handleSwipeStart);
    document.addEventListener("touchmove", handleSwipeMove);
    document.addEventListener("touchend", handleSwipeEnd);
    document.addEventListener("touchcancel", handleSwipeEnd);
}

let threshold = 80;

let xStart;
let yStart;
let touchActive = false;
let block = false;

let timeout;

const handleSwipeStart = evt => {
    block = false;
    xStart = evt.touches[0].clientX;                                 
    yStart = evt.touches[0].clientY;
}

let swipeBar = document.getElementById("swipe-bar");
let swipeBarBg = document.getElementById("swipe-bar-bg");

const handleSwipeMove = evt => {
    if (block) {
        resetSwipe();
        return;
    }
    let xNow = evt.touches[0].clientX;
    let yNow = evt.touches[0].clientY;

    let yDistance = yStart - yNow;
    let xDistance = xStart - xNow;

    //console.log(xNow);

    if (Math.abs(xDistance) < 15) {
        return
    }

    //setBar(xDistance);

    if (Math.abs(xDistance) < threshold) {
        return;
    }
    clearTimeout(timeout);
    xStart = xNow;
    yStart = yNow;
    if (xDistance > 0) {
        goNext();
    } else {
        goPrev();
    }
    block = true;
    resetBar();
}

export function blockSwipe() {
    block = true;
}

export function unblockSwipe() {
    block = false;
}

const setBar = dist => {
    let perc;
    perc = (Math.abs(dist) / threshold) * 100;
    if (dist < 0) {
        swipeBarBg.classList.remove("neg");
    } else {
        swipeBarBg.classList.add("neg");
        perc = 100 - perc;
    }
    swipeBar.style.width = perc + "%";
}

const resetBar = () => {
    /*
    swipeBar.style.width = "0";
    swipeBarBg.classList.remove("neg");
    */
}

const handleSwipeEnd = evt => {
    block = false;
    resetBar();
    let xStart = evt.touches[0].clientX;
    let yStart = evt.touches[0].clientY;
    console.log("swipe end");
    /*
    block = false;
    */
}
