package core

import (
    "strings"
	"time"
    "fmt"
)

func SeparateTextAndHashtags(note string) (textPart string, hashtagParts []string) {
    parts := strings.Split(note, "#")

    textPart = strings.TrimSpace(parts[0])

    for _, part := range parts[1:] {
        trimmedPart := strings.TrimSpace(part)
        if trimmedPart != "" {
            hashtagParts = append(hashtagParts, trimmedPart)
        }
    }

    return
}

func TimeConvert(timeStr string) (string) {
    noteTime, _ := time.Parse("2006-01-02 15:04:05", timeStr)

    now := time.Now()

    elapsed := now.Sub(noteTime)
    var timeAgo string

    if elapsed < time.Minute {
        timeAgo = fmt.Sprintf("%d seconds ago", int(elapsed.Seconds()))
    } else if elapsed < time.Hour {
        timeAgo = fmt.Sprintf("%d minutes ago", int(elapsed.Minutes()))
    } else if elapsed < 24*time.Hour {
        timeAgo = fmt.Sprintf("%d hours ago", int(elapsed.Hours()))
    } else {
        timeAgo = fmt.Sprintf("%d days ago", int(elapsed.Hours()/24))
    }

    return timeAgo
}
