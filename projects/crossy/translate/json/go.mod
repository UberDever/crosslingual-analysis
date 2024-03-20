module translate-json

require translate/shared v0.0.0

require github.com/nsf/jsondiff v0.0.0-20230430225905-43f6cf3098c1 // indirect

replace translate/shared v0.0.0 => ../shared

go 1.21.3
