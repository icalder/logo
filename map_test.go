package main

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

// https://github.com/samber/lo#map
func TestMapCities(t *testing.T) {
	//assert := assert.New(t)
	uniqueCountryNames := lo.Uniq(lo.Map([]city{
		{"London", "UK"},
		{"Bristol", "UK"},
		{"Paris", "France"},
		{"Rouen", "France"},
	}, func(c city, _ int) string { return c.country },
	))
	assert.ElementsMatch(t, uniqueCountryNames, []string{"UK", "France"})
}
