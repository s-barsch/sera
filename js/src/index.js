import initCycle from "./graph-cycle/init";
import initOverview from "./graph-overview/init";
import initText from "./text";
import initAudio from "./audio";
import initGraphMore from "./graph-overview/more";

const stInit = page => {
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
  initServiceWorker();
  if (typeof pageVars !== "undefined") {
    stInit(pageVars);
  }
});

const initServiceWorker = () => {
  if ("serviceWorker" in navigator) {
    window.addEventListener("load", function() {
      navigator.serviceWorker.register("/service-worker.js").then(function(registration) {
        console.log("ServiceWorker registration successful with scope: ", registration.scope);
      }, function(err) {
        console.log("ServiceWorker registration failed: ", err);
      });
    });
  }
}
