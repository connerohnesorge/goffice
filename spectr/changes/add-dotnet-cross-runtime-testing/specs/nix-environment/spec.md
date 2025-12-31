# Nix Environment Specification

**Related Documents**:
- Full proposal context: See `../../proposal.md`
- Implementation design: See `../../design.md` (section: Nix Integration)
- Task breakdown: See `../../tasks.md` (Phase 1: Nix Environment Setup)
- E2E testing spec: See `../e2e-testing/spec.md`

## ADDED Requirements

### Requirement: Multi-Runtime Development Shell
The system SHALL provide a Nix development shell with both Go and .NET SDK environments for cross-runtime testing.

#### Scenario: Enter shell with all tools
- GIVEN flake.nix is configured
- WHEN developer runs `nix develop`
- THEN shell activates with Go 1.25, .NET SDK 9, and all test tools

#### Scenario: .NET SDK available in shell
- GIVEN developer is in Nix shell
- WHEN `dotnet --version` is executed
- THEN .NET 9.0 version is displayed

#### Scenario: LibreOffice available for conversion
- GIVEN developer is in Nix shell
- WHEN `soffice --version` is executed
- THEN LibreOffice version is displayed

#### Scenario: Ghostscript available for rendering
- GIVEN developer is in Nix shell
- WHEN `gs --version` is executed
- THEN Ghostscript version is displayed

---

### Requirement: Test Execution Scripts
The system SHALL provide Nix shell scripts for running different test suites.

#### Scenario: Run Go tests only
- GIVEN developer is in Nix shell
- WHEN `test-go` script is executed
- THEN gotestsum runs Go unit tests
- AND test results are displayed

#### Scenario: Run .NET tests only
- GIVEN developer is in Nix shell
- WHEN `test-dotnet` script is executed
- THEN `dotnet test` runs on Open-XML-SDK
- AND test results are displayed

#### Scenario: Run E2E tests
- GIVEN developer is in Nix shell
- WHEN `test-e2e` script is executed
- THEN C# bridge is built with `dotnet build`
- AND E2E test runner is invoked with Go tests
- AND test results are displayed

#### Scenario: Run all tests
- GIVEN developer is in Nix shell
- WHEN `test-all` script is executed
- THEN Go tests run
- AND .NET tests run
- AND E2E tests run
- AND summary shows results from all suites

#### Scenario: Update baselines
- GIVEN developer is in Nix shell
- WHEN `update-baselines` script is executed
- THEN C# bridge is built
- AND baseline generation runs
- AND git diff command is suggested for review

---

### Requirement: Reproducible Builds
The system SHALL pin all dependency versions for reproducible cross-runtime testing.

#### Scenario: .NET SDK version pinned
- GIVEN flake.lock is committed
- WHEN `nix develop` runs on different machine
- THEN same .NET SDK version (9.0) is provided

#### Scenario: Flake lock updates
- GIVEN developer runs `nix flake update`
- WHEN flake.lock is regenerated
- THEN all package versions are updated consistently

#### Scenario: direnv auto-activation
- GIVEN .envrc is configured with `use flake`
- WHEN developer enters project directory
- THEN Nix shell activates automatically
- AND all tools become available in PATH

---

### Requirement: C# Project Build Integration
The system SHALL support building C# projects within the Nix environment.

#### Scenario: Build C# bridge project
- GIVEN C# bridge project at tests/e2e/bridges/csharp/DocxBridge/
- WHEN `dotnet build` is executed in Nix shell
- THEN project builds successfully
- AND executable is created at bin/Release/net9.0/DocxBridge

#### Scenario: Restore NuGet packages
- GIVEN C# project with package references
- WHEN `dotnet restore` is executed
- THEN NuGet packages are downloaded and restored

#### Scenario: C# project references Open-XML-SDK
- GIVEN DocxBridge.csproj with Open-XML-SDK package reference
- WHEN project is built
- THEN Open-XML-SDK assembly is available
- AND bridge can import and use SDK types

---

### Requirement: Cross-Platform Support
The system SHALL support cross-runtime testing on all major platforms via Nix.

#### Scenario: Support x86_64-linux
- GIVEN flake.nix uses eachDefaultSystem
- WHEN Nix shell activates on x86_64-linux
- THEN all tools are available

#### Scenario: Support aarch64-linux
- GIVEN flake.nix uses eachDefaultSystem
- WHEN Nix shell activates on ARM64 Linux
- THEN all tools are available

#### Scenario: Support x86_64-darwin
- GIVEN flake.nix uses eachDefaultSystem
- WHEN Nix shell activates on Intel macOS
- THEN all tools are available

#### Scenario: Support aarch64-darwin
- GIVEN flake.nix uses eachDefaultSystem
- WHEN Nix shell activates on Apple Silicon macOS
- THEN all tools are available

---

### Requirement: Cross-Platform Functional Verification
The system SHALL verify that tools not only exist but function correctly on all platforms.

#### Scenario: Verify dotnet build succeeds on Linux
- GIVEN Nix shell on Linux
- WHEN `dotnet build` is executed on C# bridge project
- THEN build completes successfully
- AND executable is created

#### Scenario: Verify dotnet build succeeds on macOS
- GIVEN Nix shell on macOS
- WHEN `dotnet build` is executed on C# bridge project
- THEN build completes successfully
- AND executable is created

#### Scenario: Verify LibreOffice conversion works on Linux
- GIVEN Nix shell on Linux
- WHEN `soffice --headless --convert-to pdf` is executed
- THEN PDF is created successfully

#### Scenario: Verify LibreOffice conversion works on macOS
- GIVEN Nix shell on macOS
- WHEN `soffice --headless --convert-to pdf` is executed
- THEN PDF is created successfully

#### Scenario: Verify Ghostscript rendering works cross-platform
- GIVEN Nix shell on any platform
- WHEN `gs -sDEVICE=png16m` is executed to convert PDF to PNG
- THEN PNG is created successfully
