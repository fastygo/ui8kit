# Security Policy

## Supported Versions

| Version | Supported |
|---------|-----------|
| latest  | Yes       |
| < latest | No       |

Security fixes are applied to the latest release only. We recommend that all users keep their dependencies up to date.

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a potential security issue in UI8Kit, please report it through GitHub's built-in private vulnerability reporting.

### How to report

1. Navigate to the **Security** tab of this repository.
2. Select **"Report a vulnerability"**.
3. Provide a clear description of the issue, including reproduction steps if applicable.

GitHub's private disclosure mechanism ensures that vulnerability details remain confidential until a fix is available. Do not open public issues, pull requests, or discussions for security-related reports.

### What to include

- Affected package and version.
- Description of the vulnerability and its potential impact.
- Steps to reproduce or a minimal proof of concept.
- Suggested remediation, if known.

### Scope

UI8Kit is a server-side HTML rendering library with no runtime JavaScript framework. The primary attack surface includes:

- HTML injection through unescaped component output.
- CSS injection via UtilityProps or class parameters.
- Path traversal in embedded asset serving (`styles.FS`).

Issues outside the library's control (e.g. misconfigured HTTP servers, Tailwind CSS CDN integrity) are out of scope.

### Disclosure Timeline

Upon receiving a valid report, we aim to:

1. Acknowledge receipt within GitHub's reporting interface.
2. Investigate and develop a fix.
3. Release a patched version.
4. Credit the reporter in the release notes (unless anonymity is requested).

Timelines depend on severity and complexity. Critical vulnerabilities are prioritized.
