/*
 * Copyright 2019 gosoon.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kubeadm

// APIServerPort is the expected default APIServerPort on the control plane node(s)
// https://kubernetes.io/docs/reference/access-authn-authz/controlling-access/#api-server-ports-and-ips
const APIServerPort = 6443

// Token defines a dummy, well known token for automating TLS bootstrap process
const Token = "abcdef.0123456789abcdef"

// ObjectName is the name every generated object will have
// I.E. `metadata:\nname: config`
const ObjectName = "config"
