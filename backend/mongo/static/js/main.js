function submitForm(event) {
  event.preventDefault();

  const formData = new FormData(form);

  console.log("isvrsim se");

  fetch("/api/form", {
    method: "POST",
    body: JSON.stringify(formData),
  });
}

var form = document.getElementById("form");

form.addEventListener("submit", submitForm);

function onNavigationHandler(event) {
  event.preventDefault();

  const parent = event.currentTarget;
  let current = event.target;
  let id = event.target.dataset.id;

  while (!id && current !== parent) {
    current = current.parentElement;
    id = current.dataset.id;
  }

  if (!id) {
    return;
  }

  const currentUrl = document.location.href;
  const baseUrl = currentUrl.slice(0, currentUrl.lastIndexOf("/"));
  let newUrl;

  switch (id) {
    case "Search": {
      newUrl = "/search.html";
      break;
    }
  }

  document.location.href = baseUrl + newUrl
}
