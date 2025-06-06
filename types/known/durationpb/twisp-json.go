package durationpb

import "google.golang.org/protobuf/encoding/protojson"

func (x *Duration) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Duration) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}