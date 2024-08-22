import Plyr from 'plyr';

export default function initVideoPlayer() {
  run();
}

async function run() {
  const videos = document.getElementsByTagName('video');

  if (videos.length > 0) {
    /*
    loadCSS('/js/plyr/plyr.css');
    await loadJS('/js/plyr/plyr.min.js');
    */
    initPlayers(videos);
  }
}

function initPlayers(videos) {
  const players = Array.from(videos).map(v => new Plyr(v, playerOptions(v)));
}

function playerOptions(video) {
  const lang = document.documentElement.lang;
  const captionsActive = displayCaptions(video, lang);
  console.log(captionsActive)
  return {
    disableContextMenu: false,
    captions: { active: captionsActive, language: lang, update: true },
    keyboard: { focused: true, global: true },
    speed: { selected: 1, options: [0.5, 0.75, 1, 1.25, 1.5, 1.75, 2] },
    settings: ['speed', 'captions', 'quality'],
    controls: ['play-large', 'play', 'progress', 'current-time', 'captions', 'settings', 'fullscreen'],
    quality: { default: 1080, options: [1080, 720, 480] },
    invertTime: false,
    storage: { enabled: false, key: 'plyr' }
  }
}

function displayCaptions(video, lang) {
  if (lang == "en") {
    return true
  }
  const children = video.children;
  for (var i = 0; i < children.length; i++) {
    if (children[i].tagName == 'TRACK' && children[i].default) {
      return true
    }
  }
  return false
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
