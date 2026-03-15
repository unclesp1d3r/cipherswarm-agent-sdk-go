# Agents

## Overview

Agents API

### Available Operations

* [GetAgent](#getagent) - Gets an instance of an agent
* [UpdateAgent](#updateagent) - Updates the agent
* [SendHeartbeat](#sendheartbeat) - Send a heartbeat for an agent
* [SubmitBenchmark](#submitbenchmark) - submit agent benchmarks
* [SubmitErrorAgent](#submiterroragent) - Submit an error for an agent
* [SetAgentShutdown](#setagentshutdown) - shutdown agent

## GetAgent

Returns an agent

### Example Usage

<!-- UsageSnippet language="go" operationID="getAgent" method="get" path="/api/v1/client/agents/{id}" -->
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

    res, err := s.Agents.GetAgent(ctx, 963394)
    if err != nil {
        log.Fatal(err)
    }
    if res.Agent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `int64`                                                  | :heavy_check_mark:                                       | id                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAgentResponse](../../models/operations/getagentresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## UpdateAgent

Updates an agent

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAgent" method="put" path="/api/v1/client/agents/{id}" -->
```go
package main

import(
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Agents.UpdateAgent(ctx, 2843, operations.UpdateAgentRequestBody{
        ID: 2843,
        HostName: "nautical-produce.com",
        ClientSignature: "<value>",
        OperatingSystem: "BeOS",
        Devices: []string{
            "<value 1>",
            "<value 2>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Agent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `id`                                                                                   | `int64`                                                                                | :heavy_check_mark:                                                                     | id                                                                                     |
| `requestBody`                                                                          | [operations.UpdateAgentRequestBody](../../models/operations/updateagentrequestbody.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateAgentResponse](../../models/operations/updateagentresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## SendHeartbeat

Send a heartbeat for an agent to keep it alive. Optionally accepts an 'activity' parameter in the request body to track the agent's current activity state.

### Example Usage

<!-- UsageSnippet language="go" operationID="sendHeartbeat" method="post" path="/api/v1/client/agents/{id}/heartbeat" -->
```go
package main

import(
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Agents.SendHeartbeat(ctx, 379816, &operations.SendHeartbeatRequestBody{
        Activity: cipherswarmagentsdkgo.Pointer("cracking"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `id`                                                                                        | `int64`                                                                                     | :heavy_check_mark:                                                                          | id                                                                                          |
| `requestBody`                                                                               | [*operations.SendHeartbeatRequestBody](../../models/operations/sendheartbeatrequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.SendHeartbeatResponse](../../models/operations/sendheartbeatresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## SubmitBenchmark

Submit a benchmark for an agent

### Example Usage

<!-- UsageSnippet language="go" operationID="submitBenchmark" method="post" path="/api/v1/client/agents/{id}/submit_benchmark" -->
```go
package main

import(
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Agents.SubmitBenchmark(ctx, 981111, operations.SubmitBenchmarkRequestBody{
        HashcatBenchmarks: []components.HashcatBenchmark{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `id`                                                                                           | `int64`                                                                                        | :heavy_check_mark:                                                                             | id                                                                                             |
| `requestBody`                                                                                  | [operations.SubmitBenchmarkRequestBody](../../models/operations/submitbenchmarkrequestbody.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.SubmitBenchmarkResponse](../../models/operations/submitbenchmarkresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 400, 401, 422         | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## SubmitErrorAgent

Submit an error for an agent

### Example Usage

<!-- UsageSnippet language="go" operationID="submitErrorAgent" method="post" path="/api/v1/client/agents/{id}/submit_error" -->
```go
package main

import(
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Agents.SubmitErrorAgent(ctx, 649742, operations.SubmitErrorAgentRequestBody{
        Message: "<value>",
        Severity: operations.SeverityMajor,
        AgentID: 600242,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `id`                                                                                             | `int64`                                                                                          | :heavy_check_mark:                                                                               | id                                                                                               |
| `requestBody`                                                                                    | [operations.SubmitErrorAgentRequestBody](../../models/operations/submiterroragentrequestbody.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.SubmitErrorAgentResponse](../../models/operations/submiterroragentresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## SetAgentShutdown

Marks the agent as shutdown and offline, freeing any assigned tasks.

### Example Usage

<!-- UsageSnippet language="go" operationID="setAgentShutdown" method="post" path="/api/v1/client/agents/{id}/shutdown" -->
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

    res, err := s.Agents.SetAgentShutdown(ctx, 709306)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `int64`                                                  | :heavy_check_mark:                                       | id                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SetAgentShutdownResponse](../../models/operations/setagentshutdownresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |