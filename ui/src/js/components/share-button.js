const shareButtonTmpl = document.createElement("template");
shareButtonTmpl.innerHTML = `
<style>
    button {
        position: relative;
        border: none;
        background-color: transparent;
        cursor: pointer;
        font-size: var(--font-size-l);
        color: var(--text-accent);
    }
    p {
        position: absolute;
        right: 0;
        bottom: 100%;
        font-size: var(--font-size-s);
        background-color: var(--surface-primary);
        color: var(--text-alt);
        line-height: 1;
        padding: var(--space-2xs);
        border-radius: 0.33em;

    }
    p::before {
        content: "";
        position: absolute;
        top: 100%;
        right: var(--space-2xs);
        width: 0;
        border-top: var(--space-2xs) solid var(--surface-primary);
        border-left: var(--space-2xs) solid transparent;
        border-right: var(--space-2xs) solid transparent;
    }
    p[aria-hidden="true"] {
        display: none;
    }
</style>
<button title="Share a link to this post">
    <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 256 256"><path fill="currentColor" d="M176 160a39.9 39.9 0 0 0-28.62 12.09l-46.1-29.63a39.8 39.8 0 0 0 0-28.92l46.1-29.63a40 40 0 1 0-8.66-13.45l-46.1 29.63a40 40 0 1 0 0 55.82l46.1 29.63A40 40 0 1 0 176 160m0-128a24 24 0 1 1-24 24a24 24 0 0 1 24-24M64 152a24 24 0 1 1 24-24a24 24 0 0 1-24 24m112 72a24 24 0 1 1 24-24a24 24 0 0 1-24 24"/></svg>

    <p aria-hidden="true">Copied!</p>
</button>
`;

class ShareButton extends HTMLElement {
  static WebShareSupported =
    navigator != null && typeof navigator.share == "function";

  /**
   * @type {HTMLButtonElement}
   */
  #buttonElem;

  /**
   * @type {HTMLParagraphElement}
   */
  #tooltipElem;

  /**
   * @type {number}
   */
  #tooltipTimeout;

  constructor() {
    super();
    let children = this.innerHTML;
    this.attachShadow({ mode: "open" }).appendChild(
      shareButtonTmpl.content.cloneNode(true)
    );
    this.shadowRoot.lastChild.innerHTML = children;
  }

  connectedCallback() {
    this.#buttonElem = this.shadowRoot.querySelector("button");
    this.#tooltipElem = this.shadowRoot.querySelector("p");

    this.#buttonElem.onclick = this.WebShareSupported
      ? () => this.share()
      : () => this.copyToClipboard();
  }

  get title() {
    return this.getAttribute("title");
  }

  get text() {
    return this.getAttribute("text");
  }

  get url() {
    return this.getAttribute("url");
  }

  async share() {
    const shareData = {
      url: this.url,
    };

    if (this.title) {
      shareData.title = this.title;
    }

    if (this.text) {
      shareData.text = this.text;
    }

    try {
      await navigator.share(shareData);
    } catch {}
  }

  copyToClipboard() {
    navigator.clipboard.writeText(this.url);
    this.#tooltipElem.ariaHidden = false;

    if (!!this.#tooltipTimeout) {
      clearTimeout(this.#tooltipTimeout);
    }

    this.#tooltipTimeout = setTimeout(
      () => (this.#tooltipElem.ariaHidden = true),
      1500
    );
  }
}

customElements.define("share-button", ShareButton);
