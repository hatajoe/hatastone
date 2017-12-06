package battle

import (
	"testing"

	"github.com/hatajoe/hatastone/battle"
)

func TestContextExec(t *testing.T) {
	ctx := NewContext(&battle.Battle{})
	if err := ctx.Exec(); err != nil {
		t.Fatal(err)
	}
	if _, ok := ctx.GetState().(*InMarigan); !ok {
		t.Fatalf("current state is not expected. expected=%T, actual=%T", &InMarigan{}, ctx.GetState())
	}
	if err := ctx.Exec(); err != nil {
		t.Fatal(err)
	}
	if _, ok := ctx.GetState().(*InMatch); !ok {
		t.Fatalf("current state is not expected. expected=%T, actual=%T", &InMatch{}, ctx.GetState())
	}
	if err := ctx.Exec(); err != nil {
		t.Fatal(err)
	}
	if _, ok := ctx.GetState().(*InResult); !ok {
		t.Fatalf("current state is not expected. expected=%T, actual=%T", &InResult{}, ctx.GetState())
	}
	if err := ctx.Exec(); err != nil {
		t.Fatal(err)
	}
	if ctx.GetState() != nil {
		t.Fatalf("current state is not expected. expected=nil, actual=%T", ctx.GetState())
	}
}
