package nftableslib

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/google/nftables"
)

// ChainsInterface defines third level interface operating with nf chains
type ChainsInterface interface {
	Chains() ChainFuncs
}

// ChainFuncs defines funcations to operate with chains
type ChainFuncs interface {
	Chain(name string) (RulesInterface, error)
	Create(name string, hookNum nftables.ChainHook, priority nftables.ChainPriority, chainType nftables.ChainType)
	Dump() ([]byte, error)
	// TODO figure out what other methods are needed and them
}

type nfChains struct {
	conn  NetNS
	table *nftables.Table
	sync.Mutex
	chains map[string]*nfChain
}

type nfChain struct {
	chainType nftables.ChainType
	chain     *nftables.Chain
	RulesInterface
}

// Chain return Rules Interface for a specified chain
func (nfc *nfChains) Chain(name string) (RulesInterface, error) {
	nfc.Lock()
	defer nfc.Unlock()
	// Check if nf table with the same family type and name  already exists
	if c, ok := nfc.chains[name]; ok {
		return c.RulesInterface, nil

	}
	return nil, fmt.Errorf("chain %s does not exist", name)
}

// Chains return a list of methods available for Chain operations
func (nfc *nfChains) Chains() ChainFuncs {
	return nfc
}

func (nfc *nfChains) Create(name string, hookNum nftables.ChainHook, priority nftables.ChainPriority, chainType nftables.ChainType) {
	nfc.Lock()
	defer nfc.Unlock()
	if _, ok := nfc.chains[name]; ok {
		delete(nfc.chains, name)
	}
	c := nfc.conn.AddChain(&nftables.Chain{
		Name:     name,
		Hooknum:  hookNum,
		Priority: priority,
		Table:    nfc.table,
		Type:     chainType,
	})
	nfc.chains[name] = &nfChain{
		chain:          c,
		chainType:      chainType,
		RulesInterface: newRules(nfc.conn, nfc.table, c),
	}
}

func (nfc *nfChains) Dump() ([]byte, error) {
	nfc.Lock()
	defer nfc.Unlock()
	var data []byte

	for _, c := range nfc.chains {
		if b, err := json.Marshal(&c.chain); err != nil {
			return nil, err
		} else {
			data = append(data, b...)
		}
		if b, err := c.Rules().Dump(); err != nil {
			return nil, err
		} else {
			data = append(data, b...)
		}
	}

	return data, nil
}

func newChains(conn NetNS, t *nftables.Table) ChainsInterface {
	return &nfChains{
		conn:   conn,
		table:  t,
		chains: make(map[string]*nfChain),
	}
}
