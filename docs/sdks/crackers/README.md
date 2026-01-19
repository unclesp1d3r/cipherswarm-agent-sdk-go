# Crackers

## Overview

Crackers API

### Available Operations

* [CheckForCrackerUpdate](#checkforcrackerupdate) - Check for Cracker Update

## CheckForCrackerUpdate

Check for a cracker update, based on the operating system and version.

### Example Usage

<!-- UsageSnippet language="go" operationID="checkForCrackerUpdate" method="get" path="/api/v1/client/crackers/check_for_cracker_update" -->
```go
package main

import(
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Crackers.CheckForCrackerUpdate(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CrackerUpdate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `operatingSystem`                                        | **string*                                                | :heavy_minus_sign:                                       | operating_system                                         |
| `version`                                                | **string*                                                | :heavy_minus_sign:                                       | version                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CheckForCrackerUpdateResponse](../../models/operations/checkforcrackerupdateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 400                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |