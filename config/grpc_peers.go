package config

type gRPCPEers map[string]string

func (m gRPCPEers) Profile() string {
	return m["profile"]
}

func (m gRPCPEers) Chat() string {
	return m["chat"]
}

func (m gRPCPEers) S3() string {
	return m["s3"]
}
