package service

// Status holds the status time
type Status struct {
	Ok  bool   `json:"ok"`
	Err string `json:"err"`
}

// RemoteApp defines the app basic details
type RemoteApp struct {
	Name    string `json:"name"`
	Source  string `json:"source"`
	Service string `json:"service"`
	Type    string `json:"apptype"`
	Status  Status `json:"status,omitempty"`
}

// Remote defines the server basic details
type Remote struct {
	Name    string      `json:"name"`
	Address string      `json:"address"`
	Apps    []RemoteApp `json:"apps"`
	Status  Status      `json:"status,omitempty"`
}
