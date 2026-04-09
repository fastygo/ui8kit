(function () {
  var namespace = window.ui8kit || {};
  window.ui8kit = namespace;
  if (namespace.tabs) {
    return;
  }

  function ready(fn) {
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", fn);
      return;
    }
    fn();
  }

  function activateTab(tabRoot, value, useFocus) {
    var triggers = tabRoot.querySelectorAll('[data-tabs-trigger]');
    var panels = tabRoot.querySelectorAll('[data-tabs-panel]');
    for (var i = 0; i < triggers.length; i += 1) {
      var trigger = triggers[i];
      var isActive = trigger.getAttribute("data-tabs-value") === value;
      trigger.setAttribute("aria-selected", isActive ? "true" : "false");
      trigger.setAttribute("tabindex", isActive ? "0" : "-1");
      if (isActive && useFocus && typeof trigger.focus === "function") {
        trigger.focus();
      }
    }

    for (i = 0; i < panels.length; i += 1) {
      var panel = panels[i];
      var active = panel.getAttribute("data-tabs-value") === value;
      panel.hidden = !active;
    }
  }

  function defaultValue(root) {
    var active = root.getAttribute("data-tabs-value");
    if (active) {
      return active;
    }
    var selected = root.querySelector('[data-tabs-trigger][aria-selected="true"]');
    if (selected && selected.getAttribute("data-tabs-value")) {
      return selected.getAttribute("data-tabs-value");
    }
    var first = root.querySelector('[data-tabs-trigger]');
    return first ? first.getAttribute("data-tabs-value") : "";
  }

  function onKeydown(event, root) {
    var trigger = event.target.closest('[data-tabs-trigger]');
    if (!trigger || !root.contains(trigger)) {
      return;
    }
    var triggers = root.querySelectorAll('[data-tabs-trigger]');
    if (!triggers.length) {
      return;
    }
    var list = Array.prototype.slice.call(triggers);
    var index = list.indexOf(trigger);
    if (index === -1) {
      return;
    }
    if (event.key === "ArrowRight" || event.key === "ArrowDown") {
      event.preventDefault();
      var next = (index + 1) % list.length;
      activateTab(root, list[next].getAttribute("data-tabs-value"), true);
    } else if (event.key === "ArrowLeft" || event.key === "ArrowUp") {
      event.preventDefault();
      var previous = (index - 1 + list.length) % list.length;
      activateTab(root, list[previous].getAttribute("data-tabs-value"), true);
    }
  }

  ready(function () {
    var roots = document.querySelectorAll('[data-ui8kit="tabs"]');
    for (var i = 0; i < roots.length; i += 1) {
      var root = roots[i];
      var value = defaultValue(root);
      if (value) {
        activateTab(root, value, false);
      }

      root.addEventListener("click", (function (activeRoot) {
        return function (event) {
        var trigger = event.target.closest('[data-tabs-trigger]');
        if (!trigger || !activeRoot) {
          return;
        }
        var target = trigger.getAttribute("data-tabs-value");
        if (!target) {
          return;
        }
        event.preventDefault();
        activateTab(activeRoot, target, false);
      };
      })(root));

      root.addEventListener("keydown", (function (activeRoot) {
        return function (event) {
          onKeydown(event, activeRoot);
        };
      })(root));
    }
  });

  namespace.tabs = { init: function () {} };
})();
