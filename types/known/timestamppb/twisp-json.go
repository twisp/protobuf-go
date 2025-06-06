package timestamppb

import (
	"google.golang.org/protobuf/encoding/protojson"
)

func (x *Timestamp) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Timestamp) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}
