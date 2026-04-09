(function () {
  var namespace = window.ui8kit || {};
  window.ui8kit = namespace;
  if (namespace.tooltip) {
    return;
  }

  function ready(fn) {
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", fn);
      return;
    }
    fn();
  }

  function openTooltip(tooltip) {
    var content = tooltip.querySelector('[role="tooltip"]');
    if (!content) {
      return;
    }
    content.removeAttribute("hidden");
    tooltip.setAttribute("data-state", "open");
    content.setAttribute("aria-hidden", "false");
  }

  function closeTooltip(tooltip) {
    var content = tooltip.querySelector('[role="tooltip"]');
    if (!content) {
      return;
    }
    content.setAttribute("hidden", "hidden");
    tooltip.setAttribute("data-state", "closed");
    content.setAttribute("aria-hidden", "true");
  }

  ready(function () {
    var tooltips = document.querySelectorAll('[data-ui8kit="tooltip"]');
    for (var i = 0; i < tooltips.length; i += 1) {
      var root = tooltips[i];
      root.addEventListener("mouseenter", function () {
        openTooltip(this);
      });
      root.addEventListener("focusin", function () {
        openTooltip(this);
      });
      root.addEventListener("mouseleave", function () {
        closeTooltip(this);
      });
      root.addEventListener("focusout", function () {
        closeTooltip(this);
      });
    }
  });

  namespace.tooltip = { init: function () {} };
})();
