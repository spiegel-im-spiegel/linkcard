package config

import (
	"encoding/json"
	"os"

	"github.com/goark/errs"
)

const configFile = "./linkcard.json" // Path of the configuration file

// Config represents toptags command options.
type Config struct {
	VersionFlag   bool   `pflag:"version,v,show version information"`
	UserAgent     string `pflag:"user-agent,u,User-Agent header value" json:"user_agent,omitempty"`
	DataPath      string `pflag:"data-path,d,linkcard data file path" json:"data_path,omitempty"`
	ImageDir      string `pflag:"image-dir,i,image directory path" json:"image_dir,omitempty"`
	ImageBasePath string `pflag:"image-base-path,b,base path for image directory" json:"image_base_path,omitempty"`
	ImageWidth    int    `pflag:"image-width,w,width of image" json:"image_width,omitempty"`
	Rating        int    `pflag:"rating,r,number of stars (1-5)" json:"rating,omitempty"`
	PageTitle     string `pflag:"page-title,t,title of the page" json:"page_title,omitempty"`
	Comment       string `pflag:"comment,c,comment for linkcard data" json:"comment,omitempty"`
}

// DefaultConfig returns a default Config instance.
func DefaultConfig() *Config {
	return &Config{
		VersionFlag:   false,
		DataPath:      "",
		ImageDir:      "",
		ImageBasePath: "",
		UserAgent:     "",
		ImageWidth:    100,
		Rating:        0,
		PageTitle:     "",
		Comment:       "",
	}
}

// ImportConfigFromFile reads the configuration from a JSON file and returns a Config instance.
func ImportConfigFromFile() (*Config, error) {
	data, err := os.ReadFile(configFile) // Read the configuration file
	if err != nil {
		if os.IsNotExist(err) {
			// If the configuration file does not exist, return the default configuration without error
			return DefaultConfig(), nil
		}
		return DefaultConfig(), errs.Wrap(err, errs.WithContext("config_file", configFile))
	}
	cfg := DefaultConfig()                            // Initialize default configuration
	if err := json.Unmarshal(data, cfg); err != nil { // Unmarshal JSON data into the Config struct
		return DefaultConfig(), errs.Wrap(err, errs.WithContext("config_file", configFile))
	}
	return cfg, nil
}
