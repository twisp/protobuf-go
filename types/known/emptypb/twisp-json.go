package emptypb

import "google.golang.org/protobuf/encoding/protojson"

func (x *Empty) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Empty) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}