# Starling Exporter

A Prometheus exporter for [Starling Bank's Developer API](https://developer.starlingbank.com/docs).

## Why is this a thing?

Starling Bank have a developer API, and I felt like hooking it up to my home monitoring system. Why not?

You can use this to aggregate metrics on your account balance and transactions, then monitor them live using Grafana dashboards.

You could even set up custom spending alerts.

## Installation

The Docker image is published to Docker Hub.

It can be run with the following environment variables:

| Name                    | Description                                                   |
| ----------------------- | ------------------------------------------------------------- |
| `STARLING_ACCESS_TOKEN` | Your Starling Personal Access Token                           |
| `STARLING_SANDBOX`      | Use the sandbox environment. Example: `STARLING_SANDBOX=true` |