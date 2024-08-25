import styles from './layout.module.css';

import { Captions, ChapterTitle, Controls, Gesture, Spinner } from '@vidstack/react';

import * as Buttons from './buttons';
import * as Menus from './menus';
import * as Sliders from './sliders';
import { TimeGroup } from './time-group';
import type { Source } from '../main';

export interface VideoLayoutProps {
  thumbnails?: string;
  selectSource(size: string): void;
  sources: Source[];
  currentSize: string;
}

function BufferingIndicator() {
  return (
    <div className="vds-buffering-indicator">
      <Spinner.Root className="vds-buffering-spinner">
        <Spinner.Track className="vds-buffering-track" />
        <Spinner.TrackFill className="vds-buffering-track-fill" />
      </Spinner.Root>
    </div>
  );
}

export function VideoLayout({ thumbnails, selectSource, sources, currentSize}: VideoLayoutProps) {
  return (
    <>
      <Gestures />
      <BufferingIndicator />
      <Captions className={`${styles.captions} vds-captions`} />
      <Controls.Root className={`${styles.controls} vds-controls`}>
        <div className="vds-controls-spacer" />
        <Controls.Group className={`${styles.controlsGroup} vds-controls-group`}>
          <Sliders.Time thumbnails={thumbnails} />
        </Controls.Group>
        <Controls.Group className={`${styles.controlsGroup} vds-controls-group`}>
          <Buttons.Play tooltipPlacement="top start" />
          <TimeGroup />
          <ChapterTitle className="vds-chapter-title" />
          <div className="vds-controls-spacer" />
          <Buttons.Caption tooltipPlacement="top" />
          <Menus.Settings placement="top end" tooltipPlacement="top" selectSource={selectSource} sources={sources} currentSize={currentSize} />
          <Buttons.Fullscreen tooltipPlacement="top end" />
        </Controls.Group>
      </Controls.Root>
    </>
  );
}

function Gestures() {
  return (
    <>
      <Gesture className={styles.gesture} event="pointerup" action="toggle:paused" />
      <Gesture className={styles.gesture} event="dblpointerup" action="toggle:fullscreen" />
      <Gesture className={styles.gesture} event="pointerup" action="toggle:controls" />
      <Gesture className={styles.gesture} event="dblpointerup" action="seek:-10" />
      <Gesture className={styles.gesture} event="dblpointerup" action="seek:10" />
    </>
  );
}
