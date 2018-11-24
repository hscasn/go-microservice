package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-microservice/pkg/health"
	"go-microservice/pkg/testingtools"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	calledOnClose := false
	onClose := func() {
		calledOnClose = true
	}
	log := logrus.NewEntry(logrus.New())

	s := Create(log, health.Checks{}, onClose)
	go s.Start()
	defer func() {
		tries := 0
		for !calledOnClose {
			time.Sleep(100 * time.Millisecond)
			tries++
			if tries > 100 {
				break
			}
		}
		if !calledOnClose {
			t.Error("Should have called the onClose function")
		}
	}()
	defer func() { s.shutdown <- true }()

	for s.httpSrv == nil {
		time.Sleep(time.Millisecond * 100)
	}

	addr := fmt.Sprintf("http://%s", s.httpSrv.Addr)

	// Ready
	res, _ := testingtools.HTTPRequest(t, addr, "GET", "/ready")
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}
}
