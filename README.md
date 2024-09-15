# Sferal

This application is a flat-file content management system for handling fragmented data such as text snippets, and audio and video clips

`go build; ./sacer -a`

## CSS

Compile [Less](http://lesscss.org/) files:

`lessc --clean-css ./css/main.less ./css/dist/main.css`

## JS

Build JavaScript:

`pnpm -C js/bundle install`\
`pnpm -C js/bundle build`

`pnpm -C js/video install`\
`pnpm -C js/video build`
