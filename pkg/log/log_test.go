package log

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSetLevel(t *testing.T) {
	Create("hello world", false)
	if logrus.GetLevel() != logrus.InfoLevel {
		t.Errorf("The default level should be Info")
	}
}
