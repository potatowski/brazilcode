# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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