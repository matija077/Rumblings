import RandomData from "../../services/randomData";
import Http from "../../services/http";
import html from "./myList.html";

class MyList extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });

    const style = document.createElement("style");
    style.textContent = `
      ul {
        list-style-type: none;
        padding: 0;
        margin: 0;
      }

			ul > li + li {
				margin-top: 20px;
			}

			li > span + span {
				margin-left: 5px;
			}
    `;
    this.shadowRoot.innerHTML = html;
    this.shadowRoot.appendChild(style);
  }
  async connectedCallback() {
    const randomData = new RandomData(new Http());
    const { data, error } = await randomData.getRandomData();

    if (error) {
      return;
    }

    this.includedArr = Array.from({ length: data.length }).fill(true);
    const includedArr = this.includedArr;

    this.offset = 0;
    this.limit = 10;
    this.end = includedArr.length;

    const list = this.shadowRoot.querySelector(".my-list");
    const listTemplate = list.querySelector(".my-list-template");
    data.forEach((person, idx) => {
      const listTemplateClone = listTemplate.content.cloneNode(true);
      const personItem = listTemplateClone.querySelector(".person");
      const personItemName = personItem.querySelector('.person-name');
      const dropdown = listTemplateClone.querySelector("my-dropdown");
      const personClassName = `person-${idx}`;

      personItem.appendChild(dropdown);
      list.appendChild(personItem);

      personItem.dataset.idx = idx;
      personItem.classList.add(personClassName);
      personItem.onclick = (e) => {
        includedArr[e.currentTarget.dataset.idx] =
          !includedArr[e.currentTarget.dataset.idx];
          dropdown.isVisible = !dropdown.isVisible
      };

      personItemName.textContent = person.first;

      dropdown.classList.add(personClassName);
      // goes after because method doesnt exists until web component is connected to the actual DOM
      dropdown.addData(person, idx); 
    });

    const back = this.shadowRoot.querySelector(".backButton");
    const forward = this.shadowRoot.querySelector(".forwardButton");
    back.clickHandler = () => this.goBack();
    forward.clickHandler = () => this.goForward();

    this.shadowRoot.appendChild(list);
  }

  goBack() {
    console.log("back");
    if (this.offset === 0) {
      return;
    }

    this.offset -= 10;
  }

  goForward() {
    console.log('forward')
    console.log(this.offset);
    console.log(this.end);
    if (this.offset + this.limit >= this.end) {
      return;
    }

    this.offset += 10;
  }

  requestUpdate() {
    console.log("wtf");
  }

  update() {
    console.log("wtf");
  }

  attributeChangedCallback(name, oldValue, newValue) {
    console.log("uspjeh");
  }

  static get observedAttributes() {
    return ["offset"];
  }
}

customElements.define("my-list", MyList);

// The code below can go after the define() method

// Use the custom element in the HTML file to create an instance of the list
const list = document.createElement("my-list");

// Set any desired attributes and properties on the list element
list.setAttribute("id", "my-list");
list.classList.add("my-list");

// Add the list to the page
document.body.appendChild(list);

// Call the connectedCallback() method when the 'load' event is emitted
window.addEventListener("load", () => {
  const myList = document.querySelector("my-list");
  try {
    myList.connectedCallback();
  } catch (err) {}
});

/*const div = document.createElement('div')
const myList = document.getElementsByClassName('my-list')
div.innerHTML = myList.includedArr.map(d => `<p>${d}</p>`)*/
