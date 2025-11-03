# \DefaultAPI

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNodePool**](DefaultAPI.md#CreateNodePool) | **Post** /api/hyperfleet/v1/clusters/{cluster_id}/nodepools | Create nodepool
[**GetClusterById**](DefaultAPI.md#GetClusterById) | **Get** /api/hyperfleet/v1/clusters/{cluster_id} | Get cluster by ID
[**GetClusterStatusById**](DefaultAPI.md#GetClusterStatusById) | **Get** /api/hyperfleet/v1/clusters/{cluster_id}/statuses/{status_id} | Get adapter status by ID
[**GetClusterStatuses**](DefaultAPI.md#GetClusterStatuses) | **Get** /api/hyperfleet/v1/clusters/{cluster_id}/statuses | List all adapter statuses for cluster
[**GetClusters**](DefaultAPI.md#GetClusters) | **Get** /api/hyperfleet/v1/clusters | List clusters
[**GetCompatibility**](DefaultAPI.md#GetCompatibility) | **Get** /api/hyperfleet/v1/compatibility | 
[**GetNodePoolById**](DefaultAPI.md#GetNodePoolById) | **Get** /api/hyperfleet/v1/nodepools/{nodepool_id} | Get nodepool by ID
[**GetNodePoolStatusById**](DefaultAPI.md#GetNodePoolStatusById) | **Get** /api/hyperfleet/v1/clusters/{cluster_id}/nodepools/{nodepool_id}/statuses/{status_id} | Get adapter status by ID
[**GetNodePools**](DefaultAPI.md#GetNodePools) | **Get** /api/hyperfleet/v1/nodepools | List all nodepools for cluster
[**GetNodePoolsByClusterId**](DefaultAPI.md#GetNodePoolsByClusterId) | **Get** /api/hyperfleet/v1/clusters/{cluster_id}/nodepools | List all nodepools for cluster
[**GetNodePoolsStatuses**](DefaultAPI.md#GetNodePoolsStatuses) | **Get** /api/hyperfleet/v1/clusters/{cluster_id}/nodepools/{nodepool_id}/statuses | List all adapter statuses for nodepools
[**PostCluster**](DefaultAPI.md#PostCluster) | **Post** /api/hyperfleet/v1/clusters | Create cluster
[**PostClusterStatuses**](DefaultAPI.md#PostClusterStatuses) | **Post** /api/hyperfleet/v1/clusters/{cluster_id}/statuses | Create adapter status
[**PostNodePoolStatuses**](DefaultAPI.md#PostNodePoolStatuses) | **Post** /api/hyperfleet/v1/clusters/{cluster_id}/nodepools/{nodepool_id}/statuses | Create adapter status



## CreateNodePool

> NodePoolCreateResponse CreateNodePool(ctx, clusterId).NodePoolCreateRequest(nodePoolCreateRequest).Execute()

Create nodepool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterId := "clusterId_example" // string | Cluster ID
	nodePoolCreateRequest := *openapiclient.NewNodePoolCreateRequest(map[string]string{"key": "Inner_example"}, time.Now(), time.Now(), "Name_example", map[string]interface{}{"key": interface{}(123)}) // NodePoolCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateNodePool(context.Background(), clusterId).NodePoolCreateRequest(nodePoolCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNodePool`: NodePoolCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodePoolCreateRequest** | [**NodePoolCreateRequest**](NodePoolCreateRequest.md) |  | 

### Return type

[**NodePoolCreateResponse**](NodePoolCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterById

> GetClusterById200Response GetClusterById(ctx, clusterId).Execute()

Get cluster by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterId := "clusterId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetClusterById(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetClusterById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterById`: GetClusterById200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetClusterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetClusterById200Response**](GetClusterById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterStatusById

> AdapterStatus GetClusterStatusById(ctx, clusterId, statusId).Execute()

Get adapter status by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterId := "clusterId_example" // string | Cluster ID
	statusId := "statusId_example" // string | Status ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetClusterStatusById(context.Background(), clusterId, statusId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetClusterStatusById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterStatusById`: AdapterStatus
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetClusterStatusById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**statusId** | **string** | Status ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterStatusByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AdapterStatus**](AdapterStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterStatuses

> GetNodePoolsStatuses200Response GetClusterStatuses(ctx, clusterId).Search(search).Page(page).PageSize(pageSize).OrderBy(orderBy).Order(order).Execute()

List all adapter statuses for cluster



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterId := "clusterId_example" // string | Cluster ID
	search := "search_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	orderBy := "orderBy_example" // string |  (optional) (default to "created_at")
	order := openapiclient.OrderDirection("asc") // OrderDirection |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetClusterStatuses(context.Background(), clusterId).Search(search).Page(page).PageSize(pageSize).OrderBy(orderBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetClusterStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterStatuses`: GetNodePoolsStatuses200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetClusterStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **orderBy** | **string** |  | [default to &quot;created_at&quot;]
 **order** | [**OrderDirection**](OrderDirection.md) |  | 

### Return type

[**GetNodePoolsStatuses200Response**](GetNodePoolsStatuses200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusters

> ClusterList GetClusters(ctx).Search(search).Page(page).PageSize(pageSize).OrderBy(orderBy).Order(order).Execute()

List clusters

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	search := "search_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	orderBy := "orderBy_example" // string |  (optional) (default to "created_at")
	order := openapiclient.OrderDirection("asc") // OrderDirection |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetClusters(context.Background()).Search(search).Page(page).PageSize(pageSize).OrderBy(orderBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusters`: ClusterList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **orderBy** | **string** |  | [default to &quot;created_at&quot;]
 **order** | [**OrderDirection**](OrderDirection.md) |  | 

### Return type

[**ClusterList**](ClusterList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompatibility

> string GetCompatibility(ctx).Search(search).Page(page).PageSize(pageSize).OrderBy(orderBy).Order(order).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	search := "search_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	orderBy := "orderBy_example" // string |  (optional) (default to "created_at")
	order := openapiclient.OrderDirection("asc") // OrderDirection |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCompatibility(context.Background()).Search(search).Page(page).PageSize(pageSize).OrderBy(orderBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCompatibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompatibility`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCompatibility`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompatibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **orderBy** | **string** |  | [default to &quot;created_at&quot;]
 **order** | [**OrderDirection**](OrderDirection.md) |  | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodePoolById

> NodePool GetNodePoolById(ctx, nodepoolId).Execute()

Get nodepool by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	nodepoolId := "nodepoolId_example" // string | NodePool ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetNodePoolById(context.Background(), nodepoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetNodePoolById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodePoolById`: NodePool
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetNodePoolById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodepoolId** | **string** | NodePool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodePoolByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NodePool**](NodePool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodePoolStatusById

> AdapterStatus GetNodePoolStatusById(ctx, clusterId, nodepoolId, statusId).Execute()

Get adapter status by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterId := "clusterId_example" // string | Cluster ID
	nodepoolId := "nodepoolId_example" // string | 
	statusId := "statusId_example" // string | Status ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetNodePoolStatusById(context.Background(), clusterId, nodepoolId, statusId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetNodePoolStatusById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodePoolStatusById`: AdapterStatus
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetNodePoolStatusById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**nodepoolId** | **string** |  | 
**statusId** | **string** | Status ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodePoolStatusByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AdapterStatus**](AdapterStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodePools

> GetNodePoolsByClusterId200Response GetNodePools(ctx).Search(search).Page(page).PageSize(pageSize).OrderBy(orderBy).Order(order).Execute()

List all nodepools for cluster



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	search := "search_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	orderBy := "orderBy_example" // string |  (optional) (default to "created_at")
	order := openapiclient.OrderDirection("asc") // OrderDirection |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetNodePools(context.Background()).Search(search).Page(page).PageSize(pageSize).OrderBy(orderBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetNodePools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodePools`: GetNodePoolsByClusterId200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetNodePools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNodePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **orderBy** | **string** |  | [default to &quot;created_at&quot;]
 **order** | [**OrderDirection**](OrderDirection.md) |  | 

### Return type

[**GetNodePoolsByClusterId200Response**](GetNodePoolsByClusterId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodePoolsByClusterId

> GetNodePoolsByClusterId200Response GetNodePoolsByClusterId(ctx, clusterId).Search(search).Page(page).PageSize(pageSize).OrderBy(orderBy).Order(order).Execute()

List all nodepools for cluster



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterId := "clusterId_example" // string | Cluster ID
	search := "search_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	orderBy := "orderBy_example" // string |  (optional) (default to "created_at")
	order := openapiclient.OrderDirection("asc") // OrderDirection |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetNodePoolsByClusterId(context.Background(), clusterId).Search(search).Page(page).PageSize(pageSize).OrderBy(orderBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetNodePoolsByClusterId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodePoolsByClusterId`: GetNodePoolsByClusterId200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetNodePoolsByClusterId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodePoolsByClusterIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **orderBy** | **string** |  | [default to &quot;created_at&quot;]
 **order** | [**OrderDirection**](OrderDirection.md) |  | 

### Return type

[**GetNodePoolsByClusterId200Response**](GetNodePoolsByClusterId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodePoolsStatuses

> GetNodePoolsStatuses200Response GetNodePoolsStatuses(ctx, clusterId, nodepoolId).Search(search).Page(page).PageSize(pageSize).OrderBy(orderBy).Order(order).Execute()

List all adapter statuses for nodepools



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterId := "clusterId_example" // string | Cluster ID
	nodepoolId := "nodepoolId_example" // string | 
	search := "search_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	orderBy := "orderBy_example" // string |  (optional) (default to "created_at")
	order := openapiclient.OrderDirection("asc") // OrderDirection |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetNodePoolsStatuses(context.Background(), clusterId, nodepoolId).Search(search).Page(page).PageSize(pageSize).OrderBy(orderBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetNodePoolsStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodePoolsStatuses`: GetNodePoolsStatuses200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetNodePoolsStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**nodepoolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodePoolsStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **search** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **orderBy** | **string** |  | [default to &quot;created_at&quot;]
 **order** | [**OrderDirection**](OrderDirection.md) |  | 

### Return type

[**GetNodePoolsStatuses200Response**](GetNodePoolsStatuses200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCluster

> Error PostCluster(ctx).ClusterCreateRequest(clusterCreateRequest).Execute()

Create cluster



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterCreateRequest := *openapiclient.NewClusterCreateRequest("Kind_example", "Name_example", map[string]interface{}{"key": interface{}(123)}, map[string]string{"key": "Inner_example"}, time.Now(), time.Now()) // ClusterCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostCluster(context.Background()).ClusterCreateRequest(clusterCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCluster`: Error
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterCreateRequest** | [**ClusterCreateRequest**](ClusterCreateRequest.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClusterStatuses

> AdapterStatus PostClusterStatuses(ctx, clusterId).AdapterStatusCreateRequest(adapterStatusCreateRequest).Execute()

Create adapter status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterId := "clusterId_example" // string | Cluster ID
	adapterStatusCreateRequest := *openapiclient.NewAdapterStatusCreateRequest("AdapterName_example", int32(123), []openapiclient.Condition{*openapiclient.NewCondition("Type_example", "Status_example", int32(123), time.Now(), time.Now())}) // AdapterStatusCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostClusterStatuses(context.Background(), clusterId).AdapterStatusCreateRequest(adapterStatusCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostClusterStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClusterStatuses`: AdapterStatus
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostClusterStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClusterStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adapterStatusCreateRequest** | [**AdapterStatusCreateRequest**](AdapterStatusCreateRequest.md) |  | 

### Return type

[**AdapterStatus**](AdapterStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNodePoolStatuses

> AdapterStatus PostNodePoolStatuses(ctx, clusterId, nodepoolId).AdapterStatusCreateRequest(adapterStatusCreateRequest).Execute()

Create adapter status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterId := "clusterId_example" // string | Cluster ID
	nodepoolId := "nodepoolId_example" // string | 
	adapterStatusCreateRequest := *openapiclient.NewAdapterStatusCreateRequest("AdapterName_example", int32(123), []openapiclient.Condition{*openapiclient.NewCondition("Type_example", "Status_example", int32(123), time.Now(), time.Now())}) // AdapterStatusCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostNodePoolStatuses(context.Background(), clusterId, nodepoolId).AdapterStatusCreateRequest(adapterStatusCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostNodePoolStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNodePoolStatuses`: AdapterStatus
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostNodePoolStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**nodepoolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNodePoolStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **adapterStatusCreateRequest** | [**AdapterStatusCreateRequest**](AdapterStatusCreateRequest.md) |  | 

### Return type

[**AdapterStatus**](AdapterStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

