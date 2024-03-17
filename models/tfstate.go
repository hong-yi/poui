package models

type TfState struct {
	Version int    `json:"version,omitempty"`
	Serial  int    `json:"serial,omitempty"`
	Lineage string `json:"lineage,omitempty"`
	Backend struct {
		Type   string `json:"type,omitempty"`
		Config struct {
			Hostname     any    `json:"hostname,omitempty"`
			Organization string `json:"organization,omitempty"`
			Token        any    `json:"token,omitempty"`
			Workspaces   struct {
				Name string `json:"name,omitempty"`
				Tags any    `json:"tags,omitempty"`
			} `json:"workspaces,omitempty"`
		} `json:"config,omitempty"`
		Hash int64 `json:"hash,omitempty"`
	} `json:"backend,omitempty"`
	Modules []struct {
		Path    []string `json:"path,omitempty"`
		Outputs struct {
		} `json:"outputs,omitempty"`
		Resources struct {
		} `json:"resources,omitempty"`
		DependsOn []any `json:"depends_on,omitempty"`
	} `json:"modules,omitempty"`
}
