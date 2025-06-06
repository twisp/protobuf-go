package fieldmaskpb

import "google.golang.org/protobuf/encoding/protojson"

func (x *FieldMask) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *FieldMask) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}