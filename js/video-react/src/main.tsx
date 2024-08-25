import React from 'react';
import ReactDOM from 'react-dom/client';
import { Player } from './player';

interface VideoData {
  sources: Source[];
  tracks: MyTrack[];
  poster: string;
}

export interface Source {
  src: string;
  size: string;
}

export interface MyTrack {
  src: string;
  label: string;
  language: string;
  kind: TextTrackKind;
  default: boolean;
}

function getVideoData(target: HTMLElement): VideoData {
  return {
    sources: getSources(target),
    tracks: getTracks(target),
    poster: target.getElementsByTagName('video')[0].poster,
  }
}

function getSources(target: HTMLElement): Source[] {
  let sources: Source[] = [];
  for (const s of target.getElementsByTagName('source')) {
    sources.push({
      src: s.src,
      size: s.dataset.size!,
    })
  }
  return sources
}

function getTracks(target: HTMLElement): MyTrack[] {
  let sources: MyTrack[] = [];
  for (const t of target.getElementsByTagName('track')) {
    sources.push({
      src: t.src,
      label: t.label,
      kind: t.kind as TextTrackKind,
      default: t.default,
      language: t.srclang
    })
  }
  return sources
}

const target = document.getElementById('video-player')!;
const videoData = getVideoData(target)

ReactDOM.createRoot(target).render(
  <React.StrictMode>
    <Player sources={videoData.sources} tracks={videoData.tracks} poster={videoData.poster} />
  </React.StrictMode>,
);