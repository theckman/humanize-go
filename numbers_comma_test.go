// Copyright (c) 2018 Tim Heckman
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file at the root of this repository.

package humanize_test

import (
	"github.com/theckman/humanize-go"
	. "gopkg.in/check.v1"
)

func (*TestSuite) TestCommaInt64(c *C) {
	var human string

	tests := []struct {
		i int64
		r string
	}{
		{1, "1"},
		{10, "10"},
		{100, "100"},
		{1000, "1,000"},
		{10000, "10,000"},
		{100000, "100,000"},
		{1000000, "1,000,000"},
		{1000000000, "1,000,000,000"},
		{10000000000, "10,000,000,000"},
		{-10000000, "-10,000,000"},
		{-1000000, "-1,000,000"},
		{-100000, "-100,000"},
		{-100, "-100"},
	}

	for _, test := range tests {
		human = humanize.CommaInt64(test.i)
		c.Check(human, Equals, test.r)
	}
}

func (*TestSuite) TestCommaInt(c *C) {
	var human string

	tests := []struct {
		i int
		r string
	}{
		{1, "1"},
		{10, "10"},
		{100, "100"},
		{1000, "1,000"},
		{10000, "10,000"},
		{100000, "100,000"},
		{1000000, "1,000,000"},
		{1000000000, "1,000,000,000"},
		{10000000000, "10,000,000,000"},
		{-10000000, "-10,000,000"},
		{-1000000, "-1,000,000"},
		{-100000, "-100,000"},
		{-100, "-100"},
	}

	for _, test := range tests {
		human = humanize.CommaInt(test.i)
		c.Check(human, Equals, test.r)
	}
}

func (*TestSuite) TestCommaUint64(c *C) {
	var human string

	tests := []struct {
		i uint64
		r string
	}{
		{1, "1"},
		{10, "10"},
		{100, "100"},
		{1000, "1,000"},
		{10000, "10,000"},
		{100000, "100,000"},
		{1000000, "1,000,000"},
		{1000000000, "1,000,000,000"},
		{10000000000, "10,000,000,000"},
	}

	for _, test := range tests {
		human = humanize.CommaUint64(test.i)
		c.Check(human, Equals, test.r)
	}
}

func (*TestSuite) TestCommaUint(c *C) {
	var human string

	tests := []struct {
		i uint
		r string
	}{
		{1, "1"},
		{10, "10"},
		{100, "100"},
		{1000, "1,000"},
		{10000, "10,000"},
		{100000, "100,000"},
		{1000000, "1,000,000"},
		{1000000000, "1,000,000,000"},
		{10000000000, "10,000,000,000"},
	}

	for _, test := range tests {
		human = humanize.CommaUint(test.i)
		c.Check(human, Equals, test.r)
	}
}
