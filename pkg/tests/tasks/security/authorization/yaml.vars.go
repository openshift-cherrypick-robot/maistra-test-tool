// Copyright 2021 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package authorizaton

import (
	"github.com/maistra/maistra-test-tool/pkg/util"
	"github.com/maistra/maistra-test-tool/pkg/util/env"
)

type SMCP struct {
	Name      string `default:"basic"`
	Namespace string `default:"istio-system"`
}

var (
	smcpName       string = env.Getenv("SMCPNAME", "basic")
	meshNamespace  string = env.Getenv("MESHNAMESPACE", "istio-system")
	smcp           SMCP   = SMCP{smcpName, meshNamespace}
	gatewayHTTP, _        = util.ShellSilent(`kubectl get routes -n %s istio-ingressgateway -o jsonpath='{.spec.host}'`, meshNamespace)
)
