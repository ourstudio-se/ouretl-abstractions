package ouretl

// WorkerPlugin defines a worker (or ETL `source`) that's always
// running, under ouretl-core. A worker plugin can choose
// to publish messages into the plugin pipeline, using
// the provided `target` function.
//
// Example:
// -----
// package main
//
// type pluginDef struct{
//      portNumber string
// }
//
// func GetWorker(_ ouretl.Config, settings ouretl.PluginSettings) (ouretl.WorkerPlugin, error) {
//     return &pluginDef{
//         portNumber: settings.Get("PortSettings"),
//     }
// }
//
// func (def *pluginDef) Start(target func([]byte)) error {
//     router := httprouter.New()
//     router.POST("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//         defer r.Body.Close()
//
//         message, err := ioutil.ReadAll(r.Body)
//         if err != nil {
//             w.WriteHeader(http.StatusBadRequest)
//             return
//         }
//
//         target(message)
//         w.WriteHeader(http.StatusAccepted)
//     })
//
//     addr := fmt.Sprintf(":%s", def.portNumber)
//     return http.ListenAndServe(addr, router)
// }
// -----
type WorkerPlugin interface {
	Start(func([]byte)) error
}
