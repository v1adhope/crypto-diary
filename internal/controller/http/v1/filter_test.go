package v1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidDate(t *testing.T) {
	tests := []struct {
		actual    string
		assertion assert.ErrorAssertionFunc
	}{
		{"", assert.Error},
		{"2023-12-99", assert.Error},
		{"2023-12-a9", assert.Error},
		{"2023-12-19", assert.NoError},
		{"2023-12-19:2022-12-19:2021-12-19", assert.Error},
		{"2023-12-19:2022-12-19", assert.NoError},
		{"2022-12-19:2023-12-19", assert.NoError},
	}

	for _, tt := range tests {
		_, err := getValidDate(tt.actual)
		tt.assertion(t, err)
	}
}

func TestValidateDate(t *testing.T) {
	tests := []struct {
		actual    string
		assertion assert.BoolAssertionFunc
	}{
		{"", assert.False},
		{"2023-12-99", assert.False},
		{"2023-12a-19", assert.False},
		{"2023-12-19", assert.True},
	}

	for _, tt := range tests {
		tt.assertion(t, validateDate(tt.actual))
	}
}

func TestGetValidPair(t *testing.T) {
	tests := []struct {
		actual    string
		assertion assert.ErrorAssertionFunc
	}{
		{"", assert.Error},
		{",", assert.Error},
		{strings.Repeat("0", 256), assert.Error},
		{"btc/usdt", assert.NoError},
		{"btc/usdt,eth/usdt", assert.NoError},
	}

	for _, tt := range tests {
		_, err := getValidPair(tt.actual)
		tt.assertion(t, err)
	}
}

func TestValidatePair(t *testing.T) {
	cases := []struct {
		actual    string
		assertion assert.BoolAssertionFunc
	}{
		{"", assert.False},
		{"0123456789333", assert.False},
		{"btc/usdt", assert.True},
	}

	for _, tt := range cases {
		tt.assertion(t, validatePair(tt.actual))
	}
}

func TestGetValidStrategically(t *testing.T) {
	tests := []struct {
		actual    string
		assertion assert.ErrorAssertionFunc
	}{
		{"", assert.Error},
		{"falsa", assert.Error},
		{"false", assert.NoError},
		{"t", assert.NoError},
	}

	for _, tt := range tests {
		_, err := getValidStrategically(tt.actual)
		tt.assertion(t, err)
	}
}

func TestValidateStrategically(t *testing.T) {
	tests := []struct {
		actual    string
		assertion assert.BoolAssertionFunc
	}{
		{"12", assert.False},
		{"turue", assert.False},
		{"true", assert.True},
		{"f", assert.True},
	}

	for _, tt := range tests {
		tt.assertion(t, validateStrategically(tt.actual))
	}
}
