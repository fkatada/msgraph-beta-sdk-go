package playprompt

import (
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// PlayPromptResponseable 
type PlayPromptResponseable interface {
    AdditionalDataHolder
    Parsable
    GetPlayPromptOperation()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlayPromptOperationable)
    SetPlayPromptOperation(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlayPromptOperationable)()
}
