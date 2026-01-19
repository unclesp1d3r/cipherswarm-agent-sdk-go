<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
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
<!-- End SDK Example Usage [usage] -->