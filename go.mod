module github.com/webitel/engine

go 1.18

require (
	github.com/go-gorp/gorp v2.2.0+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/hashicorp/consul/api v1.12.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/lib/pq v1.10.5
	github.com/nicksnyder/go-i18n v1.10.1
	github.com/pborman/uuid v1.2.1
	github.com/pkg/errors v0.9.1
	github.com/streadway/amqp v1.0.0
	github.com/webitel/call_center v0.0.0-20220506100737-7c57b182e77c
	github.com/webitel/protos/cc v0.0.0-20220512123629-c2f00de6990d
	github.com/webitel/protos/engine v0.0.0-20220512095803-4e77436dbc3d
	github.com/webitel/wlog v0.0.0-20190823170623-8cc283b29e3e
	go.uber.org/atomic v1.9.0
	go.uber.org/ratelimit v0.2.0
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4
	google.golang.org/grpc v1.46.0
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/armon/go-metrics v0.3.11 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.2.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.9.7 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/poy/onpar v1.1.2 // indirect
	github.com/ziutek/mymysql v1.5.4 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220505152158-f39f71e6c8f3 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.27.0
