package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AdministrativeUnitable 
type AdministrativeUnitable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    DirectoryObjectable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExtensions()([]Extensionable)
    GetMembers()([]DirectoryObjectable)
    GetScopedRoleMembers()([]ScopedRoleMembershipable)
    GetVisibility()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExtensions(value []Extensionable)()
    SetMembers(value []DirectoryObjectable)()
    SetScopedRoleMembers(value []ScopedRoleMembershipable)()
    SetVisibility(value *string)()
}
