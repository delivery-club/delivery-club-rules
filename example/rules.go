package example

import (
	"github.com/quasilyte/go-ruleguard/dsl"

	rules "github.com/delivery-club/delivery-club-rules"
)

func init() {
	dsl.ImportRules("", rules.Bundle)
}
