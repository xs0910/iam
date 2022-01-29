package v1

import "time"

type Object interface {
	GetID() uint64
	SetID(id uint64)
	GetName() string
	SetName(name string)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
}

type ObjectMetaAccessor interface {
	GetObjectMeta() Object
}

// Type exposes the type and APIVersion of versioned or internal API objects.
type Type interface {
	GetAPIVersion() string
	SetAPIVersion(version string)
	GetKind() string
	SetKind(kind string)
}

// ListInterface lets you work with list metadata from any of the versioned or
// internal API objects. Attempting to set or retrieve a field on an object that does
// not support that field will be a no-op and return a default value.
type ListInterface interface {
	GetTotalCount() int64
	SetTotalCount(count int64)
}

var _ Object = &ObjectMeta{}
var _ Type = &TypeMeta{}
var _ ListInterface = &ListMeta{}
