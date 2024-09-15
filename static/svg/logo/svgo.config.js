module.exports = {
    plugins: [
        { name: 'removeComments' },
        {
            name: "removeAttrs",
            params: {
                attrs: ["line-height", "font-size", "font-family", "line-height", "word-spacing", "letter-spacing", "style", "fill"],
            }
        }
    ]
}