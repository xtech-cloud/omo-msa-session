package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	_ "github.com/micro/go-plugins/registry/etcdv3/v2"
	proto "github.com/xtech-cloud/omo-msp-session/proto/session"
	"io"
	"omo.msa.session/cache"
	"omo.msa.session/config"
	"omo.msa.session/grpc"
	"os"
	"path/filepath"
	"time"
)

var (
	BuildVersion string
	BuildTime    string
	CommitID     string
)

func main() {
	config.Setup()
	err := cache.InitData()
	if err != nil {
		panic(err)
	}

	//test()

	// New Service
	service := micro.NewService(
		micro.Name("omo.msa.session"),
		micro.Version(BuildVersion),
		micro.RegisterTTL(time.Second*time.Duration(config.Schema.Service.TTL)),
		micro.RegisterInterval(time.Second*time.Duration(config.Schema.Service.Interval)),
		micro.Address(config.Schema.Service.Address),
	)
	// Initialise service
	service.Init()
	// Register Handler
	_ = proto.RegisterSessionServiceHandler(service.Server(), new(grpc.SessionService))

	app, _ := filepath.Abs(os.Args[0])

	logger.Info("-------------------------------------------------------------")
	logger.Info("- Micro Service Agent -> Run")
	logger.Info("-------------------------------------------------------------")
	logger.Infof("- version      : %s", BuildVersion)
	logger.Infof("- application  : %s", app)
	logger.Infof("- md5          : %s", md5hex(app))
	logger.Infof("- build        : %s", BuildTime)
	logger.Infof("- commit       : %s", CommitID)
	logger.Info("-------------------------------------------------------------")
	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

func test()  {
	user := "60e2a3b5547b471255e8c962"
	token,err := cache.CreateSession(user)
	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Println(token)
	}

	result,err := cache.ParseToken(token)
	fmt.Println(err)
	fmt.Println(result.Id)
}

func md5hex(_file string) string {
	h := md5.New()

	f, err := os.Open(_file)
	if err != nil {
		return ""
	}
	defer f.Close()

	io.Copy(h, f)

	return hex.EncodeToString(h.Sum(nil))
}
