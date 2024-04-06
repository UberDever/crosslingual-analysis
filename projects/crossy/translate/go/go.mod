module translate-go

require (
	golang.org/x/tools v0.19.0
	translate/shared v0.0.0
)

require (
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nsf/jsondiff v0.0.0-20230430225905-43f6cf3098c1 // indirect
	golang.org/x/mod v0.16.0 // indirect
)

replace translate/shared v0.0.0 => ../shared

go 1.21.3
