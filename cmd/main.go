package main

import (
	"context"
	"os"
	"strconv"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/runtime"

	podcontroller "github.com/k8-proxy/go-k8s-controller/pkg/controller"
)

func main() {

	ctx, _ := context.WithCancel(context.Background())
	loggerConfig := zap.NewProductionConfig()

	// general logger
	logger, err := loggerConfig.Build()
	runtime.Must(err)

	podNamespace := "icap-adaptation"

	podCountStr := os.Getenv("POD_COUNT")
	minioUser := os.Getenv("MINIO_USER")
	minioPassword := os.Getenv("MINIO_PASSWORD")
	podCount, err := strconv.Atoi(podCountStr)
	if err != nil {
		podCount = 10 // default value
	}

	rs := &podcontroller.RebuildSettings{
		PodCount:      podCount,
		MinioUser:     minioUser,
		MinioPassword: minioPassword,
	}

	ctrl, err := podcontroller.NewPodController(logger, podNamespace, rs)
	if err != nil {
		logger.Panic("Failed to initialise the controller", zap.Error(err))
	}

	ctrl.Run(ctx)

	<-ctx.Done()
}
