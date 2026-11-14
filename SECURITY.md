# Security Policy

## Supported versions

Only the latest release of the current major version (`v3`) receives security fixes.

## Reporting a vulnerability

Please **do not open a public issue** for a security problem.

Report it privately through GitHub: open the repository's **Security** tab and choose **Report a vulnerability**. Include:

- the affected function and version;
- an input that reproduces the problem;
- what goes wrong (for example, a panic, excessive CPU or memory use, or an invalid number accepted as valid).

You should get a first answer within 7 days. Once a fix is released, the advisory is published and you are credited unless you ask otherwise.

## Scope

In scope:

- inputs that make a function panic or consume excessive CPU or memory;
- validation that accepts invalid document numbers or rejects valid ones in a way that can be abused;
- generated numbers that are predictable when they are expected to be random.

Out of scope:

- using generated numbers as real documents: they are meant for tests and examples only;
- checking whether a document actually exists or belongs to someone, which this package does not do.
