import observeMonths from './months';
import initTopLink from './top-link';
import initSubnavFix from './subnav-fix';

export default function initOverview(pageType) {
    if (pageType != "graph-main" || window.innerWidth > 1024) {
        initTopLink();
    }
    if (pageType == "graph-year") {
        initSubnavFix();
        observeMonths();
    };
  initZoom();
}


const initZoom = () => {
  const l = document.getElementsByClassName("img");
  for (let img of l) {
    img.addEventListener("click", performZoom);
  }
}

const performZoom = evt => {
  evt.target.classList.toggle("expand");
}

