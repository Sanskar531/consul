// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package failover

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	svctest "github.com/hashicorp/consul/agent/grpc-external/services/resource/testing"
	"github.com/hashicorp/consul/internal/catalog/internal/mappers/failovermapper"
	"github.com/hashicorp/consul/internal/catalog/internal/types"
	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/resource"
	rtest "github.com/hashicorp/consul/internal/resource/resourcetest"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v2beta1"
	"github.com/hashicorp/consul/proto-public/pbresource"
	"github.com/hashicorp/consul/sdk/testutil"
)

type controllerSuite struct {
	suite.Suite

	ctx    context.Context
	client *rtest.Client
	rt     controller.Runtime

	failoverMapper FailoverMapper

	ctl failoverPolicyReconciler

	tenancies []*pbresource.Tenancy
}

func (suite *controllerSuite) SetupTest() {
	suite.tenancies = rtest.TestTenancies()
	client := svctest.NewResourceServiceBuilder().
		WithRegisterFns(types.Register, types.RegisterDNSPolicy).
		WithTenancies(suite.tenancies...).
		Run(suite.T())

	suite.rt = controller.Runtime{
		Client: client,
		Logger: testutil.Logger(suite.T()),
	}
	suite.client = rtest.NewClient(client)
	suite.failoverMapper = failovermapper.New()
	suite.ctx = testutil.TestContext(suite.T())
}

func (suite *controllerSuite) TestController() {
	// This test's purpose is to exercise the controller in a halfway realistic
	// way, verifying the event triggers work in the live code.

	// Run the controller manager
	mgr := controller.NewManager(suite.client, suite.rt.Logger)
	mgr.Register(FailoverPolicyController(suite.failoverMapper))
	mgr.SetRaftLeader(true)
	go mgr.Run(suite.ctx)

	suite.runTestCaseWithTenancies(func(tenancy *pbresource.Tenancy) {
		// Create an advance pointer to some services.
		apiServiceRef := resource.Reference(rtest.Resource(pbcatalog.ServiceType, "api").WithTenancy(tenancy).ID(), "")
		otherServiceRef := resource.Reference(rtest.Resource(pbcatalog.ServiceType, "other").WithTenancy(tenancy).ID(), "")

		// create a failover without any services
		failoverData := &pbcatalog.FailoverPolicy{
			Config: &pbcatalog.FailoverConfig{
				Destinations: []*pbcatalog.FailoverDestination{{
					Ref: apiServiceRef,
				}},
			},
		}
		failover := rtest.Resource(pbcatalog.FailoverPolicyType, "api").
			WithData(suite.T(), failoverData).
			WithTenancy(tenancy).
			Write(suite.T(), suite.client)

		suite.T().Cleanup(suite.deleteResourceFunc(failover.Id))

		suite.client.WaitForStatusCondition(suite.T(), failover.Id, StatusKey, ConditionMissingService)

		// Provide the service.
		apiServiceData := &pbcatalog.Service{
			Workloads: &pbcatalog.WorkloadSelector{Prefixes: []string{"api-"}},
			Ports: []*pbcatalog.ServicePort{{
				TargetPort: "http",
				Protocol:   pbcatalog.Protocol_PROTOCOL_HTTP,
			}},
		}
		svc := rtest.Resource(pbcatalog.ServiceType, "api").
			WithData(suite.T(), apiServiceData).
			WithTenancy(tenancy).
			Write(suite.T(), suite.client)

		suite.T().Cleanup(suite.deleteResourceFunc(svc.Id))

		suite.client.WaitForStatusCondition(suite.T(), failover.Id, StatusKey, ConditionOK)

		// Update the failover to reference an unknown port
		failoverData = &pbcatalog.FailoverPolicy{
			PortConfigs: map[string]*pbcatalog.FailoverConfig{
				"http": {
					Destinations: []*pbcatalog.FailoverDestination{{
						Ref:  apiServiceRef,
						Port: "http",
					}},
				},
				"admin": {
					Destinations: []*pbcatalog.FailoverDestination{{
						Ref:  apiServiceRef,
						Port: "admin",
					}},
				},
			},
		}
		svc = rtest.Resource(pbcatalog.FailoverPolicyType, "api").
			WithData(suite.T(), failoverData).
			WithTenancy(tenancy).
			Write(suite.T(), suite.client)

		suite.T().Cleanup(suite.deleteResourceFunc(svc.Id))

		suite.client.WaitForStatusCondition(suite.T(), failover.Id, StatusKey, ConditionUnknownPort("admin"))

		// update the service to fix the stray reference, but point to a mesh port
		apiServiceData = &pbcatalog.Service{
			Workloads: &pbcatalog.WorkloadSelector{Prefixes: []string{"api-"}},
			Ports: []*pbcatalog.ServicePort{
				{
					TargetPort: "http",
					Protocol:   pbcatalog.Protocol_PROTOCOL_HTTP,
				},
				{
					TargetPort: "admin",
					Protocol:   pbcatalog.Protocol_PROTOCOL_MESH,
				},
			},
		}
		svc = rtest.Resource(pbcatalog.ServiceType, "api").
			WithData(suite.T(), apiServiceData).
			WithTenancy(tenancy).
			Write(suite.T(), suite.client)

		suite.T().Cleanup(suite.deleteResourceFunc(svc.Id))

		suite.client.WaitForStatusCondition(suite.T(), failover.Id, StatusKey, ConditionUsingMeshDestinationPort(apiServiceRef, "admin"))

		// update the service to fix the stray reference to not be a mesh port
		apiServiceData = &pbcatalog.Service{
			Workloads: &pbcatalog.WorkloadSelector{Prefixes: []string{"api-"}},
			Ports: []*pbcatalog.ServicePort{
				{
					TargetPort: "http",
					Protocol:   pbcatalog.Protocol_PROTOCOL_HTTP,
				},
				{
					TargetPort: "admin",
					Protocol:   pbcatalog.Protocol_PROTOCOL_HTTP,
				},
			},
		}
		svc = rtest.Resource(pbcatalog.ServiceType, "api").
			WithData(suite.T(), apiServiceData).
			WithTenancy(tenancy).
			Write(suite.T(), suite.client)

		suite.T().Cleanup(suite.deleteResourceFunc(svc.Id))

		suite.client.WaitForStatusCondition(suite.T(), failover.Id, StatusKey, ConditionOK)

		// change failover leg to point to missing service
		failoverData = &pbcatalog.FailoverPolicy{
			PortConfigs: map[string]*pbcatalog.FailoverConfig{
				"http": {
					Destinations: []*pbcatalog.FailoverDestination{{
						Ref:  apiServiceRef,
						Port: "http",
					}},
				},
				"admin": {
					Destinations: []*pbcatalog.FailoverDestination{{
						Ref:  otherServiceRef,
						Port: "admin",
					}},
				},
			},
		}
		svc = rtest.Resource(pbcatalog.FailoverPolicyType, "api").
			WithData(suite.T(), failoverData).
			WithTenancy(tenancy).
			Write(suite.T(), suite.client)

		suite.T().Cleanup(suite.deleteResourceFunc(svc.Id))

		suite.client.WaitForStatusCondition(suite.T(), failover.Id, StatusKey, ConditionMissingDestinationService(otherServiceRef))

		// Create the missing service, but forget the port.
		otherServiceData := &pbcatalog.Service{
			Workloads: &pbcatalog.WorkloadSelector{Prefixes: []string{"other-"}},
			Ports: []*pbcatalog.ServicePort{{
				TargetPort: "http",
				Protocol:   pbcatalog.Protocol_PROTOCOL_HTTP,
			}},
		}
		svc = rtest.Resource(pbcatalog.ServiceType, "other").
			WithData(suite.T(), otherServiceData).
			WithTenancy(tenancy).
			Write(suite.T(), suite.client)

		suite.T().Cleanup(suite.deleteResourceFunc(svc.Id))

		suite.client.WaitForStatusCondition(suite.T(), failover.Id, StatusKey, ConditionUnknownDestinationPort(otherServiceRef, "admin"))

		// fix the destination leg's port
		otherServiceData = &pbcatalog.Service{
			Workloads: &pbcatalog.WorkloadSelector{Prefixes: []string{"other-"}},
			Ports: []*pbcatalog.ServicePort{
				{
					TargetPort: "http",
					Protocol:   pbcatalog.Protocol_PROTOCOL_HTTP,
				},
				{
					TargetPort: "admin",
					Protocol:   pbcatalog.Protocol_PROTOCOL_HTTP,
				},
			},
		}
		svc = rtest.Resource(pbcatalog.ServiceType, "other").
			WithData(suite.T(), otherServiceData).
			WithTenancy(tenancy).
			Write(suite.T(), suite.client)

		suite.T().Cleanup(suite.deleteResourceFunc(svc.Id))

		suite.client.WaitForStatusCondition(suite.T(), failover.Id, StatusKey, ConditionOK)

		// Update the two services to use differnet port names so the easy path doesn't work
		apiServiceData = &pbcatalog.Service{
			Workloads: &pbcatalog.WorkloadSelector{Prefixes: []string{"api-"}},
			Ports: []*pbcatalog.ServicePort{
				{
					TargetPort: "foo",
					Protocol:   pbcatalog.Protocol_PROTOCOL_HTTP,
				},
				{
					TargetPort: "bar",
					Protocol:   pbcatalog.Protocol_PROTOCOL_HTTP,
				},
			},
		}
		svc = rtest.Resource(pbcatalog.ServiceType, "api").
			WithData(suite.T(), apiServiceData).
			WithTenancy(tenancy).
			Write(suite.T(), suite.client)

		suite.T().Cleanup(suite.deleteResourceFunc(svc.Id))

		otherServiceData = &pbcatalog.Service{
			Workloads: &pbcatalog.WorkloadSelector{Prefixes: []string{"other-"}},
			Ports: []*pbcatalog.ServicePort{
				{
					TargetPort: "foo",
					Protocol:   pbcatalog.Protocol_PROTOCOL_HTTP,
				},
				{
					TargetPort: "baz",
					Protocol:   pbcatalog.Protocol_PROTOCOL_HTTP,
				},
			},
		}
		svc = rtest.Resource(pbcatalog.ServiceType, "other").
			WithData(suite.T(), otherServiceData).
			WithTenancy(tenancy).
			Write(suite.T(), suite.client)

		suite.T().Cleanup(suite.deleteResourceFunc(svc.Id))

		failoverData = &pbcatalog.FailoverPolicy{
			Config: &pbcatalog.FailoverConfig{
				Destinations: []*pbcatalog.FailoverDestination{{
					Ref: otherServiceRef,
				}},
			},
		}
		failover = rtest.Resource(pbcatalog.FailoverPolicyType, "api").
			WithData(suite.T(), failoverData).
			WithTenancy(tenancy).
			Write(suite.T(), suite.client)

		suite.T().Cleanup(suite.deleteResourceFunc(failover.Id))

		suite.client.WaitForStatusCondition(suite.T(), failover.Id, StatusKey, ConditionUnknownDestinationPort(otherServiceRef, "bar"))

		// and fix it the silly way by removing it from api+failover
		apiServiceData = &pbcatalog.Service{
			Workloads: &pbcatalog.WorkloadSelector{Prefixes: []string{"api-"}},
			Ports: []*pbcatalog.ServicePort{
				{
					TargetPort: "foo",
					Protocol:   pbcatalog.Protocol_PROTOCOL_HTTP,
				},
			},
		}
		svc = rtest.Resource(pbcatalog.ServiceType, "api").
			WithData(suite.T(), apiServiceData).
			WithTenancy(tenancy).
			Write(suite.T(), suite.client)

		suite.T().Cleanup(suite.deleteResourceFunc(svc.Id))

		suite.client.WaitForStatusCondition(suite.T(), failover.Id, StatusKey, ConditionOK)
	})
}

func TestFailoverController(t *testing.T) {
	suite.Run(t, new(controllerSuite))
}

func (suite *controllerSuite) runTestCaseWithTenancies(testCase func(tenancy *pbresource.Tenancy)) {
	for _, tenancy := range suite.tenancies {
		suite.Run(suite.appendTenancyInfo(tenancy), func() {
			testCase(tenancy)
		})
	}
}

func (suite *controllerSuite) appendTenancyInfo(tenancy *pbresource.Tenancy) string {
	return fmt.Sprintf("%s_Namespace_%s_Partition", tenancy.Namespace, tenancy.Partition)
}

func (suite *controllerSuite) deleteResourceFunc(id *pbresource.ID) func() {
	return func() {
		suite.client.MustDelete(suite.T(), id)
	}
}
