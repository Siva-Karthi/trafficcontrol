package client

import (
	"github.com/apache/trafficcontrol/lib/go-tc"
	"net/http"
)

/*

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

func (to *Session) GetDomainsWithHdr(header http.Header) ([]tc.Domain, ReqInf, error) {
	var data tc.DomainsResponse
	inf, err := get(to, apiBase+"/cdns/domains", &data, header)
	if err != nil {
		return nil, inf, err
	}
	return data.Response, inf, nil
}

// Deprecated: GetDomains will be removed in 6.0. Use GetDomainsWithHdr.
func (to *Session) GetDomains() ([]tc.Domain, ReqInf, error) {
	return to.GetDomainsWithHdr(nil)
}
