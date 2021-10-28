// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package installer

import (
	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmware-tanzu/cluster-api-provider-byoh/agent/installer/internal/algo"
)

var _ = Describe("Byohost Installer Tests", func() {
	Context("When installer is created for unsupported OS", func() {
		It("Should return error", func() {
			_, err := New("repo", "downloadPath", logr.Discard())
			Expect(err).Should((HaveOccurred()))
		})
	})
	Context("When installer is created with empty bundle repo", func() {
		It("Should return error", func() {
			_, err := New("", "downloadPath", logr.Discard())
			Expect(err).Should((HaveOccurred()))
		})
	})
	Context("When installer is created with empty download path", func() {
		It("Should return error", func() {
			_, err := New("repo", "", logr.Discard())
			Expect(err).Should((HaveOccurred()))
		})
	})
	Context("When installer is created", func() {
		It("Install/uninstall should return error for unsupported k8s", func() {
			for _, os := range ListSupportedOS() {
				i := NewPreviewInstaller(os, nil)

				err := i.Install("unsupported-k8s")
				Expect(err).Should((HaveOccurred()))

				err = i.Uninstall("unsupported-k8s")
				Expect(err).Should((HaveOccurred()))
			}
		})
	})
	Context("When installer is created", func() {
		It("Install/uninstall should call only the output builder", func() {
			for _, os := range ListSupportedOS() {
				for _, k8s := range ListSupportedK8s(os) {
					{
						ob := algo.OutputBuilderCounter{}
						i := NewPreviewInstaller(os, &ob)
						err := i.Install(k8s)
						Expect(err).ShouldNot((HaveOccurred()))
						Expect(ob.LogCalledCnt).Should(Equal(24))
					}

					{
						ob := algo.OutputBuilderCounter{}
						i := NewPreviewInstaller(os, &ob)
						err := i.Uninstall(k8s)
						Expect(err).ShouldNot((HaveOccurred()))
						Expect(ob.LogCalledCnt).Should(Equal(24))
					}
				}
			}
		})
	})
	Context("When ListSupportedOS is called", func() {
		It("Should return non-empty result", func() {
			Expect(ListSupportedOS()).ShouldNot(BeEmpty())
		})
	})
	Context("When ListSupportedK8s is called for all supported OSes", func() {
		It("Should return non-empty result", func() {
			for _, os := range ListSupportedOS() {
				Expect(ListSupportedK8s(os)).ShouldNot(BeEmpty())
			}
		})
	})
	Context("When PreviewChanges is called for all supported os and k8s", func() {
		It("Should not return error", func() {
			for _, os := range ListSupportedOS() {
				for _, k8s := range ListSupportedK8s(os) {
					_, _, err := PreviewChanges(os, k8s)
					Expect(err).ShouldNot((HaveOccurred()))
				}
			}
		})
	})
})

func NewPreviewInstaller(os string, ob algo.OutputBuilder) *installer {
	i, err := newUnchecked(os, "", "", logr.Discard(), ob)
	if err != nil {
		panic(err)
	}
	return i
}
