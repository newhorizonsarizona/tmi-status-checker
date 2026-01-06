module github.com/d2tm/tmi-status-checker/send_mail

go 1.24.0

require (
	github.com/d2tm/tmi-status-checker/util/notify v0.0.0-20260101124939-bb086ab1a8e8
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

require (
	github.com/newhorizonsarizona/tmi-status-checker/util v0.0.0-20260101123914-7ff383e2aa35 // indirect
	github.com/sashabaranov/go-openai v1.41.2 // indirect
	golang.org/x/time v0.14.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/newhorizonsarizona/tmi-status-checker/util => ../util

replace github.com/d2tm/tmi-status-checker/util/notify => ../util/notify
