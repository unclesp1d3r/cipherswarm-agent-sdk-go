# RecommendedTimeouts

Recommended timeout settings for agent HTTP connections


## Fields

| Field                              | Type                               | Required                           | Description                        |
| ---------------------------------- | ---------------------------------- | ---------------------------------- | ---------------------------------- |
| `ConnectTimeout`                   | `int64`                            | :heavy_check_mark:                 | TCP connect timeout in seconds     |
| `ReadTimeout`                      | `int64`                            | :heavy_check_mark:                 | Read timeout in seconds            |
| `WriteTimeout`                     | `int64`                            | :heavy_check_mark:                 | Write timeout in seconds           |
| `RequestTimeout`                   | `int64`                            | :heavy_check_mark:                 | Overall request timeout in seconds |