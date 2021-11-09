package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i43e25b1a2e7db024aa2c8195c093b9f95136385066d5a6b0e1a1718ccbb4941a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/tasks"
    i44c0c086502e5c42f8c0b70140fda98b421c68cee2ad57dce93ec3e17b4a1a9a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/singlevalueextendedproperties"
    i99568de587864405cbb7ec5179efa3e533e234570ca91e410212ec1d04fe089c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/multivalueextendedproperties"
    i4c149284b5e78d79fdca30504030b7dd8696f7b3a114850f0c0c50352668aeb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/tasks/item"
    i645e01811e5e9cf150ad7755ded1c4cca58e15173c57613b642c0c878c5c4224 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/multivalueextendedproperties/item"
    i7f227a30fb4ec3965178563ee0f1c3b874776270d21828095647c10a35379e25 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/singlevalueextendedproperties/item"
)

// Builds and executes requests for operations under \me\outlook\taskFolders\{outlookTaskFolder-id}
type OutlookTaskFolderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type OutlookTaskFolderRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type OutlookTaskFolderRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OutlookTaskFolderRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get taskFolders from me
type OutlookTaskFolderRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type OutlookTaskFolderRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolder;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new OutlookTaskFolderRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOutlookTaskFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskFolderRequestBuilder) {
    m := &OutlookTaskFolderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/outlook/taskFolders/{outlookTaskFolder_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new OutlookTaskFolderRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOutlookTaskFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookTaskFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property taskFolders for me
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskFolderRequestBuilder) CreateDeleteRequestInformation(options *OutlookTaskFolderRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get taskFolders from me
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskFolderRequestBuilder) CreateGetRequestInformation(options *OutlookTaskFolderRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Update the navigation property taskFolders in me
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskFolderRequestBuilder) CreatePatchRequestInformation(options *OutlookTaskFolderRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete navigation property taskFolders for me
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskFolderRequestBuilder) Delete(options *OutlookTaskFolderRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get taskFolders from me
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskFolderRequestBuilder) Get(options *OutlookTaskFolderRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolder, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOutlookTaskFolder() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolder), nil
}
func (m *OutlookTaskFolderRequestBuilder) MultiValueExtendedProperties()(*i99568de587864405cbb7ec5179efa3e533e234570ca91e410212ec1d04fe089c.MultiValueExtendedPropertiesRequestBuilder) {
    return i99568de587864405cbb7ec5179efa3e533e234570ca91e410212ec1d04fe089c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskFolders.item.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OutlookTaskFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i645e01811e5e9cf150ad7755ded1c4cca58e15173c57613b642c0c878c5c4224.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i645e01811e5e9cf150ad7755ded1c4cca58e15173c57613b642c0c878c5c4224.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Update the navigation property taskFolders in me
// Parameters:
//  - options : Options for the request
func (m *OutlookTaskFolderRequestBuilder) Patch(options *OutlookTaskFolderRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *OutlookTaskFolderRequestBuilder) SingleValueExtendedProperties()(*i44c0c086502e5c42f8c0b70140fda98b421c68cee2ad57dce93ec3e17b4a1a9a.SingleValueExtendedPropertiesRequestBuilder) {
    return i44c0c086502e5c42f8c0b70140fda98b421c68cee2ad57dce93ec3e17b4a1a9a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskFolders.item.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OutlookTaskFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7f227a30fb4ec3965178563ee0f1c3b874776270d21828095647c10a35379e25.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i7f227a30fb4ec3965178563ee0f1c3b874776270d21828095647c10a35379e25.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskFolderRequestBuilder) Tasks()(*i43e25b1a2e7db024aa2c8195c093b9f95136385066d5a6b0e1a1718ccbb4941a.TasksRequestBuilder) {
    return i43e25b1a2e7db024aa2c8195c093b9f95136385066d5a6b0e1a1718ccbb4941a.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskFolders.item.tasks.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OutlookTaskFolderRequestBuilder) TasksById(id string)(*i4c149284b5e78d79fdca30504030b7dd8696f7b3a114850f0c0c50352668aeb9.OutlookTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask_id"] = id
    }
    return i4c149284b5e78d79fdca30504030b7dd8696f7b3a114850f0c0c50352668aeb9.NewOutlookTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
