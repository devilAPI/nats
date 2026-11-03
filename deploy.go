package nats

import "time"

// Package scopes. An empty scope keeps the historical behaviour (machine-wide).
const (
	PackageScopeMachine = "machine"
	PackageScopeUser    = "user"
)

type DeployAction struct {
	AgentId         string    `json:"agentid,omitempty"`
	Action          string    `json:"action,omitempty"`
	When            time.Time `json:"when,omitempty"`
	PackageId       string    `json:"packageid,omitempty"`
	PackageName     string    `json:"packagename,omitempty"`
	PackageVersion  string    `json:"packageversion,omitempty"`
	PackageBranch   string    `json:"packagebranch,omitempty"`
	PackageBrewType string    `json:"packagebrewtype,omitempty"`
	PackageVerified bool      `json:"packageverified,omitempty"`
	PackageScope    string    `json:"packagescope,omitempty"`
	Repository      string    `json:"repository,omitempty"`
	Info            string    `json:"info,omitempty"`
	Failed          bool      `json:"failed,omitempty"`
}

// AvailableUpdate is a package installed on an endpoint for which
// the package manager reports a newer version
type AvailableUpdate struct {
	PackageId        string `json:"packageid"`
	PackageName      string `json:"packagename,omitempty"`
	PackageBranch    string `json:"packagebranch,omitempty"`
	PackageBrewType  string `json:"packagebrewtype,omitempty"`
	InstalledVersion string `json:"installedversion,omitempty"`
	AvailableVersion string `json:"availableversion,omitempty"`
	Source           string `json:"source,omitempty"`
	// Scope is PackageScopeMachine or PackageScopeUser
	Scope string `json:"scope,omitempty"`
	// Users lists the logged-in users that have a user-scoped package installed
	Users []string `json:"users,omitempty"`
}

// UpdatePackagesRequest asks an agent to update several packages, one after another.
// If Packages is empty, the agent updates every package with an available update.
type UpdatePackagesRequest struct {
	AgentId  string            `json:"agentid,omitempty"`
	Packages []AvailableUpdate `json:"packages,omitempty"`
	// RequestedAt is set for requests queued for agents that may be offline, so
	// agents can ignore requests that waited for too long
	RequestedAt time.Time `json:"requestedat,omitempty"`
}

// ListPackageUpdatesRequest asks an agent for its available updates. With StatusOnly,
// the agent only reports its last update run, which is fast and doesn't use the
// package manager.
type ListPackageUpdatesRequest struct {
	StatusOnly bool `json:"statusonly,omitempty"`
}

// UpdateResult is the outcome of updating a single package
type UpdateResult struct {
	AvailableUpdate
	Failed bool `json:"failed,omitempty"`
	// UpToDate is set if the package manager found no applicable update
	UpToDate bool   `json:"uptodate,omitempty"`
	Info     string `json:"info,omitempty"`
}

// UpdateRun records a bulk update run performed by an agent
type UpdateRun struct {
	Started  time.Time      `json:"started"`
	Finished time.Time      `json:"finished,omitempty"`
	Total    int            `json:"total"`
	Results  []UpdateResult `json:"results,omitempty"`
	// Error is set if the run couldn't be performed, e.g. updates couldn't be listed
	Error string `json:"error,omitempty"`
}

// Running reports if the update run is still in progress
func (r *UpdateRun) Running() bool {
	return r != nil && !r.Started.IsZero() && r.Finished.IsZero()
}

// Counts returns how many packages of the run were updated, were already up to date
// and failed to update
func (r *UpdateRun) Counts() (updated, upToDate, failed int) {
	if r == nil {
		return 0, 0, 0
	}
	for _, result := range r.Results {
		switch {
		case result.Failed:
			failed++
		case result.UpToDate:
			upToDate++
		default:
			updated++
		}
	}
	return updated, upToDate, failed
}

// AvailableUpdatesResponse is the agent's answer to an available updates request
type AvailableUpdatesResponse struct {
	Updates []AvailableUpdate `json:"updates,omitempty"`
	LastRun *UpdateRun        `json:"lastrun,omitempty"`
	Error   string            `json:"error,omitempty"`
}
