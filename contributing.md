# Contributing to Holidays Per Country

Thank you for your interest in contributing to this project! We welcome contributions from the community.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [How to Contribute](#how-to-contribute)
- [Adding a New Country Provider](#adding-a-new-country-provider)
- [Development Guidelines](#development-guidelines)
- [Testing](#testing)
- [Submitting Changes](#submitting-changes)

## Code of Conduct

This project adheres to a code of conduct that all contributors are expected to follow. Please be respectful and constructive in your interactions.

## Getting Started

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/holidays-per-country.git
   cd holidays-per-country
   ```
3. Add the upstream repository:
   ```bash
   git remote add upstream https://github.com/bobadilla-tech/holidays-per-country.git
   ```
4. Install Go (version 1.23 or later)
5. Run tests to ensure everything works:
   ```bash
   go test ./...
   ```

## How to Contribute

### Reporting Bugs

- Check if the issue already exists in the [issue tracker](https://github.com/bobadilla-tech/holidays-per-country/issues)
- If not, create a new issue with:
  - Clear description of the bug
  - Steps to reproduce
  - Expected vs actual behavior
  - Go version and operating system

### Suggesting Features

- Open an issue describing the feature
- Explain the use case and benefits
- Be open to discussion and feedback

### Fixing Issues

- Look for issues labeled `good first issue` or `help wanted`
- Comment on the issue to let others know you're working on it
- Follow the development guidelines below

## Adding a New Country Provider

To add support for a new country's holidays:

### 1. Create Provider File

Create a new file in `providers/` named `{countrycode}.go` (lowercase ISO 3166-1 alpha-2 code):

```go
package providers

import (
    "time"
    "github.com/bobadilla-tech/holidays-per-country/providers/internal"
)

type YourCountryProvider struct{}

func (_ YourCountryProvider) RegisterHolidays(year int) []internal.Holiday {
    return []internal.Holiday{
        // Add holidays here
    }
}
```

### 2. Register the Provider

Add your provider to the registry in the appropriate initialization file.

### 3. Implement Holidays

Use helper functions for different holiday types:

**Fixed-date holidays:**
```go
internal.NewHoliday(year, time.January, 1, "New Year's Day", nil)
```

**Movable holidays:**
```go
easter := internal.Easter(year)
internal.NewHolidayFromTime(easter, "Easter Sunday", nil)
```

**Day-of-week based holidays:**
```go
firstMonday := internal.FindDay(year, time.September, time.Monday, 1)
internal.NewHolidayFromTime(firstMonday, "Labor Day", nil)
```

**Subdivision-specific holidays:**
```go
internal.NewHoliday(year, time.June, 24, "Saint-Jean-Baptiste Day", []string{"CA-QC"})
```

### 4. Handle Special Cases

- **Historical changes**: Use conditionals for holidays that changed over time
- **Year-specific rules**: Check year ranges for holidays that were added or removed
- **Subdivisions**: Use ISO 3166-2 codes for state/province-specific holidays

### 5. Add Tests

Create `providers/{countrycode}_test.go`:

```go
package providers

import (
    "testing"
    "time"
)

func TestYourCountryProvider(t *testing.T) {
    provider := YourCountryProvider{}
    holidays := provider.RegisterHolidays(2024)

    // Test that key holidays are present
    // Test subdivision-specific holidays
    // Test year-specific rules
}
```

### 6. Document Sources

Add comments documenting where holiday information came from:

```go
// Holiday data source: https://www.officialgovernmentsite.com/holidays
// Last verified: 2024-01-15
```

## Development Guidelines

### Code Style

- Follow standard Go conventions
- Run `gofmt` before committing
- Use meaningful variable and function names
- Add comments for complex logic

### Commit Messages

Use clear, descriptive commit messages:

```
feat: add holidays for Portugal
fix: correct Easter calculation for leap years
docs: update README with subdivision codes
refactor: extract common holiday calculations
test: add tests for Chinese lunar calendar
```

Prefix types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `test`: Adding or updating tests
- `refactor`: Code refactoring
- `chore`: Maintenance tasks

### Code Quality

- Write unit tests for new code
- Maintain or improve test coverage
- Ensure all tests pass: `go test ./...`
- Run linter: `golangci-lint run`
- Check for race conditions: `go test -race ./...`

### Helper Functions

Reuse existing helper functions from `providers/internal/`:

- `internal.Easter(year)` - Calculate Easter Sunday
- `internal.GoodFriday(year)` - Good Friday
- `internal.FindDay(year, month, weekday, occurrence)` - Nth weekday of month
- `internal.FindLastDay(year, month, weekday)` - Last weekday of month
- `internal.ChineseLunarToGregorian(year, month, day)` - Lunar to Gregorian conversion

## Testing

### Run All Tests

```bash
go test ./...
```

### Run Tests with Coverage

```bash
go test -cover ./...
```

### Run Tests for Specific Provider

```bash
go test ./providers -run TestUnitedStatesProvider
```

### Generate Coverage Report

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Submitting Changes

1. Create a new branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes and commit:
   ```bash
   git add .
   git commit -m "feat: description of your changes"
   ```

3. Keep your fork updated:
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```

4. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

5. Create a Pull Request:
   - Go to the original repository
   - Click "New Pull Request"
   - Select your fork and branch
   - Fill in the PR template with:
     - Description of changes
     - Related issue (if applicable)
     - Testing performed
     - Breaking changes (if any)

### Pull Request Guidelines

- One feature/fix per PR
- Include tests for new functionality
- Update documentation as needed
- Ensure CI passes (tests, linting, build)
- Respond to review feedback promptly

## Questions?

If you have questions or need help:

- Check existing [issues](https://github.com/bobadilla-tech/holidays-per-country/issues)
- Open a new issue for discussion
- Tag maintainers if needed

Thank you for contributing! 🎉
