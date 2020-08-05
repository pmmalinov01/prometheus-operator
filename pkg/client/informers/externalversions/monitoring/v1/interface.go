// Copyright 2018 The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/prometheus-operator/prometheus-operator/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Alertmanagers returns a AlertmanagerInformer.
	Alertmanagers() AlertmanagerInformer
	// PodMonitors returns a PodMonitorInformer.
	PodMonitors() PodMonitorInformer
	// Probes returns a ProbeInformer.
	Probes() ProbeInformer
	// Prometheuses returns a PrometheusInformer.
	Prometheuses() PrometheusInformer
	// PrometheusRules returns a PrometheusRuleInformer.
	PrometheusRules() PrometheusRuleInformer
	// ServiceMonitors returns a ServiceMonitorInformer.
	ServiceMonitors() ServiceMonitorInformer
	// ThanosRulers returns a ThanosRulerInformer.
	ThanosRulers() ThanosRulerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Alertmanagers returns a AlertmanagerInformer.
func (v *version) Alertmanagers() AlertmanagerInformer {
	return &alertmanagerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PodMonitors returns a PodMonitorInformer.
func (v *version) PodMonitors() PodMonitorInformer {
	return &podMonitorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Probes returns a ProbeInformer.
func (v *version) Probes() ProbeInformer {
	return &probeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Prometheuses returns a PrometheusInformer.
func (v *version) Prometheuses() PrometheusInformer {
	return &prometheusInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PrometheusRules returns a PrometheusRuleInformer.
func (v *version) PrometheusRules() PrometheusRuleInformer {
	return &prometheusRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceMonitors returns a ServiceMonitorInformer.
func (v *version) ServiceMonitors() ServiceMonitorInformer {
	return &serviceMonitorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ThanosRulers returns a ThanosRulerInformer.
func (v *version) ThanosRulers() ThanosRulerInformer {
	return &thanosRulerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
