// Copyright (c) 2018 Tim Heckman
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file at the root of this repository.

package humanize_test

import (
	"github.com/theckman/humanize-go"
	. "gopkg.in/check.v1"
)

func (*TestSuite) TestOrdinalInt(c *C) {
	var human string

	tests := []struct {
		i int
		r string
	}{
		{i: 0, r: "0th"},
		{i: 1, r: "1st"},
		{i: 2, r: "2nd"},
		{i: 3, r: "3rd"},
		{i: 4, r: "4th"},
		{i: 5, r: "5th"},
		{i: 6, r: "6th"},
		{i: 7, r: "7th"},
		{i: 8, r: "8th"},
		{i: 9, r: "9th"},
		{i: 11, r: "11th"},
		{i: 12, r: "12th"},
		{i: 13, r: "13th"},
		{i: 111, r: "111th"},
		{i: 212, r: "212th"},
		{i: 313, r: "313th"},
		{i: -13, r: "-13th"},
	}

	for _, test := range tests {
		human = humanize.OrdinalInt(test.i)
		c.Check(human, Equals, test.r)
	}
}

func (*TestSuite) TestOrdinalInt64(c *C) {
	var human string

	tests := []struct {
		i int64
		r string
	}{
		{i: 0, r: "0th"},
		{i: 1, r: "1st"},
		{i: 2, r: "2nd"},
		{i: 3, r: "3rd"},
		{i: 4, r: "4th"},
		{i: 5, r: "5th"},
		{i: 6, r: "6th"},
		{i: 7, r: "7th"},
		{i: 8, r: "8th"},
		{i: 9, r: "9th"},
		{i: 11, r: "11th"},
		{i: 12, r: "12th"},
		{i: 13, r: "13th"},
		{i: 111, r: "111th"},
		{i: 212, r: "212th"},
		{i: 313, r: "313th"},
		{i: -13, r: "-13th"},
	}

	for _, test := range tests {
		human = humanize.OrdinalInt64(test.i)
		c.Check(human, Equals, test.r)
	}
}

func (*TestSuite) TestOrdinaUint(c *C) {
	var human string

	tests := []struct {
		i uint
		r string
	}{
		{i: 0, r: "0th"},
		{i: 1, r: "1st"},
		{i: 2, r: "2nd"},
		{i: 3, r: "3rd"},
		{i: 4, r: "4th"},
		{i: 5, r: "5th"},
		{i: 6, r: "6th"},
		{i: 7, r: "7th"},
		{i: 8, r: "8th"},
		{i: 9, r: "9th"},
		{i: 11, r: "11th"},
		{i: 12, r: "12th"},
		{i: 13, r: "13th"},
		{i: 111, r: "111th"},
		{i: 212, r: "212th"},
		{i: 313, r: "313th"},
	}

	for _, test := range tests {
		human = humanize.OrdinalUint(test.i)
		c.Check(human, Equals, test.r)
	}
}

func (*TestSuite) TestOrdinalUint64(c *C) {
	var human string

	tests := []struct {
		i uint64
		r string
	}{
		{i: 0, r: "0th"},
		{i: 1, r: "1st"},
		{i: 2, r: "2nd"},
		{i: 3, r: "3rd"},
		{i: 4, r: "4th"},
		{i: 5, r: "5th"},
		{i: 6, r: "6th"},
		{i: 7, r: "7th"},
		{i: 8, r: "8th"},
		{i: 9, r: "9th"},
		{i: 11, r: "11th"},
		{i: 12, r: "12th"},
		{i: 13, r: "13th"},
		{i: 111, r: "111th"},
		{i: 212, r: "212th"},
		{i: 313, r: "313th"},
	}

	for _, test := range tests {
		human = humanize.OrdinalUint64(test.i)
		c.Check(human, Equals, test.r)
	}
}
