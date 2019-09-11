export default function initTopLink() {
    let obs = new IntersectionObserver(
        intersectionCallback,
        { threshold: 0.1 }
    );
    obs.observe(document.getElementById("subnav"));
};

const intersectionCallback = (entries, observer) => {
    entries.forEach(function(entry) {
        if (entry.isIntersecting) {
            document.getElementById("top-link").classList.add("hide");
        } else {
            document.getElementById("top-link").classList.remove("hide");
        };
    });
};

