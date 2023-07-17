# # ProductCreateProductRequest
The CreateProductRequest message is used to create a new product within the system. It contains various fields that allow specifying the details and attributes of the product.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** | Represents the ID of the tenant associated with the product.  | [optional]
**EntityType**| **string** | Specifies the type of entity for the product.  | [optional]
**EntityCode**| **string** | Indicates the code of the entity associated with the product.  | [optional]
**Code**| **string** | Represents the unique code or identifier for the product.  | [optional]
**IsConfigurable**| **bool** | Specifies whether the product has variants or not.  | [optional]
**VariantAttributes**| **[]string** | Contains a list of attributes specific to the product variants.  | [optional]
**IsVirtual**| **bool** | Indicates whether the product is virtual or not.  | [optional]
**IsGiftcard**| **bool** | Specifies whether the product is a gift card or not.  | [optional]
**UrlKey**| [**ProductLocalizedText**](ProductLocalizedText.md) |   | [optional]
**MaxSaleableQuantity**| **int64** | Specifies the maximum quantity that can be sold for the product in each order.  | [optional]
**Attributes**| [**map[string]ProtobufAny**](ProtobufAny.md) | Contains a map of additional attributes associated with the product, where the key is the attribute name and the value is any type of value.  | [optional]
**Variants**| [**map[string]ProductProductVariant**](ProductProductVariant.md) | Represents a map of product variants associated with the product, where the key is the variant ID or code, and the value is a ProductVariant message.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

