# Go API client for product

API for managing products

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import product "github.com/gemini-commerce/go-client-product"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), product.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), product.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), product.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), product.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://product.api.gogemini.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ProductAPI* | [**ProductAddMediaGalleryEntry**](docs/ProductAPI.md#productaddmediagalleryentry) | **Post** /product.Product/AddMediaGalleryEntry | 
*ProductAPI* | [**ProductBulkCreateAttribute**](docs/ProductAPI.md#productbulkcreateattribute) | **Post** /product.Product/BulkCreateAttribute | 
*ProductAPI* | [**ProductBulkDeleteProducts**](docs/ProductAPI.md#productbulkdeleteproducts) | **Post** /product.Product/BulkDeleteProducts | 
*ProductAPI* | [**ProductBulkUpdate**](docs/ProductAPI.md#productbulkupdate) | **Post** /product.Product/BulkUpdate | 
*ProductAPI* | [**ProductBulkUpdateV2**](docs/ProductAPI.md#productbulkupdatev2) | **Post** /product.Product/BulkUpdateV2 | 
*ProductAPI* | [**ProductCreateAttributeGroup**](docs/ProductAPI.md#productcreateattributegroup) | **Post** /product.Product/CreateAttributeGroup | 
*ProductAPI* | [**ProductCreateAttributeOptions**](docs/ProductAPI.md#productcreateattributeoptions) | **Post** /product.Product/CreateAttributeOptions | 
*ProductAPI* | [**ProductCreateEntity**](docs/ProductAPI.md#productcreateentity) | **Post** /product.Product/CreateEntity | 
*ProductAPI* | [**ProductCreateOptionsList**](docs/ProductAPI.md#productcreateoptionslist) | **Post** /product.Product/CreateOptionsList | 
*ProductAPI* | [**ProductCreateProduct**](docs/ProductAPI.md#productcreateproduct) | **Post** /product.Product/CreateProduct | 
*ProductAPI* | [**ProductCreateProductV2**](docs/ProductAPI.md#productcreateproductv2) | **Post** /product.Product/CreateProductV2 | 
*ProductAPI* | [**ProductDelete**](docs/ProductAPI.md#productdelete) | **Post** /product.Product/Delete | 
*ProductAPI* | [**ProductDeleteAttribute**](docs/ProductAPI.md#productdeleteattribute) | **Post** /product.Product/DeleteAttribute | 
*ProductAPI* | [**ProductDeleteProduct**](docs/ProductAPI.md#productdeleteproduct) | **Post** /product.Product/DeleteProduct | 
*ProductAPI* | [**ProductGetAttributeGroup**](docs/ProductAPI.md#productgetattributegroup) | **Post** /product.Product/GetAttributeGroup | 
*ProductAPI* | [**ProductGetAttributeOption**](docs/ProductAPI.md#productgetattributeoption) | **Post** /product.Product/GetAttributeOption | 
*ProductAPI* | [**ProductGetAttributeOptionByCode**](docs/ProductAPI.md#productgetattributeoptionbycode) | **Post** /product.Product/GetAttributeOptionByCode | 
*ProductAPI* | [**ProductGetAttributeOptions**](docs/ProductAPI.md#productgetattributeoptions) | **Post** /product.Product/GetAttributeOptions | 
*ProductAPI* | [**ProductGetEntity**](docs/ProductAPI.md#productgetentity) | **Post** /product.Product/GetEntity | 
*ProductAPI* | [**ProductGetOptionsList**](docs/ProductAPI.md#productgetoptionslist) | **Post** /product.Product/GetOptionsList | 
*ProductAPI* | [**ProductGetProduct**](docs/ProductAPI.md#productgetproduct) | **Post** /product.Product/GetProduct | 
*ProductAPI* | [**ProductGetProductByCode**](docs/ProductAPI.md#productgetproductbycode) | **Post** /product.Product/GetProductByCode | 
*ProductAPI* | [**ProductGetProductByUrlKey**](docs/ProductAPI.md#productgetproductbyurlkey) | **Post** /product.Product/GetProductByUrlKey | 
*ProductAPI* | [**ProductListAttributeGroups**](docs/ProductAPI.md#productlistattributegroups) | **Post** /product.Product/ListAttributeGroups | Attribute Groups endpoints
*ProductAPI* | [**ProductListAttributeOptions**](docs/ProductAPI.md#productlistattributeoptions) | **Post** /product.Product/ListAttributeOptions | 
*ProductAPI* | [**ProductListEntities**](docs/ProductAPI.md#productlistentities) | **Post** /product.Product/ListEntities | 
*ProductAPI* | [**ProductListOptionsLists**](docs/ProductAPI.md#productlistoptionslists) | **Post** /product.Product/ListOptionsLists | 
*ProductAPI* | [**ProductListProducts**](docs/ProductAPI.md#productlistproducts) | **Post** /product.Product/ListProducts | 
*ProductAPI* | [**ProductListProductsByIds**](docs/ProductAPI.md#productlistproductsbyids) | **Post** /product.Product/ListProductsByIds | 
*ProductAPI* | [**ProductListProductsBySku**](docs/ProductAPI.md#productlistproductsbysku) | **Post** /product.Product/ListProductsBySku | 
*ProductAPI* | [**ProductListVariantsBySku**](docs/ProductAPI.md#productlistvariantsbysku) | **Post** /product.Product/ListVariantsBySku | 
*ProductAPI* | [**ProductRemoveMediaGalleryEntry**](docs/ProductAPI.md#productremovemediagalleryentry) | **Post** /product.Product/RemoveMediaGalleryEntry | 
*ProductAPI* | [**ProductUpdateAttribute**](docs/ProductAPI.md#productupdateattribute) | **Post** /product.Product/UpdateAttribute | 
*ProductAPI* | [**ProductUpdateAttributeGroup**](docs/ProductAPI.md#productupdateattributegroup) | **Post** /product.Product/UpdateAttributeGroup | 
*ProductAPI* | [**ProductUpdateAttributeOptions**](docs/ProductAPI.md#productupdateattributeoptions) | **Post** /product.Product/UpdateAttributeOptions | 
*ProductAPI* | [**ProductUpdateMediaGalleryEntry**](docs/ProductAPI.md#productupdatemediagalleryentry) | **Post** /product.Product/UpdateMediaGalleryEntry | 
*ProductAPI* | [**ProductUpdateOptionsList**](docs/ProductAPI.md#productupdateoptionslist) | **Post** /product.Product/UpdateOptionsList | 
*ProductAPI* | [**ProductUpdateProduct**](docs/ProductAPI.md#productupdateproduct) | **Post** /product.Product/UpdateProduct | 
*ProductAPI* | [**ProductUpdateProductV2**](docs/ProductAPI.md#productupdateproductv2) | **Post** /product.Product/UpdateProductV2 | 


## Documentation For Models

 - [EntitymanagerAttribute](docs/EntitymanagerAttribute.md)
 - [EntitymanagerAttributeGroup](docs/EntitymanagerAttributeGroup.md)
 - [EntitymanagerAttributeOption](docs/EntitymanagerAttributeOption.md)
 - [EntitymanagerAttributeOptionErrors](docs/EntitymanagerAttributeOptionErrors.md)
 - [EntitymanagerAttributeOptionSwatch](docs/EntitymanagerAttributeOptionSwatch.md)
 - [EntitymanagerAttributeWriteError](docs/EntitymanagerAttributeWriteError.md)
 - [EntitymanagerAttributeWriteErrors](docs/EntitymanagerAttributeWriteErrors.md)
 - [EntitymanagerBulkCreateAttributeRequest](docs/EntitymanagerBulkCreateAttributeRequest.md)
 - [EntitymanagerBulkCreateAttributeResponse](docs/EntitymanagerBulkCreateAttributeResponse.md)
 - [EntitymanagerCreateAttributeGroupRequest](docs/EntitymanagerCreateAttributeGroupRequest.md)
 - [EntitymanagerCreateAttributeOptionsRequest](docs/EntitymanagerCreateAttributeOptionsRequest.md)
 - [EntitymanagerCreateAttributeOptionsResponse](docs/EntitymanagerCreateAttributeOptionsResponse.md)
 - [EntitymanagerCreateEntityResponse](docs/EntitymanagerCreateEntityResponse.md)
 - [EntitymanagerCreateOptionsListRequest](docs/EntitymanagerCreateOptionsListRequest.md)
 - [EntitymanagerCreateOptionsListResponse](docs/EntitymanagerCreateOptionsListResponse.md)
 - [EntitymanagerDeleteAttributeRequest](docs/EntitymanagerDeleteAttributeRequest.md)
 - [EntitymanagerEntity](docs/EntitymanagerEntity.md)
 - [EntitymanagerEntityIdentifier](docs/EntitymanagerEntityIdentifier.md)
 - [EntitymanagerEntityRequest](docs/EntitymanagerEntityRequest.md)
 - [EntitymanagerGetAttributeGroupRequest](docs/EntitymanagerGetAttributeGroupRequest.md)
 - [EntitymanagerGetAttributeOptionByCodeRequest](docs/EntitymanagerGetAttributeOptionByCodeRequest.md)
 - [EntitymanagerGetAttributeOptionByCodeResponse](docs/EntitymanagerGetAttributeOptionByCodeResponse.md)
 - [EntitymanagerGetAttributeOptionRequest](docs/EntitymanagerGetAttributeOptionRequest.md)
 - [EntitymanagerGetAttributeOptionResponse](docs/EntitymanagerGetAttributeOptionResponse.md)
 - [EntitymanagerGetAttributeOptionsRequest](docs/EntitymanagerGetAttributeOptionsRequest.md)
 - [EntitymanagerGetAttributeOptionsRequestOption](docs/EntitymanagerGetAttributeOptionsRequestOption.md)
 - [EntitymanagerGetAttributeOptionsResponse](docs/EntitymanagerGetAttributeOptionsResponse.md)
 - [EntitymanagerGetAttributeOptionsResponseOption](docs/EntitymanagerGetAttributeOptionsResponseOption.md)
 - [EntitymanagerGetOptionsListRequest](docs/EntitymanagerGetOptionsListRequest.md)
 - [EntitymanagerGetOptionsListResponse](docs/EntitymanagerGetOptionsListResponse.md)
 - [EntitymanagerListAttributeGroupsRequest](docs/EntitymanagerListAttributeGroupsRequest.md)
 - [EntitymanagerListAttributeGroupsResponse](docs/EntitymanagerListAttributeGroupsResponse.md)
 - [EntitymanagerListAttributeOptionsRequest](docs/EntitymanagerListAttributeOptionsRequest.md)
 - [EntitymanagerListAttributeOptionsResponse](docs/EntitymanagerListAttributeOptionsResponse.md)
 - [EntitymanagerListEntitiesRequest](docs/EntitymanagerListEntitiesRequest.md)
 - [EntitymanagerListEntitiesResponse](docs/EntitymanagerListEntitiesResponse.md)
 - [EntitymanagerListOptionsListsRequest](docs/EntitymanagerListOptionsListsRequest.md)
 - [EntitymanagerListOptionsListsResponse](docs/EntitymanagerListOptionsListsResponse.md)
 - [EntitymanagerOptionsList](docs/EntitymanagerOptionsList.md)
 - [EntitymanagerRenderAs](docs/EntitymanagerRenderAs.md)
 - [EntitymanagerTypes](docs/EntitymanagerTypes.md)
 - [EntitymanagerUpdateAttributeGroupRequest](docs/EntitymanagerUpdateAttributeGroupRequest.md)
 - [EntitymanagerUpdateAttributeGroupRequestPayload](docs/EntitymanagerUpdateAttributeGroupRequestPayload.md)
 - [EntitymanagerUpdateAttributeOptionsRequest](docs/EntitymanagerUpdateAttributeOptionsRequest.md)
 - [EntitymanagerUpdateAttributeOptionsResponse](docs/EntitymanagerUpdateAttributeOptionsResponse.md)
 - [EntitymanagerUpdateAttributeRequest](docs/EntitymanagerUpdateAttributeRequest.md)
 - [EntitymanagerUpdateAttributeRequestPayload](docs/EntitymanagerUpdateAttributeRequestPayload.md)
 - [EntitymanagerUpdateOptionsListRequest](docs/EntitymanagerUpdateOptionsListRequest.md)
 - [EntitymanagerUpdateOptionsListResponse](docs/EntitymanagerUpdateOptionsListResponse.md)
 - [ListProductsRequestFilter](docs/ListProductsRequestFilter.md)
 - [ProductAddMediaGalleryEntryRequest](docs/ProductAddMediaGalleryEntryRequest.md)
 - [ProductAddMediaGalleryEntryResponse](docs/ProductAddMediaGalleryEntryResponse.md)
 - [ProductAttributeResponseError](docs/ProductAttributeResponseError.md)
 - [ProductBulkDeleteProductsRequest](docs/ProductBulkDeleteProductsRequest.md)
 - [ProductBulkUpdateRequest](docs/ProductBulkUpdateRequest.md)
 - [ProductBulkUpdateRequestPayload](docs/ProductBulkUpdateRequestPayload.md)
 - [ProductBulkUpdateRequestV2](docs/ProductBulkUpdateRequestV2.md)
 - [ProductBulkUpdateRequestV2Payload](docs/ProductBulkUpdateRequestV2Payload.md)
 - [ProductBulkUpdateResponse](docs/ProductBulkUpdateResponse.md)
 - [ProductBulkUpdateResponseResponse](docs/ProductBulkUpdateResponseResponse.md)
 - [ProductBulkUpdateResponseV2](docs/ProductBulkUpdateResponseV2.md)
 - [ProductBulkUpdateResponseV2Response](docs/ProductBulkUpdateResponseV2Response.md)
 - [ProductCreateProductRequest](docs/ProductCreateProductRequest.md)
 - [ProductCreateProductRequestV2](docs/ProductCreateProductRequestV2.md)
 - [ProductCreateProductResponse](docs/ProductCreateProductResponse.md)
 - [ProductCreateProductResponseV2](docs/ProductCreateProductResponseV2.md)
 - [ProductDeleteProductRequest](docs/ProductDeleteProductRequest.md)
 - [ProductDeleteRequest](docs/ProductDeleteRequest.md)
 - [ProductDeleteResponse](docs/ProductDeleteResponse.md)
 - [ProductFieldMask](docs/ProductFieldMask.md)
 - [ProductGetProductByCodeRequest](docs/ProductGetProductByCodeRequest.md)
 - [ProductGetProductByCodeResponse](docs/ProductGetProductByCodeResponse.md)
 - [ProductGetProductByUrlKeyRequest](docs/ProductGetProductByUrlKeyRequest.md)
 - [ProductGetProductByUrlKeyResponse](docs/ProductGetProductByUrlKeyResponse.md)
 - [ProductGetProductRequest](docs/ProductGetProductRequest.md)
 - [ProductGetProductResponse](docs/ProductGetProductResponse.md)
 - [ProductListProductsByIdsRequest](docs/ProductListProductsByIdsRequest.md)
 - [ProductListProductsByIdsResponse](docs/ProductListProductsByIdsResponse.md)
 - [ProductListProductsBySkuRequest](docs/ProductListProductsBySkuRequest.md)
 - [ProductListProductsBySkuResponse](docs/ProductListProductsBySkuResponse.md)
 - [ProductListProductsRequest](docs/ProductListProductsRequest.md)
 - [ProductListProductsResponse](docs/ProductListProductsResponse.md)
 - [ProductListVariantsBySkuRequest](docs/ProductListVariantsBySkuRequest.md)
 - [ProductListVariantsBySkuResponse](docs/ProductListVariantsBySkuResponse.md)
 - [ProductLocalizedText](docs/ProductLocalizedText.md)
 - [ProductMediaGallery](docs/ProductMediaGallery.md)
 - [ProductMediaGalleryEntry](docs/ProductMediaGalleryEntry.md)
 - [ProductMediaGalleryEntryMetadata](docs/ProductMediaGalleryEntryMetadata.md)
 - [ProductProductEntity](docs/ProductProductEntity.md)
 - [ProductProductResponseError](docs/ProductProductResponseError.md)
 - [ProductProductVariant](docs/ProductProductVariant.md)
 - [ProductRemoveMediaGalleryEntryRequest](docs/ProductRemoveMediaGalleryEntryRequest.md)
 - [ProductUpdateMediaGalleryEntryRequest](docs/ProductUpdateMediaGalleryEntryRequest.md)
 - [ProductUpdateProductRequest](docs/ProductUpdateProductRequest.md)
 - [ProductUpdateProductRequestV2](docs/ProductUpdateProductRequestV2.md)
 - [ProductUpdateProductResponse](docs/ProductUpdateProductResponse.md)
 - [ProductentitymanagerLocalizedText](docs/ProductentitymanagerLocalizedText.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [RpcStatus](docs/RpcStatus.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### standardAuthorization


- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@gemini-commerce.com

