/*
Copyright 2019 The Tekton Authors

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

package pipeline

import (
	"fmt"
	"sort"
)

// Images holds the images reference for a number of container images used
// across tektoncd pipelines.
type Images struct {
	// EntrypointImage is container image containing our entrypoint binary.
	EntrypointImage string
	// SidecarLogResultsImage is container image containing the binary that fetches results from the steps and logs it to stdout.
	SidecarLogResultsImage string
	// NopImage is the container image used to kill sidecars.
	NopImage string
	// GitImage is the container image with Git that we use to implement the Git source step.
	GitImage string
	// ShellImage is the container image containing bash shell.
	ShellImage string
	// ShellImageWin is the container image containing powershell.
	ShellImageWin string
	// GsutilImage is the container image containing gsutil.
	GsutilImage string
	// WorkingDirInitImage is the container image containing our working dir init binary.
	WorkingDirInitImage string

	// NOTE: Make sure to add any new images to Validate below!
}

// Validate returns an error if any image is not set.
func (i Images) Validate() error {
	var unset []string
	for _, f := range []struct {
		v, name string
	}{
		{i.EntrypointImage, "entrypoint-image"},
		{i.SidecarLogResultsImage, "sidecarlogresults-image"},
		{i.NopImage, "nop-image"},
		{i.GitImage, "git-image"},
		{i.ShellImage, "shell-image"},
		{i.ShellImageWin, "shell-image-win"},
		{i.GsutilImage, "gsutil-image"},
		{i.WorkingDirInitImage, "workingdirinit-image"},
	} {
		if f.v == "" {
			unset = append(unset, f.name)
		}
	}
	if len(unset) > 0 {
		sort.Strings(unset)
		return fmt.Errorf("found unset image flags: %s", unset)
	}
	return nil
}
