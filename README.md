# Starling Exporter

A Prometheus exporter for [Starling Bank's Developer API](https://developer.starlingbank.com/docs).

## Why is this a thing?

I found out Starling Bank have a developer API, and decided to hook it up to my Pi Zero hosted monitoring stack. Why not?

You can use this to aggregate metrics on your account balance and transactions, then monitor them live using Grafana dashboards.

You could even set up custom budget alerts if you wanted to.

## Installation

The Docker image is published to Docker Hub.

```
docker run -p8080:8080 -eSTARLING_ACCESS_TOKEN="my-pat" rornic/starling-exporter:0.1.0
```

## Configuration

### Environment Variables

| Name                    | Description                                                   |
| ----------------------- | ------------------------------------------------------------- |
| `STARLING_ACCESS_TOKEN` | Your Starling Personal Access Token                           |
| `STARLING_SANDBOX`      | Use the sandbox environment. Example: `STARLING_SANDBOX=true` |

### Access Token

Starling Exporter does not currently support OAuth, so you will need to create a Personal Access Token.

Starling apply a quota on PATs of 1000 requests per day. In its default configuration Starling Exporter uses less than that, but if you use the PAT for anything else then beware of rate limiting.

It should have 'Read Personal' and 'Read Financial' scopes.

## Contribution

All and any contributions are welcomed. Please feel free to open issues or PRs with ideas and improvements.

Starling Exporter uses an Apache 2.0 license. You are free to do with this software as you wish.