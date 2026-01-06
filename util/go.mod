module github.com/newhorizonsarizona/tmi-status-checker/util

go 1.24.0

replace github.com/d2tm/tmi-status-checker/util/notify => ./notify

require (
	github.com/sashabaranov/go-openai v1.41.2
	golang.org/x/time v0.14.0
)
