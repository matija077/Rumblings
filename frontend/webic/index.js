import mainHtml from "./main.html";

import './components/dropdown/Dropdown.js'
import './pages/my-list/Mylist';
import './components/test/test';

class MyElement extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
  }

  connectedCallback() {
    this.render();
  }

  render() {
    this.shadowRoot.innerHTML = mainHtml;
  }
}

customElements.define("my-element", MyElement);

const root = document.querySelector('#root');
root.innerHTML = '<my-element></my-element>'