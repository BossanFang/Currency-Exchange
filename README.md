# Currency Exchange API

This project provides a currency exchange API that allows users to convert amounts between different currencies. It is built using Go and the Gin web framework, and it includes input validation middleware and unit tests.

## Features

- Convert amounts between TWD, JPY, and USD.
- Input validation middleware to ensure correct parameters.
- Returns converted amount formatted with comma as thousands separator and rounded to two decimal places.
- Includes unit tests for validation and conversion logic.

## Project Structure

```
Currency-Exchange/
├── Dockerfile
├── docker-compose.yml
├── main.go
├── main_test.go
├── README.md
├── middleware/
│   └── validate.go
└── service/
    └── currency_exchange_service.go
```

## Getting Started

### Prerequisites

- Go 1.22 or later
- Docker (optional, for containerized deployment)
- Docker Compose (optional, for containerized deployment)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/BossanFang/Currency-Exchange.git
   cd Currency-Exchange
   ```
2. Install dependencies:

   ```bash
   go mod tidy
   ```

## Running the Application

### Locally

1. Run the application locally:

   ```bash
   go run main.go
   ```

   The API will be available at `http://localhost:8080`.

### Using Docker

1. Build the Docker image:

   ```bash
   docker build -t currency_exchange .
   ```

2. Run the application using Docker:

   ```bash
   docker run -p 8080:8080 currency_exchange
   ```

   The API will be available at `http://localhost:8080`.

### Using Docker Compose

1. Build and run the application using Docker Compose:

   ```bash
   docker-compose up --build
   ```

   The API will be available at `http://localhost:8080`.

**Parameters:**
- `source` (string): The source currency (TWD, JPY, USD).
- `target` (string): The target currency (TWD, JPY, USD).
- `amount` (string): The amount to convert (supports comma as thousands separator).

**Example Request:**

```
GET /exchange?source=USD&target=JPY&amount=1,525
```

**Example Response:**

```json
{
  "msg": "success",
  "amount": "170,496.53"
}
```

### Running Tests

Run the unit tests using the following command:

```bash
go test
```

### License

This project is licensed under the MIT License.

### Contact

If you have any questions, please feel free to reach out.

   
