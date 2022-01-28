/*
Copyright 2021.
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

package v2

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v1 "github.com/kubeshop/testkube-operator/apis/script/v1"
)

// ConvertTo converts this Script to the Hub version (v1).
func (src *Script) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.Script)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.Type_ = src.Spec.Type_
	dst.Spec.Name = src.Spec.Name
	dst.Spec.Params = src.Spec.Params
	dst.Spec.InputType = src.Spec.InputType
	dst.Spec.Tags = src.Spec.Tags

	if src.Spec.Content != nil {
		dst.Spec.Content = src.Spec.Content.Content
	}

	if src.Spec.Content != nil && src.Spec.Content.Repository != nil {
		dst.Spec.Repository = &v1.Repository{
			Type_:    src.Spec.Content.Repository.Type_,
			Uri:      src.Spec.Content.Repository.Uri,
			Branch:   src.Spec.Content.Repository.Branch,
			Path:     src.Spec.Content.Repository.Path,
			Username: src.Spec.Content.Repository.Username,
			Token:    src.Spec.Content.Repository.Token,
		}
	}

	// Status

	return nil
}

// ConvertFrom converts Script from the Hub version (v1) to this version.
func (dst *Script) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.Script)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.Type_ = src.Spec.Type_
	dst.Spec.Name = src.Spec.Name
	dst.Spec.Params = src.Spec.Params
	dst.Spec.InputType = src.Spec.InputType
	dst.Spec.Tags = src.Spec.Tags

	dst.Spec.Content = &ScriptContent{
		Type_:   "file-string",
		Content: src.Spec.Content,
	}

	if src.Spec.Repository != nil {
		dst.Spec.Content = &ScriptContent{
			Type_: "git-dir",
			Repository: &Repository{
				Type_:    src.Spec.Repository.Type_,
				Uri:      src.Spec.Repository.Uri,
				Branch:   src.Spec.Repository.Branch,
				Path:     src.Spec.Repository.Path,
				Username: src.Spec.Repository.Username,
				Token:    src.Spec.Repository.Token,
			},
		}
	}

	// Status
	return nil
}
