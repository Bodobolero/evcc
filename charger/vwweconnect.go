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
	"fmt"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/config"
)

// WeConnect charger implementation using "Ladeziegel" (charging brick) from Volkswagen and controlled
// by the WeConnect cloud.
type WeConnect struct {
	vehicle api.Vehicle
	log     *util.Logger
}

func init() {
	registry.Add("vwweconnect", NewVWWeConnectFromConfig)
}

// NewWeConnectFromConfig creates a WeConnect charger from generic config
func NewVWWeConnectFromConfig(other map[string]interface{}) (api.Charger, error) {
	var cc struct {
		Vehiclename string
	}

	log := util.NewLogger("vwweconnect")

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.Vehiclename == "" {
		return nil, api.ErrMissingCredentials
	}

	veh, err := config.Vehicles().ByName(cc.Vehiclename)
	if err != nil {
		return nil, err
	}
	vehicle := veh.Instance()

	return NewVWConnect(log, vehicle)
}

// NewWeConnect creates a new connection with standbypower for charger
func NewVWConnect(log *util.Logger, vehicle api.Vehicle) (*WeConnect, error) {

	c := &WeConnect{
		vehicle: vehicle,
		log:     log,
	}

	return c, nil
}

func (wb *WeConnect) MaxCurrent(current int64) error {
	return fmt.Errorf("MaxCurrent() not implemented for VWWConnect charge brick")
}

// Status implements the api.Charger interface
func (c *WeConnect) Status() (api.ChargeStatus, error) {
	if stateinterface, ok := c.vehicle.(api.ChargeState); ok {
		return stateinterface.Status()
	} else {
		return api.StatusNone, nil
	}
}

// Enabled implements the api.Charger interface
func (c *WeConnect) Enabled() (bool, error) {
	if stateinterface, ok := c.vehicle.(api.ChargeState); ok {
		state, err := stateinterface.Status()
		if err != nil {
			return false, err
		}
		if state == api.StatusC {
			return true, nil
		}
		return false, nil
	} else {
		return false, fmt.Errorf("vehicle does not support charge state")
	}
}

// Enable implements the api.Charger interface
func (c *WeConnect) Enable(enable bool) error {
	if chargerIF, ok := c.vehicle.(api.ChargeController); ok {
		return chargerIF.ChargeEnable(enable)
	} else {
		return fmt.Errorf("vechicle does not support charge controller")
	}
}

// TotalEnergy implements the api.MeterEnergy interface
func (c *WeConnect) TotalEnergy() (float64, error) {
	return 0.0, fmt.Errorf("vwweconnect charger does not support TotalEnergy")
}
