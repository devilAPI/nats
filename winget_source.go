package nats

// WinGetSourceConfig asks an agent to register the console's private winget repository
// (a WinGet REST source), so packages uploaded to the console can be installed, updated
// and uninstalled with winget like any other package
type WinGetSourceConfig struct {
	// Name of the winget source, e.g. OpenUEM
	Name string `json:"name"`
	// URL of the REST source, including the agent's access token
	URL string `json:"url"`
	// TrustCA asks the agent to trust OpenUEM's CA, which issued the console's certificate,
	// as winget only accepts sources with a trusted certificate
	TrustCA bool `json:"trustca,omitempty"`
}
