module github.com/ahiho/go-pkg/logger/grpc

go 1.19

replace go-pkg/logger v0.0.0-local => ../

require (
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.5
	go-pkg/logger v0.0.0-local
	google.golang.org/grpc v1.54.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/rs/zerolog v1.29.1 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
