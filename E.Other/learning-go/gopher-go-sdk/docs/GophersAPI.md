# \GophersAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckHealth**](GophersAPI.md#CheckHealth) | **Get** /healthz | 
[**GopherDelete**](GophersAPI.md#GopherDelete) | **Delete** /gopher | 
[**GopherGet**](GophersAPI.md#GopherGet) | **Get** /gopher | 
[**GopherPost**](GophersAPI.md#GopherPost) | **Post** /gopher | Add a new Gopher
[**GopherPut**](GophersAPI.md#GopherPut) | **Put** /gopher | 
[**GophersGet**](GophersAPI.md#GophersGet) | **Get** /gophers | 



## CheckHealth

> string CheckHealth(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/scraly/gophers-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GophersAPI.CheckHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GophersAPI.CheckHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckHealth`: string
	fmt.Fprintf(os.Stdout, "Response from `GophersAPI.CheckHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckHealthRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GopherDelete

> GopherDelete(ctx).Name(name).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/scraly/gophers-go-sdk"
)

func main() {
	name := "name_example" // string | Gopher name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GophersAPI.GopherDelete(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GophersAPI.GopherDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGopherDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Gopher name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GopherGet

> Gopher GopherGet(ctx).Name(name).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/scraly/gophers-go-sdk"
)

func main() {
	name := "name_example" // string | Gopher name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GophersAPI.GopherGet(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GophersAPI.GopherGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GopherGet`: Gopher
	fmt.Fprintf(os.Stdout, "Response from `GophersAPI.GopherGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGopherGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Gopher name | 

### Return type

[**Gopher**](Gopher.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GopherPost

> Gopher GopherPost(ctx).Gopher(gopher).Execute()

Add a new Gopher

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/scraly/gophers-go-sdk"
)

func main() {
	gopher := *openapiclient.NewGopherPutRequest("Name_example", "Displayname_example", "Url_example") // GopherPutRequest | The Gopher to create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GophersAPI.GopherPost(context.Background()).Gopher(gopher).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GophersAPI.GopherPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GopherPost`: Gopher
	fmt.Fprintf(os.Stdout, "Response from `GophersAPI.GopherPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGopherPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gopher** | [**GopherPutRequest**](GopherPutRequest.md) | The Gopher to create. | 

### Return type

[**Gopher**](Gopher.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GopherPut

> Gopher GopherPut(ctx).Gopher(gopher).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/scraly/gophers-go-sdk"
)

func main() {
	gopher := *openapiclient.NewGopherPutRequest("Name_example", "Displayname_example", "Url_example") // GopherPutRequest | The Gopher to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GophersAPI.GopherPut(context.Background()).Gopher(gopher).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GophersAPI.GopherPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GopherPut`: Gopher
	fmt.Fprintf(os.Stdout, "Response from `GophersAPI.GopherPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGopherPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gopher** | [**GopherPutRequest**](GopherPutRequest.md) | The Gopher to update. | 

### Return type

[**Gopher**](Gopher.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GophersGet

> []Gopher GophersGet(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/scraly/gophers-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GophersAPI.GophersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GophersAPI.GophersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GophersGet`: []Gopher
	fmt.Fprintf(os.Stdout, "Response from `GophersAPI.GophersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGophersGetRequest struct via the builder pattern


### Return type

[**[]Gopher**](Gopher.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

