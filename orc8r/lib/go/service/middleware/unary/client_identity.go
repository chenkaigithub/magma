/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package unary

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	// Client Certificate Serial Number Header
	CLIENT_CERT_SN_KEY = "x-magma-client-cert-serial"

	// Magic Certificate Value for orc8r service clients
	ORC8R_CLIENT_CERT_VALUE = "7ZZXAF7CAETF241KL22B8YRR7B5UF401"
)

// CloudClientInterceptor sets Magic Certificate Value for orc8r service clients in the outgoing CTX
// if the CTX metadata already has a client certificate SN key, CloudClientInterceptor will overwrite it with
// the Magic Certificate Value
func CloudClientInterceptor(
	ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	md, exists := metadata.FromOutgoingContext(ctx)
	if exists {
		md = md.Copy()
		md.Set(CLIENT_CERT_SN_KEY, ORC8R_CLIENT_CERT_VALUE)
	} else {
		md = metadata.Pairs(CLIENT_CERT_SN_KEY, ORC8R_CLIENT_CERT_VALUE)
	}
	outgoingCtx := metadata.NewOutgoingContext(ctx, md)
	err := invoker(outgoingCtx, method, req, reply, cc, opts...)
	return err
}
