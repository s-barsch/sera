import initCycle from './graph-cycle/init';
import initOverview from './graph-overview/init';
import initText from './text';
import initAudio from './audio';
import initGraphMore from './graph-overview/more';
import initOptionToggles from './options';

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
  initOptionToggles();
  initServiceWorker();
  if (typeof pageVars !== "undefined") {
    stInit(pageVars);
  }
});

const initServiceWorker = () => {
  if ("serviceWorker" in navigator) {
    window.addEventListener("load", function() {
      navigator.serviceWorker.register("/sw.js").then(function(registration) {
      }, function(err) {
        console.log("ServiceWorker registration failed: ", err);
      });
    });
  }
}
