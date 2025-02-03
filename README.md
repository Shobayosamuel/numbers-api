# Number Classification API

## Overview
A Go-based web service that provides detailed classification and interesting properties of numbers using the Gin web framework.

## Features
- Determine if a number is prime
- Check if a number is perfect
- Identify number properties (even/odd, Armstrong number)
- Calculate digit sum
- Retrieve fun facts about numbers

## Prerequisites
- Go 1.16+
- Gin Web Framework

## Installation
1. Clone the repository
```bash
git clone https://github.com/Shobayosamuel/numbers-api.git
cd numbers-api
```

2. Install dependencies
```bash
go mod tidy
```

## Running the API
```bash
go run main.go
```
Default port: 8080
Configurable via PORT environment variable

## API Endpoint
`GET /api/classify-number?number={value}`

### Response Example
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```

## Configuration
- CORS enabled for all origins
- Supports GET requests
- Gin Release Mode

## Error Handling
- Returns 400 status for invalid inputs
- Provides clear error messages