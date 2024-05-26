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
![Docker Pulls](https://img.shields.io/docker/pulls/evcc/evcc)
[![OSS hosting by cloudsmith](https://img.shields.io/badge/OSS%20hosting%20by-cloudsmith-blue?logo=cloudsmith)](https://cloudsmith.io/~evcc/packages/)
[![Latest Version](https://img.shields.io/github/release/evcc-io/evcc.svg)](https://github.com/evcc-io/evcc/releases)<br/>
[![Built with Depot](https://depot.dev/badges/built-with-depot.svg)](https://depot.dev/?utm_source=evcc)

evcc is an extensible EV Charge Controller and home energy management system.

![Screenshot](assets/github/screenshot.webp)

Our goal is to provide local energy management, without relying on cloud services.
Featured in [PV Magazine](https://www.pv-magazine.de/2022/01/14/mit-open-source-lademanager-schnittstellen-zu-wallbox-und-photovoltaik-anlage-meistern/) and [c’t Magazin](https://www.youtube.com/watch?v=MoBpEXHMNjI).

## Features

- simple and clean user interface
- support for many [EV chargers](https://docs.evcc.io/en/docs/devices/chargers):
  - ABB, ABL, Alfen, Alphatec, Amperfied, Ampure, Audi, AUTEL, Autoaid, Bender, BMW, cFos, Charge Amps, Compleo, CUBOS, Cupra, Dadapower, DaheimLaden, Delta, E.ON Drive, E3/DC, Easee, Ebee, echarge, EcoHarmony, Edgetech, Elecq, eledio, Elli, EM2GO, EN+, enercab, Ensto, EntraTek, ESL, eSystems, Etrel, EVBox, Free2Move, Free2move eSolutions, Fronius, Garo, go-e, Hardy Barth, Heidelberg, Hesotec, Homecharge, Huawei, Innogy, INRO, Juice, Kathrein, KEBA, Kontron Solar, Kostal, KSE, LadeFoxx, LRT, Mennekes, NRGkick, OBO Bettermann, OpenEVSE, openWB, Optec, Orbis, PC Electric, Peblar, Phoenix Contact, Plugchoice, Porsche, Pracht, Pulsares, Pulsatrix, Qcells, Schneider, Schrack, SENEC, Siemens, Skoda, SMA, Smartfox, SolarEdge, Solax, Sonnen, Spelsberg, Stark in Strom, Sungrow, TechniSat, Tesla, Tigo, TinkerForge, Ubitricity, V2C Trydan, Vestel, Victron, Viridian EV, Volkswagen, Volt Time, Wallbe, wallbox, Walther Werke, Webasto, Weidmüller, Zaptec, ZJ Beny. [Read more.](https://docs.evcc.io/en/docs/devices/chargers)
  - **EEBus** support (Elli, PMCC)
  - **OCPP** support
  - **build-your-own:** Phoenix Contact (includes ESL Walli), EVSE DIN
  - **smart switches:** AVM, FRITZ!, Home Assistant, Homematic IP, HomeWizard, myStrom, Shelly, Tasmota, TP-Link. [Read more.](https://docs.evcc.io/en/docs/devices/smartswitches)
  - **heat pumps and electric heaters:** alpha innotec, Bosch, Buderus, Bösch, CTA All-In-One, Daikin, Elco, IDM, Junkers, Kermi, Lambda, my-PV, Nibe, Novelan, Roth, Stiebel Eltron, Tecalor, Vaillant, Viessmann, Wolf, Zewotherm. [Read more.](https://docs.evcc.io/en/docs/devices/heating)
- support for many [energy meters](https://docs.evcc.io/en/docs/devices/meters):
  - **solar inverters and battery systems:** A-Tronix, Acrel, Ads-tec, Alpha ESS, Ampere, Anker, APsystems, AVM, Axitec, BGEtech, Bosch, Bosswerk, Carlo Gavazzi, Deye, E3/DC, Eastron, Enphase, FENECON, FRITZ!, FoxESS, Fronius, Ginlong, go-e, GoodWe, Growatt, Homematic IP, HomeWizard, Hoymiles, Huawei, IAMMETER, IGEN Tech, Kostal, LG, Loxone, M-TEC, Marstek, myStrom, OpenEMS, Powerfox, Qcells, RCT, SAJ, SAX, SENEC, Senergy, Shelly, Siemens, Sigenergy, SMA, Smartfox, SofarSolar, Solaranzeige, SolarEdge, SolarMax, Solarwatt, Solax, Solinteng, Sonnen, St-ems, Steca, Sungrow, Sunsynk, Sunway, Tasmota, Tesla, TP-Link, VARTA, Victron, Wattsonic, Youless, ZCS Azzurro, Zendure. [Read more.](https://docs.evcc.io/en/docs/devices/meters)
  - **general energy meters:** A-Tronix, ABB, Acrel, Alpha ESS, Ampere, AVM, Axitec, Bernecker Engineering, BGEtech, Bosch, Carlo Gavazzi, cFos, Deye, DSMR, DZG, E3/DC, Eastron, Enphase, ESPHome, FENECON, FoxESS, FRITZ!, Fronius, Ginlong, go-e, GoodWe, Growatt, Homematic IP, HomeWizard, Huawei, IAMMETER, inepro, IOmeter, Janitza, KEBA, Kostal, LG, Loxone, M-TEC, mhendriks, my-PV, myStrom, OpenEMS, ORNO, P1Monitor, Powerfox, Qcells, RCT, Saia-Burgess Controls (SBC), SAJ, SAX, Schneider Electric, SENEC, Shelly, Siemens, Sigenergy, SMA, Smartfox, SofarSolar, Solaranzeige, SolarEdge, SolarMax, Solarwatt, Solax, Solinteng, Sonnen, St-ems, Sungrow, Sunsynk, Sunway, Tasmota, Tesla, Tibber, TQ, VARTA, Victron, Volkszähler, Wago, Wattsonic, Weidmüller, Youless, ZCS Azzurro, Zuidwijk. [Read more.](https://docs.evcc.io/en/docs/devices/meters)
  - **integrated systems**: SMA Sunny Home Manager and Energy Meter, KOSTAL Smart Energy Meter (KSEM, EMxx)
  - **sunspec**-compatible inverter or home battery devices
  - **mbmd**-compatible devices, see [volkszaehler/mbmd](https://github.com/volkszaehler/mbmd#supported-devices) for a complete list
- [vehicle](https://docs.evcc.io/en/docs/devices/vehicles) integrations (state of charge, remote charge, battery and preconditioning status):
  - Aiways, Audi, BMW, Citroën, Dacia, DS, Fiat, Ford, Hyundai, Jeep, Kia, Mercedes-Benz, MG, Mini, Nissan, NIU, Opel, Peugeot, Polestar, Renault, Seat, Skoda, Smart, Tesla, Toyota, Volkswagen, Volvo, Zero Motorcycles. [Read more.](https://docs.evcc.io/en/docs/devices/vehicles)
  - **services:** OVMS, Tronity, evNotify, ioBroker.bmw, mg2mqtt, mz2mqtt, TeslaLogger, TeslaMate, Tessi, volvo2mqtt
- [plugins](https://docs.evcc.io/en/docs/devices/plugins) for integrating with any charger, smartswitch, heatpump, electric heater, meter, solar- / battery-inverter or vehicle:
  - Modbus, HTTP, MQTT, JavaScript, WebSocket, Go and shell scripts
- status [notifications](https://docs.evcc.io/en/docs/reference/configuration/messaging) using [Telegram](https://telegram.org), [PushOver](https://pushover.net) and [many more](https://containrrr.dev/shoutrrr/)
- logging using [InfluxDB](https://www.influxdata.com) and [Grafana](https://grafana.com/grafana/)
- [REST](https://docs.evcc.io/en/docs/integrations/rest-api) and [MQTT](https://docs.evcc.io/en/docs/integrations/mqtt-api) APIs for integration with home automation systems
- Add-ons for [Home Assistant](https://docs.evcc.io/en/docs/integrations/home-assistant) and [openHAB](https://www.openhab.org/addons/bindings/evcc) (not maintained by the evcc core team)

## Getting Started

You'll find everything you need in our [documentation](https://docs.evcc.io/en/).

## Contributing

Technical details on how to contribute, how to add translations and how to build evcc from source can be found [here](CONTRIBUTING.md).

[![Weblate Hosted](https://hosted.weblate.org/widgets/evcc/-/evcc/287x66-grey.png)](https://hosted.weblate.org/engage/evcc/)

## Sponsorship

<img src="assets/github/evcc-gopher.png" align="right" width="150" />

evcc believes in open source software. We're committed to provide best in class EV charging experience.
Maintaining evcc consumes time and effort. With the vast amount of different devices to support, we depend on community and vendor support to keep evcc alive.

While evcc is open source, we would also like to encourage vendors to provide open source hardware devices, public documentation and support open source projects like ours that provide additional value to otherwise closed hardware. Where this is not the case, evcc requires "sponsor token" to finance ongoing development and support of evcc.

<<<<<<< HEAD
Learn more about our [sponsorship model](https://docs.evcc.io/en/docs/sponsorship).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

For additional license information regarding fonts, icons, and other assets, please see the [LICENSES](LICENSES/) folder.

**Note:** All sponsor-required components are excluded from the MIT License.
See file license header for details.
If you want to use them in your own project, one evcc sponsorship token is required per evcc instance.
Custom licensing agreements are available - please [contact us](mailto:info@evcc.io) to discuss your specific requirements.
=======
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


### install vite

```bash
sudo npm install -g vite
```



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

### Clone and build evcc fork

```bash
git clone https://github.com/Bodobolero/evcc.git
cd evcc
npm install
make
```

#### Errors during build

running out of resources on small raspberrys:

```bash
github.com/enbility/eebus-go/spine: /usr/local/go/pkg/tool/linux_arm/compile: signal: killed
```

Create swap and retry make
```
sudo fallocate -l 1G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
```

then when the make completes:

```
sudo swapoff /swapfile
sudo rm /swapfile
```


copy evcc.yaml and evcc.db to raspberry

```bash
scp evcc.yaml pi@192.168.178.17:/home/pi/evcc/
scp evcc.db pi@192.168.178.17:/home/pi/.evcc/evcc.db
```

#### check config

./evcc checkconfig

#### convert into a system service
sudo nano /etc/systemd/system/evcc.service

```
[Unit]
Description=EVCC PV suplus EV charging Service
After=network.target

[Service]
ExecStart=/home/pi/evcc/evcc
WorkingDirectory=/home/pi/evcc
Restart=always
User=pi
Group=pi

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl enable evcc.service
sudo systemctl start evcc.service
sudo systemctl status evcc.service

#troubleshooting
sudo journalctl -u evcc.service
```

#### Access to web interface from local network or VPN

http://192.168.178.17:7070/#/


>>>>>>> 4a9971795 (building locally)
