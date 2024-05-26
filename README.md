# evcc

This is a fork of https://github.com/evcc-io/evcc

## Added capability in this fork:

see discussion in https://github.com/evcc-io/evcc/issues/12046

### Using Solaredge Rest API as a meter

My meter is connected to the converter over a (twisted pair?) cable S0.
The SEK5 thus has the consumption values. It also transmits it to the SolarEdge cloud. However so far I did not yet find a way how it re-exports those values over Modbus. So I can not access the consumption locally. My meter is an Eltako DSZ12E
Verbrauchszähler S0 mit 800 Imp/Kwh
https://www.eltako.com/fileadmin/downloads/de/_bedienung/DSZ12E-3x80A_28380611-3_dt.pdf
Eltako DSZ12E 3x80A

I have no local API/access to it, however the Solaredge REST API has the information it collects

I installed the go tool from Ulrich Schreiner https://gitlab.com/ulrichSchreiner/solaredge referenced under https://github.com/evcc-io/evcc/discussions/2941

Thus I did not have to modify evcc but could use the existing custom meter / plugin capability.

However I have a private fork on https://github.com/Bodobolero/solaredge-go-library where I made the following changes
- polling interval during the night is 15 minutes
- polling interval during the day (when solarpower is active) is 3 minutes
- this is to ensure I stay within the Solaredge API request quota of 300/day

### Using Weconnect API with Charge brick (Ladeziegel) as a charger

My use case is:

I have a 5.04 kWp (kilo-watt peak) PV.
Because I want to use PV energy to charge my car I do not use a wallbox, instead I have a Green-up socket which supports my 1 phase 10-16 A 240 V charging device that comes with my car.
Since this is our second car we will use it for short distances, usually 2-20 km and the 2.3 kw charging using the green-up socket will be enough to support our usage of the car.

I want an automated solution that decides when to start charging and when to stop charging.
The decision will be based on the SOC of the car (I will charge at any time until I reach a SOC of x % (e.g. 30 % to support my local commute).
I want to charge between 30 % and 80 % when there is enough PV production to support the 2.4 kw charging (with 10 A * 240 V), which means when my production is higher than <current consumption> + 2.4 kw.

For that purpose the S0 800 Imp/kWh meter is totally sufficient.

My car has a mileage of 10.000 km per year, at 20 kwh/100 km that is 2000 kwh / year.
My reimbursement for PV is 13 Eurocent. My energy tarif is 27 cent/kwh.

The total possible savings per year by optimizating the charge time is 27 cent - 13 cent = 14 cent * 2000 kwh.
That means I can save a total of 280 Euro/year.

Now comes the investment argument:
I can buy a cheap e-go wallbox that can be remotely controlled (around 300 Euro) and I can also replace the smart meter.
To modify the house installation I need a certified electrician that will cost me a few hundred Euro.

Which means to get the current implementation of evcc working I need an investment of > 1000 Euro to save 280 Euro per year.

With a few modifications to the software I can achieve the same with just investment in code.

support REST API for SolarEdge in the metering area would save me to replace the meter

support VW We.connect to control the charging instead of requiring a remote-controllable charger would save me to buy a wallbox

Since evcc project owners didn't want my contributions I created a fork and implemented the
vehicle WeConnect API as a charger to start/stop charging and get the current charging state

## Links
 🚘☀️

[![Build](https://github.com/evcc-io/evcc/actions/workflows/nightly.yml/badge.svg)](https://github.com/evcc-io/evcc/actions/workflows/nightly.yml)
[![Translation](https://hosted.weblate.org/widgets/evcc/-/evcc/svg-badge.svg)](https://hosted.weblate.org/engage/evcc/)
[![Open in Visual Studio Code](https://img.shields.io/static/v1?logo=visualstudiocode&label=&message=Open%20in%20VS%20Code&labelColor=2c2c32&color=007acc&logoColor=007acc)](https://open.vscode.dev/evcc-io/evcc)
[![OSS hosting by cloudsmith](https://img.shields.io/badge/OSS%20hosting%20by-cloudsmith-blue?logo=cloudsmith)](https://cloudsmith.io/~evcc/packages/)
[![Latest Version](https://img.shields.io/github/release/evcc-io/evcc.svg)](https://github.com/evcc-io/evcc/releases)

evcc is an extensible EV Charge Controller and home energy management system. Featured in [PV magazine](https://www.pv-magazine.de/2021/01/15/selbst-ist-der-groeoenlandhof-wallbox-ladesteuerung-selbst-gebaut/).

![Screenshot](docs/screenshot.png)

## Features

- simple and clean user interface
- wide range of supported [chargers](https://docs.evcc.io/docs/devices/chargers):
  - ABL eMH1, Alfen (Eve), Bender (CC612/613), cFos (PowerBrain), Daheimladen, Ebee (Wallbox), Ensto (Chago Wallbox), [EVSEWifi/ smartWB](https://www.evse-wifi.de), Garo (GLB, GLB+, LS4), go-eCharger, HardyBarth (eCB1, cPH1, cPH2), Heidelberg (Energy Control), Innogy (eBox), Juice (Charger Me), KEBA/BMW, Mennekes (Amedio, Amtron Premium/Xtra, Amtron ChargeConrol), older NRGkicks (before 2022/2023), [openWB (includes Pro)](https://openwb.de/), Optec (Mobility One), PC Electric (includes Garo), Siemens, TechniSat (Technivolt), [Tinkerforge Warp Charger](https://www.warp-charger.com), Ubitricity (Heinz), Vestel, Wallbe, Webasto (Live), Mobile Charger Connect and many more
  - experimental EEBus support (Elli, PMCC)
  - experimental OCPP support
  - Build-your-own: Phoenix Contact (includes ESL Walli), [EVSE DIN](http://evracing.cz/simple-evse-wallbox)
  - Smart-Home outlets: FritzDECT, Shelly, Tasmota, TP-Link
- wide range of supported [meters](https://docs.evcc.io/docs/devices/meters) for grid, pv, battery and charger:
  - ModBus: Eastron SDM, MPM3PM, ORNO WE, SBC ALE3 and many more, see <https://github.com/volkszaehler/mbmd#supported-devices> for a complete list
  - Integrated systems: SMA Sunny Home Manager and Energy Meter, KOSTAL Smart Energy Meter (KSEM, EMxx)
  - Sunspec-compatible inverter or home battery devices: Fronius, SMA, SolarEdge, KOSTAL, STECA, E3DC, ...
  - and various others: Discovergy, Tesla PowerWall, LG ESS HOME, OpenEMS (FENECON)
- [vehicle](https://docs.evcc.io/docs/devices/vehicles) integration (state of charge, remote charge, battery and preconditioning status):
  - Audi, BMW, Citroën, Dacia, Fiat, Ford, Hyundai, Jaguar, Kia, Landrover, ~~Mercedes~~, Mini, Nissan, Opel, Peugeot, Porsche, Renault, Seat, Smart, Skoda, Tesla, Volkswagen, Volvo, ...
  - Services: OVMS, Tronity
  - Scooters: Niu, ~~Silence~~
- [plugins](https://docs.evcc.io/docs/reference/plugins) for integrating with any charger/ meter/ vehicle:
  - Modbus, HTTP, MQTT, Javascript, WebSockets and shell scripts
- status [notifications](https://docs.evcc.io/docs/reference/configuration/messaging) using [Telegram](https://telegram.org), [PushOver](https://pushover.net) and [many more](https://containrrr.dev/shoutrrr/)
- logging using [InfluxDB](https://www.influxdata.com) and [Grafana](https://grafana.com/grafana/)
- granular charge power control down to mA steps with supported chargers (labeled by e.g. smartWB as [OLC](https://board.evse-wifi.de/viewtopic.php?f=16&t=187))
- REST and MQTT [APIs](https://docs.evcc.io/docs/reference/api) for integration with home automation systems
- Add-ons for [Home Assistant](https://github.com/evcc-io/evcc-hassio-addon) and [OpenHAB](https://www.openhab.org/addons/bindings/evcc) (not maintained by the evcc core team)

## Getting Started

You'll find everything you need in our [documentation](https://docs.evcc.io/).

## Contributing

Technical details on how to contribute, how to add translations and how to build evcc from source can be found [here](CONTRIBUTING.md).

[![Weblate Hosted](https://hosted.weblate.org/widgets/evcc/-/evcc/287x66-grey.png)](https://hosted.weblate.org/engage/evcc/)

## Sponsorship

<img src="docs/logo.png" align="right" width="150" />

evcc believes in open source software. We're committed to provide best in class EV charging experience.
Maintaining evcc consumes time and effort. With the vast amount of different devices to support, we depend on community and vendor support to keep evcc alive.

While evcc is open source, we would also like to encourage vendors to provide open source hardware devices, public documentation and support open source projects like ours that provide additional value to otherwise closed hardware. Where this is not the case, evcc requires "sponsor token" to finance ongoing development and support of evcc.

The personal sponsor token requires a [Github Sponsorship](https://github.com/sponsors/evcc-io) and can be requested at [sponsor.evcc.io](https://sponsor.evcc.io/).


## rebasing this repo from upstream repo

```bash
git remote add upstream git@github.com:evcc-io/evcc.git
git fetch upstream
git checkout master
git rebase upstream/master
# Resolve any conflicts, then
git rebase --continue
git push origin master --force
```

## Deployment on Raspberry

### Install go

```bash
wget https://go.dev/dl/go1.22.3.linux-armv6l.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.3.linux-armv6l.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```

### Install npm

```bash
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo bash -
sudo apt-get install -y nodejs
npm --version
node --version
```

npm: 10.5.2
node: v20.13.1

### Clone my fork

#### enable github access from my pi

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
eval "$(ssh-agent -s)"ssh-add ~/.ssh/id_ed25519
cat ~/.ssh/id_ed25519.pub
```
Add the rsa key to your github account which has the private repo for the Solaredge go library


#### build solaredge api 

git clone git@github.com:Bodobolero/solaredge-go-library.git
cd solaredge-go-library/
make

#### convert into a system service
sudo nano /etc/systemd/system/solaredge.service

```
[Unit]
Description=SolarEdge API Service
After=network.target

[Service]
Environment="SOLAREDGE_APIKEY=<addyourkey>"
Environment="SOLAREDGE_SITEID=<addyoursite>"
ExecStart=/home/pi/solaredge-go-library/bin/solaredge serve
WorkingDirectory=/home/pi/solaredge-go-library/bin/
Restart=always
User=pi
Group=pi

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl enable solaredge.service
sudo systemctl start solaredge.service
sudo systemctl status solaredge.service

#troubleshooting
sudo journalctl -u solaredge.service
curl http://localhost:7777/flow
```





