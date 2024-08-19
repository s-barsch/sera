
export default function initPreload(html) {
    const l = document.getElementById("body").getElementsByTagName("img");
    if (l.length > 0) {
        console.log("found");
        l[0].addEventListener("load", () => {
            console.log("preload next image");
            document.getElementById("preload").innerHTML = html;
        });
    }
}
