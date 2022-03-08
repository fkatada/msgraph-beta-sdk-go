package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivilegedRoleable 
type PrivilegedRoleable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignments()([]PrivilegedRoleAssignmentable)
    GetName()(*string)
    GetSettings()(PrivilegedRoleSettingsable)
    GetSummary()(PrivilegedRoleSummaryable)
    SetAssignments(value []PrivilegedRoleAssignmentable)()
    SetName(value *string)()
    SetSettings(value PrivilegedRoleSettingsable)()
    SetSummary(value PrivilegedRoleSummaryable)()
}
