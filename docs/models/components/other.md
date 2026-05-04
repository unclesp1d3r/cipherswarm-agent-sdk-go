# Other

Structured error context from the agent


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Category`                                         | `*string`                                          | :heavy_minus_sign:                                 | Error category (e.g. hash_format, device, network) |
| `Retryable`                                        | `*bool`                                            | :heavy_minus_sign:                                 | Whether the operation can be retried               |
| `ErrorType`                                        | `*string`                                          | :heavy_minus_sign:                                 | Machine-readable error type slug                   |
| `Terminal`                                         | `*bool`                                            | :heavy_minus_sign:                                 | Whether the error is definitively unrecoverable    |
| `AffectedCount`                                    | `*int64`                                           | :heavy_minus_sign:                                 | Number of hashes affected by the error             |
| `TotalCount`                                       | `*int64`                                           | :heavy_minus_sign:                                 | Total number of hashes in the hash list            |
| `AdditionalProperties`                             | map[string]`any`                                   | :heavy_minus_sign:                                 | N/A                                                |