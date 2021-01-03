import Plyr from 'plyr';

export default function initVideoPlayer() {
  run();
  /*
  */
}

async function run() {
  const videos = document.getElementsByTagName('video');

  if (videos.length > 0) {
    loadCSS('/js/plyr/plyr.css');
    await loadJS('/js/plyr/plyr.min.js');
    initPlayers(videos);
  }
}

function initPlayers(videos) {
  const players = Array.from(videos).map(v => new Plyr(v, playerOptions(v)));
}

function playerOptions(video) {
  console.log(video)
  const lang = document.documentElement.lang;
  return {
    disableContextMenu: false,
    captions: { active: false, language: lang, update: false },
    settings: ['captions', 'quality', 'loop'],
    controls: ['play-large', 'play', 'progress', 'current-time', 'mute', 'volume', 'captions', 'settings', 'fullscreen'],
    quality: { default: 1080, options: [1080, 720] },
    invertTime: false
  }
}

const loadCSS = src => {
  return new Promise((resolve, reject) => {
    const script = document.createElement('link')
    script.rel = 'stylesheet'
    script.onload = resolve
    script.onerror = reject
    script.href = src
    document.head.prepend(script)
  })
}

const loadJS = src => {
  return new Promise((resolve, reject) => {
    const script = document.createElement('script')
    script.type = 'text/javascript'
    script.onload = resolve
    script.onerror = reject
    script.src = src
    document.head.append(script)
  })
}
