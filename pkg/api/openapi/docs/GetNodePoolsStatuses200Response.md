# GetNodePoolsStatuses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** |  | 
**Size** | **int32** |  | 
**Total** | **int32** |  | 
**Items** | [**[]AdapterStatus**](AdapterStatus.md) |  | 

## Methods

### NewGetNodePoolsStatuses200Response

`func NewGetNodePoolsStatuses200Response(page int32, size int32, total int32, items []AdapterStatus, ) *GetNodePoolsStatuses200Response`

NewGetNodePoolsStatuses200Response instantiates a new GetNodePoolsStatuses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNodePoolsStatuses200ResponseWithDefaults

`func NewGetNodePoolsStatuses200ResponseWithDefaults() *GetNodePoolsStatuses200Response`

NewGetNodePoolsStatuses200ResponseWithDefaults instantiates a new GetNodePoolsStatuses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetNodePoolsStatuses200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetNodePoolsStatuses200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetNodePoolsStatuses200Response) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *GetNodePoolsStatuses200Response) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetNodePoolsStatuses200Response) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetNodePoolsStatuses200Response) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTotal

`func (o *GetNodePoolsStatuses200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetNodePoolsStatuses200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetNodePoolsStatuses200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *GetNodePoolsStatuses200Response) GetItems() []AdapterStatus`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetNodePoolsStatuses200Response) GetItemsOk() (*[]AdapterStatus, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetNodePoolsStatuses200Response) SetItems(v []AdapterStatus)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


