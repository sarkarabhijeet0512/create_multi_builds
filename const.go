package main

const (
	ourPath          = "github.com/theQRL/zond" // Path to our module)
	clientIdentifier = "zond"                   // Client identifier to advertise over the network
)

const (
	VersionMajor = 1        // Major version component of the current release
	VersionMinor = 11       // Minor version component of the current release
	VersionPatch = 0        // Patch version component of the current release
	VersionMeta  = "stable" // Version metadata to append to the version string
)

// In go 1.18 and beyond, the go tool embeds VCS information into the build.

const (
	govcsTimeLayout = "2006-01-02T15:04:05Z"
	ourTimeLayout   = "20060102"
)
