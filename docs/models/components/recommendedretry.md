# RecommendedRetry

Recommended retry settings for agent HTTP requests


## Fields

| Field                            | Type                             | Required                         | Description                      |
| -------------------------------- | -------------------------------- | -------------------------------- | -------------------------------- |
| `MaxAttempts`                    | `int64`                          | :heavy_check_mark:               | Maximum number of retry attempts |
| `InitialDelay`                   | `int64`                          | :heavy_check_mark:               | Initial retry delay in seconds   |
| `MaxDelay`                       | `int64`                          | :heavy_check_mark:               | Maximum retry delay in seconds   |