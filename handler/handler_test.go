package handler_test

import (
	"testing"

	"github.com/kompit-recruitment/backend/generated/api"
	"github.com/kompit-recruitment/backend/handler"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	handler := handler.New()

	t.Run("should implement ServerInterface", func(t *testing.T) {
		require.Implements(t, (*api.ServerInterface)(nil), handler)
	})
}
