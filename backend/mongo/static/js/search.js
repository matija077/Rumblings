const search = document.getElementById("search");
search.addEventListener("input", handleSearch);

var isDebouncing = false;
var timerId;

function setDebouncer(value) {
  if (timerId) {
    window.clearTimeout(timerId);
  }

  timerId = setTimeout(() => {
    isDebouncing = false;
    makeRequest(value);
  }, 2000);
}

function handleSearch(event) {
  console.log(event);
  setDebouncer(event.target.value);
}

function makeRequest(value) {
  fetch(`http://localhost:3000/api/search?key=SSN&value=${value}`, {method: 'GET'});
}
