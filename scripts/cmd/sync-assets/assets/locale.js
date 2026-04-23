(function () {
  var namespace = window.ui8kit || {};
  window.ui8kit = namespace;
  if (namespace.languageSwitch) {
    return;
  }

  function ready(fn) {
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", fn);
      return;
    }
    fn();
  }

  function parseResponse(html) {
    var parser = new DOMParser();
    return parser.parseFromString(html, "text/html");
  }

  function replaceMainContent(button, html) {
    var targetSelector = button.getAttribute("data-spa-target") || "main";
    var parsed = parseResponse(html);

    var currentTarget = document.querySelector(targetSelector);
    var nextTarget = parsed.querySelector(targetSelector);
    if (currentTarget && nextTarget) {
      currentTarget.innerHTML = nextTarget.innerHTML;
    }

    var parsedTitle = parsed.querySelector("title");
    if (parsedTitle && parsedTitle.textContent) {
      document.title = parsedTitle.textContent;
    }

    var nextButton = button.id ? parsed.getElementById(button.id) : null;
    if (nextButton) {
      if (nextButton.getAttribute("href")) {
        button.setAttribute("href", nextButton.getAttribute("href"));
      }
      if (nextButton.dataset.currentLocale) {
        button.dataset.currentLocale = nextButton.dataset.currentLocale;
      }
      if (nextButton.dataset.nextLocale) {
        button.dataset.nextLocale = nextButton.dataset.nextLocale;
      }
      button.textContent = nextButton.textContent;
    }

    var locale = parsed.documentElement && parsed.documentElement.getAttribute("lang");
    if (locale) {
      document.documentElement.setAttribute("lang", locale);
    }
  }

  function bindLanguageSwitch(button) {
    button.addEventListener("click", function (event) {
      event.preventDefault();
      var href = button.getAttribute("href");
      if (!href) {
        return;
      }

      fetch(href, {
        credentials: "same-origin",
        headers: {
          "X-Locale-Switch": "1",
        },
      })
        .then(function (response) {
          if (!response.ok) {
            throw new Error("locale switch request failed");
          }
          return response.text();
        })
        .then(function (html) {
          replaceMainContent(button, html);
          history.pushState({}, "", href);
        })
        .catch(function () {
          window.location.href = href;
        });
    });
  }

  ready(function () {
    var toggles = document.querySelectorAll("[data-ui8kit-spa-lang]");
    for (var i = 0; i < toggles.length; i += 1) {
      bindLanguageSwitch(toggles[i]);
    }
  });

  namespace.languageSwitch = { init: function () {} };
})();
