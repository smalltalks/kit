package meta

// Version represent the version of a software
type Version string

var pkgVersion = Version("v0.0.1-alpha")

// GetVersion return the current version of KIT
func GetVersion() Version {
	return pkgVersion
}
