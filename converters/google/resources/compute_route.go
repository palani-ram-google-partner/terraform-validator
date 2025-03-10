// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const ComputeRouteAssetType string = "compute.googleapis.com/Route"

func resourceConverterComputeRoute() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeRouteAssetType,
		Convert:   GetComputeRouteCaiObject,
	}
}

func GetComputeRouteCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/routes/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeRouteApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeRouteAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "Route",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeRouteApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	destRangeProp, err := expandComputeRouteDestRange(d.Get("dest_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dest_range"); !isEmptyValue(reflect.ValueOf(destRangeProp)) && (ok || !reflect.DeepEqual(v, destRangeProp)) {
		obj["destRange"] = destRangeProp
	}
	descriptionProp, err := expandComputeRouteDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeRouteName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeRouteNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	priorityProp, err := expandComputeRoutePriority(d.Get("priority"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("priority"); ok || !reflect.DeepEqual(v, priorityProp) {
		obj["priority"] = priorityProp
	}
	tagsProp, err := expandComputeRouteTags(d.Get("tags"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tags"); !isEmptyValue(reflect.ValueOf(tagsProp)) && (ok || !reflect.DeepEqual(v, tagsProp)) {
		obj["tags"] = tagsProp
	}
	nextHopGatewayProp, err := expandComputeRouteNextHopGateway(d.Get("next_hop_gateway"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("next_hop_gateway"); !isEmptyValue(reflect.ValueOf(nextHopGatewayProp)) && (ok || !reflect.DeepEqual(v, nextHopGatewayProp)) {
		obj["nextHopGateway"] = nextHopGatewayProp
	}
	nextHopInstanceProp, err := expandComputeRouteNextHopInstance(d.Get("next_hop_instance"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("next_hop_instance"); !isEmptyValue(reflect.ValueOf(nextHopInstanceProp)) && (ok || !reflect.DeepEqual(v, nextHopInstanceProp)) {
		obj["nextHopInstance"] = nextHopInstanceProp
	}
	nextHopIpProp, err := expandComputeRouteNextHopIp(d.Get("next_hop_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("next_hop_ip"); !isEmptyValue(reflect.ValueOf(nextHopIpProp)) && (ok || !reflect.DeepEqual(v, nextHopIpProp)) {
		obj["nextHopIp"] = nextHopIpProp
	}
	nextHopVpnTunnelProp, err := expandComputeRouteNextHopVpnTunnel(d.Get("next_hop_vpn_tunnel"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("next_hop_vpn_tunnel"); !isEmptyValue(reflect.ValueOf(nextHopVpnTunnelProp)) && (ok || !reflect.DeepEqual(v, nextHopVpnTunnelProp)) {
		obj["nextHopVpnTunnel"] = nextHopVpnTunnelProp
	}
	nextHopIlbProp, err := expandComputeRouteNextHopIlb(d.Get("next_hop_ilb"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("next_hop_ilb"); !isEmptyValue(reflect.ValueOf(nextHopIlbProp)) && (ok || !reflect.DeepEqual(v, nextHopIlbProp)) {
		obj["nextHopIlb"] = nextHopIlbProp
	}

	return obj, nil
}

func expandComputeRouteDestRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRoutePriority(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteTags(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v.(*schema.Set).List(), nil
}

func expandComputeRouteNextHopGateway(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == "default-internet-gateway" {
		return replaceVars(d, config, "projects/{{project}}/global/gateways/default-internet-gateway")
	} else {
		return v, nil
	}
}

func expandComputeRouteNextHopInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == "" {
		return v, nil
	}
	val, err := parseZonalFieldValue("instances", v.(string), "project", "next_hop_instance_zone", d, config, true)
	if err != nil {
		return nil, err
	}

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return nil, err
	}

	nextInstance, err := config.NewComputeClient(userAgent).Instances.Get(val.Project, val.Zone, val.Name).Do()
	if err != nil {
		return nil, err
	}
	return nextInstance.SelfLink, nil
}

func expandComputeRouteNextHopIp(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteNextHopVpnTunnel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("vpnTunnels", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for next_hop_vpn_tunnel: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRouteNextHopIlb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
