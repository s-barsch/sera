export default function initAudio() {
    
    let audio = document.getElementsByTagName("audio");
    for (let i = 0; i < audio.length; i++) {
        let captions = audio[i].parentElement.getElementsByClassName("captions")[0];
        //console.log(captions);
        for (let j = 0; j < audio[i].textTracks.length; j++) {
            audio[i].onplay = function() {
              captions.classList.add("showing");
            };
            audio[i].textTracks[j].oncuechange = function() {
                // assuming there is only one active cue
                var cue = this.activeCues[0];
                if (cue) {
                    captions.innerHTML = "<mark>" + cue.text + "</mark>";

                }
            }
        }
    }
}
