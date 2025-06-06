package apipb

import "google.golang.org/protobuf/encoding/protojson"

func (x *Api) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Api) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *Method) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Method) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *Mixin) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Mixin) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}