import styles from './layout.module.css';

import { Captions, ChapterTitle, Controls, Gesture } from '@vidstack/react';

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

export function VideoLayout({ thumbnails, selectSource, sources, currentSize}: VideoLayoutProps) {
  return (
    <>
      <Gestures />
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
