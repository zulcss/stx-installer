package constants

const (
	// Immutable directory storage
	OEMDir      = "/oem/"
	OEMImageDir = "/oem/images"
)

// StarlingX Version map to stx-installer
func StarlingxVersion() map[string]string {
	return map[string]string{
		"current": "current",
	}
}
