import html from './button.html';

class Button extends HTMLElement {
    constructor() {
      // Always call super first in constructor
      super();
  
      // Create a shadow root
      const shadow = this.attachShadow({mode: 'open'});

      this.shadowRoot.innerHTML = html;
  
      // Create a button element
      const button = this.shadowRoot.querySelector('.my-button');
      button.type = this.getAttribute('type');

      // Add an event listener to the button
      button.addEventListener('click', (e) => {
        this.clickHandler?.(e)
      });
  
      // Append the button to the shadow root
      shadow.appendChild(button);
    }

    set clickHandler(clickHandler) {
        console.log('setano')
        this.clickHandler = clickHandler;
    }
  }

 
  
  // Define the new element
  customElements.define('my-button', Button);
  