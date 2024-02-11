package version

import (
	"fmt"
	"strconv"
	"strings"
)

func FromString(v string) Version {
	version := Version{}
	if strings.HasPrefix(v, "v") {
		v = strings.ReplaceAll(v, `v`, ``)
		p := strings.Split(v, `.`)
		if len(p) == 3 {
			major, mmIntErr := strconv.Atoi(p[0])
			if mmIntErr == nil {
				version.Major = major
			}
			minor, mIntErr := strconv.Atoi(p[1])
			if mIntErr == nil {
				version.Minor = minor
			}
			patch, pIntErr := strconv.Atoi(p[2])
			if pIntErr == nil {
				version.Patch = patch
			} else {
				version.Patch = 1
			}
		}
	}
	return version
}

type Version struct {
	Major int `json:"ma"`
	Minor int `json:"mi"`
	Patch int `json:"pa"`
}

// Copy returns a new struct of Version with the Major Minor and Patch preserved
func (v *Version) Copy() Version {
	return Version{
		Major: v.Major,
		Minor: v.Minor,
		Patch: v.Patch,
	}
}

// String formats the struct using vMajor.Minor.Patch starting with v0.0.1
func (v *Version) String() string {
	if v == nil {
		return fmt.Sprintf("v0.0.1")
	}
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// BumpMajor sets the Major +1 and then the Minor and Patch to 0
func (v *Version) BumpMajor() bool {
	v.Major += 1
	v.Minor = 0
	v.Patch = 0
	return true
}

// BumpMinor sets the Minor +1 and then the Patch to 0 while leaving the Major untouched
func (v *Version) BumpMinor() bool {
	v.Minor += 1
	v.Patch = 0
	return true
}

// BumpPatch sets the Patch +1 and leaves the Minor and Major untouched
func (v *Version) BumpPatch() bool {
	v.Patch += 1
	return true
}
