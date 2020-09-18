[<img src="https://raw.githubusercontent.com/stefba/stferal/master/static/svg/stferal-logo-compressed.svg" alt="Saint Feral Logo" width=900>](https://en.stferal.com/)


This application is the website and archive system of Saint Feral. It uses a flat-file database and makes heavy use of a directory tree structure. I share it here for educational purposes only.

## CSS

Compile [Less](http://lesscss.org/) files:

`lessc --clean-css ./css/main.less ./css/dist/main.css`

## JS

Build JavaScript bundle:

`yarn --cwd ./js build`
