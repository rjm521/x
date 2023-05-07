package strx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractMiddle(t *testing.T) {
	for _, test := range []struct {
		in   string
		want string
	}{
		{"", ""},
		{"abc|dec,", "dec"},
		{"abc|abc,", "abc"},
	} {
		got := ExtractMiddle(test.in, "|", ",")
		assert.Equal(t, test.want, got, test.in)
	}
}

func TestExtractQuotedWords(t *testing.T) {
	for _, test := range []struct {
		in   []string
		want []string
	}{
		{[]string{"\"hello\""}, []string{"hello"}},
		{[]string{`"how you doing"`, `"this is not my fault"`},
			[]string{"how you doing", "this is not my fault"}},
	} {
		got := ExtractQuotedWords(test.in)
		assert.Equal(t, test.want, got, test.in)
	}
}

func TestExtractKeywords(t *testing.T) {
	for _, test := range []struct {
		in   string
		want []string
	}{
		{`
this is the content>
"i"
"am"
"what"
"i"
"am"
<over`, []string{"am", "am", "i", "i", "what"}},
	} {
		got := ExtractKeywords(test.in, ">", "<")
		assert.Equal(t, test.want, got, test.in)
	}

}
