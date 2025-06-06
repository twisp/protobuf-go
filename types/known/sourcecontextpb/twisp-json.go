package sourcecontextpb

import "google.golang.org/protobuf/encoding/protojson"

func (x *SourceContext) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *SourceContext) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}