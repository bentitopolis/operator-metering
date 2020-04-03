// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/bentitopolis/operator-metering/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// HiveTables returns a HiveTableInformer.
	HiveTables() HiveTableInformer
	// MeteringConfigs returns a MeteringConfigInformer.
	MeteringConfigs() MeteringConfigInformer
	// PrestoTables returns a PrestoTableInformer.
	PrestoTables() PrestoTableInformer
	// Reports returns a ReportInformer.
	Reports() ReportInformer
	// ReportDataSources returns a ReportDataSourceInformer.
	ReportDataSources() ReportDataSourceInformer
	// ReportQueries returns a ReportQueryInformer.
	ReportQueries() ReportQueryInformer
	// StorageLocations returns a StorageLocationInformer.
	StorageLocations() StorageLocationInformer
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

// HiveTables returns a HiveTableInformer.
func (v *version) HiveTables() HiveTableInformer {
	return &hiveTableInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MeteringConfigs returns a MeteringConfigInformer.
func (v *version) MeteringConfigs() MeteringConfigInformer {
	return &meteringConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PrestoTables returns a PrestoTableInformer.
func (v *version) PrestoTables() PrestoTableInformer {
	return &prestoTableInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Reports returns a ReportInformer.
func (v *version) Reports() ReportInformer {
	return &reportInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ReportDataSources returns a ReportDataSourceInformer.
func (v *version) ReportDataSources() ReportDataSourceInformer {
	return &reportDataSourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ReportQueries returns a ReportQueryInformer.
func (v *version) ReportQueries() ReportQueryInformer {
	return &reportQueryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StorageLocations returns a StorageLocationInformer.
func (v *version) StorageLocations() StorageLocationInformer {
	return &storageLocationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
