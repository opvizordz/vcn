/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 * Built on top of CLI (MIT license)
 * https://github.com/urfave/cli#overview
 */

package main

import (
	"math/big"
	"os"
	"time"

	"github.com/fatih/color"
)

func DashboardURL() string {
	switch StageEnvironment() {
	case StageProduction:
		return "https://dashboard.codenotary.io"
	case StageStaging:
		return "https://dashboard.staging.codenotary.io"
	case StageTest:
		return "https://dashboard.test.codenotary.io"
	default:
		return "https://dashboard.codenotary.io"
	}
}

func MainNetEndpoint() string {
	switch StageEnvironment() {
	case StageProduction:
		return "https://main.codenotary.io"
	case StageStaging:
		return "https://main.staging.codenotary.io"
	case StageTest:
		return "https://main.test.codenotary.io"
	default:
		return "https://main.codenotary.io"
	}
}

func FoundationEndpoint() string {
	switch StageEnvironment() {
	case StageProduction:
		return "https://api.codenotary.io/foundation"
	case StageStaging:
		return "https://api.staging.codenotary.io/foundation"
	case StageTest:
		return "https://api.test.codenotary.io/foundation"
	default:
		return "https://api.codenotary.io/foundation"
	}
}

func AssetsRelayContractAddress() string {
	switch StageEnvironment() {
	case StageProduction:
		return "0x495021fe1a48a5b0c85ef1abd68c42cdfc7cda08"
	case StageStaging:
		return "0xf1d4b9fe8290bb5718db5d46c313e7b266570c21"
	case StageTest:
		return "0x0"
	default:
		return "0x495021fe1a48a5b0c85ef1abd68c42cdfc7cda08"
	}
}

func StyleAffordance() (color.Attribute, color.Attribute) {
	return color.FgHiBlue, color.Bold
}

func StyleError() (color.Attribute, color.Attribute) {
	return color.FgRed, color.Bold
}

func StyleWarning() (color.Attribute, color.Attribute) {
	return color.FgYellow, color.Bold
}

func StyleSuccess() (color.Attribute, color.Attribute) {
	return color.FgGreen, color.Bold
}

func ErrorWikiURL() string {
	return "https://github.com/vchain-us/vcn/wiki/Errors#"
}

func TrackingEvent() string {
	return FoundationEndpoint() + "/v1/tracking-event"
}

func TokenCheckEndpoint() string {
	return PublisherEndpoint() + "/auth/check"
}

func PublisherEndpoint() string {
	return FoundationEndpoint() + "/v1/publisher"
}

func WalletEndpoint() string {
	return FoundationEndpoint() + "/v1/wallet"
}

func ArtifactEndpoint() string {
	return FoundationEndpoint() + "/v1/artifact"
}

func ArtifactEndpointForWallet(walletAddress string) string {
	return FoundationEndpoint() + "/v1/artifact?wallet-address=" + walletAddress
}

func TxVerificationRounds() uint64 {
	return 10
}

func PollInterval() time.Duration {
	return 1 * time.Second
}

func VcnDirectory() string {
	return os.Getenv("HOME") + "/.vcn"
}

func WalletDirectory() string {
	return VcnDirectory() + "/wallets"
}

func TokenFile() string {
	return VcnDirectory() + "/t"
}

func GasPrice() *big.Int {
	return big.NewInt(0)
}

func GasLimit() uint64 {
	return 20000000
}
