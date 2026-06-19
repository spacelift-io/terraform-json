// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package sanitize

type Sanitizer func(old interface{}) interface{}

func NewValueSanitizer(value interface{}) Sanitizer {
	return func(old interface{}) interface{} {
		return value
	}
}
