(function () {
  var namespace = window.ui8kit || {};
  window.ui8kit = namespace;
  if (namespace.accordion) {
    return;
  }

  function ready(fn) {
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", fn);
      return;
    }
    fn();
  }

  function getAccordionRoots() {
    return document.querySelectorAll('[data-ui8kit="accordion"]');
  }

  function isMultiple(root) {
    return (root.getAttribute("data-accordion-type") || "single") === "multiple";
  }

  function setItemState(item, open) {
    var trigger = item.querySelector('[data-ui8kit-accordion-trigger]');
    var panel = item.querySelector('[data-ui8kit-accordion-content]');
    if (!trigger || !panel) {
      return;
    }
    item.setAttribute("data-state", open ? "open" : "closed");
    if (trigger) {
      trigger.setAttribute("aria-expanded", open ? "true" : "false");
    }
    if (panel) {
      if (open) {
        panel.removeAttribute("hidden");
      } else {
        panel.setAttribute("hidden", "hidden");
      }
    }
  }

  function closeOthers(root, currentItem) {
    var items = root.querySelectorAll('[data-accordion-item]');
    for (var i = 0; i < items.length; i += 1) {
      if (items[i] !== currentItem) {
        setItemState(items[i], false);
      }
    }
  }

  function toggle(root, trigger, panel, item, open) {
    var current = item.getAttribute("data-state") === "open";
    var next = typeof open === "boolean" ? open : !current;
    if (!isMultiple(root)) {
      closeOthers(root, item);
    }
    setItemState(item, next);
  }

  ready(function () {
    var roots = getAccordionRoots();
    for (var i = 0; i < roots.length; i += 1) {
      var root = roots[i];
      var items = root.querySelectorAll('[data-accordion-item]');
      for (var j = 0; j < items.length; j += 1) {
        var item = items[j];
        var trigger = item.querySelector('[data-ui8kit-accordion-trigger]');
        var panel = item.querySelector('[data-ui8kit-accordion-content]');
        if (!trigger || !panel) {
          continue;
        }
        var openByDefault = item.getAttribute("data-state") === "open";
        setItemState(item, openByDefault);
        trigger.setAttribute("type", "button");
        trigger.addEventListener("click", function (evt) {
          evt.preventDefault();
          var itemNode = evt.currentTarget.closest('[data-accordion-item]');
          if (!itemNode) {
            return;
          }
          var accordionRoot = evt.currentTarget.closest('[data-ui8kit="accordion"]');
          toggle(accordionRoot, evt.currentTarget, itemNode.querySelector('[data-ui8kit-accordion-content]'), itemNode);
        });
        trigger.addEventListener("keydown", function (evt) {
          if (evt.key !== "Enter" && evt.key !== " ") {
            return;
          }
          evt.preventDefault();
          evt.currentTarget.click();
        });
      }
    }
  });

  namespace.accordion = { init: function () {} };
})();
