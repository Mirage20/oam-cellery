/*
 * Copyright (c) 2020 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package informers

import (
	"fmt"
	"reflect"
	"time"

	meshinformers "cellery.io/cellery-controller/pkg/generated/informers/externalversions"
	meshv1alpha2 "cellery.io/cellery-controller/pkg/generated/informers/externalversions/mesh/v1alpha2"
	"cellery.io/oam-cellery/pkg/clients"
	oamcoreinformers "cellery.io/oam-cellery/pkg/generated/informers/externalversions"
	oamcorev1alpha1 "cellery.io/oam-cellery/pkg/generated/informers/externalversions/core/v1alpha1"
	kubeinformers "k8s.io/client-go/informers"
)

type Interface interface {
	// Cellery mesh informers
	Cells() meshv1alpha2.CellInformer
	Composites() meshv1alpha2.CompositeInformer

	// OAM core informers
	ComponentSchematics() oamcorev1alpha1.ComponentSchematicInformer
}

type informers struct {
	kubeInformerFactory    kubeinformers.SharedInformerFactory
	meshInformerFactory    meshinformers.SharedInformerFactory
	oamCoreInformerFactory oamcoreinformers.SharedInformerFactory
}

func New(clients clients.Interface, resync time.Duration) *informers {
	return &informers{
		kubeInformerFactory:    kubeinformers.NewSharedInformerFactory(clients.Kubernetes(), resync),
		meshInformerFactory:    meshinformers.NewSharedInformerFactory(clients.Mesh(), resync),
		oamCoreInformerFactory: oamcoreinformers.NewSharedInformerFactory(clients.Oam(), resync),
	}
}

func (i *informers) Start(stopCh <-chan struct{}) error {
	i.kubeInformerFactory.Start(stopCh)
	i.meshInformerFactory.Start(stopCh)
	i.oamCoreInformerFactory.Start(stopCh)
	return func(results ...map[reflect.Type]bool) error {
		for i, _ := range results {
			for t, ok := range results[i] {
				if !ok {
					return fmt.Errorf("failed to wait for cache with type %s", t)
				}
			}
		}
		return nil
	}(i.kubeInformerFactory.WaitForCacheSync(stopCh), i.meshInformerFactory.WaitForCacheSync(stopCh), i.oamCoreInformerFactory.WaitForCacheSync(stopCh))
}

func (i *informers) Cells() meshv1alpha2.CellInformer {
	return i.meshInformerFactory.Mesh().V1alpha2().Cells()
}

func (i *informers) Composites() meshv1alpha2.CompositeInformer {
	return i.meshInformerFactory.Mesh().V1alpha2().Composites()
}

func (i *informers) ComponentSchematics() oamcorev1alpha1.ComponentSchematicInformer {
	return i.oamCoreInformerFactory.Core().V1alpha1().ComponentSchematics()
}
