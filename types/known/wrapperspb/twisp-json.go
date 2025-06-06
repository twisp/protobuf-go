package wrapperspb

import "google.golang.org/protobuf/encoding/protojson"

func (x *DoubleValue) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *DoubleValue) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *FloatValue) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *FloatValue) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *Int64Value) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Int64Value) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *UInt64Value) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *UInt64Value) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *Int32Value) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *Int32Value) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *UInt32Value) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *UInt32Value) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *BoolValue) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *BoolValue) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *StringValue) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *StringValue) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}

func (x *BytesValue) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(x)
}

func (x *BytesValue) UnmarshalJSON(b []byte) error {
	return protojson.Unmarshal(b, x)
}