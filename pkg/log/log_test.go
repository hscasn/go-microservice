package log

import (
	"testing"
)

func TestSetLevel(t *testing.T) {
	Create("hello world", true)
	if GetLevel() != InfoLevel {
		t.Errorf("The default level should be Info")
	}
	SetLevel(DebugLevel)
	if GetLevel() != DebugLevel {
		t.Errorf("The level should be Debug")
	}
	SetLevel(InfoLevel)
	if GetLevel() != InfoLevel {
		t.Errorf("The level should be Info")
	}
	SetLevel(WarnLevel)
	if GetLevel() != WarnLevel {
		t.Errorf("The level should be Warn")
	}
	SetLevel(ErrorLevel)
	if GetLevel() != ErrorLevel {
		t.Errorf("The level should be Error")
	}
	SetLevel(FatalLevel)
	if GetLevel() != FatalLevel {
		t.Errorf("The level should be Fatal")
	}
}

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name string
		l    Level
		want string
	}{
		{name: "Debug", l: DebugLevel, want: "debug"},
		{name: "Info", l: InfoLevel, want: "info"},
		{name: "Warn", l: WarnLevel, want: "warning"},
		{name: "Error", l: ErrorLevel, want: "error"},
		{name: "Fatal", l: FatalLevel, want: "fatal"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.String(); got != tt.want {
				t.Errorf("Level.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
