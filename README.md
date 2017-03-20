## Cryptocurrency trading bot written in Golang

[![Build Status](https://travis-ci.org/thrasher-/gocryptotrader.svg?branch=master)](https://travis-ci.org/thrasher-/gocryptotrader)
[![Test Coverage](https://codecov.io/github/thrasher-/gocryptotrader/coverage.svg?branch=master)](https://codecov.io/github/thrasher-/gocryptotrader?branch=master)

A cryptocurrency trading bot supporting multiple exchanges written in Golang. 

**Please note that this bot is under development and is not ready for production!**

Join our Slack! https://gocryptotrader.slack.com

## Exchange Support Table

| Exchange | REST API | Streaming API | FIX API |
|----------|------|-----------|-----|
| Alphapoint | Yes  | Yes        | NA  |
| ANXPRO | Yes  | No        | NA  |
| Bitfinex | Yes  | Yes        | NA  |
| Bitstamp | Yes  | Yes       | NA  |
| BTCC | Yes  | Yes     | No  |
| BTCE     | Yes  | NA        | NA  |
| BTCMarkets | Yes | NA       | NA  |
| GDAX(Coinbase) | Yes | Yes | No|
| Gemini | Yes | NA | NA |
| Huobi | Yes | Yes |No |
| ItBit | Yes | NA | NA |
| Kraken | Yes | NA | NA |
| LakeBTC | Yes | NA | NA | 
| Liqui | Yes | No | NA |
| LocalBitcoins | Yes | NA | NA |
| OKCoin (both) | Yes | Yes | No |
| Poloniex | Yes | Yes | NA |

** NA means not applicable as the Exchange does not support the feature.

## Current Features
+ Support for all Exchange fiat and digital currencies, with the ability to individually toggle them on/off.
+ REST API support for all exchanges.
+ Websocket support for applicable exchanges.
+ Ability to turn off/on certain exchanges.
+ Ability to adjust manual polling timer for exchanges.
+ SMS notification support via SMS Gateway.
+ Basic event trigger system.

## Planned Features
+ WebGUI.
+ FIX support.
+ Expanding event trigger system.
+ TALib.
+ Trade history summary generation for tax purposes.

Please feel free to submit any pull requests or suggest any desired features to be added.

## Compiling instructions
Download Go from https://golang.org/dl/  
Using a terminal, type go get github.com/thrasher-/gocryptotrader  
Change directory to the package directory, then type go install.  
Copy config_example.dat to config.dat.  
Make any neccessary changes to the config file.  
Run the application!  

## Binaries
Binaries will be published once the codebase reaches a stable condition.

