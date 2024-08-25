import {
  Menu,
  MenuInstance,
  Tooltip,
  useCaptionOptions,
  type IconComponent,
  type MenuPlacement,
  type TooltipPlacement,
} from '@vidstack/react';
import {
  ArrowLeftIcon,
  ChevronRightIcon,
  ClosedCaptionsIcon,
  SettingsIcon,
  SettingsMenuIcon,
} from '@vidstack/react/icons';

import React from 'react';

import type { Source } from '../main';
import { useRef } from 'react';
import type { RadioSelectEvent } from 'vidstack';

export interface SettingsProps {
  placement: MenuPlacement;
  tooltipPlacement: TooltipPlacement;
  selectSource(size: string): void;
  sources: Source[];
  currentSize: string;
}

export function Settings({ placement, tooltipPlacement, selectSource, sources, currentSize }: SettingsProps) {
  const menu = useRef<MenuInstance>(null);
  return (
    <Menu.Root className="vds-menu" ref={menu}>
      <Tooltip.Root>
        <Tooltip.Trigger asChild>
          <Menu.Button className="vds-menu-button vds-button">
            <SettingsIcon className="vds-rotate-icon" />
          </Menu.Button>
        </Tooltip.Trigger>
        <Tooltip.Content className="vds-tooltip-content" placement={tooltipPlacement}>
          Settings
        </Tooltip.Content>
      </Tooltip.Root>
      <Menu.Content className="vds-menu-items" placement={placement}>
        <CaptionSubmenu />
        <VideoSourceSubmenu parent={menu} sources={sources} currentSize={currentSize} selectSource={selectSource} />
      </Menu.Content>
    </Menu.Root>
  );
}

function CaptionSubmenu() {
  const options = useCaptionOptions(),
    hint = options.selectedTrack?.label ?? 'Off';
  return (
    <Menu.Root>
      <SubmenuButton
        label="Captions"
        hint={hint}
        disabled={options.disabled}
        icon={ClosedCaptionsIcon}
      />
      <Menu.Content className="vds-menu-items">
        <Menu.RadioGroup className="vds-captions-radio-group vds-radio-group" value={options.selectedValue}>
          {options.map(({ label, value, select }) => (
            <RadioItem value={value} label={label} select={select} /> 
          ))}
        </Menu.RadioGroup>
      </Menu.Content>
    </Menu.Root>
  );
}

function RadioItem({ select, label, value }: { value: string, label: string, select: (trigger?: Event) => void;}) {
  return (
    <Menu.Radio className="vds-radio" value={value} onSelect={select} key={value}>
      <div className="vds-radio-check" />
      <span className="vds-radio-label">{label}</span>
    </Menu.Radio>
  );
}

function VideoSourceSubmenu({selectSource, sources, currentSize, parent}: { parent: React.RefObject<MenuInstance>, selectSource(size: string): void; sources: Source[]; currentSize: string}) {
  function select(event?: Event) {
    let rEv = event as RadioSelectEvent;
    selectSource(rEv.target.$props.value())
    parent.current?.close();
  }
  const menu = useRef<MenuInstance>(null);
  return (
    <Menu.Root ref={menu}>
      <SubmenuButton
        label="Quality"
        hint={currentSize}
        disabled={false}
        icon={SettingsMenuIcon}
      />
      <Menu.Content className="vds-menu-items">
        <Menu.RadioGroup className="vds-radio-group" value={currentSize}>
          {sources.map(({ size }) => (
            <RadioItem value={size} label={size} select={select} />
          ))}
        </Menu.RadioGroup>
      </Menu.Content>
    </Menu.Root>
  );
}

export interface SubmenuButtonProps {
  label: string;
  hint: string;
  disabled?: boolean;
  icon: IconComponent;//ReactNode;
}

function SubmenuButton({ label, hint, icon: Icon, disabled }: SubmenuButtonProps) {
  return (
    <Menu.Button className="vds-menu-button" disabled={disabled}>
      <ArrowLeftIcon className="vds-menu-button-close-icon" />
      <Icon className="vds-menu-button-icon" />
      <span className="vds-menu-button-label">{label}</span>
      <span className="vds-menu-button-hint">{hint}</span>
      <ChevronRightIcon className="vds-menu-button-open-icon" />
    </Menu.Button>
  );
}
