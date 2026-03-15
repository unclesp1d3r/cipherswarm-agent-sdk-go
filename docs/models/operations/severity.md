# Severity

The severity of the error:
                       * `info` - Informational message, no action required.
                       * `warning` - Non-critical error, no action required. Anticipated, but not necessarily problematic.
                       * `minor` - Minor error, no action required. Should be investigated, but the task can continue.
                       * `major` - Major error, action required. The task should be investigated and possibly restarted.
                       * `critical` - Critical error, action required. The task should be stopped and investigated.
                        * `fatal` - Fatal error, action required. The agent cannot continue with the task and should not be reattempted.

## Example Usage

```go
import (
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/operations"
)

value := operations.SeverityInfo
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `SeverityInfo`     | info               |
| `SeverityWarning`  | warning            |
| `SeverityMinor`    | minor              |
| `SeverityMajor`    | major              |
| `SeverityCritical` | critical           |
| `SeverityFatal`    | fatal              |