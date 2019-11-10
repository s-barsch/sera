import initCycle from "./graph-cycle/init";
import initOverview from "./graph-overview/init";
import initText from "./text";
import initAudio from "./audio";
import initGraphMore from "./graph-overview/more";

let stInit = page => {
    switch (page.type) {
        case "graph-main":
            initGraphMore();
            break;
        case "graph-year":
            initOverview(page.type);
            break;
        case "graph-el":
            initCycle(page);
            break;
    }
}

document.addEventListener("DOMContentLoaded", function() {
    initText();
    initAudio();
    initLargeFont();
    if (typeof pageVars !== "undefined") {
        stInit(pageVars);
    }
});

let initLargeFont = () => {
    document.getElementById("large-type").addEventListener("click", evt => {
        evt.preventDefault();
        document.body.classList.toggle("large");
    })
}
