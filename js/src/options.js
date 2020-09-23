
const initOptionToggles = () => {
  const options = ["colors", "type"];
  //const options = ["colors"];
  for (const opt of options) {
    document.getElementById(opt + "-switch-link").addEventListener("click", (evt) => {
      evt.preventDefault();
      toggleOption(opt);
    });
  }
}

const opposite = (option, value) => {
  console.log(option);
  console.log(value);
  switch (option) {
    case "type":
      if (value == "large") {
        return "small"
      }
      return "large"
    case "colors":
      if (value == undefined || value == "light") {
        return "dark"
      }
      return "light"
  }
  return ""
}

const switchOption = option => {
  const value = document.body.dataset[option];
  return opposite(option, value);
}

const toggleOption = option => {
  const value = switchOption(option);
  fetch("/opt/" + option + "/" + value);
  document.body.dataset[option] = value;
}

export default initOptionToggles;
