// The MIT License (MIT)
//
// Copyright (c) 2014 Matt Ho
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package template

import (
	"github.com/pborman/uuid"
	"regexp"
	"strings"
	text_template "text/template"
)

var funcMap = text_template.FuncMap{
	"upper":        Upper, // all uppercase letters
	"uppercase":    Upper,
	"lower":        Lower, // all lowercase letters
	"lowercase":    Lower,
	"cap":          Capitalize, // uppercase first letter
	"capitalize":   Capitalize,
	"decap":        Decapitalize, // lowercase first letter
	"decapitalize": Decapitalize,
	"start":        Start, // uppercase the first letter of each word
	// "start-case":   Start,
	"word": Word, // remove all non-word letters (only a-zA-Z0-9_)
	// "word-only":    Word,
	// "space":			Space,//replace all non-word letters (only a-zA-Z0-9) with a whitespace
	// "word-space":		Space,
	"Camel": Camel, // upper camel case (start-case, word-only)
	// "upper-camel": Camel,
	"camel": CamelLower, // lower camel case (start-case, word-only, decapitalize)
	// "lower-camel": CamelLower,
	"hyphen":    Hyphenate, // replace spaces with hyphens
	"hyphenate": Hyphenate,
	"norm":      Normalize, // all lowercase with hyphens (lowercase, hyphenate)
	"normalize": Normalize,
	"snake":     Snake, // replace spaces and dots with underscores
	// "snake-case":  Snake,
	// "package":         Package,//replace spaces with dots
	// "package-naming":  Package,
	"packaged": Packaged, // replace dots with slashes (net.databinder -> net/databinder)
	// "packaged-dir":    Packaged,
	"random": Random, // appends random characters to the given string
	// "generate-random": Random,
}

var (
	wordRe       = regexp.MustCompile(`\W`)
	whitespaceRe = regexp.MustCompile(`\s`)
	dotRe        = regexp.MustCompile(`\.`)
	snakeRe      = regexp.MustCompile(`\.|\s`)
)

func Upper(value string) string {
	return strings.ToUpper(value)
}

func Lower(value string) string {
	return strings.ToLower(value)
}

func Word(value string) string {
	return wordRe.ReplaceAllString(value, "")
}

func Capitalize(value string) string {
	switch len(value) {
	case 0:
		return ""
	case 1:
		return strings.ToUpper(value[:1])
	default:
		return strings.ToUpper(value[:1]) + value[1:]
	}
}

func Decapitalize(value string) string {
	switch len(value) {
	case 0:
		return ""
	case 1:
		return strings.ToLower(value[:1])
	default:
		return strings.ToLower(value[:1]) + value[1:]
	}
}

func Start(value string) string {
	parts := strings.Split(value, " ")
	capped := []string{}

	for _, part := range parts {
		capped = append(capped, Capitalize(part))
	}

	return strings.Join(capped, " ")
}

func Camel(value string) string {
	value = Start(value)
	return Word(value)
}

func CamelLower(value string) string {
	value = Camel(value)
	return Decapitalize(value)
}

func Normalize(value string) string {
	value = Hyphenate(value)
	return strings.ToLower(value)
}

func Hyphenate(value string) string {
	return whitespaceRe.ReplaceAllString(value, "-")
}

func Packaged(value string) string {
	return dotRe.ReplaceAllString(value, "/")
}

func Snake(value string) string {
	return snakeRe.ReplaceAllString(value, "_")
}

func Random(value string) string {
	return value + Word(uuid.NewRandom().String())
}
