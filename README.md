# PlayStation Portal Scraper

## Overview

The PlayStation Portal Scraper is a Go-based tool designed to monitor the availability of products on the PlayStation Direct website. It checks the page for the "Currently Unavailable" label and notifies the user through a webhook when the product appears to be available.

## Prerequisites

- Go (1.15 or later)
- Docker
- Basic knowledge of Go programming and Docker
- An IFTTT account for setting up the webhook

## Webhook Setup with IFTTT

Before running the scraper, set up a webhook with IFTTT:

1. **Create an IFTTT Account**: If you don't already have one, sign up at [IFTTT](https://ifttt.com/).

2. **Create a New Applet**:
   - Go to "Create" and then click on "If This Then That".
   - For the "This" part, choose "Webhooks" and then "Receive a web request".
   - Name the event (e.g., `portal_available`).
   - For the "That" part, choose the service you want to use for notifications (e.g., Email, SMS).
   - Configure the action to your preference (e.g., set the message text, subject).

3. **Get Your Webhook Key**: Go to the [Webhooks service page](https://ifttt.com/maker_webhooks), click on "Documentation", and note your unique webhook key.

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

To run the scraper in a Docker container, pass the IFTTT webhook key as an environment variable:

```bash
docker run --restart unless-stopped -e IFTTT_WEBHOOK_KEY=your_webhook_key playstation-scraper
```

Replace `your_webhook_key` with your actual IFTTT webhook key. The scraper is set to run automatically every hour to check the product availability.

## Customization

You can customize the scraper by editing the `main.go` file. For instance, you can change the URL to monitor a different product, adjust the notification logic, or change the frequency of checks.

## Troubleshooting

- Ensure that Go and Docker are correctly installed and configured on your system.
- If using Docker, ensure the `IFTTT_WEBHOOK_KEY` environment variable is set correctly.
- If the scraper does not send notifications, check the configuration of your webhook service.

## Contributing

Contributions to the PlayStation Portal Scraper are welcome. Please ensure that your code adheres to the existing style and that all tests pass.
