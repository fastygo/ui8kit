(function () {
  function ready(fn) {
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", fn);
      return;
    }
    fn();
  }

  function normalizeLocale(value, allowedLocales) {
    var locale = (value || "").toLowerCase();
    return allowedLocales.indexOf(locale) !== -1 ? locale : "";
  }

  function parseLocales(raw) {
    return (raw || "")
      .split(",")
      .map(function (item) {
        return item.trim().toLowerCase();
      })
      .filter(function (item, index, list) {
        return item && list.indexOf(item) === index;
      });
  }

  function browserLocales() {
    var values = [];

    if (Array.isArray(navigator.languages)) {
      values = values.concat(navigator.languages);
    }

    if (typeof navigator.language === "string") {
      values.push(navigator.language);
    }

    return values
      .map(function (item) {
        return String(item || "")
          .trim()
          .toLowerCase()
          .replace(/_/g, "-");
      })
      .filter(function (item, index, list) {
        return item && list.indexOf(item) === index;
      });
  }

  function detectPreferredLocale(allowedLocales, defaultLocale) {
    var locales = browserLocales();

    for (var i = 0; i < locales.length; i += 1) {
      var locale = locales[i];
      if (allowedLocales.indexOf(locale) !== -1) {
        return locale;
      }

      var baseLocale = locale.split("-")[0];
      if (allowedLocales.indexOf(baseLocale) !== -1) {
        return baseLocale;
      }
    }

    return defaultLocale;
  }

  function buttonDefaultLocale(toggle) {
    var available = parseLocales(toggle.dataset.availableLocales);
    return normalizeLocale(toggle.dataset.defaultLocale, available) || available[0] || "";
  }

  function buttonCurrentLocale(toggle, availableLocales) {
    return normalizeLocale(toggle.dataset.currentLocale, availableLocales) || buttonDefaultLocale(toggle);
  }

  function buttonNextLocale(toggle, currentLocale) {
    var availableLocales = parseLocales(toggle.dataset.availableLocales);
    var next = currentLocale;

    for (var i = 0; i < availableLocales.length; i += 1) {
      if (availableLocales[i] !== currentLocale) {
        next = availableLocales[i];
        break;
      }
    }

    return next;
  }

  function readStoredLocale(toggle) {
    var availableLocales = parseLocales(toggle.dataset.availableLocales);
    try {
      return normalizeLocale(localStorage.getItem("framework-language"), availableLocales);
    } catch (_) {
      return "";
    }
  }

  function writeStoredLocale(locale) {
    try {
      localStorage.setItem("framework-language", locale);
    } catch (_) {}
  }

  function localeURL(toggle, locale) {
    var next = new URL(window.location.href);
    if (locale === buttonDefaultLocale(toggle)) {
      next.searchParams.delete("lang");
    } else {
      next.searchParams.set("lang", locale);
    }
    return next.toString();
  }

  ready(function () {
    var toggle = document.getElementById("web-language-toggle");

    if (!toggle) {
      return;
    }

    var availableLocales = parseLocales(toggle.dataset.availableLocales);
    var defaultLocale = buttonDefaultLocale(toggle);
    var currentLocale = buttonCurrentLocale(toggle, availableLocales);
    var hasExplicitLocaleParam = new URL(window.location.href).searchParams.has("lang");
    var storedLocale = readStoredLocale(toggle);

    if (hasExplicitLocaleParam) {
      writeStoredLocale(currentLocale);
    } else {
      var preferredLocale = storedLocale || detectPreferredLocale(availableLocales, defaultLocale);
      if (preferredLocale) {
        writeStoredLocale(preferredLocale);
      }

      if (
        preferredLocale &&
        preferredLocale !== defaultLocale &&
        preferredLocale !== currentLocale
      ) {
        var targetURL = localeURL(toggle, preferredLocale);
        if (targetURL !== window.location.href) {
          window.location.replace(targetURL);
          return;
        }
      }

      if (preferredLocale) {
        currentLocale = preferredLocale;
      }
    }

    toggle.addEventListener("click", function () {
      var nextLocale = buttonNextLocale(toggle, currentLocale);
      if (!nextLocale || nextLocale === currentLocale) {
        return;
      }
      writeStoredLocale(nextLocale);
      window.location.assign(localeURL(toggle, nextLocale));
    });
  });
})();
