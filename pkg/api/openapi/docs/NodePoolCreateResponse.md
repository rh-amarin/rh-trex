# NodePoolCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Resource identifier | [optional] 
**Kind** | Pointer to **string** | Resource kind | [optional] 
**Href** | Pointer to **string** | Resource URI | [optional] 
**Labels** | **map[string]string** | labels for the API resource as pairs of name:value strings | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Name** | **string** | NodePool name (unique in a cluster) | 
**Spec** | **map[string]interface{}** | Cluster specification CLM doesn&#39;t know how to unmarshall the spec, it only stores and forwards to adapters to do their job But CLM will validate the schema before accepting the request | 
**OwnerReferences** | [**ObjectReference**](ObjectReference.md) |  | 
**Status** | [**NodePoolStatus**](NodePoolStatus.md) |  | 

## Methods

### NewNodePoolCreateResponse

`func NewNodePoolCreateResponse(labels map[string]string, createdAt time.Time, updatedAt time.Time, name string, spec map[string]interface{}, ownerReferences ObjectReference, status NodePoolStatus, ) *NodePoolCreateResponse`

NewNodePoolCreateResponse instantiates a new NodePoolCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolCreateResponseWithDefaults

`func NewNodePoolCreateResponseWithDefaults() *NodePoolCreateResponse`

NewNodePoolCreateResponseWithDefaults instantiates a new NodePoolCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodePoolCreateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodePoolCreateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodePoolCreateResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NodePoolCreateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *NodePoolCreateResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NodePoolCreateResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NodePoolCreateResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NodePoolCreateResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *NodePoolCreateResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NodePoolCreateResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NodePoolCreateResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *NodePoolCreateResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLabels

`func (o *NodePoolCreateResponse) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NodePoolCreateResponse) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NodePoolCreateResponse) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetCreatedAt

`func (o *NodePoolCreateResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NodePoolCreateResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NodePoolCreateResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NodePoolCreateResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NodePoolCreateResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NodePoolCreateResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *NodePoolCreateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodePoolCreateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodePoolCreateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSpec

`func (o *NodePoolCreateResponse) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NodePoolCreateResponse) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NodePoolCreateResponse) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.


### GetOwnerReferences

`func (o *NodePoolCreateResponse) GetOwnerReferences() ObjectReference`

GetOwnerReferences returns the OwnerReferences field if non-nil, zero value otherwise.

### GetOwnerReferencesOk

`func (o *NodePoolCreateResponse) GetOwnerReferencesOk() (*ObjectReference, bool)`

GetOwnerReferencesOk returns a tuple with the OwnerReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerReferences

`func (o *NodePoolCreateResponse) SetOwnerReferences(v ObjectReference)`

SetOwnerReferences sets OwnerReferences field to given value.


### GetStatus

`func (o *NodePoolCreateResponse) GetStatus() NodePoolStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodePoolCreateResponse) GetStatusOk() (*NodePoolStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodePoolCreateResponse) SetStatus(v NodePoolStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


