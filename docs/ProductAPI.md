# GeminiCommerce\Product\ProductAPI

All URIs are relative to *https://product.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMediaGalleryEntry**](ProductAPI.md#AddMediaGalleryEntry) | **Post** /product.Product/AddMediaGalleryEntry | Add Media Gallery Entry
[**BulkAddAssetsEntries**](ProductAPI.md#BulkAddAssetsEntries) | **Post** /product.Product/BulkAddAssetsEntries | Bulk Add Assets Entries
[**BulkDeleteProducts**](ProductAPI.md#BulkDeleteProducts) | **Post** /product.Product/BulkDeleteProducts | Bulk Delete Products
[**BulkEnhanceProductDataWithAI**](ProductAPI.md#BulkEnhanceProductDataWithAI) | **Post** /product.Product/BulkEnhanceProductDataWithAI | Bulk Enhance Product Data With AI
[**BulkRemoveAssetsEntries**](ProductAPI.md#BulkRemoveAssetsEntries) | **Post** /product.Product/BulkRemoveAssetsEntries | Bulk Remove Assets Entries
[**BulkUpdateAssetsEntries**](ProductAPI.md#BulkUpdateAssetsEntries) | **Post** /product.Product/BulkUpdateAssetsEntries | Bulk Update Assets Entries
[**BulkUpdateV2**](ProductAPI.md#BulkUpdateV2) | **Post** /product.Product/BulkUpdateV2 | Bulk Update Products
[**CreateAttributeOptions**](ProductAPI.md#CreateAttributeOptions) | **Post** /product.Product/CreateAttributeOptions | Create Attribute Options
[**CreateEntity**](ProductAPI.md#CreateEntity) | **Post** /product.Product/CreateEntity | Create Entity
[**CreateOptionsList**](ProductAPI.md#CreateOptionsList) | **Post** /product.Product/CreateOptionsList | Create Options List
[**CreateProductWithAI**](ProductAPI.md#CreateProductWithAI) | **Post** /product.Product/CreateProductWithAI | Create Product With AI
[**GetAttributeOption**](ProductAPI.md#GetAttributeOption) | **Post** /product.Product/GetAttributeOption | Get Attribute Option
[**GetAttributeOptions**](ProductAPI.md#GetAttributeOptions) | **Post** /product.Product/GetAttributeOptions | Get Attribute Options
[**GetEntity**](ProductAPI.md#GetEntity) | **Post** /product.Product/GetEntity | Get Entity Details
[**GetOptionsList**](ProductAPI.md#GetOptionsList) | **Post** /product.Product/GetOptionsList | Get Options List
[**GetProduct**](ProductAPI.md#GetProduct) | **Post** /product.Product/GetProduct | Get Product
[**GetProductByCode**](ProductAPI.md#GetProductByCode) | **Post** /product.Product/GetProductByCode | Get Product By Code
[**GetProductByUrlKey**](ProductAPI.md#GetProductByUrlKey) | **Post** /product.Product/GetProductByUrlKey | Get Product By Url Key
[**GetProductDataInReview**](ProductAPI.md#GetProductDataInReview) | **Post** /product.Product/GetProductDataInReview | Get Product Data In Review
[**ListAttributeOptions**](ProductAPI.md#ListAttributeOptions) | **Post** /product.Product/ListAttributeOptions | List Attribute Options
[**ListEntities**](ProductAPI.md#ListEntities) | **Post** /product.Product/ListEntities | List Entities
[**ListOptionsLists**](ProductAPI.md#ListOptionsLists) | **Post** /product.Product/ListOptionsLists | List Options Lists
[**ListProducts**](ProductAPI.md#ListProducts) | **Post** /product.Product/ListProducts | List Products
[**ListProductsByIds**](ProductAPI.md#ListProductsByIds) | **Post** /product.Product/ListProductsByIds | List Products By Ids
[**ListProductsBySku**](ProductAPI.md#ListProductsBySku) | **Post** /product.Product/ListProductsBySku | List Products By Sku
[**ListVariantsBySku**](ProductAPI.md#ListVariantsBySku) | **Post** /product.Product/ListVariantsBySku | List Product Variants By Sku
[**ProductBulkCreateAttribute**](ProductAPI.md#ProductBulkCreateAttribute) | **Post** /product.Product/BulkCreateAttribute | 
[**ProductBulkUpdate**](ProductAPI.md#ProductBulkUpdate) | **Post** /product.Product/BulkUpdate | 
[**ProductCreateAttributeGroup**](ProductAPI.md#ProductCreateAttributeGroup) | **Post** /product.Product/CreateAttributeGroup | 
[**ProductCreateProduct**](ProductAPI.md#ProductCreateProduct) | **Post** /product.Product/CreateProduct | 
[**ProductCreateProductV2**](ProductAPI.md#ProductCreateProductV2) | **Post** /product.Product/CreateProductV2 | 
[**ProductDelete**](ProductAPI.md#ProductDelete) | **Post** /product.Product/Delete | 
[**ProductDeleteAttribute**](ProductAPI.md#ProductDeleteAttribute) | **Post** /product.Product/DeleteAttribute | 
[**ProductDeleteProduct**](ProductAPI.md#ProductDeleteProduct) | **Post** /product.Product/DeleteProduct | 
[**ProductGetAttributeGroup**](ProductAPI.md#ProductGetAttributeGroup) | **Post** /product.Product/GetAttributeGroup | 
[**ProductListAttributeGroups**](ProductAPI.md#ProductListAttributeGroups) | **Post** /product.Product/ListAttributeGroups | 
[**ProductUpdateAttribute**](ProductAPI.md#ProductUpdateAttribute) | **Post** /product.Product/UpdateAttribute | 
[**ProductUpdateAttributeGroup**](ProductAPI.md#ProductUpdateAttributeGroup) | **Post** /product.Product/UpdateAttributeGroup | 
[**ProductUpdateProduct**](ProductAPI.md#ProductUpdateProduct) | **Post** /product.Product/UpdateProduct | 
[**ProductUpdateProductV2**](ProductAPI.md#ProductUpdateProductV2) | **Post** /product.Product/UpdateProductV2 | 
[**RemoveMediaGalleryEntry**](ProductAPI.md#RemoveMediaGalleryEntry) | **Post** /product.Product/RemoveMediaGalleryEntry | Remove Media Gallery Entry
[**UpdateAttributeOptions**](ProductAPI.md#UpdateAttributeOptions) | **Post** /product.Product/UpdateAttributeOptions | Update Attribute Options
[**UpdateDataToBeReviewed**](ProductAPI.md#UpdateDataToBeReviewed) | **Post** /product.Product/GetEnhanceProductDataWithAIStatus | Get Enhance Product Data With AI Status
[**UpdateDataToBeReviewed_0**](ProductAPI.md#UpdateDataToBeReviewed_0) | **Post** /product.Product/UpdateDataToBeReviewed | Update Data To Be Reviewed
[**UpdateMediaGalleryEntry**](ProductAPI.md#UpdateMediaGalleryEntry) | **Post** /product.Product/UpdateMediaGalleryEntry | Update Media Gallery Entry
[**UpdateOptionsList**](ProductAPI.md#UpdateOptionsList) | **Post** /product.Product/UpdateOptionsList | Update Options List
[**UpdateProductWithAI**](ProductAPI.md#UpdateProductWithAI) | **Post** /product.Product/UpdateProductWithAI | Update Product With AI



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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductAddMediaGalleryEntryRequest() // ProductAddMediaGalleryEntryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.AddMediaGalleryEntry(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.AddMediaGalleryEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMediaGalleryEntry`: ProductAddMediaGalleryEntryResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.AddMediaGalleryEntry`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkAddAssetsEntries

> ProductBulkAddAssetsEntriesResponse BulkAddAssetsEntries(ctx).Body(body).Execute()

Bulk Add Assets Entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductBulkAddAssetsEntriesRequest() // ProductBulkAddAssetsEntriesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.BulkAddAssetsEntries(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.BulkAddAssetsEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkAddAssetsEntries`: ProductBulkAddAssetsEntriesResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.BulkAddAssetsEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkAddAssetsEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductBulkAddAssetsEntriesRequest**](ProductBulkAddAssetsEntriesRequest.md) |  | 

### Return type

[**ProductBulkAddAssetsEntriesResponse**](ProductBulkAddAssetsEntriesResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDeleteProducts

> map[string]interface{} BulkDeleteProducts(ctx).Body(body).Execute()

Bulk Delete Products



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductBulkDeleteProductsRequest() // ProductBulkDeleteProductsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.BulkDeleteProducts(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.BulkDeleteProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkDeleteProducts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.BulkDeleteProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductBulkDeleteProductsRequest**](ProductBulkDeleteProductsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkEnhanceProductDataWithAI

> map[string]interface{} BulkEnhanceProductDataWithAI(ctx).Body(body).Execute()

Bulk Enhance Product Data With AI



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductBulkEnhanceProductDataWithAIRequest() // ProductBulkEnhanceProductDataWithAIRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.BulkEnhanceProductDataWithAI(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.BulkEnhanceProductDataWithAI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkEnhanceProductDataWithAI`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.BulkEnhanceProductDataWithAI`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkEnhanceProductDataWithAIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductBulkEnhanceProductDataWithAIRequest**](ProductBulkEnhanceProductDataWithAIRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkRemoveAssetsEntries

> map[string]interface{} BulkRemoveAssetsEntries(ctx).Body(body).Execute()

Bulk Remove Assets Entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductBulkRemoveAssetsEntriesRequest() // ProductBulkRemoveAssetsEntriesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.BulkRemoveAssetsEntries(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.BulkRemoveAssetsEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkRemoveAssetsEntries`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.BulkRemoveAssetsEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkRemoveAssetsEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductBulkRemoveAssetsEntriesRequest**](ProductBulkRemoveAssetsEntriesRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateAssetsEntries

> ProductBulkUpdateAssetsEntriesResponse BulkUpdateAssetsEntries(ctx).Body(body).Execute()

Bulk Update Assets Entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductBulkUpdateAssetsEntriesRequest() // ProductBulkUpdateAssetsEntriesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.BulkUpdateAssetsEntries(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.BulkUpdateAssetsEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUpdateAssetsEntries`: ProductBulkUpdateAssetsEntriesResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.BulkUpdateAssetsEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateAssetsEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductBulkUpdateAssetsEntriesRequest**](ProductBulkUpdateAssetsEntriesRequest.md) |  | 

### Return type

[**ProductBulkUpdateAssetsEntriesResponse**](ProductBulkUpdateAssetsEntriesResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateV2

> ProductBulkUpdateResponse BulkUpdateV2(ctx).Body(body).Execute()

Bulk Update Products



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductBulkUpdateRequestV2() // ProductBulkUpdateRequestV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.BulkUpdateV2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.BulkUpdateV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUpdateV2`: ProductBulkUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.BulkUpdateV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductBulkUpdateRequestV2**](ProductBulkUpdateRequestV2.md) |  | 

### Return type

[**ProductBulkUpdateResponse**](ProductBulkUpdateResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerCreateAttributeOptionsRequest() // EntitymanagerCreateAttributeOptionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.CreateAttributeOptions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.CreateAttributeOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAttributeOptions`: EntitymanagerCreateAttributeOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.CreateAttributeOptions`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerEntity() // EntitymanagerEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.CreateEntity(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.CreateEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEntity`: EntitymanagerCreateEntityResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.CreateEntity`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerCreateOptionsListRequest() // EntitymanagerCreateOptionsListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.CreateOptionsList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.CreateOptionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOptionsList`: EntitymanagerCreateOptionsListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.CreateOptionsList`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProductWithAI

> ProductCreateProductWithAIResponse CreateProductWithAI(ctx).Body(body).Execute()

Create Product With AI



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductCreateProductWithAIRequest() // ProductCreateProductWithAIRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.CreateProductWithAI(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.CreateProductWithAI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProductWithAI`: ProductCreateProductWithAIResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.CreateProductWithAI`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductWithAIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductCreateProductWithAIRequest**](ProductCreateProductWithAIRequest.md) |  | 

### Return type

[**ProductCreateProductWithAIResponse**](ProductCreateProductWithAIResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerGetAttributeOptionRequest() // EntitymanagerGetAttributeOptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.GetAttributeOption(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetAttributeOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttributeOption`: EntitymanagerGetAttributeOptionResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetAttributeOption`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerGetAttributeOptionsRequest() // EntitymanagerGetAttributeOptionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.GetAttributeOptions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetAttributeOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttributeOptions`: EntitymanagerGetAttributeOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetAttributeOptions`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntity

> EntitymanagerEntity GetEntity(ctx).Body(body).Execute()

Get Entity Details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerEntityRequest() // EntitymanagerEntityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.GetEntity(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntity`: EntitymanagerEntity
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetEntity`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerGetOptionsListRequest() // EntitymanagerGetOptionsListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.GetOptionsList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetOptionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptionsList`: EntitymanagerGetOptionsListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetOptionsList`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductGetProductRequest() // ProductGetProductRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.GetProduct(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProduct`: ProductGetProductResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetProduct`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductGetProductByCodeRequest() // ProductGetProductByCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.GetProductByCode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetProductByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductByCode`: ProductGetProductByCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetProductByCode`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductGetProductByUrlKeyRequest() // ProductGetProductByUrlKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.GetProductByUrlKey(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetProductByUrlKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductByUrlKey`: ProductGetProductByUrlKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetProductByUrlKey`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductDataInReview

> ProductGetProductDataInReviewResponse GetProductDataInReview(ctx).Body(body).Execute()

Get Product Data In Review



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductGetProductDataInReviewRequest() // ProductGetProductDataInReviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.GetProductDataInReview(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetProductDataInReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductDataInReview`: ProductGetProductDataInReviewResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetProductDataInReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductDataInReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductGetProductDataInReviewRequest**](ProductGetProductDataInReviewRequest.md) |  | 

### Return type

[**ProductGetProductDataInReviewResponse**](ProductGetProductDataInReviewResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerListAttributeOptionsRequest() // EntitymanagerListAttributeOptionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ListAttributeOptions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ListAttributeOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAttributeOptions`: EntitymanagerListAttributeOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ListAttributeOptions`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerListEntitiesRequest() // EntitymanagerListEntitiesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ListEntities(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ListEntities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntities`: EntitymanagerListEntitiesResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ListEntities`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerListOptionsListsRequest() // EntitymanagerListOptionsListsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ListOptionsLists(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ListOptionsLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOptionsLists`: EntitymanagerListOptionsListsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ListOptionsLists`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductListProductsRequest() // ProductListProductsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ListProducts(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ListProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProducts`: ProductListProductsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ListProducts`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductListProductsByIdsRequest() // ProductListProductsByIdsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ListProductsByIds(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ListProductsByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductsByIds`: ProductListProductsByIdsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ListProductsByIds`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductListProductsBySkuRequest() // ProductListProductsBySkuRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ListProductsBySku(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ListProductsBySku``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductsBySku`: ProductListProductsBySkuResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ListProductsBySku`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductListVariantsBySkuRequest() // ProductListVariantsBySkuRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ListVariantsBySku(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ListVariantsBySku``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVariantsBySku`: ProductListVariantsBySkuResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ListVariantsBySku`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductBulkCreateAttribute

> EntitymanagerBulkCreateAttributeResponse ProductBulkCreateAttribute(ctx).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerBulkCreateAttributeRequest() // EntitymanagerBulkCreateAttributeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductBulkCreateAttribute(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductBulkCreateAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductBulkCreateAttribute`: EntitymanagerBulkCreateAttributeResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductBulkCreateAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductBulkCreateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerBulkCreateAttributeRequest**](EntitymanagerBulkCreateAttributeRequest.md) |  | 

### Return type

[**EntitymanagerBulkCreateAttributeResponse**](EntitymanagerBulkCreateAttributeResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductBulkUpdate

> ProductBulkUpdateResponse ProductBulkUpdate(ctx).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductBulkUpdateRequest() // ProductBulkUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductBulkUpdate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductBulkUpdate`: ProductBulkUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductBulkUpdateRequest**](ProductBulkUpdateRequest.md) |  | 

### Return type

[**ProductBulkUpdateResponse**](ProductBulkUpdateResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCreateAttributeGroup

> EntitymanagerAttributeGroup ProductCreateAttributeGroup(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerCreateAttributeGroupRequest() // EntitymanagerCreateAttributeGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductCreateAttributeGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductCreateAttributeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCreateAttributeGroup`: EntitymanagerAttributeGroup
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductCreateAttributeGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCreateAttributeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerCreateAttributeGroupRequest**](EntitymanagerCreateAttributeGroupRequest.md) |  | 

### Return type

[**EntitymanagerAttributeGroup**](EntitymanagerAttributeGroup.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCreateProduct

> ProductCreateProductResponse ProductCreateProduct(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductCreateProductRequest() // ProductCreateProductRequest | The CreateProductRequest message is used to create a new product within the system. It contains various fields that allow specifying the details and attributes of the product.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductCreateProduct(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductCreateProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCreateProduct`: ProductCreateProductResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductCreateProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductCreateProductRequest**](ProductCreateProductRequest.md) | The CreateProductRequest message is used to create a new product within the system. It contains various fields that allow specifying the details and attributes of the product. | 

### Return type

[**ProductCreateProductResponse**](ProductCreateProductResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCreateProductV2

> ProductCreateProductResponseV2 ProductCreateProductV2(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductCreateProductRequestV2() // ProductCreateProductRequestV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductCreateProductV2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductCreateProductV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCreateProductV2`: ProductCreateProductResponseV2
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductCreateProductV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCreateProductV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductCreateProductRequestV2**](ProductCreateProductRequestV2.md) |  | 

### Return type

[**ProductCreateProductResponseV2**](ProductCreateProductResponseV2.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductDelete

> ProductDeleteResponse ProductDelete(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductDeleteRequest() // ProductDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductDelete(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductDelete`: ProductDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductDeleteRequest**](ProductDeleteRequest.md) |  | 

### Return type

[**ProductDeleteResponse**](ProductDeleteResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductDeleteAttribute

> map[string]interface{} ProductDeleteAttribute(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerDeleteAttributeRequest() // EntitymanagerDeleteAttributeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductDeleteAttribute(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductDeleteAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductDeleteAttribute`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductDeleteAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductDeleteAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerDeleteAttributeRequest**](EntitymanagerDeleteAttributeRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductDeleteProduct

> map[string]interface{} ProductDeleteProduct(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductDeleteProductRequest() // ProductDeleteProductRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductDeleteProduct(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductDeleteProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductDeleteProduct`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductDeleteProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductDeleteProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductDeleteProductRequest**](ProductDeleteProductRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetAttributeGroup

> EntitymanagerAttributeGroup ProductGetAttributeGroup(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerGetAttributeGroupRequest() // EntitymanagerGetAttributeGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductGetAttributeGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductGetAttributeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductGetAttributeGroup`: EntitymanagerAttributeGroup
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductGetAttributeGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetAttributeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerGetAttributeGroupRequest**](EntitymanagerGetAttributeGroupRequest.md) |  | 

### Return type

[**EntitymanagerAttributeGroup**](EntitymanagerAttributeGroup.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductListAttributeGroups

> EntitymanagerListAttributeGroupsResponse ProductListAttributeGroups(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerListAttributeGroupsRequest() // EntitymanagerListAttributeGroupsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductListAttributeGroups(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductListAttributeGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductListAttributeGroups`: EntitymanagerListAttributeGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductListAttributeGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductListAttributeGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerListAttributeGroupsRequest**](EntitymanagerListAttributeGroupsRequest.md) |  | 

### Return type

[**EntitymanagerListAttributeGroupsResponse**](EntitymanagerListAttributeGroupsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductUpdateAttribute

> EntitymanagerAttribute ProductUpdateAttribute(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerUpdateAttributeRequest() // EntitymanagerUpdateAttributeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductUpdateAttribute(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductUpdateAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductUpdateAttribute`: EntitymanagerAttribute
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductUpdateAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductUpdateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerUpdateAttributeRequest**](EntitymanagerUpdateAttributeRequest.md) |  | 

### Return type

[**EntitymanagerAttribute**](EntitymanagerAttribute.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductUpdateAttributeGroup

> EntitymanagerAttributeGroup ProductUpdateAttributeGroup(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerUpdateAttributeGroupRequest() // EntitymanagerUpdateAttributeGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductUpdateAttributeGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductUpdateAttributeGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductUpdateAttributeGroup`: EntitymanagerAttributeGroup
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductUpdateAttributeGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductUpdateAttributeGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerUpdateAttributeGroupRequest**](EntitymanagerUpdateAttributeGroupRequest.md) |  | 

### Return type

[**EntitymanagerAttributeGroup**](EntitymanagerAttributeGroup.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductUpdateProduct

> ProductUpdateProductResponse ProductUpdateProduct(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductUpdateProductRequest() // ProductUpdateProductRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductUpdateProduct(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductUpdateProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductUpdateProduct`: ProductUpdateProductResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductUpdateProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductUpdateProductRequest**](ProductUpdateProductRequest.md) |  | 

### Return type

[**ProductUpdateProductResponse**](ProductUpdateProductResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductUpdateProductV2

> map[string]interface{} ProductUpdateProductV2(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductUpdateProductRequestV2() // ProductUpdateProductRequestV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductUpdateProductV2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductUpdateProductV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductUpdateProductV2`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductUpdateProductV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductUpdateProductV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductUpdateProductRequestV2**](ProductUpdateProductRequestV2.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductRemoveMediaGalleryEntryRequest() // ProductRemoveMediaGalleryEntryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.RemoveMediaGalleryEntry(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.RemoveMediaGalleryEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveMediaGalleryEntry`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.RemoveMediaGalleryEntry`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerUpdateAttributeOptionsRequest() // EntitymanagerUpdateAttributeOptionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.UpdateAttributeOptions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdateAttributeOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAttributeOptions`: EntitymanagerUpdateAttributeOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdateAttributeOptions`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataToBeReviewed

> ProductGetEnhanceProductDataWithAIStatusResponse UpdateDataToBeReviewed(ctx).Body(body).Execute()

Get Enhance Product Data With AI Status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductGetEnhanceProductDataWithAIStatusRequest() // ProductGetEnhanceProductDataWithAIStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.UpdateDataToBeReviewed(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdateDataToBeReviewed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataToBeReviewed`: ProductGetEnhanceProductDataWithAIStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdateDataToBeReviewed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataToBeReviewedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductGetEnhanceProductDataWithAIStatusRequest**](ProductGetEnhanceProductDataWithAIStatusRequest.md) |  | 

### Return type

[**ProductGetEnhanceProductDataWithAIStatusResponse**](ProductGetEnhanceProductDataWithAIStatusResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataToBeReviewed_0

> map[string]interface{} UpdateDataToBeReviewed_0(ctx).Body(body).Execute()

Update Data To Be Reviewed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductUpdateDataToBeReviewedRequest() // ProductUpdateDataToBeReviewedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.UpdateDataToBeReviewed_0(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdateDataToBeReviewed_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataToBeReviewed_0`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdateDataToBeReviewed_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataToBeReviewed_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductUpdateDataToBeReviewedRequest**](ProductUpdateDataToBeReviewedRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductUpdateMediaGalleryEntryRequest() // ProductUpdateMediaGalleryEntryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.UpdateMediaGalleryEntry(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdateMediaGalleryEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMediaGalleryEntry`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdateMediaGalleryEntry`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

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
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewEntitymanagerUpdateOptionsListRequest() // EntitymanagerUpdateOptionsListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.UpdateOptionsList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdateOptionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOptionsList`: EntitymanagerUpdateOptionsListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdateOptionsList`: %v\n", resp)
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

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProductWithAI

> ProductUpdateProductWithAIResponse UpdateProductWithAI(ctx).Body(body).Execute()

Update Product With AI



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product"
)

func main() {
	body := *openapiclient.NewProductUpdateProductWithAIRequest() // ProductUpdateProductWithAIRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.UpdateProductWithAI(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.UpdateProductWithAI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProductWithAI`: ProductUpdateProductWithAIResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.UpdateProductWithAI`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductWithAIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductUpdateProductWithAIRequest**](ProductUpdateProductWithAIRequest.md) |  | 

### Return type

[**ProductUpdateProductWithAIResponse**](ProductUpdateProductWithAIResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

