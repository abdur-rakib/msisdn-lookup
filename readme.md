# MSISDN Lookup REST API

This is a REST API project built to perform MSISDN (Mobile Station International Subscriber Directory Number) lookups from 20 Redis sets. Each set contains a massive amount of data (approximately 20 million records). The project is developed using Golang, the Fiber framework, and Docker for containerization.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)

## Features

- MSISDN lookup from 20 Redis sets.
- High-performance REST API using Golang and Fiber.
- Containerized using Docker for easy deployment and scaling.

## Prerequisites

Before you get started, ensure you have the following dependencies installed:

- Golang
- Docker

## Installation

1. Clone this repository:

```bash
git clone https://github.com/abdur-rakib/msisdn-lookup.git
cd msisdn-lookup
```

2. Set up your Redis instances with the 20 data sets containing the MSISDN records.

3. Run the project using docker compose:

```bash
docker compose up -d --build 
```

## Usage

Once the API is running, you can make HTTP requests to perform MSISDN lookups. Detailed API documentation can be found in the [API Endpoints](#api-endpoints) section below.

## API Endpoints

- `GET /customer-segments/?msisdn={msisdn}`: Lookup the MSISDN in the Redis sets.

Replace `{msisdn}` with the MSISDN you want to look up.

Example request:

```http
GET /customer-segments/?msisdn=880172601495
```

Example response:

```json
{
    "data": {
        "customer_segment_1": 0,
        "customer_segment_10": 0,
        "customer_segment_11": {
            "0": "bngQs",
            "1": "HojRNK",
            "2": "q4KPA"
        },
        "customer_segment_12": 0,
        "customer_segment_13": {
            "0": "ktQkg",
            "1": "DJSgJ0",
            "2": "KCbKr"
        },
        "customer_segment_14": 0,
        "customer_segment_15": 0,
        "customer_segment_16": {
            "0": "dSiaB",
            "1": "zUecq6",
            "2": "rwMt8"
        },
        "customer_segment_17": 0,
        "customer_segment_18": 0,
        "customer_segment_19": 0,
        "customer_segment_2": 0,
        "customer_segment_20": 0,
        "customer_segment_3": 0,
        "customer_segment_4": {
            "0": "FHH6C",
            "1": "3ybr8L",
            "2": "tvCQz"
        },
        "customer_segment_5": 0,
        "customer_segment_6": 0,
        "customer_segment_7": 0,
        "customer_segment_8": 0,
        "customer_segment_9": 0
    },
    "msisdn": "880172601495"
}
```

## Contributing

If you'd like to contribute to this project, please follow our [Contributing Guidelines](CONTRIBUTING.md).

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

 - [Fiber Recipes](https://github.com/gofiber/recipes)
 - [For Redis Instance](https://github.com/a-h-abid/docker-commons)

