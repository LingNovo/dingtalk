package main

import (
	"net"

	"github.com/LingNovo/dingtalk/protos/oapi"
	"github.com/LingNovo/dingtalk/protos/tapi"
	"github.com/LingNovo/dingtalk/server"
	"github.com/Unknwon/goconfig"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	log      = logrus.New()
	confPath = "./conf/app.conf"
)

func loadConfig(section string, cf *goconfig.ConfigFile) server.DTConfig {
	cfg := server.DTConfig{}
	cfg.CorpId = cf.MustValue(section, "CorpId")
	cfg.CorpSecret = cf.MustValue(section, "CorpSecret")
	cfg.AgentId = cf.MustValue(section, "AgentId")
	cfg.SSOSecret = cf.MustValue(section, "SSOSecret")
	cfg.SNSAppId = cf.MustValue(section, "SNSAppId")
	cfg.SNSSecret = cf.MustValue(section, "SNSSecret")
	return cfg
}

func main() {
	conf, err := goconfig.LoadConfigFile(confPath)
	if err != nil {
		log.Fatalf("load config: %s", err)
	}
	cfg := loadConfig("app", conf)
	address := conf.MustValue("server", "address")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
	cos := server.NewCompanyOapiSrv(&cfg)
	srv := grpc.NewServer()
	oapi.RegisterOapiWarpperServer(srv, cos)
	tapi.RegisterTapiWarpperServer(srv, cos)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
