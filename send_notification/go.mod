module github.com/newhorizonsarizona/tmi-status-checker/send_notification

go 1.22.5

require (
	github.com/newhorizonsarizona/tmi-status-checker/util v0.0.0-20241002175354-e397c6147997
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/newhorizonsarizona/tmi-status-checker/util => ../util

require (
	github.com/sashabaranov/go-openai v1.31.0 // indirect
	golang.org/x/time v0.6.0 // indirect
)
