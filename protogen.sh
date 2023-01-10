protoc \
  --go_out=. \
  --go_opt=module=github.com/aligang/YandexPracticumGoAdvanced \
  --go-grpc_out=. \
  --go-grpc_opt=module=github.com/aligang/YandexPracticumGoAdvanced \
  proto/service.proto \
  proto/metric.proto
