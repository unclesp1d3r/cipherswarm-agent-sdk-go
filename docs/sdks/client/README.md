# Client

## Overview

Client API

### Available Operations

* [GetHealth](#gethealth) - Health check (unauthenticated)
* [GetConfiguration](#getconfiguration) - Get Agent Configuration
* [Authenticate](#authenticate) - Authenticate Client

## GetHealth

Returns server health status. Does not require authentication. Agents use this endpoint to verify server reachability before attempting authenticated requests.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealth" method="get" path="/api/v1/client/health" -->
```go
package main

import(
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := cipherswarmagentsdkgo.New()

    res, err := s.Client.GetHealth(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetHealthResponse](../../models/operations/gethealthresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.GetHealthResponseBody | 503                             | application/json                |
| sdkerrors.SDKError              | 4XX, 5XX                        | \*/\*                           |

## GetConfiguration

Returns the configuration for the agent.

### Example Usage

<!-- UsageSnippet language="go" operationID="getConfiguration" method="get" path="/api/v1/client/configuration" -->
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

    res, err := s.Client.GetConfiguration(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConfigurationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetConfigurationResponse](../../models/operations/getconfigurationresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## Authenticate

Authenticates the client. This is used to verify that the client is able to connect to the server.

### Example Usage

<!-- UsageSnippet language="go" operationID="authenticate" method="get" path="/api/v1/client/authenticate" -->
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

    res, err := s.Client.Authenticate(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthenticationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AuthenticateResponse](../../models/operations/authenticateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |