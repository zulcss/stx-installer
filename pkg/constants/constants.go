package constants

const (
	// Immutable directory storage
	OemDir	= "/oem/"

)

// StarlingX Version map to stx-installer
func StarlingxVersion() map[string]string {
	return map[string]string {
		"current": "current",
	}
}
