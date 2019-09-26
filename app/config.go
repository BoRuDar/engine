package app

import (
	"flag"
	"fmt"
	"github.com/webitel/engine/model"
)

var (
	consulHost     = flag.String("consul", "172.0.0.1:8500", "Host to consul")
	websocketHost  = flag.String("ws", ":10025", "WebSocket server address")
	dataSource     = flag.String("data_source", "postgres://webitel:webitel@localhost:5432/webitel?fallback_application_name=engine&sslmode=disable&connect_timeout=10&search_path=call_center", "WebSocket server address")
	amqpSource     = flag.String("amqp", "amqp://webitel:webitel@localhost:5672?heartbeat=10", "AMQP connection")
	grpcServerPort = flag.Int("grpc_port", 0, "GRPC port")
)

func (app *App) Config() *model.Config {
	return app.config
}

func loadConfig() (*model.Config, error) {
	flag.Parse()
	config := &model.Config{
		NodeName: fmt.Sprintf("engine-%s", model.NewId()),
		DiscoverySettings: model.DiscoverySettings{
			Url: *consulHost,
		},
		WebSocketSettings: model.WebSocketSettings{
			Address: *websocketHost,
		},
		ServerSettings: model.ServerSettings{
			Address: "",
			Port:    *grpcServerPort,
			Network: "tcp",
		},
		SqlSettings: model.SqlSettings{
			DriverName:                  model.NewString("postgres"),
			DataSource:                  dataSource,
			MaxIdleConns:                model.NewInt(5),
			MaxOpenConns:                model.NewInt(5),
			ConnMaxLifetimeMilliseconds: model.NewInt(3600000),
			Trace:                       true,
		},
		MessageQueueSettings: model.MessageQueueSettings{
			Url: *amqpSource,
		},
	}
	return config, nil
}
