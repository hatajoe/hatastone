package battle

import "testing"

func TestContextExec(t *testing.T) {
	ctx := Context{}
	if err := ctx.Exec(); err != nil {
		t.Fatal(err)
	}
}
