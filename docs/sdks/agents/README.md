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
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Agents.UpdateAgent(ctx, 2843, components.UpdateAgentRequest{
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `id`                                                                           | `int64`                                                                        | :heavy_check_mark:                                                             | id                                                                             |
| `updateAgentRequest`                                                           | [components.UpdateAgentRequest](../../models/components/updateagentrequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

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
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Agents.SendHeartbeat(ctx, 379816, &components.AgentHeartbeatRequest{
        Activity: cipherswarmagentsdkgo.Pointer("cracking"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HeartbeatResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `id`                                                                                  | `int64`                                                                               | :heavy_check_mark:                                                                    | id                                                                                    |
| `agentHeartbeatRequest`                                                               | [*components.AgentHeartbeatRequest](../../models/components/agentheartbeatrequest.md) | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |

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
	"log"
)

func main() {
    ctx := context.Background()

    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Agents.SubmitBenchmark(ctx, 981111, components.SubmitBenchmarkRequest{
        HashcatBenchmarks: []components.HashcatBenchmark{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BenchmarkReceipt != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `id`                                                                                   | `int64`                                                                                | :heavy_check_mark:                                                                     | id                                                                                     |
| `submitBenchmarkRequest`                                                               | [components.SubmitBenchmarkRequest](../../models/components/submitbenchmarkrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

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
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Agents.SubmitErrorAgent(ctx, 649742, components.SubmitErrorRequest{
        Message: "<value>",
        Severity: components.SeverityMajor,
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `id`                                                                           | `int64`                                                                        | :heavy_check_mark:                                                             | id                                                                             |
| `submitErrorRequest`                                                           | [components.SubmitErrorRequest](../../models/components/submiterrorrequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

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