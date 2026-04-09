(function () {
  var namespace = "ui8kit";
  var existing = window[namespace];
  if (!existing) {
    existing = {};
    window[namespace] = existing;
  }
  if (existing.ready) {
    return;
  }

  function ready(fn) {
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", fn);
      return;
    }
    fn();
  }

  function byAttr(name, root) {
    var scope = root || document;
    return scope.querySelectorAll("[data-" + name + "]");
  }

  existing.core = {
    ready: ready,
    byAttr: byAttr,
  };
  existing.ready = function (fn) {
    ready(fn);
  };
})();
