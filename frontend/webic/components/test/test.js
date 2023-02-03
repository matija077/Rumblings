function testDecorator() {
    console.log('ola')
}

class Test extends HTMLElement {
  constructor() {
    super();

    this.attachShadow({ mode: "open" });
    this.shadowRoot.innerHTML = `
        <style>
          /* CSS styles for the component go here */
        </style>
        <div class="container" prop1="" prop2="" prop3="">
        <button type="button" class="btn-prop1">Change prop1</button>
        <button type="button" class="btn-prop2">Change prop2</button>
        <button type="button" class="btn-prop3">Change prop3</button>
        </div>
      `;

    this.container = this.shadowRoot.querySelector(".container");


    this.prop1 = 0;
    this.prop2 = 0;
    this.prop3 = 0;

    this.prevProp1 = -1;
    this.prevProp2 = -1;
    this.prevProp3 = -1;

    this.shadowRoot
      .querySelector(".btn-prop1")
      .addEventListener("click", () => (this.prop1 += 1));
    this.shadowRoot
      .querySelector(".btn-prop2")
      .addEventListener("click", () => (this.prop2 += 1));
    this.shadowRoot
      .querySelector(".btn-prop3")
      .addEventListener("click", () => (this.prop3 += 1));
  }

  static get observedAttributes() {
    return ["prop1", "prop2", "prop3"];
  }

  attributeChangedCallback(name, _, newValue) {
    this[name] = newValue;
    this.render();
  }

  render() {
    if (this.prevProp1 !== this.prop1) {
      this.updateProp1();
    }

    if (this.prevProp2 !== this.prop2) {
      this.updateProp2();
    }

    if (this.prevProp3 !== this.prop3) {
      this.updateProp3();
    }

    this.prevProp1 = this.prop1;
    this.prevProp2 = this.prop2;
    this.prevProp3 = this.prop3;
  }

  updateProp1() {
    
    console.log('prop 1 update')
  }

  updateProp2() {
    console.log('prop 2 update')
  }

  updateProp3() {
    console.log('prop 3 update')
  }
}

customElements.define("my-test", Test);
