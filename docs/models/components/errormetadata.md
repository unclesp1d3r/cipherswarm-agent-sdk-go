# ErrorMetadata

Additional metadata about an agent error


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ErrorDate`                                           | [time.Time](https://pkg.go.dev/time#Time)             | :heavy_check_mark:                                    | The date of the error                                 |
| `Other`                                               | [*components.Other](../../models/components/other.md) | :heavy_minus_sign:                                    | Structured error context from the agent               |