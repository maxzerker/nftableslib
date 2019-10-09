module github.com/sbezverk/nftableslib

go 1.12

require (
	github.com/google/nftables v0.0.0-20190906062827-5d14089d2edc
	github.com/google/uuid v1.1.1
	github.com/jsimonetti/rtnetlink v0.0.0-20190830100107-3784a6c7c552 // indirect
	github.com/mdlayher/netlink v0.0.0-20190828143259-340058475d09 // indirect
	github.com/sbezverk/nftableslib/e2e/setenv v0.0.0-20191009154549-4fe065fe4e96 // indirect
	github.com/vishvananda/netns v0.0.0-20190625233234-7109fa855b0f
	golang.org/x/crypto v0.0.0-20190829043050-9756ffdc2472 // indirect
	golang.org/x/sys v0.0.0-20191008105621-543471e840be
	golang.org/x/text v0.3.2 // indirect
	golang.org/x/tools v0.0.0-20190903025054-afe7f8212f0d // indirect
)

replace github.com/sbezverk/nftableslib/e2e => ./e2e
