# Security Policy

## Reporting a Vulnerability

**Do not open a public issue.** Disclose security vulnerabilities privately via
GitHub's built-in vulnerability reporting:

1. Go to the [Security Advisories](https://github.com/Muxcore-Media/contracts-metadata/security/advisories) tab
2. Click **New draft security advisory**
3. Fill in the affected version and a clear description of the issue

You will receive an acknowledgement within **72 hours**.

### What to Include

- Which contract interface(s) are affected
- How the interface design enables privilege escalation, data leakage, or capability bypass
- A concrete example of how a module implementing this contract could be exploited
- Whether you plan to disclose publicly and your preferred timeline

## Scope

This repository contains **interface definitions only** — there is no runtime code.
The security boundary is the contract design itself.

### In Scope

- **Capability leakage** — interface methods that could allow a module to exceed its declared capabilities
- **Contract ambiguity** — underspecified interfaces that could lead to insecure implementations
- **Breaking changes** — silent semantic changes that weaken security guarantees
- **AuthZ/AuthN gaps** — missing identity propagation or authorization hooks in the interface

### Out of Scope

- Implementation bugs in modules that implement these contracts (report to the module repo)
- Issues in MuxCore core fabric (report to the [core repo](https://github.com/Muxcore-Media/core/security/advisories))
- Infrastructure not controlled by MuxCore
- Social engineering, phishing, or physical attacks

## Security Model

MuxCore's security boundary is **between modules**. Contract repos define the
interfaces that sit at those boundaries. A well-designed contract prevents
capability confusion; a poorly-designed one can create privilege escalation
paths even when every module is implemented correctly.

When reporting, focus on the interface design — not hypothetical implementation
bugs. The question is: *does this contract, as specified, allow a module to do
something it shouldn't?*

## Safe Harbor

We will not pursue legal action against security researchers who act in good
faith and follow this policy's reporting and disclosure process.
