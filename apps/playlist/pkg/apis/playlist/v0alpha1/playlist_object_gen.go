//
// Code generated by grafana-app-sdk. DO NOT EDIT.
//

package v0alpha1

import (
	"fmt"
	"github.com/grafana/grafana-app-sdk/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"time"
)

// +k8s:openapi-gen=true
type Playlist struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`

	// Spec is the spec of the Playlist
	Spec PlaylistSpec `json:"spec" yaml:"spec"`

	Status PlaylistStatus `json:"status" yaml:"status"`
}

func (o *Playlist) GetSpec() any {
	return o.Spec
}

func (o *Playlist) SetSpec(spec any) error {
	cast, ok := spec.(PlaylistSpec)
	if !ok {
		return fmt.Errorf("cannot set spec type %#v, not of type Spec", spec)
	}
	o.Spec = cast
	return nil
}

func (o *Playlist) GetSubresources() map[string]any {
	return map[string]any{
		"status": o.Status,
	}
}

func (o *Playlist) GetSubresource(name string) (any, bool) {
	switch name {
	case "status":
		return o.Status, true
	default:
		return nil, false
	}
}

func (o *Playlist) SetSubresource(name string, value any) error {
	switch name {
	case "status":
		cast, ok := value.(PlaylistStatus)
		if !ok {
			return fmt.Errorf("cannot set status type %#v, not of type PlaylistStatus", value)
		}
		o.Status = cast
		return nil
	default:
		return fmt.Errorf("subresource '%s' does not exist", name)
	}
}

func (o *Playlist) GetStaticMetadata() resource.StaticMetadata {
	gvk := o.GroupVersionKind()
	return resource.StaticMetadata{
		Name:      o.ObjectMeta.Name,
		Namespace: o.ObjectMeta.Namespace,
		Group:     gvk.Group,
		Version:   gvk.Version,
		Kind:      gvk.Kind,
	}
}

func (o *Playlist) SetStaticMetadata(metadata resource.StaticMetadata) {
	o.Name = metadata.Name
	o.Namespace = metadata.Namespace
	o.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   metadata.Group,
		Version: metadata.Version,
		Kind:    metadata.Kind,
	})
}

func (o *Playlist) GetCommonMetadata() resource.CommonMetadata {
	dt := o.DeletionTimestamp
	var deletionTimestamp *time.Time
	if dt != nil {
		deletionTimestamp = &dt.Time
	}
	// Legacy ExtraFields support
	extraFields := make(map[string]any)
	if o.Annotations != nil {
		extraFields["annotations"] = o.Annotations
	}
	if o.ManagedFields != nil {
		extraFields["managedFields"] = o.ManagedFields
	}
	if o.OwnerReferences != nil {
		extraFields["ownerReferences"] = o.OwnerReferences
	}
	return resource.CommonMetadata{
		UID:               string(o.UID),
		ResourceVersion:   o.ResourceVersion,
		Generation:        o.Generation,
		Labels:            o.Labels,
		CreationTimestamp: o.CreationTimestamp.Time,
		DeletionTimestamp: deletionTimestamp,
		Finalizers:        o.Finalizers,
		UpdateTimestamp:   o.GetUpdateTimestamp(),
		CreatedBy:         o.GetCreatedBy(),
		UpdatedBy:         o.GetUpdatedBy(),
		ExtraFields:       extraFields,
	}
}

func (o *Playlist) SetCommonMetadata(metadata resource.CommonMetadata) {
	o.UID = types.UID(metadata.UID)
	o.ResourceVersion = metadata.ResourceVersion
	o.Generation = metadata.Generation
	o.Labels = metadata.Labels
	o.CreationTimestamp = metav1.NewTime(metadata.CreationTimestamp)
	if metadata.DeletionTimestamp != nil {
		dt := metav1.NewTime(*metadata.DeletionTimestamp)
		o.DeletionTimestamp = &dt
	} else {
		o.DeletionTimestamp = nil
	}
	o.Finalizers = metadata.Finalizers
	if o.Annotations == nil {
		o.Annotations = make(map[string]string)
	}
	if !metadata.UpdateTimestamp.IsZero() {
		o.SetUpdateTimestamp(metadata.UpdateTimestamp)
	}
	if metadata.CreatedBy != "" {
		o.SetCreatedBy(metadata.CreatedBy)
	}
	if metadata.UpdatedBy != "" {
		o.SetUpdatedBy(metadata.UpdatedBy)
	}
	// Legacy support for setting Annotations, ManagedFields, and OwnerReferences via ExtraFields
	if metadata.ExtraFields != nil {
		if annotations, ok := metadata.ExtraFields["annotations"]; ok {
			if cast, ok := annotations.(map[string]string); ok {
				o.Annotations = cast
			}
		}
		if managedFields, ok := metadata.ExtraFields["managedFields"]; ok {
			if cast, ok := managedFields.([]metav1.ManagedFieldsEntry); ok {
				o.ManagedFields = cast
			}
		}
		if ownerReferences, ok := metadata.ExtraFields["ownerReferences"]; ok {
			if cast, ok := ownerReferences.([]metav1.OwnerReference); ok {
				o.OwnerReferences = cast
			}
		}
	}
}

func (o *Playlist) GetCreatedBy() string {
	if o.ObjectMeta.Annotations == nil {
		o.ObjectMeta.Annotations = make(map[string]string)
	}

	return o.ObjectMeta.Annotations["grafana.com/createdBy"]
}

func (o *Playlist) SetCreatedBy(createdBy string) {
	if o.ObjectMeta.Annotations == nil {
		o.ObjectMeta.Annotations = make(map[string]string)
	}

	o.ObjectMeta.Annotations["grafana.com/createdBy"] = createdBy
}

func (o *Playlist) GetUpdateTimestamp() time.Time {
	if o.ObjectMeta.Annotations == nil {
		o.ObjectMeta.Annotations = make(map[string]string)
	}

	parsed, _ := time.Parse(time.RFC3339, o.ObjectMeta.Annotations["grafana.com/updateTimestamp"])
	return parsed
}

func (o *Playlist) SetUpdateTimestamp(updateTimestamp time.Time) {
	if o.ObjectMeta.Annotations == nil {
		o.ObjectMeta.Annotations = make(map[string]string)
	}

	o.ObjectMeta.Annotations["grafana.com/updateTimestamp"] = updateTimestamp.Format(time.RFC3339)
}

func (o *Playlist) GetUpdatedBy() string {
	if o.ObjectMeta.Annotations == nil {
		o.ObjectMeta.Annotations = make(map[string]string)
	}

	return o.ObjectMeta.Annotations["grafana.com/updatedBy"]
}

func (o *Playlist) SetUpdatedBy(updatedBy string) {
	if o.ObjectMeta.Annotations == nil {
		o.ObjectMeta.Annotations = make(map[string]string)
	}

	o.ObjectMeta.Annotations["grafana.com/updatedBy"] = updatedBy
}

func (o *Playlist) Copy() resource.Object {
	return resource.CopyObject(o)
}

func (o *Playlist) DeepCopyObject() runtime.Object {
	return o.Copy()
}

func (o *Playlist) DeepCopy() *Playlist {
	cpy := &Playlist{}
	o.DeepCopyInto(cpy)
	return cpy
}

func (o *Playlist) DeepCopyInto(dst *Playlist) {
	dst.TypeMeta.APIVersion = o.TypeMeta.APIVersion
	dst.TypeMeta.Kind = o.TypeMeta.Kind
	o.ObjectMeta.DeepCopyInto(&dst.ObjectMeta)
	o.Spec.DeepCopyInto(&dst.Spec)
	o.Status.DeepCopyInto(&dst.Status)
}

// Interface compliance compile-time check
var _ resource.Object = &Playlist{}

// +k8s:openapi-gen=true
type PlaylistList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []Playlist `json:"items" yaml:"items"`
}

func (o *PlaylistList) DeepCopyObject() runtime.Object {
	return o.Copy()
}

func (o *PlaylistList) Copy() resource.ListObject {
	cpy := &PlaylistList{
		TypeMeta: o.TypeMeta,
		Items:    make([]Playlist, len(o.Items)),
	}
	o.ListMeta.DeepCopyInto(&cpy.ListMeta)
	for i := 0; i < len(o.Items); i++ {
		if item, ok := o.Items[i].Copy().(*Playlist); ok {
			cpy.Items[i] = *item
		}
	}
	return cpy
}

func (o *PlaylistList) GetItems() []resource.Object {
	items := make([]resource.Object, len(o.Items))
	for i := 0; i < len(o.Items); i++ {
		items[i] = &o.Items[i]
	}
	return items
}

func (o *PlaylistList) SetItems(items []resource.Object) {
	o.Items = make([]Playlist, len(items))
	for i := 0; i < len(items); i++ {
		o.Items[i] = *items[i].(*Playlist)
	}
}

func (o *PlaylistList) DeepCopy() *PlaylistList {
	cpy := &PlaylistList{}
	o.DeepCopyInto(cpy)
	return cpy
}

func (o *PlaylistList) DeepCopyInto(dst *PlaylistList) {
	resource.CopyObjectInto(dst, o)
}

// Interface compliance compile-time check
var _ resource.ListObject = &PlaylistList{}

// Copy methods for all subresource types

// DeepCopy creates a full deep copy of Spec
func (s *PlaylistSpec) DeepCopy() *PlaylistSpec {
	cpy := &PlaylistSpec{}
	s.DeepCopyInto(cpy)
	return cpy
}

// DeepCopyInto deep copies Spec into another Spec object
func (s *PlaylistSpec) DeepCopyInto(dst *PlaylistSpec) {
	resource.CopyObjectInto(dst, s)
}

// DeepCopy creates a full deep copy of PlaylistStatus
func (s *PlaylistStatus) DeepCopy() *PlaylistStatus {
	cpy := &PlaylistStatus{}
	s.DeepCopyInto(cpy)
	return cpy
}

// DeepCopyInto deep copies PlaylistStatus into another PlaylistStatus object
func (s *PlaylistStatus) DeepCopyInto(dst *PlaylistStatus) {
	resource.CopyObjectInto(dst, s)
}
