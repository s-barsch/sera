let months = [];
let current = "";

export default function obeserveMonths() {
    months = getMonths();
    window.addEventListener("scroll", handleScroll);
}

let stopInterval;

const handleScroll = evt => {
    maintainInterval();
    const month = scrollOverMonth();
    if (current != month) {
        current = month;
        highlightSubnav(month);
    }
}

const isScrollAtBottom = () => {
    return window.innerHeight + window.scrollY >= document.body.offsetHeight;
}

const scrollOverMonth = () => {
    if (isScrollAtBottom()) {
        return months[months.length - 1].id;
    }
    for (let i = 0, m; m = months[i]; i++) { 
        if (window.scrollY < m.pos) {
            if (i == 0) {
                return months[i].id;
            } else {
                return months[i-1].id;
            }
        }
        if (i == months.length - 1) {
            return months[i].id;
        }
    }
};

const getMonths = () => {
    let months = [];
    const l = document.getElementsByClassName("month");
    for (let i = l.length - 1, f; f = l[i]; i--) {
        months.push({
            id:  f.id,
            pos: f.offsetTop
        })
    };
    return months.reverse();
} 

const getLastSubnav = () => {
    const subnav = document.getElementById("subnav-fixed");
    const l = subnav.getElementsByClassName("subnav");
    return l[l.length - 1];
}

const highlightSubnav = id => {
    const l = getLastSubnav().children;
    for (let i = l.length - 1, a; a = l[i]; i--) {
        if (a.dataset.id == id) {
            a.classList.add("active");
        } else {
            a.classList.remove("active");
        }
    }
}

/*
 * The positions of the months is updated
 * every 750ms as long as the user scrolls.
 * Because when the user scrolls images are
 * loaded and the positions change.
 */

let isLoading = true;
let intervalRunning = false;
let interval;

const maintainInterval = () => {
    if (isLoading && !intervalRunning) {
        installInterval();
    };
    clearTimeout(stopInterval);
    stopInterval = setTimeout(function() {
        deinstallInterval(); 
    }, 300);
}

const installInterval = () => {
    intervalRunning = true;
    interval = setInterval(function() {
        // console.log("tick");
        isLoading = checkLoading();
        if (!isLoading) {
            deinstallInterval();
        } else {
            months = getMonths();
        }
    }, 750);
}

const deinstallInterval = () => {
    clearInterval(interval);
    intervalRunning = false;
}

const checkLoading = () => {
    const l = document.getElementsByClassName("lazy");
    if (l.length > 0) {
        return true;
    }
    return false;
}

