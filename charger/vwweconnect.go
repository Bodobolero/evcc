package charger

// MIT LICENSE

// Copyright (c) 2024 bodobolero

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

import (
	"errors"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
)

func init() {
	registry.Add("vwweconnect", NewVWWeConnectFromConfig)
}

func NewVWWeConnectFromConfig(other map[string]interface{}) (api.Charger, error) {
	cc := struct {
		Token string
		URI   string
		Cache time.Duration
	}{
		Cache: time.Second,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if (cc.URI != "") == (cc.Token != "") {
		return nil, errors.New("must have either uri or token")
	}

	return NewGoE(cc.URI, cc.Token, cc.Cache)
}
