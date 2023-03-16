package strx

import "strings"

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
