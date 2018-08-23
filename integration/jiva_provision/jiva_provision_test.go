package jiva_provision

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
	//. "github.com/AmitKumarDas/decide/integration"
)

// This test case tests basic provisioning for jiva based openebs storage engine
var _ = Describe("Provision Jiva", func() {
	Context("with openebs operator components", func() {
		It("should create and delete jiva volume", func() {

			// Iter - 1
			// Take(map[string]interface{}{}).Run("").Expect(Success())
			// Kubectl().Get("svc").Name("").Should(HaveClusterIP(""))
			// Kubectl().Apply(yaml).Should(Succeed())

			// Iter - 2
			// k := Kubectl().Apply(perconaYaml)
			// Expect(k.Build()).To(Equal("true"))

			// Iter - 2.1
			// Expect(Kubectl().Apply(perconaYaml).Build()).To(Equal("true"))
		})
	})

})
