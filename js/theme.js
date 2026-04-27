(function () {
  var root = document.documentElement;
  var themeStorageKey = "ui8kit-theme";

  function ready(fn) {
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", fn);
      return;
    }
    fn();
  }

  function readStoredTheme() {
    try {
      return localStorage.getItem(themeStorageKey);
    } catch (_) {
      return null;
    }
  }

  function writeStoredTheme(value) {
    try {
      localStorage.setItem(themeStorageKey, value);
    } catch (_) {}
  }

  function resolvePreferredTheme() {
    var storedTheme = readStoredTheme();
    if (storedTheme === "dark" || storedTheme === "light") {
      return storedTheme;
    }

    var prefersDark =
      window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches;
    return prefersDark ? "dark" : "light";
  }

  function applyTheme(theme) {
    root.classList.toggle("dark", theme === "dark");
  }

  function ensureThemeIcon(button) {
    var icon = document.getElementById("theme-toggle-icon");
    if (!icon && button) {
      icon = document.createElement("span");
      icon.id = "theme-toggle-icon";
      icon.setAttribute("aria-hidden", "true");
      button.appendChild(icon);
    }
    return icon;
  }

  function applyThemeButtonState() {
    var button = document.getElementById("ui8kit-theme-toggle");
    var icon = ensureThemeIcon(button);
    var dark = root.classList.contains("dark");
    var switchToDark =
      button && button.dataset.switchToDarkLabel
        ? button.dataset.switchToDarkLabel
        : "Switch to dark theme";
    var switchToLight =
      button && button.dataset.switchToLightLabel
        ? button.dataset.switchToLightLabel
        : "Switch to light theme";

    if (icon) {
      icon.className = dark
        ? "ui-theme-icon latty latty-sun"
        : "ui-theme-icon latty latty-moon";
    }

    if (button) {
      button.setAttribute("aria-pressed", dark ? "true" : "false");
      button.setAttribute("title", dark ? switchToLight : switchToDark);
      button.setAttribute("aria-label", dark ? switchToLight : switchToDark);
    }
  }

  applyTheme(resolvePreferredTheme());

  ready(function () {
    var themeButton = document.getElementById("ui8kit-theme-toggle");

    if (themeButton) {
      themeButton.addEventListener("click", function () {
        var nextTheme = root.classList.contains("dark") ? "light" : "dark";
        applyTheme(nextTheme);
        writeStoredTheme(nextTheme);
        applyThemeButtonState();
      });
    }

    applyThemeButtonState();
  });
})();
