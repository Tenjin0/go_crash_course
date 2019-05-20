package go_crash_course_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoCrashCourse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoCrashCourse Suite")
}
