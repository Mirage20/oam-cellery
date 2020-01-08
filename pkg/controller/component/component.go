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

package component

import (
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	kubeclient "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"reflect"

	meshclient "cellery.io/cellery-controller/pkg/generated/clientset/versioned"
	"cellery.io/oam-cellery/pkg/apis/core/v1alpha1"
	"cellery.io/oam-cellery/pkg/clients"
	"cellery.io/oam-cellery/pkg/config"
	"cellery.io/oam-cellery/pkg/controller"
	oamclient "cellery.io/oam-cellery/pkg/generated/clientset/versioned"
	v1alpha1listers "cellery.io/oam-cellery/pkg/generated/listers/core/v1alpha1"
	"cellery.io/oam-cellery/pkg/informers"
)

type reconciler struct {
	kubeClient               kubeclient.Interface
	meshClient               meshclient.Interface
	oamClient                oamclient.Interface
	componentSchematicLister v1alpha1listers.ComponentSchematicLister
	cfg                      config.Interface
	logger                   *zap.SugaredLogger
	recorder                 record.EventRecorder
}

func NewController(
	clientset clients.Interface,
	informerset informers.Interface,
	cfg config.Interface,
	logger *zap.SugaredLogger,
) *controller.Controller {
	r := &reconciler{
		kubeClient:               clientset.Kubernetes(),
		meshClient:               clientset.Mesh(),
		oamClient:                clientset.Oam(),
		componentSchematicLister: informerset.ComponentSchematics().Lister(),
		cfg:                      cfg,
		logger:                   logger.Named("component-controller"),
	}
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(r.logger.Named("events").Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: r.kubeClient.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: "component-controller"})
	r.recorder = recorder
	c := controller.New(r, r.logger, "Component")

	r.logger.Info("Setting up event handlers")
	informerset.ComponentSchematics().Informer().AddEventHandler(informers.HandleAll(c.Enqueue))

	return c
}

func (r *reconciler) Reconcile(key string) error {
	r.logger.Infof("Reconcile called with %s", key)
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		r.logger.Errorf("invalid resource key: %s", key)
		return nil
	}
	original, err := r.componentSchematicLister.ComponentSchematics(namespace).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			r.logger.Errorf("component '%s' in work queue no longer exists", key)
			return nil
		}
		return err
	}

	component := original.DeepCopy()

	if err = r.reconcile(component); err != nil {
		r.recorder.Eventf(component, corev1.EventTypeWarning, "InternalError", "Failed to update cluster: %v", err)
		return err
	}

	if equality.Semantic.DeepEqual(original.Status, component.Status) {
		return nil
	}

	if _, err = r.updateStatus(component); err != nil {
		r.recorder.Eventf(component, corev1.EventTypeWarning, "UpdateFailed", "Failed to update status: %v", err)
		return err
	}
	r.recorder.Eventf(component, corev1.EventTypeNormal, "Updated", "Updated Component status %q", component.GetName())
	return nil
}

func (r *reconciler) reconcile(component *v1alpha1.ComponentSchematic) error {
	return nil
}

func (r *reconciler) updateStatus(desired *v1alpha1.ComponentSchematic) (*v1alpha1.ComponentSchematic, error) {
	component, err := r.componentSchematicLister.ComponentSchematics(desired.Namespace).Get(desired.Name)
	if err != nil {
		return nil, err
	}
	if !reflect.DeepEqual(component.Status, desired.Status) {
		latest := component.DeepCopy()
		latest.Status = desired.Status
		return r.oamClient.CoreV1alpha1().ComponentSchematics(desired.Namespace).UpdateStatus(latest)
	}
	return desired, nil
}
