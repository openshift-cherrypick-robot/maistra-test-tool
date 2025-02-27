package test

import (
	"os"
	"testing"
)

type SetupFunc func()

type TestSuite interface {
	Run()
	Setup(SetupFunc) TestSuite
}

type testSuite struct {
	m        *testing.M
	setupFns []SetupFunc
}

func (s *testSuite) Run() {
	for _, setupFn := range s.setupFns {
		setupFn()
	}

	exitCode := s.m.Run()
	os.Exit(exitCode)
}

func (s *testSuite) Setup(f SetupFunc) TestSuite {
	s.setupFns = append(s.setupFns, f)
	return s
}

func NewSuite(m *testing.M) TestSuite {
	return &testSuite{m: m}
}
