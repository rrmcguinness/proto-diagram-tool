
# google.retail.common.pb

## Comments


```plantuml
@startuml

package google.retail.common.pb {

struct .I18nImage {
  + Locale locale
  + []Image images
}

struct .Currency {
  + Currency code
  + Number value
  + RoundingRule rounding_rule
}
struct .CountrySubdivision {
  + string id
  + string country_code
  + string subdivision_name
  + string code
}
struct .VersionID {
  + string id
  + int64 version
  + Timestamp created
  + Timestamp effective
}
struct .BusinessKey {
  + string name
  + []string value
}
struct .I18nText {
  + Locale locale
  + string value
}
struct .I18nResource {
  + Locale locale
  + string url
}
struct .Image {
  + string url
  + string alt
  + int32 height
  + int32 width
  + string classifier
  + string frame
}
struct .VersionIDRequest {
  + string id
  + int32 version
}
struct .VersionIDEffectiveRequest {
  + VersionIDRequest id
  + Timestamp on_or_before
}
struct .AuditRecordSearchRequest {
  + AuditRecord criteria
  + TimeBoundRequest bounds
  + string action
  + string principal
}
struct .RoundingRule {
  + int32 relevant_decimal
  + bool trim_insignificant_digits
  + bool round_half_up
}
struct .Contact {
  + string platform
  + string platform_id
  + string id
  + string contact_purpose
  + ContactMethod contact_method
  + Timestamp effective_date
  + Timestamp expiration_date
  + Address address
  + string email_address
  + string telephone
  + string website
  + []SocialMediaNetworkIdentity social_network_id
  + string status
}
struct .StatusResponse {
  + Timestamp ts
  + Type type
  + string id
  + string message
  + Struct payload
}
struct .AuditRecord {
  + string id
  + Timestamp created
  + string action
  + string context
  + string principal
  + map<string, string> context_variables
}
struct .NamedMeasure {
  + string name
  + Area area
  + Capacity capacity
  + Distance distance
  + Weight weight
  + Count count
  + Packaging package
  + Time time
  + double size
}
struct .Country {
  + string id
  + string name
  + string alpha2
  + string alpha3
  + string code
  + string iso2
  + string region
  + string sub_region
  + string intermediate_region
  + string region_code
  + string sub_region_code
  + string intermediate_region_code
}
struct .ICAOCode {
  + string id
  + string country_id
  + string country_subdivision_id
  + string location_name
  + string station_name
  + string type
  + string station_key
  + string status
  + string icao
  + string national_id
  + string wmo
  + string ghcn
  + string special
  + string latitude
  + string longitude
  + string elevation_in_meters
  + string time_zone
}
struct .TimeBoundRequest {
  + Timestamp start_date
  + Timestamp end_date
}
struct .Number {
  + int32 whole
  + int32 decimal
}
struct .Address {
  + string geo_segment_id
  + string line_1
  + string line_2
  + string line_3
  + string line_4
  + string city
  + string territory
  + string postal_code
  + string iso_3166_2_country_sub_division_code
}
struct .IDRequest {
  + string id
}
}
@enduml
```

