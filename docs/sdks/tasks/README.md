# Tasks

## Overview

Tasks API

### Available Operations

* [GetNewTask](#getnewtask) - Request a new task from server
* [GetTask](#gettask) - Request the task information
* [SendCrack](#sendcrack) - Submit a cracked hash result for a task
* [SendStatus](#sendstatus) - Submit a status update for a task
* [SetTaskAccepted](#settaskaccepted) - Accept Task
* [SetTaskExhausted](#settaskexhausted) - Notify of Exhausted Task
* [SetTaskAbandoned](#settaskabandoned) - Abandon Task
* [GetTaskZaps](#gettaskzaps) - Get Completed Hashes

## GetNewTask

Request a new task from the server, if available.

### Example Usage

<!-- UsageSnippet language="go" operationID="getNewTask" method="get" path="/api/v1/client/tasks/new" -->
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

    res, err := s.Tasks.GetNewTask(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Task != nil {
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

**[*operations.GetNewTaskResponse](../../models/operations/getnewtaskresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## GetTask

Request the task information from the server.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTask" method="get" path="/api/v1/client/tasks/{id}" -->
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

    res, err := s.Tasks.GetTask(ctx, 640281)
    if err != nil {
        log.Fatal(err)
    }
    if res.Task != nil {
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

**[*operations.GetTaskResponse](../../models/operations/gettaskresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401, 404              | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## SendCrack

Submit a cracked hash result for a task.

### Example Usage

<!-- UsageSnippet language="go" operationID="sendCrack" method="post" path="/api/v1/client/tasks/{id}/submit_crack" -->
```go
package main

import(
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/types"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Tasks.SendCrack(ctx, 170496, components.HashcatResult{
        Timestamp: types.MustTimeFromString("2024-07-05T20:43:21.093Z"),
        Hash: "<value>",
        PlainText: "<value>",
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `id`                                                                 | `int64`                                                              | :heavy_check_mark:                                                   | id                                                                   |
| `hashcatResult`                                                      | [components.HashcatResult](../../models/components/hashcatresult.md) | :heavy_check_mark:                                                   | N/A                                                                  |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.SendCrackResponse](../../models/operations/sendcrackresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401, 404              | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## SendStatus

Submit a status update for a task. This includes the status of the current guess and the devices.

### Example Usage

<!-- UsageSnippet language="go" operationID="sendStatus" method="post" path="/api/v1/client/tasks/{id}/submit_status" -->
```go
package main

import(
	"context"
	cipherswarmagentsdkgo "github.com/unclesp1d3r/cipherswarm-agent-sdk-go"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/types"
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := cipherswarmagentsdkgo.New(
        cipherswarmagentsdkgo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.Tasks.SendStatus(ctx, 448338, components.HashcatStatusUpdate{
        OriginalLine: "<value>",
        Time: types.MustTimeFromString("2026-04-23T23:51:40.894Z"),
        Session: "<value>",
        HashcatGuess: components.HashcatGuess{
            GuessBase: "<value>",
            GuessBaseCount: 795653,
            GuessBaseOffset: 302009,
            GuessBasePercentage: 169.59,
            GuessMod: "<value>",
            GuessModCount: 200306,
            GuessModOffset: 691593,
            GuessModPercentage: 5018.12,
            GuessMode: 413091,
        },
        Status: 995383,
        Target: "<value>",
        Progress: []int64{
            920457,
            49769,
        },
        RestorePoint: 343573,
        RecoveredHashes: []int64{
            599902,
            745620,
        },
        RecoveredSalts: []int64{
            524798,
            857056,
        },
        Rejected: 905714,
        DeviceStatuses: []components.DeviceStatus{
            components.DeviceStatus{
                DeviceID: 919497,
                DeviceName: "<value>",
                DeviceType: components.DeviceTypeCPU,
                Speed: 937286,
                Utilization: 556052,
                Temperature: 173236,
            },
        },
        TimeStart: types.MustTimeFromString("2025-04-04T05:12:25.093Z"),
        EstimatedStop: types.MustTimeFromString("2026-07-13T12:38:59.112Z"),
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `id`                                                                             | `int64`                                                                          | :heavy_check_mark:                                                               | id                                                                               |
| `hashcatStatusUpdate`                                                            | [components.HashcatStatusUpdate](../../models/components/hashcatstatusupdate.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.SendStatusResponse](../../models/operations/sendstatusresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401, 422              | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## SetTaskAccepted

Accept an offered task from the server.

### Example Usage

<!-- UsageSnippet language="go" operationID="setTaskAccepted" method="post" path="/api/v1/client/tasks/{id}/accept_task" -->
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

    res, err := s.Tasks.SetTaskAccepted(ctx, 34650)
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

**[*operations.SetTaskAcceptedResponse](../../models/operations/settaskacceptedresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401, 404, 422         | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## SetTaskExhausted

Notify the server that the task is exhausted. This will mark the task as completed.

### Example Usage

<!-- UsageSnippet language="go" operationID="setTaskExhausted" method="post" path="/api/v1/client/tasks/{id}/exhausted" -->
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

    res, err := s.Tasks.SetTaskExhausted(ctx, 994947)
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

**[*operations.SetTaskExhaustedResponse](../../models/operations/settaskexhaustedresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401, 404              | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## SetTaskAbandoned

Abandon a task. This will mark the task as abandoned. Usually used when the client is unable to complete the task.

### Example Usage

<!-- UsageSnippet language="go" operationID="setTaskAbandoned" method="post" path="/api/v1/client/tasks/{id}/abandon" -->
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

    res, err := s.Tasks.SetTaskAbandoned(ctx, 878589)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskAbandonResponse != nil {
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

**[*operations.SetTaskAbandonedResponse](../../models/operations/settaskabandonedresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.ErrorObject      | 401, 404                   | application/json           |
| sdkerrors.TaskAbandonError | 422                        | application/json           |
| sdkerrors.SDKError         | 4XX, 5XX                   | \*/\*                      |

## GetTaskZaps

Gets the completed hashes for a task. This is a text file that should be added to the monitored directory to remove the hashes from the list during runtime.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTaskZaps" method="get" path="/api/v1/client/tasks/{id}/get_zaps" -->
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

    res, err := s.Tasks.GetTaskZaps(ctx, 349297)
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

**[*operations.GetTaskZapsResponse](../../models/operations/gettaskzapsresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorObject | 401, 404, 422         | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |