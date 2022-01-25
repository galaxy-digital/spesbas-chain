package eventcheck

import (
	"github.com/galaxy-team/spesbas-chain/eventcheck/basiccheck"
	"github.com/galaxy-team/spesbas-chain/eventcheck/epochcheck"
	"github.com/galaxy-team/spesbas-chain/eventcheck/gaspowercheck"
	"github.com/galaxy-team/spesbas-chain/eventcheck/heavycheck"
	"github.com/galaxy-team/spesbas-chain/eventcheck/parentscheck"
	"github.com/galaxy-team/spesbas-chain/inter"
)

// Checkers is collection of all the checkers
type Checkers struct {
	Basiccheck    *basiccheck.Checker
	Epochcheck    *epochcheck.Checker
	Parentscheck  *parentscheck.Checker
	Gaspowercheck *gaspowercheck.Checker
	Heavycheck    *heavycheck.Checker
}

// Validate runs all the checks except Poset-related
func (v *Checkers) Validate(e inter.EventPayloadI, parents inter.EventIs) error {
	if err := v.Basiccheck.Validate(e); err != nil {
		return err
	}
	if err := v.Epochcheck.Validate(e); err != nil {
		return err
	}
	if err := v.Parentscheck.Validate(e, parents); err != nil {
		return err
	}
	var selfParent inter.EventI
	if e.SelfParent() != nil {
		selfParent = parents[0]
	}
	if err := v.Gaspowercheck.Validate(e, selfParent); err != nil {
		return err
	}
	if err := v.Heavycheck.Validate(e); err != nil {
		return err
	}
	return nil
}
