package mashschema_test

import (
	"context"
	"github.com/aliakseiyanchuk/mashery-v3-go-client/masherytypes"
	"github.com/stretchr/testify/assert"
	"terraform-provider-mashery/mashschema"
	"testing"
	"time"
)

func TestV3EmailTemplateSetIdToResourceData_DataSource(t *testing.T) {
	d := mashschema.EmailTemplateSetMapper.NewResourceData()

	now := masherytypes.MasheryJSONTime(time.Now())

	orig := masherytypes.MasheryEmailTemplateSet{
		AddressableV3Object: masherytypes.AddressableV3Object{
			Id:      "setId",
			Name:    "name",
			Created: &now,
			Updated: &now,
		},
		Type:           "Type",
		EmailTemplates: nil,
	}

	dr := mashschema.EmailTemplateSetMapper.PersistTyped(context.TODO(), &orig, d)
	LogErrorDiagnostics(t, "email template set", &dr)

	assert.Equal(t, orig.Name, d.Get(mashschema.MashObjName).(string))
	assert.Equal(t, orig.Type, d.Get(mashschema.MashEmailTemplateSetType).(string))
}
