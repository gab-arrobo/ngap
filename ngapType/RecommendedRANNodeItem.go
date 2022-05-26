// Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University (free5gc.org)
//
// SPDX-License-Identifier: Apache-2.0

package ngapType

// Need to import "github.com/omec-project/aper" if it uses "aper"

type RecommendedRANNodeItem struct {
	AMFPagingTarget AMFPagingTarget                                         `aper:"valueLB:0,valueUB:2"`
	IEExtensions    *ProtocolExtensionContainerRecommendedRANNodeItemExtIEs `aper:"optional"`
}
