package anypb

import "google.golang.org/protobuf/encoding/protojson"

func (x *Any) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Any) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}
