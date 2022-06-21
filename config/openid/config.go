package openid

import (
	"github.com/crossplane/terrajet/pkg/config"
	"github.com/maxnovawind/provider-jet-keycloak/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("keycloak_openid_client", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "github"
		r.ShortGroup = "openid"
		r.ExternalName = config.IdentifierFromProvider
		r.Version = common.VersionV1alpha2
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["client_id"].(string); ok {
				conn["keycloak_client_id"] = []byte(a)
			}
			if a, ok := attr["client_secret"].(string); ok {
				conn["keycloak_client_secret"] = []byte(a)
			}
			return conn, nil
		}
	})
}
