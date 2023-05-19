# strongo-validation

Validation errors & helpers.

Purpose of this class to define a generic set of validation errors.

- Validation error
    - Bad request
        - Missing field
        - Invalid field value
    - Bad record
        - Missing field
        - Invalid field value
    - Bad field
        - Missing field
        - Invalid field value

You can check if an `err error` is a validation error or a specific validation error with next functions:

- `IsValidationError(err error) bool`
- `IsBadRequestError(err error) bool`
- `IsBadRecordError(err error) bool`
- `IsBadFieldValueError(err error) bool`

## Documentation

https://pkg.go.dev/github.com/strongo/validation

## Usage

This package is known to be used in next open source projects:

- https://github.com/datatug/datatug-agent-go

Please submit a pull request to add your project here if you use this package in an open source project.

## Versioning
The version is auto-incremented by CI/CD pipeline on push to main branch.

## LICENSE

MIT License
