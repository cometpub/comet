// src/js/ff-xslt-fix.js
document.querySelectorAll(".p-content, .p-summary").forEach(function(elem) {
  if (elem.firstChild?.nodeName == "#text") {
    elem.innerHTML = elem.innerText;
  }
});
