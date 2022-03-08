package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemPhoneable 
type ItemPhoneable interface {
    ItemFacetable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetNumber()(*string)
    GetType()(*PhoneType)
    SetDisplayName(value *string)()
    SetNumber(value *string)()
    SetType(value *PhoneType)()
}
