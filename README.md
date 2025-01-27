# HTTP Uptime

HTTP Uptime is a tool designed to monitor the uptime of your websites and APIs. It periodically checks the availability of your endpoints. The application exposes a `/metrics` endpoint for Prometheus. You can configure Prometheus to scrape this endpoint to collect metrics about the uptime and performance of your monitored endpoints.

## Features

- Monitor multiple endpoints
- Expose metrics to prometheus

## Installation

We recommend using Docker to run HTTP Uptime. You can pull the Docker image from Docker Hub and run it with the following commands:
```bash
docker pull jesusvico/http-uptime:latest
docker run -d \
  -v /path/to/your/config.yaml:/etc/http-uptime/config.yaml \
  -e PORT=8080 \
  -p 8080:8080 \
  jesusvico/http-uptime:latest
```

You also can use Docker Compose for running the application:
```bash
docker-compose up -d
```

## Configuration

[More information](docs/config.md) about how to create a `config.yaml` file.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or suggestions, please open an issue.
