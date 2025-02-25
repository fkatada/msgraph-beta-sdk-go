package models
import (
    "math"
    "strings"
)
// Flag enum to determine whether to delay software updates for macOS.
type MacOSSoftwareUpdateDelayPolicy int

const (
    // Software update delays will not be enforced.
    NONE_MACOSSOFTWAREUPDATEDELAYPOLICY = 1
    // Force delays for OS software updates.
    DELAYOSUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY = 2
    // Force delays for app software updates.
    DELAYAPPUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY = 4
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_MACOSSOFTWAREUPDATEDELAYPOLICY = 8
    // Force delays for major OS software updates.
    DELAYMAJOROSUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY = 16
)

func (i MacOSSoftwareUpdateDelayPolicy) String() string {
    var values []string
    options := []string{"none", "delayOSUpdateVisibility", "delayAppUpdateVisibility", "unknownFutureValue", "delayMajorOsUpdateVisibility"}
    for p := 0; p < 5; p++ {
        mantis := MacOSSoftwareUpdateDelayPolicy(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseMacOSSoftwareUpdateDelayPolicy(v string) (any, error) {
    var result MacOSSoftwareUpdateDelayPolicy
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_MACOSSOFTWAREUPDATEDELAYPOLICY
            case "delayOSUpdateVisibility":
                result |= DELAYOSUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
            case "delayAppUpdateVisibility":
                result |= DELAYAPPUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_MACOSSOFTWAREUPDATEDELAYPOLICY
            case "delayMajorOsUpdateVisibility":
                result |= DELAYMAJOROSUPDATEVISIBILITY_MACOSSOFTWAREUPDATEDELAYPOLICY
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeMacOSSoftwareUpdateDelayPolicy(values []MacOSSoftwareUpdateDelayPolicy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MacOSSoftwareUpdateDelayPolicy) isMultiValue() bool {
    return true
}
