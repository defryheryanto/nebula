package logs_test

import (
	"testing"

	"github.com/defryheryanto/nebula/internal/logs"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	t.Run("empty page", func(t *testing.T) {
		filter := &logs.Filter{}
		isPagination, page, pageSize := filter.GetPagination()
		assert.False(t, isPagination)
		assert.Empty(t, page)
		assert.Empty(t, pageSize)
	})

	t.Run("empty page size", func(t *testing.T) {
		filter := &logs.Filter{
			Page: 1,
		}
		isPagination, page, pageSize := filter.GetPagination()
		assert.True(t, isPagination)
		assert.Equal(t, int64(1), page)
		assert.Equal(t, int64(10), pageSize)
	})

	t.Run("page & page size exists", func(t *testing.T) {
		filter := &logs.Filter{
			Page:     5,
			PageSize: 50,
		}
		isPagination, page, pageSize := filter.GetPagination()
		assert.True(t, isPagination)
		assert.Equal(t, int64(5), page)
		assert.Equal(t, int64(50), pageSize)
	})
}
