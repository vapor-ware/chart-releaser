package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vapor-ware/chart-releaser/pkg/v1/ctx"
)

func Test_RenderTemplate(t *testing.T) {
	context := &ctx.Context{
		Token: "foobar",
	}

	rendered, err := RenderTemplate(context, "test-tmpl", "token:{{ .Token }}")
	assert.NoError(t, err)
	assert.Equal(t, "token:foobar", rendered)
}

func Test_RenderTemplateParseError(t *testing.T) {
	context := &ctx.Context{
		Token: "foobar",
	}

	rendered, err := RenderTemplate(context, "test-tmpl", "{{")
	assert.EqualError(t, err, "template: test-tmpl:1: unclosed action")
	assert.Equal(t, "", rendered)
}

func Test_RenderTemplateParseErrorDryRun(t *testing.T) {
	context := &ctx.Context{
		DryRun: true,
		Token:  "foobar",
	}

	rendered, err := RenderTemplate(context, "test-tmpl", "{{")
	assert.NoError(t, err)
	assert.Equal(t, "dry-run", rendered)
}

func Test_RenderTemplateExecuteError(t *testing.T) {
	context := &ctx.Context{
		Token: "foobar",
	}

	rendered, err := RenderTemplate(context, "test-tmpl", "{{ .Not.Exist }}")
	assert.EqualError(t, err, "template: test-tmpl:1:7: executing \"test-tmpl\" at <.Not.Exist>: can't evaluate field Not in type *ctx.Context")
	assert.Equal(t, "", rendered)
}

func Test_RenderTemplateExecuteErrorDryRun(t *testing.T) {
	context := &ctx.Context{
		DryRun: true,
		Token:  "foobar",
	}

	rendered, err := RenderTemplate(context, "test-tmpl", "{{ .Not.Exist }}")
	assert.NoError(t, err)
	assert.Equal(t, "dry-run", rendered)
}
