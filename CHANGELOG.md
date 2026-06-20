# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [3.0.0] - 2026-04-13

### Added
- Functional options pattern for `Generate` method (`WithUF` for VoterRegistration).
- Validation for identical-digit documents (e.g., `111.111.111-11` for CPF, `11.111.111/1111-11` for CNPJ).
- `internal/digit` package with pre-compiled regex and `strings.Builder` optimizations.
- Benchmarks for all document types (validation, formatting, generation).
- `Document` interface with `Option` type alias for cross-package compatibility.
- Exported document variables (`brazilcode.CPF`, `brazilcode.CNPJ`, etc.) for direct use.
- `ErrDocTypeNotSupported` exported error for better error handling.
- Sub-package usage: each document type can be imported individually.
- Migration guide in README.

### Changed
- **BREAKING**: Module path changed from `v2` to `v3`.
- **BREAKING**: `Generate` signature changed from `Generate(params map[string]string)` to `Generate(opts ...Option)`.
- **BREAKING**: Removed `src/` directory — packages moved to root-level sub-packages (`cpf/`, `cnpj/`, `cnh/`, `renavam/`, `voter/`).
- **BREAKING**: Removed `iface` package — `Document` interface now lives in root `brazilcode` package.
- **BREAKING**: `voterRegistration` package renamed to `voter`.
- Replaced `regexp.MustCompile` per-call with pre-compiled package-level regex.
- Replaced string concatenation (`+=`) with `strings.Builder` for document generation.
- Replaced `fmt.Sprintf("%d", ...)` with `strconv.Itoa(...)` for digit conversion.
- Replaced `rand.Seed` (deprecated) — uses global `math/rand` source.
- Consolidated utility functions into `internal/digit` (unexported to consumers).
- CNH `Generate` now retries on rare edge cases where check digits exceed single digit.

### Removed
- `src/` directory (Go anti-pattern).
- `iface` package (unnecessary indirection).
- `utils` package (consolidated into `internal/digit`).
- Debug `fmt.Println` in voter registration generation.

### Fixed
- CPF/CNPJ validation now correctly rejects documents with all identical digits.
- CNH `Generate` no longer silently produces invalid documents.
- Voter registration `Generate` now selects UFs with uniform distribution.

## [2.1.0] - 2025-03-31

### Added
- Implement RENAVAM validation and generation functions.
- Add RENAVAM instance for vehicle registration management.
- Add unit tests for RENAVAM validation, formatting, and generation.
- Add unit tests for Document interface using a mock implementation.
- Add documentation for CPF, CNPJ, CNH, and VoterRegistration.

### Changed

- Update README to reflect module path change to v2.
- Update README to include RENAVAM in the list of features.
- Update CHANGELOG with new RENAVAM features and unit tests.

## [2.0.0] - 2025-03-26

### Added
- `Document` interface for document operations.
- `IsValid`, `Format`, and `Generate` methods for document validation and formatting.

### Changed
- Converted voter registration functions to methods of `VoterRegistration` struct implementing the `Document` interface.
- Converted CPF functions to methods of `CPF` struct implementing the `Document` interface.
- Converted CNPJ functions to methods of `CNPJ` struct implementing the `Document` interface.
- Converted CNH functions to methods of `CNH` struct implementing the `Document` interface.
- Consolidated document operations into a unified interface with `IsValid`, `Format`, and `Generate` methods.
- Updated README to reflect method changes for `CNPJ` and `VoterRegistration`.
- Enhanced clarity on package functionality and usage examples in README.

### Fixed
- Updated voter registration test to ignore voter output.