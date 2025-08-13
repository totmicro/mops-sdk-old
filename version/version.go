package version

import (
	"fmt"
	"strconv"
	"strings"
)

// Version represents a semantic version
type Version struct {
	Major int
	Minor int
	Patch int
	Pre   string // Pre-release identifier
	Build string // Build metadata
}

// MOPS API version constants
const (
	CurrentAPIVersion = "1.0.0"
	MinAPIVersion     = "1.0.0"
	MaxAPIVersion     = "1.99.99"
)

// Parse parses a version string into a Version struct
func Parse(s string) (*Version, error) {
	if s == "" {
		return nil, fmt.Errorf("version string cannot be empty")
	}

	// Split build metadata
	parts := strings.Split(s, "+")
	versionPart := parts[0]
	var buildPart string
	if len(parts) > 1 {
		buildPart = parts[1]
	}

	// Split pre-release
	parts = strings.Split(versionPart, "-")
	corePart := parts[0]
	var prePart string
	if len(parts) > 1 {
		prePart = strings.Join(parts[1:], "-")
	}

	// Parse core version
	coreComponents := strings.Split(corePart, ".")
	if len(coreComponents) != 3 {
		return nil, fmt.Errorf("invalid version format: %s", s)
	}

	major, err := strconv.Atoi(coreComponents[0])
	if err != nil {
		return nil, fmt.Errorf("invalid major version: %s", coreComponents[0])
	}

	minor, err := strconv.Atoi(coreComponents[1])
	if err != nil {
		return nil, fmt.Errorf("invalid minor version: %s", coreComponents[1])
	}

	patch, err := strconv.Atoi(coreComponents[2])
	if err != nil {
		return nil, fmt.Errorf("invalid patch version: %s", coreComponents[2])
	}

	return &Version{
		Major: major,
		Minor: minor,
		Patch: patch,
		Pre:   prePart,
		Build: buildPart,
	}, nil
}

// String returns the string representation of the version
func (v *Version) String() string {
	s := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.Pre != "" {
		s += "-" + v.Pre
	}
	if v.Build != "" {
		s += "+" + v.Build
	}
	return s
}

// Compare compares two versions
// Returns: -1 if v < other, 0 if v == other, 1 if v > other
func (v *Version) Compare(other *Version) int {
	if v.Major != other.Major {
		if v.Major < other.Major {
			return -1
		}
		return 1
	}

	if v.Minor != other.Minor {
		if v.Minor < other.Minor {
			return -1
		}
		return 1
	}

	if v.Patch != other.Patch {
		if v.Patch < other.Patch {
			return -1
		}
		return 1
	}

	// Handle pre-release versions
	if v.Pre == "" && other.Pre != "" {
		return 1 // Release version > pre-release
	}
	if v.Pre != "" && other.Pre == "" {
		return -1 // Pre-release < release version
	}
	if v.Pre != "" && other.Pre != "" {
		if v.Pre < other.Pre {
			return -1
		}
		if v.Pre > other.Pre {
			return 1
		}
	}

	return 0
}

// IsCompatible checks if this version is compatible with the target version range
func (v *Version) IsCompatible(minVersion, maxVersion string) (bool, error) {
	if minVersion != "" {
		min, err := Parse(minVersion)
		if err != nil {
			return false, fmt.Errorf("invalid min version: %w", err)
		}
		if v.Compare(min) < 0 {
			return false, nil
		}
	}

	if maxVersion != "" {
		max, err := Parse(maxVersion)
		if err != nil {
			return false, fmt.Errorf("invalid max version: %w", err)
		}
		if v.Compare(max) > 0 {
			return false, nil
		}
	}

	return true, nil
}

// CheckAPICompatibility checks if a plugin API version is compatible with MOPS
func CheckAPICompatibility(pluginAPIVersion string) error {
	pluginVersion, err := Parse(pluginAPIVersion)
	if err != nil {
		return fmt.Errorf("invalid plugin API version: %w", err)
	}

	minVersion, _ := Parse(MinAPIVersion)
	maxVersion, _ := Parse(MaxAPIVersion)

	if pluginVersion.Compare(minVersion) < 0 {
		return fmt.Errorf("plugin API version %s is too old (minimum: %s)", pluginAPIVersion, MinAPIVersion)
	}

	if pluginVersion.Compare(maxVersion) > 0 {
		return fmt.Errorf("plugin API version %s is too new (maximum: %s)", pluginAPIVersion, MaxAPIVersion)
	}

	return nil
}

// GetCurrentAPIVersion returns the current MOPS API version
func GetCurrentAPIVersion() string {
	return CurrentAPIVersion
}
