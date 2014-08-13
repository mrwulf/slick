package deployer

import (
	"fmt"
	"strings"
)

type DeployParams struct {
	Environment string
	Branch      string
	Tags        string
	InitiatedBy string
	From        string
}

func (p *DeployParams) ParsedTags() string {
	return strings.Replace(p.Tags, " ", "", -1)
}

func (p *DeployParams) String() string {
	return fmt.Sprintf("env=%s branch=%s tags=%s by=%s from=%s", p.Environment, p.Branch, p.ParsedTags(), p.InitiatedBy, p.From)
}
