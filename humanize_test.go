// Copyright (c) 2018 Tim Heckman
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file at the root of this repository.

package humanize_test

import (
	"regexp"
	"testing"

	"github.com/theckman/humanize-go"
	. "gopkg.in/check.v1"
)

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (*TestSuite) TestVersion(c *C) {
	// matches:
	// 1.2.3
	// 1.2.4-rc1
	reg, err := regexp.Compile(`^(?:\d+)\.(?:\d+)\.(?:\d+)(?:-[a-zA-Z0-9\.]+)?$`)
	c.Assert(err, IsNil)
	c.Check(reg.MatchString(humanize.Version), Equals, true)
}
