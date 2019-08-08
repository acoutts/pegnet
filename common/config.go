// Copyright (c) of parts are held by the various contributors (see the CLA)
// Licensed under the MIT License. See LICENSE file in the project root for full license information.
package common

import (
	"fmt"

	"github.com/zpatrick/go-config"

	"github.com/go-ini/ini"
)

// Config names
const (
	ConfigCoordinatorListen            = "Miner.MiningCoordinatorPort"
	ConfigCoordinatorLocation          = "Miner.MiningCoordinatorHost"
	ConfigCoordinatorSecret            = "Miner.CoordinatorSecret"
	ConfigCoordinatorUseAuthentication = "Miner.UseCoordinatorAuthentication"

	ConfigCoinbaseAddress = "Miner.CoinbaseAddress"
	ConfigPegnetNetwork   = "Miner.Network"

	ConfigCoinMarketCapKey = "Oracle.CoinMarketCapKey"
)

func NewUnitTestConfig() *config.Config {
	return config.NewConfig([]config.Provider{NewUnitTestConfigProvider()})
}

// UnitTestConfigProvider is only used in unit tests.
//	This way we don't have to deal with pathing to find the
//	`defaultconfig.ini`.
type UnitTestConfigProvider struct {
	Data string
}

func NewUnitTestConfigProvider() *UnitTestConfigProvider {
	d := new(UnitTestConfigProvider)
	d.Data = `
[Debug]
# Randomize adds a random factor +/- the give percent.  3.1 for 3.1%
  Randomize=0.1
# Turns on logging so the user can see the OPRs and mining balances as they update
  Logging=true
# Puts the logs in a file.  If not specified, logs are written to stdout
  LogFile=

[Miner]
  NetworkType=LOCAL
  NumberOfMiners=15
# The number of records to submit per block. The top N records are chosen, where N is the config value
  RecordsPerBlock=10
  Protocol=PegNet 
  Network=TestNet

  # For LOCAL network testing, EC private key is
  # Es2XT3jSxi1xqrDvS5JERM3W3jh1awRHuyoahn3hbQLyfEi1jvbq
  ECAddress=EC3TsJHUs8bzbbVnratBafub6toRYdgzgbR7kWwCW4tqbmyySRmg

  # For LOCAL network testing, FCT private key is
  # Fs3E9gV6DXsYzf7Fqx1fVBQPQXV695eP3k5XbmHEZVRLkMdD9qCK
  FCTAddress=FA2jK2HcLnRdS94dEcU27rF3meoJfpUcZPSinpb7AwQvPRY6RL1Q

  CoinbasePNTAddress=tPNT_mEU1i4M5rn7bnrxNKdVVf4HXLG15Q798oaVAMrXq7zdbhQ9pv
  IdentityChain=prototype
[Oracle]
  APILayerKey=CHANGEME
  OpenExchangeRatesKey=CHANGEME
  CoinMarketCapKey=CHANGEME


[OracleDataSources]
  APILayer=1
  ExchangeRates=3
  OpenExchangeRates=2

  # Crypto
  CoinMarketCap=3
  CoinCap=4

  # Commodities
  Kitco=10

`
	return d
}

func (this *UnitTestConfigProvider) Load() (map[string]string, error) {
	settings := map[string]string{}

	file, err := ini.Load([]byte(this.Data))
	if err != nil {
		return nil, err
	}

	for _, section := range file.Sections() {
		for _, key := range section.Keys() {
			token := fmt.Sprintf("%s.%s", section.Name(), key.Name())
			settings[token] = key.String()
		}
	}

	return settings, nil
}
