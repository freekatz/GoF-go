package adapter_case1

import "testing"

func Test(t *testing.T) {
	adapter := NewAdapter()
	adapter.PrintWeak("Print weak with praen.")
	adapter.PrintStrong("Print strong with aster.")
}
