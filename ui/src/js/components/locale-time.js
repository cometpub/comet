const localeDateTmpl = document.createElement("template");
localeDateTmpl.innerHTML = `
<time></time>
`;

class LocaleDate extends HTMLElement {
  /**
   * @type {HTMLTimeElement}
   */
  #timeElem;

  constructor() {
    super();
    this.attachShadow({ mode: "open" }).appendChild(
      localeDateTmpl.content.cloneNode(true)
    );
  }

  connectedCallback() {
    this.#timeElem = this.shadowRoot.querySelector("time");
    this.#timeElem.dateTime = this.datetime;
    this.#timeElem.textContent = this.localize();
  }

  get datetime() {
    return this.getAttribute("datetime");
  }

  get dateObj() {
    return new Date(this.datetime);
  }

  localize() {
    return this.dateObj.toLocaleDateString(navigator.language, {
      year: "numeric",
      month: "short",
      day: "numeric",
    });
  }
}

customElements.define("locale-date", LocaleDate);
