package endpoints

const (
	// APIGetStatus status check only
	APIGetStatus = "/check"
	// APIGetRemotes remotes list fetch
	APIGetRemotes = "/remotes"
	// APIGetRemote fetch remote detail by name
	APIGetRemote = "/remote/:name"
	// APIGetRemotesRights remotes list with rights check
	APIGetRemotesRights = "/remotes/rights"
)
