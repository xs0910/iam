package v1

// ListOptions is the query options to a standard REST list call.
type ListOptions struct {
	TypeMeta `json:",inline"`

	// LabelSelector is used to find matching REST resources.
	LabelSelector string `json:"labelSelector,omitempty" form:"labelSelector"`

	// FieldSelector restricts the list of returned objects by their fields. Defaults to everything.
	FieldSelector string `json:"fieldSelector,omitempty" form:"fieldSelector"`

	// TimeoutSeconds specifies the seconds of ClientIP type session sticky time.
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty"`

	// Offset specify the number of records to skip before starting to return the records.
	Offset *int64 `json:"offset,omitempty" form:"offset"`

	// Limit specify the number of records to be retrieved.
	Limit *int64 `json:"limit,omitempty" form:"limit"`
}

// ExportOptions is the query options to the standard REST get call.
// Deprecated. Planned for removal in 1.18.
type ExportOptions struct {
	TypeMeta `json:",inline"`

	// Should this value be exported.  Export strips fields that a user can not specify.
	// Deprecated. Planned for removal in 1.18.
	Export bool `json:"export"`
	// Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'.
	// Deprecated. Planned for removal in 1.18.
	Exact bool `json:"exact"`
}

// GetOptions is the standard query options to the standard REST get call.
type GetOptions struct {
	TypeMeta `json:",inline"`
}

// DeleteOptions may be provided when deleting an API object.
type DeleteOptions struct {
	TypeMeta `json:",inline"`

	// +optional
	Unscoped bool `json:"unscoped"`
}

// CreateOptions may be provided when creating an API object.
type CreateOptions struct {
	TypeMeta `json:",inline"`

	// When present, indicates that modifications should not be
	// persisted. An invalid or unrecognized dryRun directive will
	// result in an error response and no further processing of the
	// request. Valid values are:
	// - All: all dry run stages will be processed
	// +optional
	DryRun []string `json:"dryRun,omitempty"`
}

// PatchOptions may be provided when patching an API object.
// PatchOptions is meant to be a superset of UpdateOptions.
type PatchOptions struct {
	TypeMeta `json:",inline"`

	// When present, indicates that modifications should not be
	// persisted. An invalid or unrecognized dryRun directive will
	// result in an error response and no further processing of the
	// request. Valid values are:
	// - All: all dry run stages will be processed
	// +optional
	DryRun []string `json:"dryRun,omitempty"`

	// Force is going to "force" Apply requests. It means user will
	// re-acquire conflicting fields owned by other people. Force
	// flag must be unset for non-apply patch requests.
	// +optional
	Force bool `json:"force,omitempty"`
}

// UpdateOptions may be provided when updating an API object.
// All fields in UpdateOptions should also be present in PatchOptions.
type UpdateOptions struct {
	TypeMeta `json:",inline"`

	// When present, indicates that modifications should not be
	// persisted. An invalid or unrecognized dryRun directive will
	// result in an error response and no further processing of the
	// request. Valid values are:
	// - All: all dry run stages will be processed
	// +optional
	DryRun []string `json:"dryRun,omitempty"`
}

// AuthorizeOptions may be provided when authorize an API object.
type AuthorizeOptions struct {
	TypeMeta `json:",inline"`
}

// TableOptions are used when a Table is requested by the caller.
type TableOptions struct {
	TypeMeta `json:",inline"`

	// NoHeaders is only exposed for internal callers. It is not included in our OpenAPI definitions
	// and may be removed as a field in a future release.
	NoHeaders bool `json:"-"`
}
