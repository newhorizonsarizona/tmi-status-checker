module github.com/d2tm/tmi-status-checker/util/notify

go 1.25.0

replace github.com/newhorizonsarizona/tmi-status-checker/util => ../../util

require (
	github.com/newhorizonsarizona/tmi-status-checker/util v0.0.0-20260401131909-8a64e43d824a
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/sashabaranov/go-openai v1.41.2 // indirect
	golang.org/x/time v0.15.0 // indirect
)
