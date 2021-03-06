package testutil

// TestProto test protocol description
var baseProto3 = `syntax = "proto3";

package proto3;

option go_package = "github.com/golang/protobuf/protoc-gen-go/testdata/proto3";

message Request {
    enum Flavour {
        SWEET = 0;
        SOUR = 1;
        UMAMI = 2;
        GOPHERLICIOUS = 3;
    }
    string name = 1;
    repeated int64 key = 2;
    Flavour taste = 3;
    Book book = 4;
    repeated int64 unpacked = 5 [packed=false];
}

message Book {
    string title = 1;
    bytes raw_data = 2;
}`

