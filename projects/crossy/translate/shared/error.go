package shared

import "encoding/json"

type errorReason int

const (
	errorParse errorReason = iota
)

type translateError struct {
	reason  errorReason
	message *string
}

func newError(reason errorReason, message *string) translateError {
	return translateError{
		reason:  reason,
		message: message,
	}
}

func (e translateError) String() string {
	reason := func() string {
		switch e.reason {
		case errorParse:
			return "Error while parsing the argument"
		}
		return "Unknown reason"
	}()
	body := make(map[string]any)
	err := make(map[string]string)
	if e.message != nil {
		err["message"] = *e.message
	}
	err["reason"] = reason
	body["error"] = err
	data, _ := json.Marshal(body)
	return string(data)
}
