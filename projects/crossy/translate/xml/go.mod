module translate-xml

go 1.21.3

require aqwari.net/xml v0.0.0-20210331023308-d9421b293817

require translate/shared v0.0.0

replace translate/shared v0.0.0 => ../shared

require (
	github.com/nsf/jsondiff v0.0.0-20230430225905-43f6cf3098c1 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
