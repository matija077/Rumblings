import DropHtml from "./dropdown.html";

export class Dropdown extends HTMLElement {
  _isVisible = false;
  
  constructor() {
    super();
    this.attachShadow({ mode: "open" });

    const style = document.createElement("style");
    style.textContent = `
			ul {
				display: block;
				height: 100px
			}
    `;

    this.shadowRoot.innerHTML = DropHtml;
    this.shadowRoot.appendChild(style);
  }

  get isVisible() {
    return this._isVisible;
  }

  set isVisible(isVisible) {
    this._isVisible = isVisible;
    this.changeVisibility(this.isVisible);
  }

  connectedCallback() {
    this.style.display = "none";
  }

  render() {}

  addData(data, idx) {
    const ul = document.createElement("ul");
    this.className = `dropdown-${idx}`;

    Object.entries(data).forEach(([key, value]) => {
      const li = document.createElement("li");
      // TODO li.onclick = () => {}
      const span = document.createElement("span");
      const span2 = document.createElement("span");
      span.textContent = key + ": ";
      span2.textContent = value;
      li.appendChild(span);
      li.appendChild(span2);
      ul.appendChild(li);
    });

    this.shadowRoot.appendChild(ul);
  }

  changeVisibility(value) {
    if (value) {
      this.style.display = "block";
    } else {
      this.style.display = "none";
    }
  }
}

customElements.define("my-dropdown", Dropdown);
