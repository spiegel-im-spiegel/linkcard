package facade

import "strings"

const (
	appNameShort  = "linkcard"                                      // Name of the application
	appName       = "github.com/spiegel-im-spiegel/" + appNameShort // Long name of the application
	repositoryURL = "https://" + appName                            // Repository URL of the application
)

// versionString is a type for version information of the application
type versionString []string

// NewVersionString returns a new versionString instance with the given version string
func NewVersionString(v string) versionString {
	return versionString([]string{
		strings.Join([]string{appNameShort, v}, " "),
		repositoryURL,
	})
}

// String returns the string representation of the versionString
func (v versionString) String() string {
	return strings.Join(v, "\n")
}

// UsageString returns the usage string of header part"
func usageString() string {
	return strings.Join([]string{
		appNameShort,
		"",
		"Usage:",
		"  " + appNameShort + " [flags] <url> [<url> ...]",
		"    <url> : URL(s) to generate link cards for",
		"",
		"Flags",
	}, "\n")
}
