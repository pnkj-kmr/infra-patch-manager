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
	// APIGetRemoteRights fetch remote rights by name
	APIGetRemoteRights = "/remote/:name/rights"
	// APIUploadPatch upload to remote
	APIUploadPatch = "/upload/patch/:name"
	// APIApplyPatch apply patch to remotes apps
	APIApplyPatch = "/apply/patch"
	// APIVerifyPatch verify patch to remotes apps
	APIVerifyPatch = "/verify/patch"
	// APIPushCMDToRemote helps to push few commands to remote
	// APIPushCMDToRemote = "/push/cmd/:name"
)

type patchbody struct {
	Remote  string `json:"remote" xml:"remote" form:"remote"`
	Apptype string `json:"apptype" xml:"apptype" form:"apptype"`
}
