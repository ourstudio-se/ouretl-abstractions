package ouretl

// DataHandlerPlugin defines a handler for a data message,
// where any type of handling can occur. A DataHandlerPlugin
// can be used either as a data transformation step or as a
// ETL `sink`. When the plugin has completed it's job, it
// should always call the provided `next` function with
// current state of the data message - unless an error occur.
// The `next` function, available as a function parameter,
// is a caller for the next plugin in line.
//
// Example:
// -----
// package main
//
// type pluginDef struct{}
//
// type jsonMessage struct {
//     DataMessage string `json:"data"`
//     ReceivedAt string `json:"received_at"`
// }
//
// func GetHandler(_ ouretl.Config, _ ouretl.PluginSettings) (ouretl.DataHandlerPlugin, error) {
//     return &pluginDef{}
// }
//
// func (def *pluginDef) Handle(data ouretl.DataMessage, next func([]byte) error) error {
//     s := string(data.Content())
//     m := &jsonMessage{
//         DataMessage: s,
//         ReceivedAt: time.Now().Local().Format("2006-01-02 15:04:05"),
//     }
//
//     output, err := json.Marshal(m)
//     if err != nil {
//         return err
//     }
//
//     return next(output)
// }
// -----
type DataHandlerPlugin interface {
	Handle(dm DataMessage, next func([]byte) error) error
}
