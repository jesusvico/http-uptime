# Configuration documentation

The http-uptime config file is a YAML. It is read once at the start of the application. It contains a list with all the endpoits that http-uptime will monitor.

## Example

```yaml
endpoints:
    - name: "Example endpoint"
      url: "https://example.com"
      method: "GET" 
```

## Structure

### endpoints

A list of endpoints to monitor. Each endpoint has the following fields:

1. **`name`**: Unique identifier of the endpoint.
2. **`url`**: The full URL of the endpoint.
3. **`method`**: The HTTP method used to monitor the endpoint. Can be `GET`, `POST`, `PUT` and `DELETE` (Default: `GET`).

Every endpoint must have a unique `name`. The `url` is mandatory and the default `method` is GET.