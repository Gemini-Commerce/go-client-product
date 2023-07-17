# \ProductApi

All URIs are relative to *https://product.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMediaGalleryEntry**](ProductApi.md#AddMediaGalleryEntry) | **Post** /product.Product/AddMediaGalleryEntry | Add Media Gallery Entry
[**BulkCreateAttribute**](ProductApi.md#BulkCreateAttribute) | **Post** /product.Product/BulkCreateAttribute | Bulk Create Attribute
[**BulkUpdate**](ProductApi.md#BulkUpdate) | **Post** /product.Product/BulkUpdate | Bulk Update Products
[**CreateAttributeOptions**](ProductApi.md#CreateAttributeOptions) | **Post** /product.Product/CreateAttributeOptions | Create Attribute Options
[**CreateEntity**](ProductApi.md#CreateEntity) | **Post** /product.Product/CreateEntity | Create Entity
[**CreateOptionsList**](ProductApi.md#CreateOptionsList) | **Post** /product.Product/CreateOptionsList | Create Options List
[**CreateProduct**](ProductApi.md#CreateProduct) | **Post** /product.Product/CreateProduct | Create Product
[**DeleteProduct**](ProductApi.md#DeleteProduct) | **Post** /product.Product/Delete | Delete Product
[**GetAttributeOption**](ProductApi.md#GetAttributeOption) | **Post** /product.Product/GetAttributeOption | Get Attribute Option
[**GetAttributeOptionByCode**](ProductApi.md#GetAttributeOptionByCode) | **Post** /product.Product/GetAttributeOptionByCode | Get Attribute Option By Code
[**GetAttributeOptions**](ProductApi.md#GetAttributeOptions) | **Post** /product.Product/GetAttributeOptions | Get Attribute Options
[**GetEntity**](ProductApi.md#GetEntity) | **Post** /product.Product/GetEntity | Get Entity
[**GetOptionsList**](ProductApi.md#GetOptionsList) | **Post** /product.Product/GetOptionsList | Get Options List
[**GetProduct**](ProductApi.md#GetProduct) | **Post** /product.Product/GetProduct | Get Product
[**GetProductByCode**](ProductApi.md#GetProductByCode) | **Post** /product.Product/GetProductByCode | Get Product By Code
[**GetProductByUrlKey**](ProductApi.md#GetProductByUrlKey) | **Post** /product.Product/GetProductByUrlKey | Get Product By Url Key
[**ListAttributeOptions**](ProductApi.md#ListAttributeOptions) | **Post** /product.Product/ListAttributeOptions | List Attribute Options
[**ListEntities**](ProductApi.md#ListEntities) | **Post** /product.Product/ListEntities | List Entities
[**ListOptionsLists**](ProductApi.md#ListOptionsLists) | **Post** /product.Product/ListOptionsLists | List Options Lists
[**ListProducts**](ProductApi.md#ListProducts) | **Post** /product.Product/ListProducts | List Products
[**ListProductsByIds**](ProductApi.md#ListProductsByIds) | **Post** /product.Product/ListProductsByIds | List Products By Ids
[**ListProductsBySku**](ProductApi.md#ListProductsBySku) | **Post** /product.Product/ListProductsBySku | List Products By Sku
[**ListVariantsBySku**](ProductApi.md#ListVariantsBySku) | **Post** /product.Product/ListVariantsBySku | List Product Variants By Sku
[**RemoveMediaGalleryEntry**](ProductApi.md#RemoveMediaGalleryEntry) | **Post** /product.Product/RemoveMediaGalleryEntry | Remove Media Gallery Entry
[**UpdateAttributeOptions**](ProductApi.md#UpdateAttributeOptions) | **Post** /product.Product/UpdateAttributeOptions | Update Attribute Options
[**UpdateMediaGalleryEntry**](ProductApi.md#UpdateMediaGalleryEntry) | **Post** /product.Product/UpdateMediaGalleryEntry | Update Media Gallery Entry
[**UpdateOptionsList**](ProductApi.md#UpdateOptionsList) | **Post** /product.Product/UpdateOptionsList | Update Options List
[**UpdateProduct**](ProductApi.md#UpdateProduct) | **Post** /product.Product/UpdateProduct | Update Product



## AddMediaGalleryEntry

> ProductAddMediaGalleryEntryResponse AddMediaGalleryEntry(ctx).Body(body).Execute()

Add Media Gallery Entry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductAddMediaGalleryEntryRequest() // ProductAddMediaGalleryEntryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.AddMediaGalleryEntry(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.AddMediaGalleryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMediaGalleryEntry`: ProductAddMediaGalleryEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.AddMediaGalleryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMediaGalleryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductAddMediaGalleryEntryRequest**](ProductAddMediaGalleryEntryRequest.md) |  | 

### Return type

[**ProductAddMediaGalleryEntryResponse**](ProductAddMediaGalleryEntryResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCreateAttribute

> EntitymanagerBulkCreateAttributeResponse BulkCreateAttribute(ctx).Body(body).Execute()

Bulk Create Attribute

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerBulkCreateAttributeRequest() // EntitymanagerBulkCreateAttributeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.BulkCreateAttribute(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.BulkCreateAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateAttribute`: EntitymanagerBulkCreateAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.BulkCreateAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerBulkCreateAttributeRequest**](EntitymanagerBulkCreateAttributeRequest.md) |  | 

### Return type

[**EntitymanagerBulkCreateAttributeResponse**](EntitymanagerBulkCreateAttributeResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdate

> ProductBulkUpdateResponse BulkUpdate(ctx).Body(body).Execute()

Bulk Update Products

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductBulkUpdateRequest() // ProductBulkUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.BulkUpdate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.BulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdate`: ProductBulkUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.BulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductBulkUpdateRequest**](ProductBulkUpdateRequest.md) |  | 

### Return type

[**ProductBulkUpdateResponse**](ProductBulkUpdateResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAttributeOptions

> EntitymanagerCreateAttributeOptionsResponse CreateAttributeOptions(ctx).Body(body).Execute()

Create Attribute Options

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerCreateAttributeOptionsRequest() // EntitymanagerCreateAttributeOptionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.CreateAttributeOptions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.CreateAttributeOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAttributeOptions`: EntitymanagerCreateAttributeOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.CreateAttributeOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttributeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerCreateAttributeOptionsRequest**](EntitymanagerCreateAttributeOptionsRequest.md) |  | 

### Return type

[**EntitymanagerCreateAttributeOptionsResponse**](EntitymanagerCreateAttributeOptionsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEntity

> EntitymanagerCreateEntityResponse CreateEntity(ctx).Body(body).Execute()

Create Entity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerEntity() // EntitymanagerEntity | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.CreateEntity(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.CreateEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEntity`: EntitymanagerCreateEntityResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.CreateEntity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerEntity**](EntitymanagerEntity.md) |  | 

### Return type

[**EntitymanagerCreateEntityResponse**](EntitymanagerCreateEntityResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOptionsList

> EntitymanagerCreateOptionsListResponse CreateOptionsList(ctx).Body(body).Execute()

Create Options List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerCreateOptionsListRequest() // EntitymanagerCreateOptionsListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.CreateOptionsList(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.CreateOptionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOptionsList`: EntitymanagerCreateOptionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.CreateOptionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOptionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerCreateOptionsListRequest**](EntitymanagerCreateOptionsListRequest.md) |  | 

### Return type

[**EntitymanagerCreateOptionsListResponse**](EntitymanagerCreateOptionsListResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProduct

> ProductCreateProductResponse CreateProduct(ctx).Body(body).Execute()

Create Product



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductCreateProductRequest() // ProductCreateProductRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.CreateProduct(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.CreateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProduct`: ProductCreateProductResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.CreateProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductCreateProductRequest**](ProductCreateProductRequest.md) |  | 

### Return type

[**ProductCreateProductResponse**](ProductCreateProductResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProduct

> ProductDeleteResponse DeleteProduct(ctx).Body(body).Execute()

Delete Product



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductDeleteRequest() // ProductDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.DeleteProduct(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.DeleteProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProduct`: ProductDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.DeleteProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductDeleteRequest**](ProductDeleteRequest.md) |  | 

### Return type

[**ProductDeleteResponse**](ProductDeleteResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributeOption

> EntitymanagerGetAttributeOptionResponse GetAttributeOption(ctx).Body(body).Execute()

Get Attribute Option

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerGetAttributeOptionRequest() // EntitymanagerGetAttributeOptionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.GetAttributeOption(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetAttributeOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttributeOption`: EntitymanagerGetAttributeOptionResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetAttributeOption`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerGetAttributeOptionRequest**](EntitymanagerGetAttributeOptionRequest.md) |  | 

### Return type

[**EntitymanagerGetAttributeOptionResponse**](EntitymanagerGetAttributeOptionResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributeOptionByCode

> EntitymanagerGetAttributeOptionByCodeResponse GetAttributeOptionByCode(ctx).Body(body).Execute()

Get Attribute Option By Code

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerGetAttributeOptionByCodeRequest() // EntitymanagerGetAttributeOptionByCodeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.GetAttributeOptionByCode(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetAttributeOptionByCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttributeOptionByCode`: EntitymanagerGetAttributeOptionByCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetAttributeOptionByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeOptionByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerGetAttributeOptionByCodeRequest**](EntitymanagerGetAttributeOptionByCodeRequest.md) |  | 

### Return type

[**EntitymanagerGetAttributeOptionByCodeResponse**](EntitymanagerGetAttributeOptionByCodeResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributeOptions

> EntitymanagerGetAttributeOptionsResponse GetAttributeOptions(ctx).Body(body).Execute()

Get Attribute Options

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerGetAttributeOptionsRequest() // EntitymanagerGetAttributeOptionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.GetAttributeOptions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetAttributeOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttributeOptions`: EntitymanagerGetAttributeOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetAttributeOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerGetAttributeOptionsRequest**](EntitymanagerGetAttributeOptionsRequest.md) |  | 

### Return type

[**EntitymanagerGetAttributeOptionsResponse**](EntitymanagerGetAttributeOptionsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntity

> EntitymanagerEntity GetEntity(ctx).Body(body).Execute()

Get Entity

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerEntityRequest() // EntitymanagerEntityRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.GetEntity(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntity`: EntitymanagerEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetEntity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerEntityRequest**](EntitymanagerEntityRequest.md) |  | 

### Return type

[**EntitymanagerEntity**](EntitymanagerEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOptionsList

> EntitymanagerGetOptionsListResponse GetOptionsList(ctx).Body(body).Execute()

Get Options List

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerGetOptionsListRequest() // EntitymanagerGetOptionsListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.GetOptionsList(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetOptionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOptionsList`: EntitymanagerGetOptionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetOptionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerGetOptionsListRequest**](EntitymanagerGetOptionsListRequest.md) |  | 

### Return type

[**EntitymanagerGetOptionsListResponse**](EntitymanagerGetOptionsListResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProduct

> ProductGetProductResponse GetProduct(ctx).Body(body).Execute()

Get Product



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductGetProductRequest() // ProductGetProductRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.GetProduct(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProduct`: ProductGetProductResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductGetProductRequest**](ProductGetProductRequest.md) |  | 

### Return type

[**ProductGetProductResponse**](ProductGetProductResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductByCode

> ProductGetProductByCodeResponse GetProductByCode(ctx).Body(body).Execute()

Get Product By Code



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductGetProductByCodeRequest() // ProductGetProductByCodeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.GetProductByCode(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetProductByCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductByCode`: ProductGetProductByCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetProductByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductGetProductByCodeRequest**](ProductGetProductByCodeRequest.md) |  | 

### Return type

[**ProductGetProductByCodeResponse**](ProductGetProductByCodeResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductByUrlKey

> ProductGetProductByUrlKeyResponse GetProductByUrlKey(ctx).Body(body).Execute()

Get Product By Url Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductGetProductByUrlKeyRequest() // ProductGetProductByUrlKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.GetProductByUrlKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetProductByUrlKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductByUrlKey`: ProductGetProductByUrlKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetProductByUrlKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductByUrlKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductGetProductByUrlKeyRequest**](ProductGetProductByUrlKeyRequest.md) |  | 

### Return type

[**ProductGetProductByUrlKeyResponse**](ProductGetProductByUrlKeyResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAttributeOptions

> EntitymanagerListAttributeOptionsResponse ListAttributeOptions(ctx).Body(body).Execute()

List Attribute Options

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerListAttributeOptionsRequest() // EntitymanagerListAttributeOptionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ListAttributeOptions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ListAttributeOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAttributeOptions`: EntitymanagerListAttributeOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ListAttributeOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAttributeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerListAttributeOptionsRequest**](EntitymanagerListAttributeOptionsRequest.md) |  | 

### Return type

[**EntitymanagerListAttributeOptionsResponse**](EntitymanagerListAttributeOptionsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntities

> EntitymanagerListEntitiesResponse ListEntities(ctx).Body(body).Execute()

List Entities

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerListEntitiesRequest() // EntitymanagerListEntitiesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ListEntities(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ListEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntities`: EntitymanagerListEntitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ListEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerListEntitiesRequest**](EntitymanagerListEntitiesRequest.md) |  | 

### Return type

[**EntitymanagerListEntitiesResponse**](EntitymanagerListEntitiesResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOptionsLists

> EntitymanagerListOptionsListsResponse ListOptionsLists(ctx).Body(body).Execute()

List Options Lists



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerListOptionsListsRequest() // EntitymanagerListOptionsListsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ListOptionsLists(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ListOptionsLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOptionsLists`: EntitymanagerListOptionsListsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ListOptionsLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOptionsListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerListOptionsListsRequest**](EntitymanagerListOptionsListsRequest.md) |  | 

### Return type

[**EntitymanagerListOptionsListsResponse**](EntitymanagerListOptionsListsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProducts

> ProductListProductsResponse ListProducts(ctx).Body(body).Execute()

List Products



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductListProductsRequest() // ProductListProductsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ListProducts(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ListProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProducts`: ProductListProductsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ListProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductListProductsRequest**](ProductListProductsRequest.md) |  | 

### Return type

[**ProductListProductsResponse**](ProductListProductsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductsByIds

> ProductListProductsByIdsResponse ListProductsByIds(ctx).Body(body).Execute()

List Products By Ids



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductListProductsByIdsRequest() // ProductListProductsByIdsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ListProductsByIds(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ListProductsByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProductsByIds`: ProductListProductsByIdsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ListProductsByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductsByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductListProductsByIdsRequest**](ProductListProductsByIdsRequest.md) |  | 

### Return type

[**ProductListProductsByIdsResponse**](ProductListProductsByIdsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductsBySku

> ProductListProductsBySkuResponse ListProductsBySku(ctx).Body(body).Execute()

List Products By Sku

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductListProductsBySkuRequest() // ProductListProductsBySkuRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ListProductsBySku(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ListProductsBySku``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProductsBySku`: ProductListProductsBySkuResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ListProductsBySku`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductsBySkuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductListProductsBySkuRequest**](ProductListProductsBySkuRequest.md) |  | 

### Return type

[**ProductListProductsBySkuResponse**](ProductListProductsBySkuResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVariantsBySku

> ProductListVariantsBySkuResponse ListVariantsBySku(ctx).Body(body).Execute()

List Product Variants By Sku

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductListVariantsBySkuRequest() // ProductListVariantsBySkuRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ListVariantsBySku(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ListVariantsBySku``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVariantsBySku`: ProductListVariantsBySkuResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ListVariantsBySku`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVariantsBySkuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductListVariantsBySkuRequest**](ProductListVariantsBySkuRequest.md) |  | 

### Return type

[**ProductListVariantsBySkuResponse**](ProductListVariantsBySkuResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMediaGalleryEntry

> map[string]interface{} RemoveMediaGalleryEntry(ctx).Body(body).Execute()

Remove Media Gallery Entry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductRemoveMediaGalleryEntryRequest() // ProductRemoveMediaGalleryEntryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.RemoveMediaGalleryEntry(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.RemoveMediaGalleryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveMediaGalleryEntry`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.RemoveMediaGalleryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMediaGalleryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductRemoveMediaGalleryEntryRequest**](ProductRemoveMediaGalleryEntryRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttributeOptions

> EntitymanagerUpdateAttributeOptionsResponse UpdateAttributeOptions(ctx).Body(body).Execute()

Update Attribute Options

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerUpdateAttributeOptionsRequest() // EntitymanagerUpdateAttributeOptionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.UpdateAttributeOptions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.UpdateAttributeOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAttributeOptions`: EntitymanagerUpdateAttributeOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.UpdateAttributeOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttributeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerUpdateAttributeOptionsRequest**](EntitymanagerUpdateAttributeOptionsRequest.md) |  | 

### Return type

[**EntitymanagerUpdateAttributeOptionsResponse**](EntitymanagerUpdateAttributeOptionsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMediaGalleryEntry

> map[string]interface{} UpdateMediaGalleryEntry(ctx).Body(body).Execute()

Update Media Gallery Entry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductUpdateMediaGalleryEntryRequest() // ProductUpdateMediaGalleryEntryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.UpdateMediaGalleryEntry(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.UpdateMediaGalleryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMediaGalleryEntry`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.UpdateMediaGalleryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMediaGalleryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductUpdateMediaGalleryEntryRequest**](ProductUpdateMediaGalleryEntryRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOptionsList

> EntitymanagerUpdateOptionsListResponse UpdateOptionsList(ctx).Body(body).Execute()

Update Options List

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerUpdateOptionsListRequest() // EntitymanagerUpdateOptionsListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.UpdateOptionsList(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.UpdateOptionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOptionsList`: EntitymanagerUpdateOptionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.UpdateOptionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOptionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerUpdateOptionsListRequest**](EntitymanagerUpdateOptionsListRequest.md) |  | 

### Return type

[**EntitymanagerUpdateOptionsListResponse**](EntitymanagerUpdateOptionsListResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProduct

> ProductUpdateProductResponse UpdateProduct(ctx).Body(body).Execute()

Update Product



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductUpdateProductRequest() // ProductUpdateProductRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.UpdateProduct(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.UpdateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProduct`: ProductUpdateProductResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.UpdateProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductUpdateProductRequest**](ProductUpdateProductRequest.md) |  | 

### Return type

[**ProductUpdateProductResponse**](ProductUpdateProductResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

