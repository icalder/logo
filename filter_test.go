package main

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

type city struct {
	name    string
	country string
}

// https://github.com/samber/lo#filter
func TestFilterUKCities(t *testing.T) {
	//assert := assert.New(t)
	ukCities := lo.Filter([]city{
		{"London", "UK"},
		{"Bristol", "UK"},
		{"Paris", "France"},
		{"Rouen", "France"},
	}, func(c city, _ int) bool {
		return c.country == "UK"
	},
	)
	assert.ElementsMatch(t, ukCities, []city{{"London", "UK"}, {"Bristol", "UK"}})
}
