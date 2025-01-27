# HTTP Uptime

HTTP Uptime is a tool designed to monitor the uptime of your websites and APIs. It periodically checks the availability of your endpoints. The application exposes a `/metrics` endpoint for Prometheus. You can configure Prometheus to scrape this endpoint to collect metrics about the uptime and performance of your monitored endpoints.

## Features

- Monitor multiple endpoints
- Expose metrics to prometheus

## Installation

To install HTTP Uptime, clone the repository and install the dependencies:

```bash
git clone https://github.com/yourusername/http-uptime.git
```

## Configuration

[More information](docs/config.md) about how to create a `config.yaml` file.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or suggestions, please open an issue.