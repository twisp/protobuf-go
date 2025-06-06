package typepb

import "google.golang.org/protobuf/encoding/protojson"

func (x *Type) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Type) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *Field) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Field) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *Enum) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Enum) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *EnumValue) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *EnumValue) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *Option) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Option) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}