package mashery

import (
	"context"
	"github.com/aliakseiyanchuk/mashery-v3-go-client/v3client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-mashery/mashschema"
)

func dataSourceMasherySystemDomains() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSystemDomains,
		Schema:      mashschema.DomainsMapper.TerraformSchema(),
	}
}

func readSystemDomains(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	v3cl := m.(v3client.Client)

	if rv, err := v3cl.GetSystemDomains(ctx); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId("system_domains")
		doLogf("received %d system domains", len(rv))
		return mashschema.DomainsMapper.PersistTyped(ctx, rv, d)
	}
}
