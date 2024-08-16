// https://bugzilla.mozilla.org/show_bug.cgi?id=98168
document.querySelectorAll(".p-content, .p-summary").forEach(function (elem) {
  if (elem.firstChild?.nodeName == "#text") {
    elem.innerHTML = elem.innerText;
  }
});
