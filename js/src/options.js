
const initOptionToggles = () => {
  const l = document.getElementsByClassName("option");
  for (const a of l) {
    a.addEventListener("click", evt => {
      evt.preventDefault();
      const o = getOption(evt.target.pathname);
      setOption(o.name, o.value);
    })
  }
}

// split path like this "/opt/colors/dark" to "colors", "dark"
const getOption = href => {
  const code = href.substr(5)
  const x = code.indexOf('/')
  return {
    name: code.substr(0, x),
    value: code.substr(x + 1)
  };
}

const setOption = (option, value) => {
  fetch("/opt/" + option + "/" + value);
  document.body.dataset[option] = value;
}

export default initOptionToggles;
