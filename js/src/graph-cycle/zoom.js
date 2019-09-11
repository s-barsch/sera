/**
 * Zooming images on /graph/2018/09/el
 **/

export default function initZoom() {
  const b = document.getElementById("part");
  const l = b.getElementsByTagName("img");

    for (let i = 0, img; img = l[i]; i++) { 
        img.addEventListener("click", zoomClickHandler);
        if (localStorage["expand"] == "1") {
            img.classList.add("expand");
        };
    }
}

const zoomClickHandler = evt => {
  zoom(evt.target);
}

const zoom = el => {
  el.classList.toggle("expand");
  localStorage["expand"] = localStorage["expand"] == "0" ? "1" : "0";
};
