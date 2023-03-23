package postgresql

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoWithAttempts(t *testing.T) {
	attempts := 3
	delay := time.Millisecond * 100

	err := DoWithAttempts(func() error {
		return fmt.Errorf("database error")
	}, attempts, delay)

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())

	err = DoWithAttempts(func() error {
		return nil
	}, attempts, delay)

	assert.NoError(t, err)

	err = DoWithAttempts(func() error {
		return errors.New("run error")
	}, attempts, delay)

	assert.Error(t, err)
	assert.Equal(t, "run error", err.Error())
}
