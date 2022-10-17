package mashschema

import (
	"github.com/aliakseiyanchuk/mashery-v3-go-client/masherytypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const MashEmailTemplateSetType = "type"

const MashEmailTemplateSetId = "set_id"

var EmailTemplateSetMapper *emailTemplateSetMapperImpl

type emailTemplateSetMapperImpl struct {
	DataSourceMapperImpl
}

func (etsm *emailTemplateSetMapperImpl) PersistTyped(set masherytypes.EmailTemplateSet, d *schema.ResourceData) diag.Diagnostics {
	data := map[string]interface{}{
		MashObjCreated:           set.Created.ToString(),
		MashObjUpdated:           set.Updated.ToString(),
		MashObjName:              set.Name,
		MashEmailTemplateSetType: set.Type,
	}

	return SetResourceFields(data, d)
}

func initEmailTemplateSetSchemaBoilerplate() {
	addComputedString(&EmailTemplateSetMapper.schema, MashObjCreated, "Date/time this email template set was created")
	addComputedString(&EmailTemplateSetMapper.schema, MashObjUpdated, "Date/time this email template set was updated")
	addRequiredString(&EmailTemplateSetMapper.schema, MashObjName, "Email data set name")
	addComputedOptionalString(&EmailTemplateSetMapper.schema, MashEmailTemplateSetType, "Email template set type")
}

func init() {
	EmailTemplateSetMapper = &emailTemplateSetMapperImpl{
		DataSourceMapperImpl{
			schema: map[string]*schema.Schema{
				MashDataSourceSearch: {
					Type:        schema.TypeMap,
					Required:    true,
					Description: "Search conditions for this email set, typically name = value",
					Elem:        stringElem(),
				},
				MashDataSourceRequired: {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "If true (default), then email template set must exist. If an element doesn't exist, the error is generated",
				},
			},

			persistOne: func(rv interface{}, d *schema.ResourceData) diag.Diagnostics {
				return EmailTemplateSetMapper.PersistTyped(rv.(masherytypes.EmailTemplateSet), d)
			},
		},
	}

	initEmailTemplateSetSchemaBoilerplate()
	addComputedString(&EmailTemplateSetMapper.schema, MashEmailTemplateSetId, "Email set Id")
}