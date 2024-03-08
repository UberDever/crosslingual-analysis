module translate-xml

go 1.21.3

require aqwari.net/xml v0.0.0-20210331023308-d9421b293817

require translate/shared v0.0.0

replace translate/shared v0.0.0 => ../shared

require (
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
