module github.com/pawpawchat/core

go 1.22.2

replace github.com/pawpawchat/chat => /home/amicie/dev/pawpawchat/chat

replace github.com/pawpawchat/profile => /home/amicie/dev/pawpawchat/profile

require (
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	github.com/pawpawchat/chat v0.0.0-00010101000000-000000000000
	github.com/pawpawchat/profile v0.0.0-20240826042740-916b5a899f74
	github.com/stretchr/testify v1.9.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240827150818-7e3bb234dfed
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
)
