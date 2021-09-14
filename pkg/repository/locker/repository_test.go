package locker

import (
	"testing"

	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/lock"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	r := NewRepository(lock.Config{})
	assert.NotNil(t, r)
}

func TestNewMock(t *testing.T) {
	m := NewMock(nil, 0)
	assert.NotNil(t, m)
}