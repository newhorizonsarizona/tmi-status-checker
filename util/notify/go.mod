module github.com/d2tm/tmi-status-checker/util/notify

go 1.24.0

replace github.com/newhorizonsarizona/tmi-status-checker/util => ../../util

require (
	github.com/newhorizonsarizona/tmi-status-checker/util v0.0.0-20251201124030-96e3e551883a
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/sashabaranov/go-openai v1.41.2 // indirect
	golang.org/x/time v0.14.0 // indirect
)
