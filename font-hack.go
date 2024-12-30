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

package font

import (
	_ "embed"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed data/hack/Hack-Bold.ttf
var hackBold []byte

//go:embed data/hack/Hack-BoldItalic.ttf
var hackBoldItalic []byte

//go:embed data/hack/Hack-Italic.ttf
var hackItalic []byte

//go:embed data/hack/Hack-Regular.ttf
var hackRegular []byte

type hackFont struct{}

var Hack Font = hackFont{}

func (f hackFont) Regular(opts *truetype.Options) font.Face    { return face(hackRegular, opts) }
func (f hackFont) Bold(opts *truetype.Options) font.Face       { return face(hackBold, opts) }
func (f hackFont) Italic(opts *truetype.Options) font.Face     { return face(hackItalic, opts) }
func (f hackFont) BoldItalic(opts *truetype.Options) font.Face { return face(hackBoldItalic, opts) }
