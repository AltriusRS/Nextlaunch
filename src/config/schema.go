package config

// Configuration is the main configuration struct
type Configuration struct {
	Spec          string        `toml:"spec" json:"spec" comment:"Spec version - Do not edit this - It will cause the application to overwrite the config files\nCurrent spec version: 1"`
	General       General       `toml:"General" json:"General" comment:"Welcome to the NextLaunch configuration file!\nThis file was automatically generated by NextLaunch, and whilst it is human-readable, it is not recommended to edit it directly.\nInstead, please use the built-in configuration editor to make changes, so that your changes are saved correctly.\nFor more information, see https://git.archimedia.uk/Altrius/nextlaunch/-/wikis/Configuration\n\nGeneral settings"`
	Database      Database      `toml:"Database" json:"Database" comment:"Database settings\nNextLaunch will cache information to the database, so it can be accessed quickly,\neven when the application is not able to connect to the internet."`
	LaunchLibrary LaunchLibrary `toml:"LaunchLibrary" json:"LaunchLibrary" comment:"Launch Library settings\nLaunch Library is a public service that provides information about upcoming and historical launches.\nIt is free to use, however you will need an API key to have higher limits. You can get one here: https://www.patreon.com/TheSpaceDevs"`
	Snapi         Snapi         `toml:"Snapi" json:"Snapi" comment:"SNAPI - The Spaceflight News API - is a public service that provides spaceflight news.\nIt is also operated by The Space Devs and is free to use."`
	ErrorLogging  ErrorLogging  `toml:"ErrorLogging" json:"ErrorLogging" comment:"Error logging settings\nError logging is used to provide information about errors that occur in the application.\nThis information is used to help me to fix bugs and improve the application.\nIt is completely anonymous and will not contain any personal or sensitive information."`
	Telemetry     Telemetry     `toml:"Telemetry" json:"Telemetry" comment:"Telemetry settings\nTelemetry is used to collect anonymous usage data and helps me to improve the application.\nThis data is completely anonymous and will not contain any personal or sensitive information."`
}

type General struct {
	LogLevel string `toml:"log_level" json:"log_level" comment:"Log level - Debug, Info, Warning, Error, Fatal - Default is Info"`
	Language string `toml:"language" json:"language" comment:"Language to use for the application - See https://git.archimedia.uk/Altrius/nextlaunch/-/wikis/Supported-Languages for a list of supported languages"`
}

type Database struct {
	DatabaseLocation string `toml:"database_location" json:"database_location" comment:"Location of the database to use"`
}

type LaunchLibrary struct {
	LaunchLibraryKey string `toml:"launch_library_key" json:"launch_library_key" comment:"LL2 API Key"`
	CacheToDatabase  bool   `toml:"cache_to_database" json:"cache_to_database" comment:"Cache the launches to the database - This allows NextLaunch to work without an internet connection (but it will be less accurate)"`
	UpdateFrequency  uint16 `toml:"update_frequency" json:"update_frequency" comment:"How often to check for updates in minutes (0 = disabled | 15 = every 15 minutes | minimum = 5 | maximum = 60)"`
}

type Snapi struct {
	EnableSnapi     bool   `toml:"enable_snapi" json:"enable_snapi" comment:"Enable Snapi"`
	UpdateFrequency uint16 `toml:"update_frequency" json:"update_frequency" comment:"How often to check for updates in minutes (0 = disabled | 15 = every 15 minutes | minimum = 5 | maximum = 60)"`
}

type ErrorLogging struct {
	LogErrors       bool   `toml:"log_errors" json:"log_errors" comment:"Log errors"`
	LogFileLocation string `toml:"log_file_location" json:"log_file_location" comment:"Log file location"`
	AutoShareErrors bool   `toml:"auto_share_errors" json:"auto_share_errors" comment:"Auto share errors"`
}

type Telemetry struct {
	EnableTelemetry bool   `toml:"enable_telemetry" json:"enable_telemetry" comment:"Enables telemetry - This will send anonymous usage data to the devs"`
	TelemetryLevel  uint16 `toml:"telemetry_level" json:"telemetry_level" comment:"0 = None, 1 = Basic, 2 = Verbose"`
}
