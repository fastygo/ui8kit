(function () {
  var namespace = window.ui8kit || {};
  window.ui8kit = namespace;
  if (namespace.dialog) {
    return;
  }

  var OPEN_ATTR = "data-state";
  var OPEN_VALUE = "open";
  var CLOSED_VALUE = "closed";

  function ready(fn) {
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", fn);
      return;
    }
    fn();
  }

  function toDialog(node) {
    if (!node) {
      return null;
    }
    if (typeof node === "string") {
      return document.querySelector('[data-ui8kit-dialog][id="' + node + '"]');
    }
    if (node.matches('[data-ui8kit="dialog"], [data-ui8kit="sheet"], [data-ui8kit="alertdialog"]')) {
      return node;
    }
    var target = node.getAttribute && node.getAttribute("data-ui8kit-dialog-target");
    if (target) {
      return document.getElementById(target);
    }
  if (node.closest) {
    return node.closest('[data-ui8kit="dialog"], [data-ui8kit="sheet"], [data-ui8kit="alertdialog"]');
  }
    return null;
  }

  function setDialogState(dialog, open) {
    if (!dialog) {
      return;
    }
    var overlay = dialog.querySelector("[data-ui8kit-dialog-overlay]");
    var closeable = open ? true : false;
    dialog.setAttribute(OPEN_ATTR, open ? OPEN_VALUE : CLOSED_VALUE);

    if (open) {
      dialog.removeAttribute("hidden");
      if (overlay) {
        overlay.removeAttribute("hidden");
      }
      dialog.dataset.wasHidden = "true";
      trapFocus(dialog);
    } else {
      dialog.setAttribute("hidden", "hidden");
      if (overlay) {
        overlay.setAttribute("hidden", "hidden");
      }
      if (closeable) {
        releaseFocus(dialog);
      }
      if (dialog.dataset.lastFocus) {
        var last = document.getElementById(dialog.dataset.lastFocus);
        if (last && typeof last.focus === "function") {
          last.focus();
        }
      }
      delete dialog.dataset.lastFocus;
    }
  }

  function openDialog(dialog, trigger) {
    if (!dialog) {
      return;
    }
    if (trigger && trigger.getAttribute) {
      var id = trigger.getAttribute("id");
      if (id) {
        dialog.dataset.lastFocus = id;
      }
    }
    setDialogState(dialog, true);
  }

  function closeDialog(dialog) {
    setDialogState(dialog, false);
  }

  function trapFocus(dialog) {
    if (!dialog || dialog.dataset.trapped === "1") {
      return;
    }
    dialog.dataset.trapped = "1";
    var first = findFocusable(dialog)[0];
    if (first && typeof first.focus === "function") {
      first.focus();
    }
  }

  function releaseFocus(dialog) {
    delete dialog.dataset.trapped;
  }

  function findFocusable(dialog) {
    return dialog.querySelectorAll(
      'button:not([disabled]), [href], input:not([disabled]), select:not([disabled]), textarea:not([disabled]), button:not([disabled]), [tabindex]:not([tabindex="-1"])'
    );
  }

  function handleTabKey(event, dialog) {
    if (event.key !== "Tab" || !dialog || dialog.getAttribute(OPEN_ATTR) !== OPEN_VALUE) {
      return;
    }
    var focusable = findFocusable(dialog);
    if (!focusable.length) {
      event.preventDefault();
      return;
    }
    var first = focusable[0];
    var last = focusable[focusable.length - 1];
    var target = document.activeElement;
    if (event.shiftKey) {
      if (target === first || target === dialog) {
        last.focus();
        event.preventDefault();
      }
    } else if (target === last || target === dialog) {
      first.focus();
      event.preventDefault();
    }
  }

  function isOpen(dialog) {
    return dialog && dialog.getAttribute(OPEN_ATTR) === OPEN_VALUE;
  }

  function onDocumentClick(event) {
    var openButton = event.target.closest("[data-ui8kit-dialog-open], [data-ui8kit-dialog-target]");
    if (openButton && openButton.matches("[data-ui8kit-dialog-open]")) {
      var dialog = toDialog(openButton);
      if (dialog) {
        openDialog(dialog, openButton);
      }
      return;
    }
    var closeButton = event.target.closest("[data-ui8kit-dialog-close]");
    if (closeButton) {
      closeDialog(toDialog(closeButton));
      return;
    }
    if (event.target.closest("[data-ui8kit-dialog-overlay]")) {
      var overlayDialog = event.target.closest("[data-ui8kit-dialog], [data-ui8kit=sheet], [data-ui8kit=alertdialog]");
      closeDialog(overlayDialog);
    }
  }

  function onDocumentKeydown(event) {
    if (event.key === "Escape") {
      var dialogs = document.querySelectorAll('[data-ui8kit="dialog"], [data-ui8kit="sheet"], [data-ui8kit="alertdialog"]');
      for (var i = 0; i < dialogs.length; i += 1) {
        var dialog = dialogs[i];
        if (isOpen(dialog)) {
          closeDialog(dialog);
          event.preventDefault();
          return;
        }
      }
      return;
    }

    var dialogs = document.querySelectorAll('[data-ui8kit="dialog"], [data-ui8kit="sheet"], [data-ui8kit="alertdialog"]');
    for (var i = 0; i < dialogs.length; i += 1) {
      if (isOpen(dialogs[i]) && dialogs[i].contains(event.target)) {
        handleTabKey(event, dialogs[i]);
      }
    }
  }

  ready(function () {
    var dialogs = document.querySelectorAll('[data-ui8kit="dialog"], [data-ui8kit="sheet"], [data-ui8kit="alertdialog"]');
    for (var i = 0; i < dialogs.length; i += 1) {
      setDialogState(dialogs[i], dialogs[i].getAttribute(OPEN_ATTR) === OPEN_VALUE);
    }
    document.addEventListener("click", onDocumentClick);
    document.addEventListener("keydown", onDocumentKeydown);
  });

  namespace.dialog = {
    open: function (id) {
      openDialog(toDialog(id));
    },
    close: function (id) {
      closeDialog(toDialog(id));
    },
  };
  namespace.ready = function (fn) {
    ready(fn);
  };
})();
