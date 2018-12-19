package health

import (
	"testing"
)

type dummyCheckThatSucceeds struct{}

func (d *dummyCheckThatSucceeds) Ping() bool {
	return true
}

type dummyCheckThatFails struct{}

func (d *dummyCheckThatFails) Ping() bool {
	return false
}

func TestCreate(t *testing.T) {
	t.Parallel()

	var c map[string]string

	c = New(Checks{})
	if len(c) != 0 {
		t.Errorf("should have returned an empty map")
	}
	c = New(Checks{
		"client1": &dummyCheckThatSucceeds{},
		"client2": &dummyCheckThatFails{},
	})
	if len(c) != 2 {
		t.Errorf("should have returned a map with two checks")
	}
	if c["client1"] != HEALTHY {
		t.Errorf("client1 should be healthy")
	}
	if c["client2"] != UNHEALTHY {
		t.Errorf("client2 should be unhealthy")
	}
}

func TestCreateSummarized(t *testing.T) {
	t.Parallel()

	var c string

	c = CreateSummarized(Checks{})
	if c != HEALTHY {
		t.Errorf("should be healthy for no clients")
	}

	c = CreateSummarized(Checks{
		"client1": &dummyCheckThatSucceeds{},
		"client2": &dummyCheckThatSucceeds{},
	})
	if c != HEALTHY {
		t.Errorf("should be healthy for two clients that ping")
	}

	c = CreateSummarized(Checks{
		"client1": &dummyCheckThatFails{},
		"client2": &dummyCheckThatFails{},
	})
	if c != UNHEALTHY {
		t.Errorf("should be unhealthy for two clients that dont ping")
	}

	c = CreateSummarized(Checks{
		"client1": &dummyCheckThatSucceeds{},
		"client2": &dummyCheckThatFails{},
	})
	if c != UNHEALTHY {
		t.Errorf("should be unhealthy if one client fails")
	}
}

func TestSummarize(t *testing.T) {
	t.Parallel()

	var c map[string]string

	c = New(Checks{})
	if Summarize(c) != HEALTHY {
		t.Errorf("should be healthy for no clients")
	}

	c = New(Checks{
		"client1": &dummyCheckThatSucceeds{},
		"client2": &dummyCheckThatSucceeds{},
	})
	if Summarize(c) != HEALTHY {
		t.Errorf("should be healthy for two clients that ping")
	}

	c = New(Checks{
		"client1": &dummyCheckThatFails{},
		"client2": &dummyCheckThatFails{},
	})
	if Summarize(c) != UNHEALTHY {
		t.Errorf("should be unhealthy for two clients that dont ping")
	}

	c = New(Checks{
		"client1": &dummyCheckThatSucceeds{},
		"client2": &dummyCheckThatFails{},
	})
	if Summarize(c) != UNHEALTHY {
		t.Errorf("should be unhealthy if one client fails")
	}
}
