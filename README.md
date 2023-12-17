# PlayStation Portal Scraper

## Overview

The PlayStation Portal Scraper is a Go-based tool designed to monitor the availability of products on the PlayStation Direct website. It checks the page for the "Currently Unavailable" label and notifies the user through a desktop notification when the product appears to be available.

## Prerequisites

-   Go (1.15 or later)
-   Docker
-   Basic knowledge of Go programming and Docker

## Setup

1. **Clone the Repository**

    Clone this repository to your local machine:

    ```bash
    git clone https://github.com/yourusername/playstation-portal-scraper.git
    cd playstation-portal-scraper
    ```

2. **Build the Docker Image**

    Build the Docker image using the provided Dockerfile:

    ```bash
    docker build -t playstation-scraper .
    ```

## Running the Scraper

You can run the scraper either directly as a Go program or inside a Docker container.

### Option 1: Running as a Go Program

To run the scraper directly, execute:

```bash
go run main.go
```

### Option 2: Running in Docker

To run the scraper in a Docker container, execute:

```bash
docker run --rm playstation-scraper
```

The scraper is set to run automatically every hour to check the product availability.

## Customization

You can customize the scraper by editing the `main.go` file. For instance, you can change the URL to monitor a different product or modify the notification logic.

## Troubleshooting

-   Ensure that Go and Docker are correctly installed and configured on your system.
-   If the scraper does not send notifications, check that your system allows desktop notifications from applications.

## Contributing

Contributions to the PlayStation Portal Scraper are welcome. Please ensure that your code adheres to the existing style and that all tests pass.

## License

This project is licensed under the [MIT License](LICENSE).
