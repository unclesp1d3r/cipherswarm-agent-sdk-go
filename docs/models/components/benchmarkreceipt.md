# BenchmarkReceipt

Receipt returned after benchmark submission indicating how many entries were accepted


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ReceivedCount`                                                 | `int64`                                                         | :heavy_check_mark:                                              | Number of benchmark entries received in the request             |
| `ProcessedCount`                                                | `int64`                                                         | :heavy_check_mark:                                              | Number of benchmark entries successfully stored                 |
| `FailedCount`                                                   | `int64`                                                         | :heavy_check_mark:                                              | Number of benchmark entries rejected due to validation failures |
| `Message`                                                       | `*string`                                                       | :heavy_minus_sign:                                              | Human-readable summary of the submission result                 |