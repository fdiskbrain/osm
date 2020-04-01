package ads

import (
	"context"
	"reflect"

	xds "github.com/envoyproxy/go-control-plane/envoy/api/v2"

	"github.com/open-service-mesh/osm/pkg/catalog"
	"github.com/open-service-mesh/osm/pkg/envoy"
	"github.com/open-service-mesh/osm/pkg/logger"
	"github.com/open-service-mesh/osm/pkg/smi"
	"github.com/open-service-mesh/osm/pkg/utils"
)

type empty struct{}

var (
	packageName = utils.GetLastChunkOfSlashed(reflect.TypeOf(empty{}).PkgPath())
	log         = logger.New(packageName)
)

//Server implements the Envoy xDS Aggregate Discovery Services
type Server struct {
	ctx         context.Context
	catalog     catalog.MeshCataloger
	meshSpec    smi.MeshSpec
	xdsHandlers map[envoy.TypeURI]func(context.Context, catalog.MeshCataloger, smi.MeshSpec, *envoy.Proxy, *xds.DiscoveryRequest) (*xds.DiscoveryResponse, error)
}
