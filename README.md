[<img src="https://raw.githubusercontent.com/stefba/stferal/master/static/svg/stferal-logo-compressed.svg" alt="Stef Feral Logo" width=900>](https://en.stferal.com/)

Source code of the website [stferal.com](https://en.stferal.com/).

## JS

Build JavaScript bundle:

`yarn --cwd ./js build`

## CSS

Compile [Less](http://lesscss.org/) files:

`lessc --clean-css ./css/main.less ./css/dist/main.css`

## Go

Run application:

`go build; ./stferal -a`
