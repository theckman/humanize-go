// Copyright (c) 2018 Tim Heckman
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file at the root of this repository.

package humanize

import "fmt"

var ords = []string{"th", "st", "nd", "rd"}

func abs(i int64) uint64 {
	if i < 0 {
		return uint64(i * -1)
	}

	return uint64(i)
}

func ordinality(i uint64) string {
	// numbers like 211, 212, and 213 are special
	if v := i % 100; v == 11 || v == 12 || v == 13 {
		return ords[0]
	}

	if lastNum := i % 10; lastNum < 4 {
		return ords[lastNum]
	}

	return ords[0]
}

// OrdinalInt is a function that takes an int and returns a string with its
// ordinal value. For example, 2 would become "2nd".
func OrdinalInt(i int) string {
	return fmt.Sprintf("%d%s", i, ordinality(abs(int64(i))))
}

// OrdinalInt64 is a function that takes an int64 and returns a string with its
// ordinal value. For example, 3 would become "3rd".
func OrdinalInt64(i int64) string {
	return fmt.Sprintf("%d%s", i, ordinality(abs(i)))
}

// OrdinalUint is a function that takes an uint and returns a string with its
// ordinal value. For example, 4 would become "4th".
func OrdinalUint(i uint) string {
	return fmt.Sprintf("%d%s", i, ordinality(uint64(i)))
}

// OrdinalUint64 is a function that takes an uint64 and returns a string with its
// ordinal value. For example, 42 would become "42nd".
func OrdinalUint64(i uint64) string {
	return fmt.Sprintf("%d%s", i, ordinality(i))
}
