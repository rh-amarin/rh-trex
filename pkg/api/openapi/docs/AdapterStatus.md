# AdapterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Resource identifier | [optional] 
**Kind** | Pointer to **string** | Resource kind | [optional] 
**Href** | Pointer to **string** | Resource URI | [optional] 
**Labels** | **map[string]string** | labels for the API resource as pairs of name:value strings | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**AdapterName** | **string** | Name of the adapter that generated this status (Validator, DNS...) | 
**OwnerReferences** | [**ObjectReference**](ObjectReference.md) | Reference to the cluster to which this adapter status pertains | 
**ObservedGeneration** | **int32** | Which generation for an entity (Clusters, NodePools) was current at the time of creating this status | 
**Conditions** | [**[]Condition**](Condition.md) | Kubernetes-style conditions tracking adapter state | 
**Data** | Pointer to **map[string]interface{}** | Adapter-specific data (structure varies by adapter type) | [optional] 
**Metadata** | Pointer to [**AdapterStatusMetadata**](AdapterStatusMetadata.md) |  | [optional] 

## Methods

### NewAdapterStatus

`func NewAdapterStatus(labels map[string]string, createdAt time.Time, updatedAt time.Time, adapterName string, ownerReferences ObjectReference, observedGeneration int32, conditions []Condition, ) *AdapterStatus`

NewAdapterStatus instantiates a new AdapterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterStatusWithDefaults

`func NewAdapterStatusWithDefaults() *AdapterStatus`

NewAdapterStatusWithDefaults instantiates a new AdapterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdapterStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdapterStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdapterStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdapterStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AdapterStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AdapterStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AdapterStatus) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AdapterStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *AdapterStatus) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AdapterStatus) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AdapterStatus) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AdapterStatus) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLabels

`func (o *AdapterStatus) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AdapterStatus) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AdapterStatus) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetCreatedAt

`func (o *AdapterStatus) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdapterStatus) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdapterStatus) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AdapterStatus) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdapterStatus) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdapterStatus) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAdapterName

`func (o *AdapterStatus) GetAdapterName() string`

GetAdapterName returns the AdapterName field if non-nil, zero value otherwise.

### GetAdapterNameOk

`func (o *AdapterStatus) GetAdapterNameOk() (*string, bool)`

GetAdapterNameOk returns a tuple with the AdapterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterName

`func (o *AdapterStatus) SetAdapterName(v string)`

SetAdapterName sets AdapterName field to given value.


### GetOwnerReferences

`func (o *AdapterStatus) GetOwnerReferences() ObjectReference`

GetOwnerReferences returns the OwnerReferences field if non-nil, zero value otherwise.

### GetOwnerReferencesOk

`func (o *AdapterStatus) GetOwnerReferencesOk() (*ObjectReference, bool)`

GetOwnerReferencesOk returns a tuple with the OwnerReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerReferences

`func (o *AdapterStatus) SetOwnerReferences(v ObjectReference)`

SetOwnerReferences sets OwnerReferences field to given value.


### GetObservedGeneration

`func (o *AdapterStatus) GetObservedGeneration() int32`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *AdapterStatus) GetObservedGenerationOk() (*int32, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *AdapterStatus) SetObservedGeneration(v int32)`

SetObservedGeneration sets ObservedGeneration field to given value.


### GetConditions

`func (o *AdapterStatus) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AdapterStatus) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AdapterStatus) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.


### GetData

`func (o *AdapterStatus) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdapterStatus) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdapterStatus) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AdapterStatus) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMetadata

`func (o *AdapterStatus) GetMetadata() AdapterStatusMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AdapterStatus) GetMetadataOk() (*AdapterStatusMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AdapterStatus) SetMetadata(v AdapterStatusMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AdapterStatus) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


