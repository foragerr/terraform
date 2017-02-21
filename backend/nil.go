package backend

import (
	"github.com/hashicorp/terraform/state"
	"github.com/hashicorp/terraform/terraform"
)

// Nil is a no-op implementation of Backend.
//
// This is useful to embed within another struct to implement all of the
// backend interface for testing.
type Nil struct{}

func (Nil) Input(
	ui terraform.UIInput,
	c *terraform.ResourceConfig) (*terraform.ResourceConfig, error) {
	return c, nil
}

func (Nil) Validate(*terraform.ResourceConfig) ([]string, []error) {
	return nil, nil
}

func (Nil) Configure(*terraform.ResourceConfig) error {
	return nil
}

func (Nil) State() (state.State, error) {
	// We have to return a non-nil state to adhere to the interface
	return &state.InmemState{}, nil
}

func (Nil) States() ([]string, string, error) {
	// The default state always exists
	return []string{DefaultStateName}, DefaultStateName, nil
}

func (Nil) ChangeState(string) error {
	return nil
}
