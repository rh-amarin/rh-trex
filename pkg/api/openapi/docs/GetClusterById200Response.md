# GetClusterById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Resource kind | 
**Name** | **string** | Cluster name (unique) | 
**Spec** | **map[string]map[string]interface{}** | Cluster specification CLM doesn&#39;t know how to unmarshall the spec, it only stores and forwards to adapters to do their job But CLM will validate the schema before accepting the request | 
**Id** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** | Resource URI | [optional] 
**Labels** | **map[string]string** | labels for the API resource as pairs of name:value strings | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Generation** | **int32** | Generation field is updated on customer updates, reflecting the version of the \&quot;intent\&quot; of the customer | 
**Status** | [**ClusterStatus**](ClusterStatus.md) |  | 
**CreatedBy** | **string** |  | 
**UpdatedBy** | **string** |  | 
**Code** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**OperationId** | Pointer to **string** |  | [optional] 

## Methods

### NewGetClusterById200Response

`func NewGetClusterById200Response(kind string, name string, spec map[string]map[string]interface{}, labels map[string]string, createdAt time.Time, updatedAt time.Time, generation int32, status ClusterStatus, createdBy string, updatedBy string, ) *GetClusterById200Response`

NewGetClusterById200Response instantiates a new GetClusterById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterById200ResponseWithDefaults

`func NewGetClusterById200ResponseWithDefaults() *GetClusterById200Response`

NewGetClusterById200ResponseWithDefaults instantiates a new GetClusterById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *GetClusterById200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetClusterById200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetClusterById200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *GetClusterById200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterById200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterById200Response) SetName(v string)`

SetName sets Name field to given value.


### GetSpec

`func (o *GetClusterById200Response) GetSpec() map[string]map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *GetClusterById200Response) GetSpecOk() (*map[string]map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *GetClusterById200Response) SetSpec(v map[string]map[string]interface{})`

SetSpec sets Spec field to given value.


### GetId

`func (o *GetClusterById200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterById200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterById200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterById200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *GetClusterById200Response) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GetClusterById200Response) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GetClusterById200Response) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *GetClusterById200Response) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLabels

`func (o *GetClusterById200Response) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetClusterById200Response) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetClusterById200Response) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetCreatedAt

`func (o *GetClusterById200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetClusterById200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetClusterById200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetClusterById200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetClusterById200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetClusterById200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetGeneration

`func (o *GetClusterById200Response) GetGeneration() int32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *GetClusterById200Response) GetGenerationOk() (*int32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *GetClusterById200Response) SetGeneration(v int32)`

SetGeneration sets Generation field to given value.


### GetStatus

`func (o *GetClusterById200Response) GetStatus() ClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetClusterById200Response) GetStatusOk() (*ClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetClusterById200Response) SetStatus(v ClusterStatus)`

SetStatus sets Status field to given value.


### GetCreatedBy

`func (o *GetClusterById200Response) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetClusterById200Response) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetClusterById200Response) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedBy

`func (o *GetClusterById200Response) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetClusterById200Response) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetClusterById200Response) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetCode

`func (o *GetClusterById200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetClusterById200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetClusterById200Response) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetClusterById200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *GetClusterById200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetClusterById200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetClusterById200Response) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetClusterById200Response) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetOperationId

`func (o *GetClusterById200Response) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *GetClusterById200Response) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *GetClusterById200Response) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *GetClusterById200Response) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


