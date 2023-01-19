package influxdbv2

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func DataReady() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataGetReady,
		Schema: map[string]*schema.Schema{
			"output": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func DataGetReady(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	influx := meta.(influxdb2.Client)
	ready, err := influx.Ready(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("server is not ready: %v", err))
	}
	if ready != nil {
		log.Printf("Server is ready !")
	}

	output := map[string]string{
		"url": influx.ServerURL(),
	}

	id := ""
	id = influx.ServerURL()
	d.SetId(id)
	err = d.Set("output", output)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
