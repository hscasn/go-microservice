package server

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"go-microservice/pkg/health"
	"go-microservice/pkg/testingtools"

	"github.com/sirupsen/logrus"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	calledOnClose := false
	onClose := func() {
		calledOnClose = true
	}
	log := logrus.NewEntry(logrus.New())

	s := Create(log, health.Checks{}, 8001, onClose)
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
	defer func() { s.Shutdown <- true }()

	for s.httpSrv == nil {
		time.Sleep(time.Millisecond * 100)
	}

	addr := fmt.Sprintf("http://%s", s.httpSrv.Addr)

	// Ready
	res, _, err := testingtools.HTTPRequest(addr, "GET", "/ready")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("should get OK status")
	}
}

func TestCreateThatFails(t *testing.T) {
	t.Parallel()
	calledOnClose := false
	onClose := func() {
		calledOnClose = true
	}
	log := logrus.NewEntry(logrus.New())

	s := Create(log, health.Checks{}, 1, onClose)
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
	defer func() { s.Shutdown <- true }()

	for s.httpSrv == nil {
		time.Sleep(time.Millisecond * 100)
	}
}
