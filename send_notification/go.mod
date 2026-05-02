module github.com/newhorizonsarizona/tmi-status-checker/send_notification

go 1.25.0

replace github.com/newhorizonsarizona/tmi-status-checker/util => ../util

replace github.com/d2tm/tmi-status-checker/util/notify => ../util/notify

require github.com/d2tm/tmi-status-checker/util/notify v0.0.0-20260501132513-96e89a330495

require (
	github.com/newhorizonsarizona/tmi-status-checker/util v0.0.0-20260401131909-8a64e43d824a // indirect
	github.com/sashabaranov/go-openai v1.41.2 // indirect
	golang.org/x/time v0.15.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
