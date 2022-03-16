package _00_499

import (
    "strings"
)

func countSegments(s string) int {
    ans := strings.Fields(s)
    return len(ans)
}

