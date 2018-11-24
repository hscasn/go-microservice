package env

import (
	"os"
	"testing"
)

func TestStringWithDefault_UseDefault(t *testing.T) {
	os.Setenv("_SOME_VARIABLE", "")
	r := StringWithDefault("_SOME_VARIABLE", "my variable")
	if r != "my variable" {
		t.Errorf("should have given a default value")
	}
}

func TestStringWithDefault_NoDefault(t *testing.T) {
	os.Setenv("_SOME_VARIABLE", "hello")
	r := StringWithDefault("_SOME_VARIABLE", "my variable")
	if r != "hello" {
		t.Errorf("should not have given a default value")
	}
}

func TestString_WithPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("should have panicked")
		}
	}()
	os.Setenv("_SOME_VARIABLE", "")
	String("_SOME_VARIABLE")
}

func TestString_NoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("should not have panicked")
		}
	}()
	os.Setenv("_SOME_VARIABLE", "hello")
	r := String("_SOME_VARIABLE")
	if r != "hello" {
		t.Errorf("should have the correct value")
	}
}

func TestIntWithDefault_UseDefault(t *testing.T) {
	os.Setenv("_SOME_VARIABLE", "")
	r := IntWithDefault("_SOME_VARIABLE", 5)
	if r != 5 {
		t.Errorf("should have given a default value")
	}
}

func TestIntWithDefault_NoDefault(t *testing.T) {
	os.Setenv("_SOME_VARIABLE", "hello")
	r := IntWithDefault("_SOME_VARIABLE", 5)
	if r != 5 {
		t.Errorf("should not have given a default value")
	}
}

func TestInt_WithPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("should have panicked")
		}
	}()
	os.Setenv("_SOME_VARIABLE", "")
	Int("_SOME_VARIABLE")
}

func TestInt_NoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("should not have panicked")
		}
	}()
	os.Setenv("_SOME_VARIABLE", "5")
	r := Int("_SOME_VARIABLE")
	if r != 5 {
		t.Errorf("should have the correct value")
	}
}

func TestFloatWithDefault_UseDefault(t *testing.T) {
	os.Setenv("_SOME_VARIABLE", "")
	r := FloatWithDefault("_SOME_VARIABLE", 5)
	if r != 5 {
		t.Errorf("should have given a default value")
	}
}

func TestFloatWithDefault_NoDefault(t *testing.T) {
	os.Setenv("_SOME_VARIABLE", "hello")
	r := FloatWithDefault("_SOME_VARIABLE", 5)
	if r != 5 {
		t.Errorf("should not have given a default value")
	}
}

func TestFloat_WithPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("should have panicked")
		}
	}()
	os.Setenv("_SOME_VARIABLE", "")
	Float("_SOME_VARIABLE")
}

func TestFloat_NoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("should not have panicked")
		}
	}()
	os.Setenv("_SOME_VARIABLE", "5")
	r := Float("_SOME_VARIABLE")
	if r != 5 {
		t.Errorf("should have the correct value")
	}
}

func TestBoolWithDefault_UseDefault(t *testing.T) {
	os.Setenv("_SOME_VARIABLE", "")
	r := BoolWithDefault("_SOME_VARIABLE", true)
	if r != true {
		t.Errorf("should have given a default value")
	}

	r = BoolWithDefault("_SOME_VARIABLE", false)
	if r != false {
		t.Errorf("should have given a default value")
	}
}

func TestBoolWithDefault_NoDefault(t *testing.T) {
	os.Setenv("_SOME_VARIABLE", "false")
	r := BoolWithDefault("_SOME_VARIABLE", true)
	if r != false {
		t.Errorf("should not have given a default value")
	}
	os.Setenv("_SOME_VARIABLE", "true")
	r = BoolWithDefault("_SOME_VARIABLE", false)
	if r != true {
		t.Errorf("should not have given a default value")
	}
}

func TestBool_WithPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("should have panicked")
		}
	}()
	os.Setenv("_SOME_VARIABLE", "")
	Bool("_SOME_VARIABLE")
}

func TestBool_NoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("should not have panicked")
		}
	}()
	os.Setenv("_SOME_VARIABLE", "true")
	r := Bool("_SOME_VARIABLE")
	if r != true {
		t.Errorf("should have the correct value")
	}
}
