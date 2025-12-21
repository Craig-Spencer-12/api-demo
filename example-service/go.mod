module example-service

go 1.25.5

replace common => ../common

require common v0.0.0

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/segmentio/kafka-go v0.4.49 // indirect
)
