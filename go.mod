module github.com/sofiastockholmcode/payment-service

go 1.19

require (
	buf.build/gen/go/buftestorg/itemapis/bufbuild/connect-go v1.9.0-20230704110622-8e3e8cde4079.1
	buf.build/gen/go/buftestorg/itemapis/protocolbuffers/go v1.31.0-20230704110622-8e3e8cde4079.1
	buf.build/gen/go/buftestorg/user-service/bufbuild/connect-go v1.9.0-20230704090155-a75bd819adf4.1
	buf.build/gen/go/buftestorg/user-service/protocolbuffers/go v1.31.0-20230704090155-a75bd819adf4.1
	github.com/bufbuild/connect-go v1.9.0
	github.com/herrberk/go-http2-streaming v0.0.0-20230224235236-7741e51bed90
)

require (
	buf.build/gen/go/buftestorg/itemapis/grpc/go v1.3.0-20230704110622-8e3e8cde4079.1 // indirect
	buf.build/gen/go/buftestorg/user-service/grpc/go v1.3.0-20230704090155-a75bd819adf4.1 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230629202037-9506855d4529 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
