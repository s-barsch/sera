module.exports = {
    plugins: [
        { name: 'removeComments' },
        {
            name: "removeAttrs",
            params: {
                attrs: ["xmlns", "xmlns:inkscape", "line-height", "font-size", "font-family", "line-height", "word-spacing", "letter-spacing", "style", "fill"],
            }
        }
    ]
}