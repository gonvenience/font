// Copyright © 2024 The Homeport Team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package font_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/gonvenience/font"
)

var _ = Describe("Font", func() {
	Context("Go", func() {
		It("should be able to get the Go font face", func() {
			Expect(Go.Regular(nil)).ToNot(BeNil())
			Expect(Go.Bold(nil)).ToNot(BeNil())
			Expect(Go.Italic(nil)).ToNot(BeNil())
			Expect(Go.BoldItalic(nil)).ToNot(BeNil())
		})
	})

	Context("Hack", func() {
		It("should be able to get the Hack font face", func() {
			Expect(Hack.Regular(nil)).ToNot(BeNil())
			Expect(Hack.Bold(nil)).ToNot(BeNil())
			Expect(Hack.Italic(nil)).ToNot(BeNil())
			Expect(Hack.BoldItalic(nil)).ToNot(BeNil())
		})
	})
})
