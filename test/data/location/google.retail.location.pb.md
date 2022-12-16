
# google.retail.location.pb

## Comments


```plantuml
package google.retail.location.pb {
struct .Site {
  + Address address
  + Climate climate
  + int32 longitude_degrees
  + int32 longitude_minutes
  + int32 longitude_seconds
  + int32 latitude_degrees
  + int32 latitude_minutes
  + int32 latitude_seconds
  + string latitude_direction_code
  + double altitude_meters
  + VersionID id
  + string site_type
  + PhysicalLocation location
  + string operational_party_id
  + double north_angle_offset
  + string icao_code
  + TimeZone time_zone
  + []google.retail.common.pb.Contact contacts
  + []string location_ids
}
struct .LocationMeasure {
  + string name
  + Distance distance
  + Count count
  + Capacity capacity
  + Area area
  + Weight weight
  + double value
}
struct .Location {
  + VersionID id
  + string location_type
  + string location_function_type
  + string inventory_location_security_type
  + bool stock_ledger_control_flag
  + []LocationMeasure measures
  + string merchandise_hierarchy_id
  + Timestamp effective_date
  + double performance_score
  + []MerchGroup merchandise_groups
  + double angle_offset_degrees
  + string coordinate_x
  + string coordinate_y
  + string coordinate_z
  + int32 ordinal
  + string coordinate_x
  + string coordinate_y
  + string coordinate_z
  + []LocationVertex location_vertices
  + LocationCoordinate location_coordinate
  + []Location children
}
struct .GeoSegmentGroup {
  + VersionID id
  + GeoSegmentGroup type
  + []string geo_segment_ids
}
struct .GeoSegment {
  + VersionID id
  + GeoSegment type
  + string abbreviation
  + string description
  + []Site sites
  + []GeoSegment children
}
}

```

