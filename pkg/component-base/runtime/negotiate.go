package runtime

import (
	"fmt"
	"github.com/xs0910/iam/pkg/component-base/json"
)

// NegotiateError is returned when a ClientNegotiator is unable to locate a serializer for the requested operation.
type NegotiateError struct {
	ContentType string
	Stream      bool
}

func (e NegotiateError) Error() string {
	if e.Stream {
		return fmt.Sprintf("no stream serializers registered for %s", e.ContentType)
	}
	return fmt.Sprintf("no serializers registered for %s", e.ContentType)
}

type apimachineryClientNegotiator struct{}

var _ ClientNegotiator = &apimachineryClientNegotiator{}

func (a *apimachineryClientNegotiator) Encoder() (Encoder, error) {
	return &apimachineryClientNegotiatorSerializer{}, nil
}

func (a *apimachineryClientNegotiator) Decoder() (Decoder, error) {
	return &apimachineryClientNegotiatorSerializer{}, nil
}

type apimachineryClientNegotiatorSerializer struct{}

func (a apimachineryClientNegotiatorSerializer) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
func (a *apimachineryClientNegotiatorSerializer) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// NewSimpleClientNegotiator will negotiate for a single serializer. This should only be used
// for testing or when the caller is taking responsibility for setting the GVK on encoded objects.
func NewSimpleClientNegotiator() ClientNegotiator {
	return &apimachineryClientNegotiator{}
}
