//import '@vidstack/react/player/styles/default/theme.css';
import './fix.css';

import { useRef, useEffect, useState } from 'react';

import styles from './player.module.css';

import {
  MediaPlayer,
  MediaProvider,
  Poster,
  Track,
  useMediaRemote,
  type MediaPlayerInstance,
} from '@vidstack/react';

import { VideoLayout } from './layout/layout';
import type { Source, MyTrack } from './main';


export function Player({ sources, tracks, poster }: { sources: Source[], tracks: MyTrack[], poster: string}) {
  let player = useRef<MediaPlayerInstance>(null);
  const [currentSrc, setSrc] = useState('');
  const [currentSize, setSize] = useState('');
  const [startTime, setStart] = useState(0);


  const remote = useMediaRemote(player)

  function onCanPlay() {
    if (startTime !== 0) {
      remote.seek(startTime)
      remote.play();
      setStart(0)
    }
  }

  function selectSource(size: string) {
    for (const s of sources) {
      if (s.size === size) {
        const currentTime = player.current?.currentTime!;
        setSrc(s.src)
        setSize(size)
        setStart(currentTime)
        return
      }
    }
    console.log('source not found')
  }

  useEffect(() => {
    selectSource('1080');
  }, [])

  return (
    <>
      <MediaPlayer
        className={`${styles.player} media-player`}
        src={currentSrc}
        crossOrigin
        playsInline
        loop={true}
        autoPlay={true}
        onCanPlay={onCanPlay}
        ref={player}
      >
        <MediaProvider>
          <Poster src={poster} className='vds-poster' />
          {tracks.map((track) => (
            <Track {...track} key={track.src} />
          ))}
        </MediaProvider>

        <VideoLayout selectSource={selectSource} sources={sources} currentSize={currentSize} />
      </MediaPlayer>

    </>
  );
}
