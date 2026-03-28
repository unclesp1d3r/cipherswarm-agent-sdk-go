# HeartbeatResponseState

The state of the agent:
                       * `pending` - The agent needs to perform the setup process again.
                       * `active` - The agent is ready to accept tasks, all is good.
                       * `error` - The agent has encountered an error and needs to be checked.
                       * `stopped` - The agent has been stopped by the user.
                       * `offline` - The agent has not checked in recently and is considered offline.

## Example Usage

```go
import (
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
)

value := components.HeartbeatResponseStateActive
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `HeartbeatResponseStateActive`  | active                          |
| `HeartbeatResponseStateError`   | error                           |
| `HeartbeatResponseStateOffline` | offline                         |
| `HeartbeatResponseStatePending` | pending                         |
| `HeartbeatResponseStateStopped` | stopped                         |