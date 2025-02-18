module github.com/newhorizonsarizona/tmi-status-checker/send_notification

go 1.24.0

replace github.com/newhorizonsarizona/tmi-status-checker/util => ../util

replace github.com/d2tm/tmi-status-checker/util/notify => ../util/notify

require (
	github.com/newhorizonsarizona/tmi-status-checker/util v0.0.0-20250202012046-4070b311deda // indirect
	github.com/sashabaranov/go-openai v1.31.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
