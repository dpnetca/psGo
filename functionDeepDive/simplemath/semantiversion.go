package simplemath

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

// This is a method on the semantic Version object
// this one doesn't need to be a pointer, but good practice too
func (sv *SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

// these ones need to be poionter to change the variable value
func (sv *SemanticVersion) IngrementMajor() {
	sv.major += 1
	sv.minor = 0
	sv.patch = 0
}

func (sv *SemanticVersion) IngrementMinor() {
	sv.minor += 1
	sv.patch = 0
}

func (sv *SemanticVersion) IngrementPatch() {
	sv.patch += 1
}
