package model

import (
	"encoding/xml"
)

type AccountBillingInfoList struct {
	XMLName xml.Name `xml:"account_billing_objmt_infos" json:"account_billing_objmt_infos"`

	Status   string `xml:"status,omitempty" json:"status,omitempty"`
	SizeUnit string `xml:"size_unit,omitempty" json:"size_unit,omitempty"`
	DateTime string `xml:"date_time,omitempty" json:"date_time,omitempty"`

	Info []AccountBillingInfo `xml:"account_billing_objmt_info" json:"account_billing_objmt_info"`
}

type AccountBillingInfo struct {
	XMLName xml.Name `xml:"account_billing_objmt_info" json:"account_billing_objmt_info"`

	AccountId      string `xml:"account_id,omitempty"`
	ConsistentTime int64  `xml:"consistent_time,omitempty"`

	TotalUserObjectMetric    []StorageClassBasedCountSize `xml:"total_user_object_metric>storage_class_counts,omitempty"`
	TotalMPUMetric           []StorageClassBasedCountSize `xml:"total_mpu_metric>storage_class_counts,omitempty"`
	TotalReplicaObjectMetric []StorageClassBasedCountSize `xml:"total_replica_object_metric>storage_class_counts,omitempty"`
}

type AccountBillingSampleList struct {
	XMLName xml.Name `xml:"account_billing_objmt_samples" json:"account_billing_objmt_samples"`

	Status    string `xml:"status,omitempty" json:"status,omitempty"`
	SizeUnit  string `xml:"size_unit,omitempty" json:"size_unit,omitempty"`
	DateTime  string `xml:"date_time,omitempty" json:"date_time,omitempty"`
	StartTime string `xml:"start_time,omitempty" json:"start_time,omitempty"`
	EndTime   string `xml:"end_time,omitempty" json:"end_time,omitempty"`

	Samples []AccountBillingSample `json:"account_billing_objmt_sample" xml:"account_billing_objmt_sample"`
}

type AccountBillingSample struct {
	XMLName xml.Name `xml:"account_billing_objmt_sample" json:"account_billing_objmt_sample"`

	AccountId       string `xml:"account_id,omitempty"`
	StartTime       string `xml:"start_time,omitempty"`
	EndTime         string `xml:"end_time,omitempty"`
	SampleTimeRange int64  `xml:"sample_time_range,omitempty"`
	ConsistentTime  int64  `xml:"consistent_time,omitempty"`

	AccountBillingInfo AccountBillingInfo `xml:"account_billing_objmt_info,omitempty"`

	UserCreationDelta    []StorageClassBasedCountSize `xml:"user_creation_delta>storage_class_counts"`
	UserDeletionDelta    []StorageClassBasedCountSize `xml:"user_deletion_delta>storage_class_counts"`
	MpuCreateDelta       []StorageClassBasedCountSize `xml:"mpu_create_delta>storage_class_counts"`
	MpuDeleteDelta       []StorageClassBasedCountSize `xml:"mpu_delete_delta>storage_class_counts"`
	ReplicaCreationDelta []StorageClassBasedCountSize `xml:"replica_creation_delta>storage_class_counts"`
	ReplicaDeletionDelta []StorageClassBasedCountSize `xml:"replica_deletion_delta>storage_class_counts"`
}

type StorageClassBasedCountSize struct {
	XMLName xml.Name `json:"storage_class_counts" xml:"storage_class_counts"`

	StorageClass       string `xml:"storage_class,omitempty"`
	Counts             int64  `xml:"count_size>counts,omitempty"`
	LogicalSize        int64  `xml:"count_size>logical_size,omitempty"`
	CreateLogicalSize  int64  `xml:"count_size>create_logical_size,omitempty"`
	DeleteLogicalSize  int64  `xml:"count_size>delete_logical_size,omitempty"`
	PhysicalSize       int64  `xml:"count_size>physical_size,omitempty"`
	CreatePhysicalSize int64  `xml:"count_size>create_physical_size,omitempty"`
	DeletePhysicalSize int64  `xml:"count_size>delete_physical_size,omitempty"`
}

type BucketBillingInfoList struct {
	XMLName xml.Name `xml:"bucket_billing_objmt_infos" json:"bucket_billing_objmt_infos"`

	Status   string `xml:"status,omitempty" json:"status,omitempty"`
	SizeUnit string `xml:"size_unit,omitempty" json:"size_unit,omitempty"`
	DateTime string `xml:"date_time,omitempty" json:"date_time,omitempty"`

	Info []BucketBillingInfo `xml:"bucket_billing_objmt_info" json:"bucket_billing_objmt_info"`
}

type BucketBillingInfo struct {
	XMLName xml.Name `xml:"bucket_billing_objmt_info" json:"bucket_billing_objmt_info"`

	BucketName       string  `xml:"bucket_name,omitempty"`
	CompressionRatio float64 `xml:"compression_ratio,omitempty"`
	ConsistentTime   int64   `xml:"consistent_time,omitempty"`

	TotalUserObjectMetric    []StorageClassBasedCountSize `xml:"total_user_object_metric>storage_class_counts"`
	TotalMPUMetric           []StorageClassBasedCountSize `xml:"total_mpu_metric>storage_class_counts"`
	TotalReplicaObjectMetric []StorageClassBasedCountSize `xml:"total_replica_object_metric>storage_class_counts"`
}

type BucketBillingSampleList struct {
	XMLName xml.Name `xml:"bucket_billing_objmt_samples" json:"bucket_billing_objmt_samples"`

	Status    string `xml:"status,omitempty" json:"status,omitempty"`
	SizeUnit  string `xml:"size_unit,omitempty" json:"size_unit,omitempty"`
	DateTime  string `xml:"date_time,omitempty" json:"date_time,omitempty"`
	StartTime string `xml:"start_time,omitempty" json:"start_time,omitempty"`
	EndTime   string `xml:"end_time,omitempty" json:"end_time,omitempty"`

	Samples []BucketBillingSample `json:"bucket_billing_objmt_sample" xml:"bucket_billing_objmt_sample"`
}

type BucketBillingSample struct {
	XMLName xml.Name `xml:"bucket_billing_objmt_sample" json:"bucket_billing_objmt_sample"`

	BucketName      string `xml:"bucket_name,omitempty"`
	SampleTimeRange int64  `xml:"sample_time_range,omitempty"`
	ConsistentTime  int64  `xml:"consistent_time,omitempty"`

	BucketBillingInfo BucketBillingInfo `xml:"bucket_billing_objmt_info,omitempty"`

	BucketBillingTags []BucketBillingTag `xml:"bucket_billing_tag,omitempty"`

	UserCreationDelta    []StorageClassBasedCountSize `xml:"user_creation_delta>storage_class_counts"`
	UserDeletionDelta    []StorageClassBasedCountSize `xml:"user_deletion_delta>storage_class_counts"`
	MpuCreateDelta       []StorageClassBasedCountSize `xml:"mpu_create_delta>storage_class_counts"`
	MpuDeleteDelta       []StorageClassBasedCountSize `xml:"mpu_delete_delta>storage_class_counts"`
	ReplicaCreationDelta []StorageClassBasedCountSize `xml:"replica_creation_delta>storage_class_counts"`
	ReplicaDeletionDelta []StorageClassBasedCountSize `xml:"replica_deletion_delta>storage_class_counts"`
}

type BucketBillingTag struct {
	XMLName xml.Name `xml:"bucket_billing_tag" json:"bucket_billing_tag"`

	BucketName string `xml:"bucket_name,omitempty"`
	Ingress    int64  `xml:"ingress,omitempty"`
	Egress     int64  `xml:"egress,omitempty"`
}

type BucketPerfDataList struct {
	XMLName xml.Name `xml:"bucket_perf_samples" json:"bucket_perf_samples"`

	Status    string `xml:"status,omitempty" json:"status,omitempty"`
	SizeUnit  string `xml:"size_unit,omitempty" json:"size_unit,omitempty"`
	DateTime  string `xml:"date_time,omitempty" json:"date_time,omitempty"`
	StartTime string `xml:"start_time,omitempty" json:"start_time,omitempty"`
	EndTime   string `xml:"end_time,omitempty" json:"end_time,omitempty"`

	Samples []BucketPerfSample `xml:"bucket_perf_sample" json:"bucket_perf_sample" `
}

type BucketPerfSample struct {
	XMLName xml.Name `xml:"bucket_perf_sample" json:"bucket_perf_sample"`

	BucketName      string `xml:"bucket_name,omitempty"`
	SampleTimeRange int64  `xml:"sample_time_range,omitempty"`
	ConsistentTime  int64  `xml:"consistent_time,omitempty"`

	IngressLatency int64 `xml:"ingress_latency,omitempty"`
	IngressBytes   int64 `xml:"ingress_bytes,omitempty"`
	IngressCounts  int64 `xml:"ingress_counts,omitempty"`
	EgressBytes    int64 `xml:"egress_bytes,omitempty"`
	EgressCounts   int64 `xml:"egress_counts,omitempty"`
}

type BucketReplicationSampleList struct {
	XMLName xml.Name `xml:"bucket_replication_samples" json:"bucket_replication_samples"`

	Status    string `xml:"status,omitempty" json:"status,omitempty"`
	SizeUnit  string `xml:"size_unit,omitempty" json:"size_unit,omitempty"`
	DateTime  string `xml:"date_time,omitempty" json:"date_time,omitempty"`
	StartTime string `xml:"start_time,omitempty" json:"start_time,omitempty"`
	EndTime   string `xml:"end_time,omitempty" json:"end_time,omitempty"`

	Samples []BucketReplicationSample `xml:"bucket_replication_sample,omitempty" json:"bucket_replication_sample,omitempty"`
}

type BucketReplicationSample struct {
	XMLName xml.Name `xml:"bucket_replication_sample" json:"bucket_replication_sample"`

	SourceBucket      string `xml:"replication_source_destination>source_bucket,omitempty"`
	DestinationBucket string `xml:"replication_source_destination>destination_bucket_arn,omitempty"`

	SampleTimeRange int64 `xml:"sample_time_range,omitempty"`
	ConsistentTime  int64 `xml:"consistent_time,omitempty"`

	ReplicationBillingInfo ReplicationBillingInfo `xml:"replication_billing_info,omitempty"`

	PendingToReplicateDelta []StorageClassBasedCountSize `xml:"pending_to_replicate_delta>storage_class_counts,omitempty"`
	ReplicatedDelta         []StorageClassBasedCountSize `xml:"replicated_delta>storage_class_counts,omitempty"`
	ReplicatedFailureDelta  []StorageClassBasedCountSize `xml:"replicated_failure_delta>storage_class_counts,omitempty"`
}

type BucketReplicationInfoList struct {
	XMLName xml.Name `xml:"replication_info_list" json:"replication_info_list"`

	Status   string `xml:"status,omitempty" json:"status,omitempty"`
	SizeUnit string `xml:"size_unit,omitempty" json:"size_unit,omitempty"`
	DateTime string `xml:"date_time,omitempty" json:"date_time,omitempty"`

	Info []ReplicationBillingInfo `xml:"replication_billing_info" json:"replication_billing_info"`
}

type ReplicationBillingInfo struct {
	XMLName xml.Name `xml:"replication_billing_info" json:"replication_billing_info"`

	SourceBucket      string `xml:"replication_source_destination>source_bucket,omitempty"`
	DestinationBucket string `xml:"replication_source_destination>destination_bucket_arn,omitempty"`
	ConsistentTime    int64  `xml:"consistent_time,omitempty"`

	PendingToReplicate []StorageClassBasedCountSize `xml:"pending_to_replicate>storage_class_counts,omitempty"`
}

type StoreBillingInfoList struct {
	XMLName xml.Name `xml:"store_billing_info_list" json:"store_billing_info_list"`

	Status   string `xml:"status,omitempty" json:"status,omitempty"`
	SizeUnit string `xml:"size_unit,omitempty" json:"size_unit,omitempty"`
	DateTime string `xml:"date_time,omitempty" json:"date_time,omitempty"`

	Info StoreBillingInfo `xml:"store_billing_info,omitempty" json:"store_billing_info,omitempty"`
}

type StoreBillingInfo struct {
	XMLName xml.Name `xml:"store_billing_info" json:"store_billing_info"`

	CompressionRatio float64 `xml:"compression_ratio,omitempty"`
	ConsistentTime   int64   `xml:"consistent_time,omitempty"`

	TotalUserObjectMetric    []StorageClassBasedCountSize `xml:"total_user_object_metric>storage_class_counts,omitempty"`
	TotalMPUMetric           []StorageClassBasedCountSize `xml:"total_mpu_metric>storage_class_counts,omitempty"`
	TotalReplicaObjectMetric []StorageClassBasedCountSize `xml:"total_replica_object_metric>storage_class_counts,omitempty"`

	TopBucketsByObjectCount []TopNBucket `xml:"top_n_buckets_by_object_count>top_n_bucket"`
	TopBucketsByObjectSize  []TopNBucket `xml:"top_n_buckets_by_object_size>top_n_bucket"`
}

type StoreBillingSampleList struct {
	XMLName xml.Name `xml:"store_billing_samples" json:"store_billing_samples"`

	Status    string `xml:"status,omitempty" json:"status,omitempty"`
	SizeUnit  string `xml:"size_unit,omitempty" json:"size_unit,omitempty"`
	DateTime  string `xml:"date_time,omitempty" json:"date_time,omitempty"`
	StartTime string `xml:"start_time,omitempty" json:"start_time,omitempty"`
	EndTime   string `xml:"end_time,omitempty" json:"end_time,omitempty"`

	Samples []StoreBillingSample `xml:"store_billing_sample,omitempty" json:"store_billing_sample,omitempty"`
}

type StoreBillingSample struct {
	XMLName xml.Name `xml:"store_billing_sample" json:"store_billing_sample"`

	SampleTimeRange int64 `xml:"sample_time_range,omitempty"`
	ConsistentTime  int64 `xml:"consistent_time,omitempty"`

	Info StoreBillingInfo `xml:"store_billing_info,omitempty" json:"store_billing_info,omitempty"`

	UserCreationDelta    []StorageClassBasedCountSize `xml:"user_creation_delta>storage_class_counts"`
	UserDeletionDelta    []StorageClassBasedCountSize `xml:"user_deletion_delta>storage_class_counts"`
	MpuCreateDelta       []StorageClassBasedCountSize `xml:"mpu_create_delta>storage_class_counts"`
	MpuDeleteDelta       []StorageClassBasedCountSize `xml:"mpu_delete_delta>storage_class_counts"`
	ReplicaCreationDelta []StorageClassBasedCountSize `xml:"replica_creation_delta>storage_class_counts"`
	ReplicaDeletionDelta []StorageClassBasedCountSize `xml:"replica_deletion_delta>storage_class_counts"`
}

type TopNBucket struct {
	XMLName xml.Name `xml:"top_n_bucket" json:"top_n_bucket"`

	BucketName   string `xml:"bucket_name,omitempty" json:"bucket_name,omitempty"`
	MetricNumber int64  `xml:"metric_number,omitempty" json:"metric_number,omitempty"`
}

type StoreReplicationData struct {
	XMLName xml.Name `xml:"store_replication_list" json:"store_replication_list"`

	Status    string `xml:"status,omitempty" json:"status,omitempty"`
	SizeUnit  string `xml:"size_unit,omitempty" json:"size_unit,omitempty"`
	DateTime  string `xml:"date_time,omitempty" json:"date_time,omitempty"`
	StartTime string `xml:"start_time,omitempty" json:"start_time,omitempty"`
	EndTime   string `xml:"end_time,omitempty" json:"end_time,omitempty"`

	Samples []StoreReplicationThroughputRto `xml:"store_replication_throughput_rto,omitempty" json:"store_replication_throughput_rto,omitempty"`
}

type StoreReplicationThroughputRto struct {
	XMLName xml.Name `xml:"store_replication_throughput_rto" json:"store_replication_throughput_rto"`

	SampleTimeRange int64 `xml:"sample_time_range,omitempty"`
	ConsistentTime  int64 `xml:"consistent_time,omitempty"`

	DestinationStore string `xml:"destination_store,omitempty" json:"destination_store,omitempty"`
	Throughput       int64  `xml:"throughput,omitempty" json:"throughput,omitempty"`
	RTO              int64  `xml:"rto,omitempty" json:"rto,omitempty"`

	PendingToReplicate []StorageClassBasedCountSize `xml:"pending_to_replicate>storage_class_counts,omitempty"`
	ReplicatedDelta    []StorageClassBasedCountSize `xml:"replicated_delta>storage_class_counts,omitempty"`
}
