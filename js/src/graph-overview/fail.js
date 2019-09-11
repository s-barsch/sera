

const observeMonths = () => {
    let obs = [];

    const l = document.getElementsByClassName("month");
    for (let i = l.length - 1, el; el = l[i]; i--) {
        obs[i] = new IntersectionObserver(
            intersectionCallback,
            { threshold: [0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0] }
        );
        obs[i].observe(el);
    };
    console.log("do this first");
}

const getLastSubnav = () => {
      const l = document.getElementsByClassName("subnav")
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

let active = {
    id: "",
    ratio: 0.0
}

const intersectionCallback = entries => {
    entries.forEach(function(el) {
        console.log(el.target.id);
        console.log(el.isIntersecting);
        console.log(el.intersectionRatio);
        if (active.id == el.target.id && active.ratio > el.intersectionRatio) {
            active.id = "";
            active.ratio = el.intersectionRatio;
            console.log("decreasing " + el.target.id)
            return;
        }
        if (el.intersectionRatio > active.ratio) {
            active.ratio = el.intersectionRatio;
            if (active.id != el.target.id) {
                active.id = el.target.id;
                highlightSubnav(el.target.id);
                console.log("active " + el.target.id);
            }
        }
        /*
        console.log(entry.target.id);
        console.log(entry.isIntersecting);
        console.log(entry.intersectionRatio);
        if (entry.isIntersecting) {
            highlightSubnav(entry.target.id);
        }
        */
    });
}

