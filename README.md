[<img src="https://raw.githubusercontent.com/stefba/stferal/master/static/svg/stferal-logo-compressed.svg" alt="Stef Feral Logo" width=900>](https://en.stferal.com/)

Source code of the website [stferal.com](https://en.stferal.com/).

---

This application is the website and archvie system of Stef Feral. It uses a flat-file database and makes heavy use of a directory tree structure.

## CSS

Compile [Less](http://lesscss.org/) files:

`lessc --clean-css ./css/main.less ./css/dist/main.css`

## JS

Build JavaScript bundle:

`yarn --cwd ./js build`
