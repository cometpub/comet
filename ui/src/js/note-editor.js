import { Editor } from "@tiptap/core";
import Bold from "@tiptap/extension-bold";
import BulletList from "@tiptap/extension-bullet-list";
import Code from "@tiptap/extension-code";
import Document from "@tiptap/extension-document";
import History from "@tiptap/extension-history";
import Italic from "@tiptap/extension-italic";
import ListItem from "@tiptap/extension-list-item";
import OrderedList from "@tiptap/extension-ordered-list";
import Paragraph from "@tiptap/extension-paragraph";
import Placeholder from "@tiptap/extension-placeholder";
import Strike from "@tiptap/extension-strike";
import TaskItem from "@tiptap/extension-task-item";
import TaskList from "@tiptap/extension-task-list";
import Text from "@tiptap/extension-text";
import Typography from "@tiptap/extension-typography";
import { Markdown } from "tiptap-markdown";

const noteEditorTmpl = document.createElement("template");
noteEditorTmpl.innerHTML = `
<style>
note-editor {
  display: flex;
  flex-direction: column;
  gap: var(--space-xs);
}

.ProseMirror p.is-editor-empty:first-child::before {
  content: attr(data-placeholder);
  float: left;
  color: var(--text-secondary);
  pointer-events: none;
  height: 0;
}

note-editor .tiptap {
  outline-offset: 5px;
}

note-editor .tiptap [data-type="taskList"] {
  padding-inline-start: 0;
}

note-editor #photos-previews {
  display: flex;
  gap: var(--space-2xs);
  flex-wrap: wrap;
}

note-editor #photos-previews:empty {
  display: none;
}

note-editor #photos-previews img {
  height: 96px;
  display: block;
}

note-editor label {
  user-select: none;
  cursor: pointer;
}

note-editor input:focus + label {
  outline: medium auto currentColor;
  outline: medium auto invert;
  outline: 5px auto -webkit-focus-ring-color;
}
note-editor [role="group"] {
  display: flex;
  gap: var(--space-2xs);
}
note-editor #clear-attachments {
  align-self: start;
  font-size: var(--font-size-s);
}

note-editor #clear-attachments:has(+ #photos-previews:empty) {
  display: none;
}
</style>
<div id="editor"></div>
<div role="group">
  <input type="checkbox" id="status" class="sr-only" />
  <label for="status">
    <span class="icon-unlocked" aria-label="Public post"></span>
  </label>

  <input type="file" id="photos" class="sr-only" accept="image/png,image/jpeg,image/gif,image/webp,image/avif"/>
  <label for="photos">
    <span class="icon-image" aria-label="Attach a photo"></span>
  </label>
</div>
<button id="clear-attachments" role="link">
  <strong><span class="sr-only">Clear</span>Attachments</strong> <span class="icon-cancel-circle"></span>
</button>
<ul id="photos-previews" role="list"></ul>
`;

class NoteEditor extends HTMLElement {
  static formAssociated = true;

  /**
   * @type {Editor}
   */
  #editor;

  /**
   * @type {ElementInternals}
   */
  #internals;

  /**
   * @type {HTMLInputElement}
   */
  #statusElem;

  /**
   * @type {HTMLSpanElement}
   */
  #statusIconElem;

  /**
   * @type {HTMLInputElem}
   */
  #photosElem;

  /**
   * @type {HTMLULListElement}
   */
  #photosPreviewElem;

  /**
   * @type {HTMLButtonElement}
   */
  #clearAttachmentsElem;

  /**
   * @type {FileList}
   */
  #files = [];

  constructor() {
    super();
    this.#internals = this.attachInternals();
    this.appendChild(noteEditorTmpl.content.cloneNode(true));
  }

  connectedCallback() {
    this.#statusElem = this.querySelector("#status");
    this.#statusIconElem = this.querySelector("label[for='status'] > span");
    this.#photosElem = this.querySelector("#photos");
    this.#photosPreviewElem = this.querySelector("#photos-previews");
    this.#clearAttachmentsElem = this.querySelector("#clear-attachments");

    const editorElem = this.querySelector("#editor");

    this.#editor = new Editor({
      element: editorElem,
      extensions: [
        Bold,
        BulletList,
        Code,
        Document,
        History,
        Italic,
        ListItem,
        OrderedList,
        Paragraph,
        Placeholder.configure({
          placeholder: "What’s on your mind?",
        }),
        Strike,
        TaskItem,
        TaskList,
        Text,
        Typography,
        Markdown.configure({
          html: false,
          tightLists: true,
          bulletListMarker: "-",
          linkify: true,
          breaks: false,
          transformPastedText: true,
          transformCopiedText: true,
        }),
      ],
    });

    this.#editor.on("update", () => {
      this.onChange();
    });

    this.#statusElem.onchange = () => {
      this.onStatusChange();
    };

    this.#photosElem.onchange = () => {
      this.#files = Array.from(this.#photosElem.files);
      this.onChange();
      this.updatePhotoPreviews();
    };

    this.#clearAttachmentsElem.onclick = (event) => {
      event.preventDefault();
      this.#files = [];
      this.onChange();
      this.updatePhotoPreviews();
    };
  }

  get name() {
    return this.getAttribute("name");
  }

  get status() {
    return this.#statusElem.checked ? "private" : "public";
  }

  onStatusChange() {
    if (this.#statusElem.checked) {
      this.#statusIconElem.classList.replace("icon-unlocked", "icon-lock");
      this.#statusIconElem.ariaLabel = "Private post";
    } else {
      this.#statusIconElem.classList.replace("icon-lock", "icon-unlocked");
      this.#statusIconElem.ariaLabel = "Public post";
    }

    this.onChange();
  }

  updatePhotoPreviews() {
    this.#photosPreviewElem.innerHTML = "";

    for (let file of this.#files) {
      let li = document.createElement("li");
      let img = document.createElement("img");
      img.src = URL.createObjectURL(file);
      li.onclick = () => {
        this.removePhoto(file);
      };
      li.appendChild(img);
      this.#photosPreviewElem.appendChild(li);
    }
  }

  onChange() {
    const formData = new FormData();

    formData.append("summary", this.#editor.storage.markdown.getMarkdown());
    formData.append("status", this.status);

    if (this.#files.length > 0) {
      formData.append("photos", this.#files[0]);
    }

    this.#internals.setFormValue(formData);
  }
}

customElements.define("note-editor", NoteEditor);
