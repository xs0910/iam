package scheme

// ObjectKind is used by serialization to set type information from the Scheme onto the serialized version of an object.
// For objects that cannot be serialized or have unique requirements, this interface may be a no-op.
type ObjectKind interface {
	// SetGroupVersionKind sets or clears the intended serialized kind of an object. Passing kind nil
	// should clear the current setting.
	SetGroupVersionKind(kind GroupVersionKind)
	// GroupVersionKind returns the stored group, version, and kind of object, or nil if the object does
	// not expose or provide these fields.
	GroupVersionKind() GroupVersionKind
}

type emptyObjectKind struct{}

// EmptyObjectKind implements the ObjectKind interface as a noop.
var EmptyObjectKind = emptyObjectKind{}

// SetGroupVersionKind implements the ObjectKind interface.
func (emptyObjectKind) SetGroupVersionKind(gvk GroupVersionKind) {}

// GroupVersionKind implements the ObjectKind interface.
func (emptyObjectKind) GroupVersionKind() GroupVersionKind { return GroupVersionKind{} }