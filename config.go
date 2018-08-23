package ouretl

// PluginDefinition is an abstraction for a plugin, and describes its
// runtime behavior and load it
type PluginDefinition interface {
	Name() string
	FilePath() string
	Version() string
	Priority() int
	IsActive() bool
	Settings() PluginSettings
}

// Config is an abstraction for configuration settings, usually loaded
// from a TOML based configuration file. It describes which plugins
// to use, in which order they run, and how to handle plugin settings.
// `inherit_settings_from_env` is a boolean that indicates if there are
// plugin settings available as environment variables. If `true` then
// any plugin setting will first try to load from it's corresponding
// environment variable, and secondly try to load from a specified
// `settings_file` for the plugin.
//
// Example of a config file:
// -----
// inherit_settings_from_env = true
//
// [[plugin]]
// name = "kafka-stream-reader"
// path = "/usr/share/ouretl/plugins/ouretl-plugin-kafka-stream-reader.so.1.0.0"
// version = "1.0.0"
// priority = 1
// settings_file = "/usr/share/ouretl/config-kafka-stream-reader.toml"
//
// [[plugin]]
// name = "elasticsearch-writer"
// path = "/usr/share/ouretl/plugins/ouretl-plugin-elasticsearch-writer.so.4.1.0"
// version = "4.1.0"
// priority = 10
// settings_file = "/usr/share/ouretl/config-elasticsearch-writer.toml"
// -----
//
//
// Example of a plugin settings file:
// -----
// string_variable = "value"
// int_variable = 1
// -----
type Config interface {
	PluginDefinitions() []PluginDefinition
	AppendPluginDefinition(pdef PluginDefinition) error
	OnPluginDefinitionAdded(func(PluginDefinition))
	OnPluginDefinitionActivated(func(PluginDefinition))
	OnPluginDefinitionDeactivated(func(PluginDefinition))
	Activate(pdef PluginDefinition)
	Deactivate(pdef PluginDefinition)
}

// PluginSettings is an abstraction to access plugin specific settings,
// simply pass the setting name to extract the correct value, either
// from an environment variable or from the plugins settings file.
// Casting to the actual type is necessary when retrieving a setting value;
// -----
// myValue, ok := psettings.Get("int_variable")
// if ok {
//     myIntValue := myValue.(int)
// }
// -----
type PluginSettings interface {
	Get(key string) (interface{}, bool)
}
