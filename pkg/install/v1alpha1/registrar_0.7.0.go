/*
Copyright 2018 The OpenEBS Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func RegisteredArtifactsFor070() (finallist ArtifactList) {
	finallist.Items = append(finallist.Items, CstorVolumeArtifactsFor070().Items...)
	finallist.Items = append(finallist.Items, JivaVolumeArtifactsFor070().Items...)

	return
}

// WithGVRCasTemplateListUpdaterFor070 updates list of given artifact
// instances with CASTemplate based Group Version Resource information
func WithGVRCasTemplateListUpdaterFor070(given []*Artifact) (updated []*Artifact) {
	gvrUpdate := WithGVRCasTemplateUpdaterFor070()
	for _, artifact := range given {
		updated = append(updated, gvrUpdate(artifact))
	}
	return
}

// WithGVRCasTemplateUpdaterFor070 updates the given artifact instance with
// CASTemplate based Group Version Resource information
func WithGVRCasTemplateUpdaterFor070() ArtifactMiddleware {
	return func(given *Artifact) (updated *Artifact) {
		given.GroupVersionResource = schema.GroupVersionResource{
			Group:    "openebs.io",
			Version:  "v1alpha1",
			Resource: "castemplates",
		}
		return given
	}
}

// WithGVRRunTaskListUpdaterFor070 updates list of given artifact
// instances with RunTask based Group Version Resource information
func WithGVRRunTaskListUpdaterFor070(given []*Artifact) (updated []*Artifact) {
	gvrUpdate := WithGVRRunTaskUpdaterFor070()
	for _, artifact := range given {
		updated = append(updated, gvrUpdate(artifact))
	}
	return
}

// WithGVRRunTaskUpdaterFor070 updates the given artifact instance with
// RunTask based Group Version Resource information
func WithGVRRunTaskUpdaterFor070() ArtifactMiddleware {
	return func(given *Artifact) (updated *Artifact) {
		given.GroupVersionResource = schema.GroupVersionResource{
			Group:    "openebs.io",
			Version:  "v1alpha1",
			Resource: "runtasks",
		}
		return given
	}
}
