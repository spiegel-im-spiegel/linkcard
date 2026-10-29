package facade

import (
	"fmt"
	"runtime/debug"
	"strings"
)

const (
	appNameShort  = "linkcard"                                      // Name of the application
	appNameLong   = "github.com/spiegel-im-spiegel/" + appNameShort // Long name of the application
	repositoryURL = "https://" + appNameLong                        // Repository URL of the application
)

// versionString is a type for version information of the application
type versionString []string

// NewVersionString returns a new versionString instance with the given version string
func NewVersionString(v string) versionString {
	return versionString([]string{
		strings.Join([]string{appNameShort, replaceVersion(v)}, " "),
		repositoryURL,
	})
}

// replaceVersion returns the version string based on the provided version and build information
func replaceVersion(v string) string {
	goVersion := getGoVersion()
	if info, ok := debug.ReadBuildInfo(); ok {
		if v != "" {
			return joinNonEmpty(v, goVersion)
		}
		if info.Main.Version != "" {
			return joinNonEmpty(info.Main.Version, goVersion)
		}
		var revision, dirty string
		for _, v := range info.Settings {
			switch v.Key {
			case "vcs.revision":
				revision = v.Value
			case "vcs.modified":
				if v.Value == "true" {
					dirty = "(dirty)"
				}
			}
		}
		if revision != "" {
			return joinNonEmpty(revision, dirty, goVersion)
		}
	} else if v != "" {
		return v
	}
	return "(version not set)"
}

var getGoVersion = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return fmt.Sprintf("(build with %v)", info.GoVersion)
	}
	return ""
}

func joinNonEmpty(parts ...string) string {
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p != "" {
			out = append(out, p)
		}
	}
	return strings.Join(out, " ")
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
