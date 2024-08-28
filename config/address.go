package config

type addresses map[string]string

func (m addresses) Profile() string {
	return m["profile"]
}

func (m addresses) Chat() string {
	return m["chat"]
}

func (m addresses) S3() string {
	return m["s3"]
}
