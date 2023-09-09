package main

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"testing"
	"time"
)

//goos: darwin
//goarch: amd64
//pkg: compare
//cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
//BenchmarkLogrus-8         350319              3451 ns/op            1365 B/op         25 allocs/op
//BenchmarkZap-8           1484893               804.8 ns/op           192 B/op          1 allocs/op
//PASS

// go test -bench .
func BenchmarkLogrus(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := logrus.New()
	logger.SetOutput(io.Discard)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.WithFields(logrus.Fields{
			"url":     "http://foo.com",
			"attempt": 3,
			"backoff": time.Second,
		}).Info("failed to fetch URL")
	}
}

func BenchmarkZap(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	)
	logger := zap.New(core)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("failed to fetch URL",
			zap.String("url", `http://foo.com`),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
	}
}
