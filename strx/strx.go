package strx

import (
	"sort"
	"strings"
)

func Cat(strs ...string) string {
	n := 0
	for _, s := range strs {
		n += len(s)
	}

	var builder strings.Builder
	builder.Grow(n)
	for _, s := range strs {
		builder.WriteString(s)
	}

	return builder.String()
}

func ExtractMiddle(str, startMarker, endMarker string) string {
	startIdx := strings.Index(str, startMarker)
	if startIdx == -1 {
		return ""
	}
	str = str[startIdx+len(startMarker):]
	endIdx := strings.Index(str, endMarker)
	if endIdx == -1 {
		return ""
	}
	return str[:endIdx]
}

func ExtractQuotedWords(strs []string) []string {
	var words []string
	for _, str := range strs {
		word := ExtractMiddle(str, "\"", "\"")
		if word == "" {
			continue
		}
		words = append(words, word)
	}
	sort.Strings(words)
	return words
}

func ExtractKeywords(content, startMarker, endMarker string) []string {
	keywordSection := ExtractMiddle(content, startMarker, endMarker)
	lines := strings.Split(keywordSection, "\n")
	return ExtractQuotedWords(lines)
}
