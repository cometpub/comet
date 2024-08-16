// src/js/utils/ff-xslt-fix.js
document.querySelectorAll(".p-content, .p-summary").forEach(function(elem) {
  if (elem.firstChild?.nodeName == "#text") {
    elem.innerHTML = elem.innerText;
  }
});

// src/js/components/locale-time.js
var localeDateTmpl = document.createElement("template");
localeDateTmpl.innerHTML = `
<time></time>
`;
var LocaleDate = class extends HTMLElement {
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
      day: "numeric"
    });
  }
};
customElements.define("locale-date", LocaleDate);

// src/js/components/theme-toggle.js
var themeToggleTempl = document.createElement("template");
themeToggleTempl.innerHTML = `
<style>
button {
  --size: var(--space-lg);

  background: none;
  color: var(--text-accent);
  border: none;
  padding: 0;
  cursor: pointer;
  display: flex;

  inline-size: var(--size);
  block-size: var(--size);
  aspect-ratio: 1;
  border-radius: 50%;

  touch-action: manipulation;
  -webkit-tap-highlight-color: transparent;

  &:is(:hover, :focus-visible) > svg > :is(.moon, .sun, .sun-beams) {
    opacity: 0.7;
  }
}

svg {
  inline-size: 100%;
  block-size: 100%;
  stroke-linecap: round;
}

@media (hover: none) {
  button {
    --size: 36px;
    padding: 6px;
  }
}

@media (hover: none) and (min-width: 350px) {
  button {
    --size: 48px;
    padding: 12px;
  }
}

svg > :is(.moon, .sun, .sun-beams) {
  transform-origin: center center;
}

  

svg > .sun-beams {
  stroke-width: 2px;
}

@media (prefers-reduced-motion: no-preference) {
  .sun {
    transition: transform 0.5s var(--ease-elastic-3);
  }

  svg > .sun-beams {
    transition: transform 0.5s var(--ease-elastic-4),
      opacity 0.5s var(--ease-3);
  }

  svg .moon > circle {
    transition: transform 0.25s var(--ease-out-5);
  }

  @supports (cx: 1) {
    svg .moon > circle {
      transition: cx 0.25s var(--ease-out-5);
    }
  }
}


:host([theme="dark"]) {
  & svg > .sun {
    transform: scale(1.75);
  }

  & svg > .sun-beams {
    opacity: 0;
  }

  & svg > .moon > circle {
    transform: translateX(-7px);
  }

  @supports (cx: 1) {
    & svg > .moon > circle {
      transform: translateX(0);
      cx: 17;
    }
  }

  & svg > .sun {
    transition-timing-function: var(--ease-3);
    transition-duration: 0.25s;
  }

  & svg > .sun-beams {
    transform: rotateZ(-25deg);
    transition-duration: 0.15s;
  }

  svg > .moon > circle {
    transition-delay: 0.25s;
    transition-duration: 0.5s;
  }
}
</style>
<button
  id="theme-toggle"
  title="Toggles light & dark"
  aria-label="auto"
  aria-live="polite"
>
  <svg aria-hidden="true" width="24" height="24" viewBox="0 0 24 24">
    <circle
      class="sun"
      cx="12"
      cy="12"
      r="6"
      mask="url(#moon-mask)"
      fill="currentColor"
    ></circle>
    <g class="sun-beams" stroke="currentColor">
      <line x1="12" y1="1" x2="12" y2="3"></line>
      <line x1="12" y1="21" x2="12" y2="23"></line>
      <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
      <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
      <line x1="1" y1="12" x2="3" y2="12"></line>
      <line x1="21" y1="12" x2="23" y2="12"></line>
      <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
      <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
    </g>
    <mask class="moon" id="moon-mask">
      <rect x="0" y="0" width="100%" height="100%" fill="white"></rect>
      <circle cx="24" cy="10" r="6" fill="black"></circle>
    </mask>
  </svg>
</button>
`;
var ThemeToggle = class extends HTMLElement {
  #storageKey = "pref:theme";
  #theme;
  /**
   * @type {HTMLButtonElement}
   */
  #button;
  constructor() {
    super();
    this.attachShadow({ mode: "open" }).appendChild(
      themeToggleTempl.content.cloneNode(true)
    );
  }
  connectedCallback() {
    this.theme = this.getColorPreference();
    window.onload = () => {
      this.reflectPreference();
      this.#button = this.shadowRoot.getElementById("theme-toggle");
      this.#button?.addEventListener("click", () => this.onClick());
    };
    window.matchMedia("(prefers-color-scheme: dark)").addEventListener("change", ({ matches: isDark }) => {
      theme.value = isDark ? "dark" : "light";
      setPreference();
    });
  }
  get theme() {
    return this.getAttribute("theme");
  }
  set theme(value) {
    this.#theme = value;
    this.setAttribute("theme", value);
    localStorage.setItem(this.#storageKey, this.#theme);
    this.reflectPreference();
  }
  getColorPreference() {
    if (localStorage.getItem(this.#storageKey))
      return localStorage.getItem(this.#storageKey);
    else
      return window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
  }
  onClick() {
    this.theme = this.theme === "light" ? "dark" : "light";
  }
  reflectPreference() {
    document.firstElementChild?.setAttribute("data-theme", this.theme);
    document.querySelector("#theme-toggle")?.setAttribute("aria-label", this.theme);
  }
};
customElements.define("theme-toggle", ThemeToggle);
