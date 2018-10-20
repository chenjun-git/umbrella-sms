package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"google.golang.org/grpc"

	"github.com/chenjun-git/umbrella-common/monitor"
	commonUtils "github.com/chenjun-git/umbrella-common/utils"

	"business/sms/common"
	"business/sms/handler"
	"business/sms/pb"
	"business/sms/server"
)

var (
	BuildTime    = "No Build Time"
	BuildGitHash = "No Build Git Hash"
	BuildGitTag  = "No Build Git Tag"
)

func main() {
	initConfig()
	initMonitor()

	go startGRPCServer()
	startHTTPServer()
}

func initConfig() {
	configPath := flag.String("config", "", "config file's path")
	flag.Parse()

	common.InitConfig(*configPath)
}

func initMonitor() {
	monitor.Init(common.Config.Monitor.Namespace, common.Config.Monitor.Subsystem)
	monitor.Monitor.SetVersion(monitor.Version{GitHash: BuildGitHash, GitTag: BuildGitTag, BuildTime: BuildTime})
}

func startGRPCServer() {
	// init tcp
	listen, err := net.Listen("tcp", common.Config.Listen)
	if err != nil {
		panic(err)
	}

	// init tracing
	// interceptors := []grpc.UnaryServerInterceptor{middlewares.Log, middlewares.Auth, caller.ExtractCallerNameUnary(), monitor.MonitorInceptorUnary()}
	// if tracingInterceptor := traceconf.GlobalTracingConfig.GrpcUnaryServerInterceptor(); tracingInterceptor != nil {
	// 	interceptors = append(interceptors, tracingInterceptor)
	// }

	// ui := grpcmiddleware.ChainUnaryServer(interceptors...)
	s := grpc.NewServer()

	smsServer := server.NewServer()
	smsServer.PatchYunzhixun(*common.Config.Yunzhixun)

	pb.RegisterSmsServer(s, smsServer)

	// log.Info("grpc: starting to serve grpc requests", zap.String("listen", common.Config.Listen))
	fmt.Printf("start grpc server, listen: %s!\n", common.Config.Listen)
	err = s.Serve(listen)
	if err != nil {
		// log.Error("grpc: service aborted", zap.Error(err))
	}
}

func startHTTPServer() {
	router := handler.RegisterSmsRouter()
	commonUtils.RegisterPProf(router)
	monitor.RegisterHandlers(router)

	httpServer := &http.Server{
		Addr:    common.Config.HTTP.Listen,
		Handler: router,
	}
	fmt.Printf("start http server, listen: %s!\n", common.Config.HTTP.Listen)
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Printf("start http server failed, err : %+v\n", err)
		panic("start http server failed")
	}
}
