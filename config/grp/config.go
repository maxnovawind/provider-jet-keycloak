package grp

import (
	"github.com/crossplane/terrajet/pkg/config"
	"github.com/maxnovawind/provider-jet-keycloak/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("keycloak_group", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "github"
		r.ShortGroup = "grp"
		r.ExternalName = config.IdentifierFromProvider
		r.Version = common.VersionV1alpha2
	})
}
