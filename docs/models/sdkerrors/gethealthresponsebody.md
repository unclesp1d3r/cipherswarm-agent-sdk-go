# GetHealthResponseBody

degraded


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Status`                                                | `string`                                                | :heavy_check_mark:                                      | Overall health status (ok or degraded)                  |
| `APIVersion`                                            | `int64`                                                 | :heavy_check_mark:                                      | API version                                             |
| `Timestamp`                                             | [time.Time](https://pkg.go.dev/time#Time)               | :heavy_check_mark:                                      | Server timestamp                                        |
| `Database`                                              | `string`                                                | :heavy_check_mark:                                      | Database health (healthy or unhealthy)                  |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_minus_sign:                                      | Raw HTTP response; suitable for custom response parsing |