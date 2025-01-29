package cyoa_test

import (
	"bytes"
	"testing"

	"github.com/FedericoBarberon/Go-Exercises/cyoa"
	approvals "github.com/approvals/go-approval-tests"
)

func TestRenderer(t *testing.T) {
	render, err := cyoa.NewBookRenderer()
	assertNoError(t, err)

	t.Run("render arc", func(t *testing.T) {
		buf := &bytes.Buffer{}

		err = render.RenderArc(buf, exampleBook["intro"])
		assertNoError(t, err)

		approvals.VerifyString(t, buf.String())
	})
	t.Run("render final arc", func(t *testing.T) {
		buf := &bytes.Buffer{}

		err = render.RenderArc(buf, exampleBook["arc-1"])
		assertNoError(t, err)

		approvals.VerifyString(t, buf.String())
	})
	t.Run("render non-existing arc", func(t *testing.T) {
		buf := &bytes.Buffer{}

		err = render.Render404(buf, "non-existing")
		assertNoError(t, err)

		approvals.VerifyString(t, buf.String())
	})
}
