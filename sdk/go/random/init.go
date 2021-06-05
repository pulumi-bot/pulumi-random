// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"fmt"
	"os"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "random:index/randomId:RandomId":
		r = &RandomId{}
	case "random:index/randomInteger:RandomInteger":
		r = &RandomInteger{}
	case "random:index/randomPassword:RandomPassword":
		r = &RandomPassword{}
	case "random:index/randomPet:RandomPet":
		r = &RandomPet{}
	case "random:index/randomShuffle:RandomShuffle":
		r = &RandomShuffle{}
	case "random:index/randomString:RandomString":
		r = &RandomString{}
	case "random:index/randomUuid:RandomUuid":
		r = &RandomUuid{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:random" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	pulumi.RegisterResourceModule(
		"random",
		"index/randomId",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"random",
		"index/randomInteger",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"random",
		"index/randomPassword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"random",
		"index/randomPet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"random",
		"index/randomShuffle",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"random",
		"index/randomString",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"random",
		"index/randomUuid",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"random",
		&pkg{version},
	)
}
