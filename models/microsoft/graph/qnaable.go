package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/search"
)

// Qnaable 
type Qnaable interface {
    SearchAnswerable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAvailabilityEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAvailabilityStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGroupIds()([]string)
    GetIsSuggested()(*bool)
    GetKeywords()(AnswerKeywordable)
    GetLanguageTags()([]string)
    GetPlatforms()([]DevicePlatformType)
    GetState()(*id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState)
    GetTargetedVariations()([]AnswerVariantable)
    SetAvailabilityEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAvailabilityStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGroupIds(value []string)()
    SetIsSuggested(value *bool)()
    SetKeywords(value AnswerKeywordable)()
    SetLanguageTags(value []string)()
    SetPlatforms(value []DevicePlatformType)()
    SetState(value *id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState)()
    SetTargetedVariations(value []AnswerVariantable)()
}
