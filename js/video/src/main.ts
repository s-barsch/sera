// Register elements.
import 'vidstack/player';
import 'vidstack/player/ui';
import 'vidstack/icons';
import 'vidstack/player/styles/default/theme.css';
import './video-player.css';
import 'vidstack/player/styles/default/layouts/video.css';
import './add.css';
//import './document.css';



import { defineCustomElement, MediaQualityRadioGroupElement } from "vidstack/elements";

defineCustomElement(MediaQualityRadioGroupElement);




const player = document.querySelector('media-player')!;

// We can listen for the `can-play` event to be notified when the player is ready.
player.addEventListener('can-play', () => {
  if (player.qualities.auto === true) {
    const firstQuality = player.qualities[0]!;
    firstQuality.selected = true;
  }
});

