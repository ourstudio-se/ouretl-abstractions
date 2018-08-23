package ouretl

// DataMessage is a wrapper abstraction for the raw message provided
// to the DataHandlerPlugins. Apart from the actual data
// content, the message has an `Origin` field containing
// a WorkerPlugin name and an `ID` field with a unique
// message ID.
type DataMessage interface {
	ID() string
	Data() []byte
	Origin() string
}
