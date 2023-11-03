package insure

import (
	"strings"
)

type stringerWithTabs interface {
	StringWithTabs(string) string
}

func stringWithTabs[S ~[]E, E stringerWithTabs](slice S, tabs string) (s string) {
	if len(slice) == 0 {
		return "\n"
	}
	sb := strings.Builder{}
	for _, ele := range slice {
		sb.WriteString(ele.StringWithTabs(tabs))
		sb.WriteByte('\n')
	}
	return sb.String()
}
