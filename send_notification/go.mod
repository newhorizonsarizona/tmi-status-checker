module github.com/newhorizonsarizona/tmi-status-checker/send_notification

go 1.24.0

replace github.com/newhorizonsarizona/tmi-status-checker/util => ../util

replace github.com/d2tm/tmi-status-checker/util/notify => ../util/notify

require github.com/d2tm/tmi-status-checker/util/notify v0.0.0-20260101124939-bb086ab1a8e8

require (
	github.com/newhorizonsarizona/tmi-status-checker/util v0.0.0-20260101123914-7ff383e2aa35 // indirect
	github.com/sashabaranov/go-openai v1.41.2 // indirect
	golang.org/x/time v0.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
