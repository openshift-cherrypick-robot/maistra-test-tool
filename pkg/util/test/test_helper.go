package test

import (
	"fmt"
	"testing"
	"time"
)

func NewTestContext(t *testing.T) TestHelper {
	ctx := &testHelper{
		t: t,
	}
	return ctx
}

type TestHelper interface {
	Name() string
	Cleanup(f func())
	Fail()
	FailNow()
	Failed() bool
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Log(args ...any)
	Logf(format string, args ...any)
	Helper()

	NewSubTest(name string) Test

	LogStep(str string)

	T() *testing.T

	WillRetry() bool
}

type testHelper struct {
	t           *testing.T
	currentStep int
}

var _ TestHelper = &testHelper{}

func (t *testHelper) Name() string {
	return t.t.Name()
}

func (t *testHelper) Fail() {
	t.t.Fail()
}

func (t *testHelper) FailNow() {
	t.t.FailNow()
}

func (t *testHelper) Failed() bool {
	return t.t.Failed()
}

func (t *testHelper) Helper() {
	t.t.Helper()
}

func (t *testHelper) Log(args ...any) {
	t.t.Helper()
	t.t.Log(args...)
}

func (t *testHelper) Logf(format string, args ...any) {
	t.t.Helper()
	t.t.Logf(format, args...)
}

func (t *testHelper) Error(args ...any) {
	t.t.Helper()
	t.Log("ERROR: " + fmt.Sprint(args...))
	t.Fail()
}

func (t *testHelper) Errorf(format string, args ...any) {
	t.t.Helper()
	t.Logf("ERROR: "+format, args...)
	t.Fail()
}

func (t *testHelper) Fatal(args ...any) {
	t.t.Helper()
	t.Log("FATAL: " + fmt.Sprint(args...))
	t.FailNow()
}

func (t *testHelper) Fatalf(format string, args ...any) {
	t.t.Helper()
	t.Logf("FATAL: "+format, args...)
	t.FailNow()
}

func (t *testHelper) Cleanup(f func()) {
	t.t.Helper()
	t.t.Cleanup(func() {
		t.T().Helper()
		start := time.Now()
		t.Log("Performing cleanup")
		f()
		t.Logf("Cleanup completed in %.2fs", time.Now().Sub(start).Seconds())
	})
}

func (t *testHelper) LogStep(str string) {
	t.t.Helper()
	t.currentStep++
	if t.currentStep > 1 {
		t.Log("")
	}
	t.Logf("STEP %d: %s", t.currentStep, str)
	t.Log("")
}

func (t *testHelper) NewSubTest(name string) Test {
	return subTest{
		t:    t.t,
		name: name,
	}
}

func (t *testHelper) T() *testing.T {
	return t.t
}

func (t *testHelper) WillRetry() bool {
	return false
}
