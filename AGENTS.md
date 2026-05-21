# ShadowCoreOS Agent Instructions

You are working on ShadowCoreOS — The Cybersecurity Research OS.

## Never Do

- Do not create fake findings.
- Do not create fake evidence.
- Do not mark AI output as validated.
- Do not allow scans without scope.
- Do not allow scans without authorization.
- Do not bypass RBAC.
- Do not run security tools directly on host.
- Do not hardcode secrets.
- Do not create duplicate modules.
- Do not create random folders.
- Do not auto-enable unknown GitHub tools.
- Do not generate reports from unvalidated findings.

## Always Do

- Use free/open-source first.
- Use Ollama local AI by default.
- Use Docker-only tool execution.
- Use validation-first finding flow.
- Store evidence before reporting.
- Enforce scope and authorization.
- Keep modules clean and isolated.
- Add docs for every major system.

## Core Rule

No permission = no scan.  
No scope = no scan.  
No evidence = no finding.  
No validation = no report.
