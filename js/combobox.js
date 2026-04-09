(function () {
  var namespace = window.ui8kit || {};
  window.ui8kit = namespace;
  if (namespace.combobox) {
    return;
  }

  function ready(fn) {
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", fn);
      return;
    }
    fn();
  }

  function hideAll(roots) {
    for (var i = 0; i < roots.length; i += 1) {
      var list = roots[i].querySelector('[role="listbox"], ul');
      if (list) {
        list.setAttribute("hidden", "hidden");
      }
    }
  }

  function optionText(item) {
    return (item.textContent || "").toLowerCase().trim();
  }

  function open(root) {
    var list = root.querySelector('[role="listbox"], ul');
    if (list) {
      list.removeAttribute("hidden");
    }
    var input = root.querySelector('input');
    var trigger = root.querySelector('[data-combobox-toggle]');
    if (trigger) {
      trigger.setAttribute("aria-expanded", "true");
    }
    if (input) {
      input.setAttribute("aria-expanded", "true");
    }
    root.setAttribute("data-state", "open");
  }

  function close(root) {
    var list = root.querySelector('[role="listbox"], ul');
    if (list) {
      list.setAttribute("hidden", "hidden");
    }
    var trigger = root.querySelector('[data-combobox-toggle]');
    if (trigger) {
      trigger.setAttribute("aria-expanded", "false");
    }
    var input = root.querySelector('input');
    if (input) {
      input.setAttribute("aria-expanded", "false");
    }
    root.setAttribute("data-state", "closed");
  }

  function filterOptions(root) {
    var input = root.querySelector('input');
    var options = root.querySelectorAll('[data-combobox-option]');
    var phrase = (input && input.value ? input.value : "").toLowerCase().trim();
    for (var i = 0; i < options.length; i += 1) {
      var option = options[i];
      var text = optionText(option);
      var visible = !phrase || text.indexOf(phrase) >= 0;
      option.style.display = visible ? "" : "none";
    }
  }

  function findFirstVisibleOption(root) {
    var options = root.querySelectorAll('[data-combobox-option]');
    for (var i = 0; i < options.length; i += 1) {
      var option = options[i];
      if (option.style.display !== "none") {
        return option;
      }
    }
    return null;
  }

  function syncFocus(root, option) {
    var options = root.querySelectorAll('[data-combobox-option]');
    for (var i = 0; i < options.length; i += 1) {
      options[i].classList.remove("ui-combobox-option-active");
      options[i].setAttribute("aria-selected", "false");
    }
    if (option) {
      option.classList.add("ui-combobox-option-active");
      option.setAttribute("aria-selected", "true");
      if (typeof option.scrollIntoView === "function") {
        option.scrollIntoView({ block: "nearest" });
      }
    }
  }

  function selectOption(root, option) {
    if (!option || option.getAttribute("aria-disabled") === "true") {
      return;
    }
    var input = root.querySelector('input');
    var value = option.getAttribute("data-combobox-value") || option.textContent || "";
    if (input) {
      input.value = value;
    }
    close(root);
  }

  function findOptionsBelow(root) {
    return root.querySelectorAll('[data-combobox-option]:not([aria-disabled="true"]):not([style*="none"])');
  }

  ready(function () {
    var roots = document.querySelectorAll('[data-ui8kit="combobox"]');
    for (var i = 0; i < roots.length; i += 1) {
      var root = roots[i];
      var input = root.querySelector("input");
      var toggle = root.querySelector('[data-combobox-toggle]');
      var options = root.querySelectorAll('[data-combobox-option]');

      if (!input || !options.length) {
        continue;
      }

      (function (activeRoot, activeInput, activeToggle, activeOptions) {
        activeInput.addEventListener("focus", function () {
          open(activeRoot);
        });
        activeInput.addEventListener("input", function () {
          open(activeRoot);
          filterOptions(activeRoot);
        });
        activeInput.addEventListener("keydown", function (event) {
          if (event.key === "ArrowDown" || event.key === "ArrowUp") {
            var visible = findOptionsBelow(activeRoot);
            if (!visible.length) {
              return;
            }
            var direction = event.key === "ArrowDown" ? 1 : -1;
            var current = activeRoot.querySelector(".ui-combobox-option-active");
            var currentIndex = current ? Array.prototype.indexOf.call(visible, current) : -1;
            if (currentIndex < 0) {
              currentIndex = direction > 0 ? -1 : 0;
            }
            var nextIndex = (currentIndex + direction + visible.length) % visible.length;
            syncFocus(activeRoot, visible[nextIndex]);
            event.preventDefault();
          } else if (event.key === "Enter" && activeRoot.querySelector(".ui-combobox-option-active")) {
            selectOption(activeRoot, activeRoot.querySelector(".ui-combobox-option-active"));
            event.preventDefault();
          } else if (event.key === "Escape") {
            close(activeRoot);
            event.preventDefault();
          }
        });

        for (var j = 0; j < activeOptions.length; j += 1) {
          activeOptions[j].addEventListener("mousedown", function (event) {
            event.preventDefault();
          });
          activeOptions[j].addEventListener("click", function (event) {
            selectOption(activeRoot, event.currentTarget);
          });
        }

        if (activeToggle) {
          activeToggle.addEventListener("click", function () {
            if (activeRoot.getAttribute("data-state") === "open") {
              close(activeRoot);
            } else {
              open(activeRoot);
            }
          });
        }

        activeRoot.dataset.hasBinding = "true";
        filterOptions(activeRoot);
      })(root, input, toggle, options);
    }

    document.addEventListener("click", function (event) {
      var current = event.target.closest('[data-ui8kit="combobox"]');
      if (!current) {
        hideAll(document.querySelectorAll('[data-ui8kit="combobox"]'));
      }
    });
  });

  namespace.combobox = { init: function () {} };
})();
