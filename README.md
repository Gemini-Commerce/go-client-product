# Go API client for product

Introducing our revolutionary Product Management Service! Designed to streamline your product inventory and elevate customer experiences, our cutting-edge protobuf service is a game-changer in the world of efficient product management.

With our service, you can effortlessly create new products, allowing you to quickly bring your ideas to life and expand your catalog. Retrieve product information in a snap, providing accurate and personalized details to your customers based on their specific needs and preferences.

Stay ahead of the competition by easily updating product information, ensuring your catalog is always up-to-date and optimized. Seamlessly remove products from your inventory when needed, maintaining a clean and relevant product selection.

Enhance the visual appeal of your products with our advanced media gallery functionalities. Effortlessly add and update captivating images and videos to showcase your products in the best possible light, engaging your customers and driving conversions.

Personalization is key in today's market, and our service enables you to offer unique options to your customers. Easily create and manage lists of customizable options for your products, providing flexibility and tailoring to individual preferences.

Attributes play a vital role in defining products, and our service empowers you to effectively manage them. From bulk attribute creation to listing and retrieving attribute options, our service ensures your product information is rich and comprehensive.

Our service extends its capabilities to entity management, allowing you to effortlessly handle different entities and create customized options lists associated with them. This provides further flexibility and customization options for your product offerings.

When it comes to bulk updates, our service has you covered. Effortlessly update multiple products simultaneously, saving you time and streamlining your operations.

Finding specific products and variants is a breeze with our service. Quickly locate products based on their unique stock keeping unit (SKU) values, ensuring efficient inventory management and smooth order fulfillment.

Experience a new level of efficiency and productivity with our Product Management Service. Unlock the full potential of streamlined product management and empower your business to thrive in today's competitive market. Try our service today and elevate your product management to new heights!

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import product "github.com/Gemini-Commerce/go-client-product"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `product.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), product.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `product.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), product.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `product.ContextOperationServerIndices` and `product.ContextOperationServerVariables` context maps.

```go
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
*ProductAPI* | [**AddMediaGalleryEntry**](docs/ProductAPI.md#addmediagalleryentry) | **Post** /product.Product/AddMediaGalleryEntry | Add Media Gallery Entry
*ProductAPI* | [**BulkAddAssetsEntries**](docs/ProductAPI.md#bulkaddassetsentries) | **Post** /product.Product/BulkAddAssetsEntries | Bulk Add Assets Entries
*ProductAPI* | [**BulkDeleteProducts**](docs/ProductAPI.md#bulkdeleteproducts) | **Post** /product.Product/BulkDeleteProducts | Bulk Delete Products
*ProductAPI* | [**BulkEnhanceProductDataWithAI**](docs/ProductAPI.md#bulkenhanceproductdatawithai) | **Post** /product.Product/BulkEnhanceProductDataWithAI | Bulk Enhance Product Data With AI
*ProductAPI* | [**BulkRemoveAssetsEntries**](docs/ProductAPI.md#bulkremoveassetsentries) | **Post** /product.Product/BulkRemoveAssetsEntries | Bulk Remove Assets Entries
*ProductAPI* | [**BulkUpdateAssetsEntries**](docs/ProductAPI.md#bulkupdateassetsentries) | **Post** /product.Product/BulkUpdateAssetsEntries | Bulk Update Assets Entries
*ProductAPI* | [**BulkUpdateV2**](docs/ProductAPI.md#bulkupdatev2) | **Post** /product.Product/BulkUpdateV2 | Bulk Update Products
*ProductAPI* | [**CreateAttributeOptions**](docs/ProductAPI.md#createattributeoptions) | **Post** /product.Product/CreateAttributeOptions | Create Attribute Options
*ProductAPI* | [**CreateEntity**](docs/ProductAPI.md#createentity) | **Post** /product.Product/CreateEntity | Create Entity
*ProductAPI* | [**CreateOptionsList**](docs/ProductAPI.md#createoptionslist) | **Post** /product.Product/CreateOptionsList | Create Options List
*ProductAPI* | [**CreateProductWithAI**](docs/ProductAPI.md#createproductwithai) | **Post** /product.Product/CreateProductWithAI | Create Product With AI
*ProductAPI* | [**GetAttributeOption**](docs/ProductAPI.md#getattributeoption) | **Post** /product.Product/GetAttributeOption | Get Attribute Option
*ProductAPI* | [**GetAttributeOptions**](docs/ProductAPI.md#getattributeoptions) | **Post** /product.Product/GetAttributeOptions | Get Attribute Options
*ProductAPI* | [**GetEntity**](docs/ProductAPI.md#getentity) | **Post** /product.Product/GetEntity | Get Entity Details
*ProductAPI* | [**GetOptionsList**](docs/ProductAPI.md#getoptionslist) | **Post** /product.Product/GetOptionsList | Get Options List
*ProductAPI* | [**GetProduct**](docs/ProductAPI.md#getproduct) | **Post** /product.Product/GetProduct | Get Product
*ProductAPI* | [**GetProductByCode**](docs/ProductAPI.md#getproductbycode) | **Post** /product.Product/GetProductByCode | Get Product By Code
*ProductAPI* | [**GetProductByUrlKey**](docs/ProductAPI.md#getproductbyurlkey) | **Post** /product.Product/GetProductByUrlKey | Get Product By Url Key
*ProductAPI* | [**GetProductDataInReview**](docs/ProductAPI.md#getproductdatainreview) | **Post** /product.Product/GetProductDataInReview | Get Product Data In Review
*ProductAPI* | [**ListAttributeOptions**](docs/ProductAPI.md#listattributeoptions) | **Post** /product.Product/ListAttributeOptions | List Attribute Options
*ProductAPI* | [**ListEntities**](docs/ProductAPI.md#listentities) | **Post** /product.Product/ListEntities | List Entities
*ProductAPI* | [**ListOptionsLists**](docs/ProductAPI.md#listoptionslists) | **Post** /product.Product/ListOptionsLists | List Options Lists
*ProductAPI* | [**ListProducts**](docs/ProductAPI.md#listproducts) | **Post** /product.Product/ListProducts | List Products
*ProductAPI* | [**ListProductsByIds**](docs/ProductAPI.md#listproductsbyids) | **Post** /product.Product/ListProductsByIds | List Products By Ids
*ProductAPI* | [**ListProductsBySku**](docs/ProductAPI.md#listproductsbysku) | **Post** /product.Product/ListProductsBySku | List Products By Sku
*ProductAPI* | [**ListVariantsBySku**](docs/ProductAPI.md#listvariantsbysku) | **Post** /product.Product/ListVariantsBySku | List Product Variants By Sku
*ProductAPI* | [**ProductBulkCreateAttribute**](docs/ProductAPI.md#productbulkcreateattribute) | **Post** /product.Product/BulkCreateAttribute | 
*ProductAPI* | [**ProductBulkUpdate**](docs/ProductAPI.md#productbulkupdate) | **Post** /product.Product/BulkUpdate | 
*ProductAPI* | [**ProductCreateAttributeGroup**](docs/ProductAPI.md#productcreateattributegroup) | **Post** /product.Product/CreateAttributeGroup | 
*ProductAPI* | [**ProductCreateProduct**](docs/ProductAPI.md#productcreateproduct) | **Post** /product.Product/CreateProduct | 
*ProductAPI* | [**ProductCreateProductV2**](docs/ProductAPI.md#productcreateproductv2) | **Post** /product.Product/CreateProductV2 | 
*ProductAPI* | [**ProductDelete**](docs/ProductAPI.md#productdelete) | **Post** /product.Product/Delete | 
*ProductAPI* | [**ProductDeleteAttribute**](docs/ProductAPI.md#productdeleteattribute) | **Post** /product.Product/DeleteAttribute | 
*ProductAPI* | [**ProductDeleteProduct**](docs/ProductAPI.md#productdeleteproduct) | **Post** /product.Product/DeleteProduct | 
*ProductAPI* | [**ProductGetAttributeGroup**](docs/ProductAPI.md#productgetattributegroup) | **Post** /product.Product/GetAttributeGroup | 
*ProductAPI* | [**ProductListAttributeGroups**](docs/ProductAPI.md#productlistattributegroups) | **Post** /product.Product/ListAttributeGroups | 
*ProductAPI* | [**ProductUpdateAttribute**](docs/ProductAPI.md#productupdateattribute) | **Post** /product.Product/UpdateAttribute | 
*ProductAPI* | [**ProductUpdateAttributeGroup**](docs/ProductAPI.md#productupdateattributegroup) | **Post** /product.Product/UpdateAttributeGroup | 
*ProductAPI* | [**ProductUpdateProduct**](docs/ProductAPI.md#productupdateproduct) | **Post** /product.Product/UpdateProduct | 
*ProductAPI* | [**ProductUpdateProductV2**](docs/ProductAPI.md#productupdateproductv2) | **Post** /product.Product/UpdateProductV2 | 
*ProductAPI* | [**RemoveMediaGalleryEntry**](docs/ProductAPI.md#removemediagalleryentry) | **Post** /product.Product/RemoveMediaGalleryEntry | Remove Media Gallery Entry
*ProductAPI* | [**UpdateAttributeOptions**](docs/ProductAPI.md#updateattributeoptions) | **Post** /product.Product/UpdateAttributeOptions | Update Attribute Options
*ProductAPI* | [**UpdateDataToBeReviewed**](docs/ProductAPI.md#updatedatatobereviewed) | **Post** /product.Product/GetEnhanceProductDataWithAIStatus | Get Enhance Product Data With AI Status
*ProductAPI* | [**UpdateDataToBeReviewed_0**](docs/ProductAPI.md#updatedatatobereviewed_0) | **Post** /product.Product/UpdateDataToBeReviewed | Update Data To Be Reviewed
*ProductAPI* | [**UpdateMediaGalleryEntry**](docs/ProductAPI.md#updatemediagalleryentry) | **Post** /product.Product/UpdateMediaGalleryEntry | Update Media Gallery Entry
*ProductAPI* | [**UpdateOptionsList**](docs/ProductAPI.md#updateoptionslist) | **Post** /product.Product/UpdateOptionsList | Update Options List
*ProductAPI* | [**UpdateProductWithAI**](docs/ProductAPI.md#updateproductwithai) | **Post** /product.Product/UpdateProductWithAI | Update Product With AI


## Documentation For Models

 - [AttributeInReviewString](docs/AttributeInReviewString.md)
 - [BulkUpdateAssetsEntriesRequestUpdateEntity](docs/BulkUpdateAssetsEntriesRequestUpdateEntity.md)
 - [EntitymanagerAiContext](docs/EntitymanagerAiContext.md)
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
 - [GetEnhanceProductDataWithAIStatusResponseJob](docs/GetEnhanceProductDataWithAIStatusResponseJob.md)
 - [ListProductsRequestFilter](docs/ListProductsRequestFilter.md)
 - [ProductAddMediaGalleryEntryRequest](docs/ProductAddMediaGalleryEntryRequest.md)
 - [ProductAddMediaGalleryEntryResponse](docs/ProductAddMediaGalleryEntryResponse.md)
 - [ProductAssetData](docs/ProductAssetData.md)
 - [ProductAssets](docs/ProductAssets.md)
 - [ProductAssetsEntry](docs/ProductAssetsEntry.md)
 - [ProductAssetsEntryMetadata](docs/ProductAssetsEntryMetadata.md)
 - [ProductAttributeInReview](docs/ProductAttributeInReview.md)
 - [ProductAttributeInReviewError](docs/ProductAttributeInReviewError.md)
 - [ProductAttributeInReviewJobStatus](docs/ProductAttributeInReviewJobStatus.md)
 - [ProductAttributeInReviewJobType](docs/ProductAttributeInReviewJobType.md)
 - [ProductAttributeInReviewSource](docs/ProductAttributeInReviewSource.md)
 - [ProductAttributeResponseError](docs/ProductAttributeResponseError.md)
 - [ProductAttributeToEnrich](docs/ProductAttributeToEnrich.md)
 - [ProductAttributeToEnrichType](docs/ProductAttributeToEnrichType.md)
 - [ProductBulkAddAssetsEntriesRequest](docs/ProductBulkAddAssetsEntriesRequest.md)
 - [ProductBulkAddAssetsEntriesResponse](docs/ProductBulkAddAssetsEntriesResponse.md)
 - [ProductBulkDeleteProductsRequest](docs/ProductBulkDeleteProductsRequest.md)
 - [ProductBulkEnhanceProductDataWithAIRequest](docs/ProductBulkEnhanceProductDataWithAIRequest.md)
 - [ProductBulkRemoveAssetsEntriesRequest](docs/ProductBulkRemoveAssetsEntriesRequest.md)
 - [ProductBulkUpdateAssetsEntriesRequest](docs/ProductBulkUpdateAssetsEntriesRequest.md)
 - [ProductBulkUpdateAssetsEntriesResponse](docs/ProductBulkUpdateAssetsEntriesResponse.md)
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
 - [ProductCreateProductWithAIRequest](docs/ProductCreateProductWithAIRequest.md)
 - [ProductCreateProductWithAIResponse](docs/ProductCreateProductWithAIResponse.md)
 - [ProductDataInReview](docs/ProductDataInReview.md)
 - [ProductDeleteProductRequest](docs/ProductDeleteProductRequest.md)
 - [ProductDeleteRequest](docs/ProductDeleteRequest.md)
 - [ProductDeleteResponse](docs/ProductDeleteResponse.md)
 - [ProductEnrichAction](docs/ProductEnrichAction.md)
 - [ProductFieldMask](docs/ProductFieldMask.md)
 - [ProductGetEnhanceProductDataWithAIStatusRequest](docs/ProductGetEnhanceProductDataWithAIStatusRequest.md)
 - [ProductGetEnhanceProductDataWithAIStatusResponse](docs/ProductGetEnhanceProductDataWithAIStatusResponse.md)
 - [ProductGetProductByCodeRequest](docs/ProductGetProductByCodeRequest.md)
 - [ProductGetProductByCodeResponse](docs/ProductGetProductByCodeResponse.md)
 - [ProductGetProductByUrlKeyRequest](docs/ProductGetProductByUrlKeyRequest.md)
 - [ProductGetProductByUrlKeyResponse](docs/ProductGetProductByUrlKeyResponse.md)
 - [ProductGetProductDataInReviewRequest](docs/ProductGetProductDataInReviewRequest.md)
 - [ProductGetProductDataInReviewResponse](docs/ProductGetProductDataInReviewResponse.md)
 - [ProductGetProductRequest](docs/ProductGetProductRequest.md)
 - [ProductGetProductResponse](docs/ProductGetProductResponse.md)
 - [ProductLanguageCode](docs/ProductLanguageCode.md)
 - [ProductListProductsByIdsRequest](docs/ProductListProductsByIdsRequest.md)
 - [ProductListProductsByIdsResponse](docs/ProductListProductsByIdsResponse.md)
 - [ProductListProductsBySkuRequest](docs/ProductListProductsBySkuRequest.md)
 - [ProductListProductsBySkuResponse](docs/ProductListProductsBySkuResponse.md)
 - [ProductListProductsRequest](docs/ProductListProductsRequest.md)
 - [ProductListProductsResponse](docs/ProductListProductsResponse.md)
 - [ProductListVariantsBySkuRequest](docs/ProductListVariantsBySkuRequest.md)
 - [ProductListVariantsBySkuResponse](docs/ProductListVariantsBySkuResponse.md)
 - [ProductLocalizedAsset](docs/ProductLocalizedAsset.md)
 - [ProductLocalizedText](docs/ProductLocalizedText.md)
 - [ProductMediaGallery](docs/ProductMediaGallery.md)
 - [ProductMediaGalleryEntry](docs/ProductMediaGalleryEntry.md)
 - [ProductMediaGalleryEntryMetadata](docs/ProductMediaGalleryEntryMetadata.md)
 - [ProductProductEntity](docs/ProductProductEntity.md)
 - [ProductProductResponseError](docs/ProductProductResponseError.md)
 - [ProductProductVariant](docs/ProductProductVariant.md)
 - [ProductRemoveMediaGalleryEntryRequest](docs/ProductRemoveMediaGalleryEntryRequest.md)
 - [ProductTranslateAction](docs/ProductTranslateAction.md)
 - [ProductUpdateAssetEntryPayload](docs/ProductUpdateAssetEntryPayload.md)
 - [ProductUpdateDataToBeReviewedRequest](docs/ProductUpdateDataToBeReviewedRequest.md)
 - [ProductUpdateMediaGalleryEntryRequest](docs/ProductUpdateMediaGalleryEntryRequest.md)
 - [ProductUpdateProductRequest](docs/ProductUpdateProductRequest.md)
 - [ProductUpdateProductRequestV2](docs/ProductUpdateProductRequestV2.md)
 - [ProductUpdateProductResponse](docs/ProductUpdateProductResponse.md)
 - [ProductUpdateProductWithAIRequest](docs/ProductUpdateProductWithAIRequest.md)
 - [ProductUpdateProductWithAIResponse](docs/ProductUpdateProductWithAIResponse.md)
 - [ProductentitymanagerLocalizedText](docs/ProductentitymanagerLocalizedText.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [RpcStatus](docs/RpcStatus.md)
 - [TranslateActionAttributeCodesToTranslate](docs/TranslateActionAttributeCodesToTranslate.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Authorization

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		product.ContextAPIKeys,
		map[string]product.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### standardAuthorization


- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: 
- **Scopes**: N/A

Example

```go
auth := context.WithValue(context.Background(), product.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, product.ContextOAuth2, tokenSource)
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

