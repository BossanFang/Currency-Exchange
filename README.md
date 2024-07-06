# Currency Exchange API

This project provides a currency exchange API that allows users to convert amounts between different currencies. It is built using Go and the Gin web framework, and it includes input validation middleware and unit tests.

## Features

- Convert amounts between TWD, JPY, and USD.
- Input validation middleware to ensure correct parameters.
- Returns converted amount formatted with comma as thousands separator and rounded to two decimal places.
- Includes unit tests for validation and conversion logic.

## Project Structure

Currency-Exchange/
├── main.go
├── middleware/
│   └── validate.go
└── service/
    └── currency_exchange.go
└── main_test.go
