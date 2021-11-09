package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2bf30f480d9d88f317b8092d5fd781f6cb4e5f67733d80223d38a86476776bbb "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/tasktriggers"
    i492b49c506e6b22f81715925a69d42d13860980b2c276bb6a340762e9963da02 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/share"
    i8043d023affe03fd84c721217abd77c800e2127dc310e8959351e97c5a7eb2dd "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/restorefactorydefaults"
    i8ba5fa3778ecee53d2bff1ec8a905d0d0378854b1508f3d8180e4d28751e427f "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/getcapabilities"
    ib03455da6d11aedbe84a78d6a60033123f8f15477e489b6577f58e98275daad5 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/shares"
    icc2a686f8838064d00a6ef0a690ee9b85a6492ce9a8ed6563aebc803aa21d09a "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/connectors"
    id0803f4cf58337d206e8f9ac19b5e26be459d54ba598819101c6635a3bfdc49d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/resetdefaults"
    i42ebbf05f8556082dcf91f341a5bdab47970319314509c1ae13698fb9a982e3d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item/tasktriggers/item"
)

// Builds and executes requests for operations under \print\printers\{printer-id}
type PrinterRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type PrinterRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type PrinterRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrinterRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The list of printers registered in the tenant.
type PrinterRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type PrinterRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Printer;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PrinterRequestBuilder) Connectors()(*icc2a686f8838064d00a6ef0a690ee9b85a6492ce9a8ed6563aebc803aa21d09a.ConnectorsRequestBuilder) {
    return icc2a686f8838064d00a6ef0a690ee9b85a6492ce9a8ed6563aebc803aa21d09a.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new PrinterRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrinterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterRequestBuilder) {
    m := &PrinterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printers/{printer_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new PrinterRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrinterRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrinterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrinterRequestBuilderInternal(urlParams, requestAdapter)
}
// The list of printers registered in the tenant.
// Parameters:
//  - options : Options for the request
func (m *PrinterRequestBuilder) CreateDeleteRequestInformation(options *PrinterRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The list of printers registered in the tenant.
// Parameters:
//  - options : Options for the request
func (m *PrinterRequestBuilder) CreateGetRequestInformation(options *PrinterRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The list of printers registered in the tenant.
// Parameters:
//  - options : Options for the request
func (m *PrinterRequestBuilder) CreatePatchRequestInformation(options *PrinterRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The list of printers registered in the tenant.
// Parameters:
//  - options : Options for the request
func (m *PrinterRequestBuilder) Delete(options *PrinterRequestBuilderDeleteOptions)(error) {
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
// The list of printers registered in the tenant.
// Parameters:
//  - options : Options for the request
func (m *PrinterRequestBuilder) Get(options *PrinterRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Printer, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrinter() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Printer), nil
}
// Builds and executes requests for operations under \print\printers\{printer-id}\microsoft.graph.getCapabilities()
func (m *PrinterRequestBuilder) GetCapabilities()(*i8ba5fa3778ecee53d2bff1ec8a905d0d0378854b1508f3d8180e4d28751e427f.GetCapabilitiesRequestBuilder) {
    return i8ba5fa3778ecee53d2bff1ec8a905d0d0378854b1508f3d8180e4d28751e427f.NewGetCapabilitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The list of printers registered in the tenant.
// Parameters:
//  - options : Options for the request
func (m *PrinterRequestBuilder) Patch(options *PrinterRequestBuilderPatchOptions)(error) {
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
func (m *PrinterRequestBuilder) ResetDefaults()(*id0803f4cf58337d206e8f9ac19b5e26be459d54ba598819101c6635a3bfdc49d.ResetDefaultsRequestBuilder) {
    return id0803f4cf58337d206e8f9ac19b5e26be459d54ba598819101c6635a3bfdc49d.NewResetDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) RestoreFactoryDefaults()(*i8043d023affe03fd84c721217abd77c800e2127dc310e8959351e97c5a7eb2dd.RestoreFactoryDefaultsRequestBuilder) {
    return i8043d023affe03fd84c721217abd77c800e2127dc310e8959351e97c5a7eb2dd.NewRestoreFactoryDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) Share()(*i492b49c506e6b22f81715925a69d42d13860980b2c276bb6a340762e9963da02.ShareRequestBuilder) {
    return i492b49c506e6b22f81715925a69d42d13860980b2c276bb6a340762e9963da02.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) Shares()(*ib03455da6d11aedbe84a78d6a60033123f8f15477e489b6577f58e98275daad5.SharesRequestBuilder) {
    return ib03455da6d11aedbe84a78d6a60033123f8f15477e489b6577f58e98275daad5.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrinterRequestBuilder) TaskTriggers()(*i2bf30f480d9d88f317b8092d5fd781f6cb4e5f67733d80223d38a86476776bbb.TaskTriggersRequestBuilder) {
    return i2bf30f480d9d88f317b8092d5fd781f6cb4e5f67733d80223d38a86476776bbb.NewTaskTriggersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.printers.item.taskTriggers.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PrinterRequestBuilder) TaskTriggersById(id string)(*i42ebbf05f8556082dcf91f341a5bdab47970319314509c1ae13698fb9a982e3d.PrintTaskTriggerRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printTaskTrigger_id"] = id
    }
    return i42ebbf05f8556082dcf91f341a5bdab47970319314509c1ae13698fb9a982e3d.NewPrintTaskTriggerRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
