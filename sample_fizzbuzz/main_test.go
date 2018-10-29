package main_test

import (
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var pathToBinary string
var session *gexec.Session

var _ = BeforeSuite(func() {
	var err error
	pathToBinary, err = gexec.Build("github.com/bborbe/sample_go/sample_fizzbuzz")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterEach(func() {
	session.Interrupt()
	Eventually(session).Should(gexec.Exit())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

var _ = Describe("FizzBuzz", func() {
	var err error
	It("prints the number", func() {
		session, err = gexec.Start(exec.Command(pathToBinary, "1"), GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		session.Wait()
		Expect(session.Out).To(Say("1"))
	})
	It("prints Fizz if the number is divisible by 3", func() {
		session, err = gexec.Start(exec.Command(pathToBinary, "3"), GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		session.Wait()
		Expect(session.Out).To(Say("Fizz"))
	})
	It("prints Buzz if the number is divisible by 5", func() {
		session, err = gexec.Start(exec.Command(pathToBinary, "5"), GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		session.Wait()
		Expect(session.Out).To(Say("Buzz"))
	})
	It("prints FizzBuzz if the number is divisible by 3 and 5", func() {
		session, err = gexec.Start(exec.Command(pathToBinary, "15"), GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		session.Wait()
		Expect(session.Out).To(Say("FizzBuzz"))
	})
})

func TestSampleFizzbuzz(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SampleFizzbuzz Suite")
}
