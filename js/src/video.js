export default function initHLSVideo() {
  let l = document.getElementsByTagName("video");
  for (let i = 0; i < l.length; i++) {
    let video = l[i];
    for (let j = 0; j < video.children.length; j++) {
      let source = video.children[j];
      if (source.src.substr(-4) == "m3u8") {
        if (!video.canPlayType('application/vnd.apple.mpegurl')) {
          let hls = new Hls();
          hls.loadSource(source.src);
          hls.attachMedia(video);
        }
      }
    }
  }
}
