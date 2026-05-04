# RecommendedCircuitBreaker

Recommended circuit breaker settings for agent connections


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `FailureThreshold`                          | `int64`                                     | :heavy_check_mark:                          | Number of failures before circuit opens     |
| `Timeout`                                   | `int64`                                     | :heavy_check_mark:                          | Seconds before circuit half-opens for retry |