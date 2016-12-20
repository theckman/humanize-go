package humanize

import (
	"bytes"
	"strconv"
)

func commaInt64(i int64) string {
	// convert the int64 to a string
	str := strconv.FormatInt(i, 10)

	isNegative := i < 0

	if isNegative {
		str = str[1:]
	}

	// get the length of the the string
	// and the number of commas we'll need
	lenStr := len(str)
	numCommas := (lenStr - 1) / 3

	// if we are not adding any commas, short-circuit
	if numCommas == 0 {
		if isNegative {
			return "-" + str
		}
		return str
	}

	// finalValue is the byte slice that will contain our
	// comma-delimited number
	finalValue := make([]byte, lenStr+numCommas)

	var (
		// counter will keep track of the numbers added
		// so that we know when to add a comma (counter%3)
		counter int

		// fvIdx is the index for final value
		fvIdx = len(finalValue) - 1
	)

	// this loop walks backwards over the string (from right to left) so that we
	// can add a comma every 3 letters
	for i := (lenStr - 1); i >= 0; i-- {
		// add the current number to the finalValue slice
		finalValue[fvIdx] = str[i]

		// decrement the index to walk backwards
		// increment the number counter
		fvIdx--
		counter++

		// if this was a number that will need a comma before it, add the comma
		// and decrement the index to further walk backwards
		if counter%3 == 0 && fvIdx > 0 {
			finalValue[fvIdx] = ','
			fvIdx--
		}
	}

	buf := &bytes.Buffer{}

	if isNegative {
		buf.WriteByte('-')
	}

	buf.Write(finalValue)

	// convert the buffer to a string and return it
	return buf.String()
}

// CommaInt64 is a function that takes an int64 and add commas to separate the
// thousands of the number (e.g., 1,000,000).
func CommaInt64(i int64) string {
	return commaInt64(i)
}

// CommaInt is a function that takes an int64 and add commas to separate the
// thousands of the number (e.g., 1,000,000).
func CommaInt(i int) string {
	return CommaInt64(int64(i))
}

func commaUint64(i uint64) string {
	// convert the int64 to a string
	str := strconv.FormatUint(i, 10)

	// get the length of the the string
	// and the number of commas we'll need
	lenStr := len(str)
	numCommas := (lenStr - 1) / 3

	// if we are not adding any commas, short-circuit
	if numCommas == 0 {
		return str
	}

	// finalValue is the byte slice that will contain our
	// comma-delimited number
	finalValue := make([]byte, lenStr+numCommas)

	var (
		// counter will keep track of the numbers added
		// so that we know when to add a comma (counter%3)
		counter int

		// fvIdx is the index for final value
		fvIdx = len(finalValue) - 1
	)

	// this loop walks backwards over the string (from right to left) so that we
	// can add a comma every 3 letters
	for i := (lenStr - 1); i >= 0; i-- {
		// add the current number to the finalValue slice
		finalValue[fvIdx] = str[i]

		// decrement the index to walk backwards
		// increment the number counter
		fvIdx--
		counter++

		// if this was a number that will need a comma before it, add the comma
		// and decrement the index to further walk backwards
		if counter%3 == 0 && fvIdx > 0 {
			finalValue[fvIdx] = ','
			fvIdx--
		}
	}

	// convert the slice to a string and return it
	return string(finalValue)
}

// CommaUint64 is a function that takes an uint64 and add commas to separate the
// thousands of the number (e.g., 1,000,000).
func CommaUint64(i uint64) string {
	return commaUint64(i)
}

// CommaUint is a function that takes an uint64 and add commas to separate the
// thousands of the number (e.g., 1,000,000).
func CommaUint(i uint) string {
	return CommaUint64(uint64(i))
}
