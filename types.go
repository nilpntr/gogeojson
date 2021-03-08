package gogeojson

import (
	"errors"
)

// TODO List:
// - Add validation to check the type
// - Check all the deepest nested coordinates slices if they hold max a length of 2
// - Check polygon multi and single ring
// - Add validation for max length in not the deepest level of coordinates like for LineString max 2 slices in coordinates

type FeatureType string

const (
	FeatureTypePoint           FeatureType = "Point"
	FeatureTypeLineString      FeatureType = "LineString"
	FeatureTypePolygon         FeatureType = "Polygon"
	FeatureTypeMultiPoint      FeatureType = "MultiPoint"
	FeatureTypeMultiLineString FeatureType = "MultiLineString"
	FeatureTypeMultiPolygon    FeatureType = "MultiPolygon"
)

var (
	ErrCoordinateTooManyValues = errors.New("a slice in the deepest level of coordinates has more then 2 values")
)

// PointField: Holds the value of a Point
type PointField struct {
	Type        FeatureType `json:"type" bson:"type"`
	Coordinates []float64   `json:"coordinates" bson:"coordinates"`
}

// NewPointField: Returns a new PointField with the provided coordinates
func NewPointField(coordinates []float64) *PointField {
	return &PointField{
		Type:        FeatureTypePoint,
		Coordinates: coordinates,
	}
}

// Validate: Returns an error if the deepest level of PointField coordinates has more then 2 values
func (f *PointField) Validate() error {
	if len(f.Coordinates) > 2 {
		return ErrCoordinateTooManyValues
	}
	return nil
}

// LineStringField: Holds the value of a LineString
type LineStringField struct {
	Type        FeatureType `json:"type" bson:"type"`
	Coordinates [][]float64 `json:"coordinates" bson:"coordinates"`
}

// NewLineStringField: Returns a new LineStringField with the provided coordinates
func NewLineStringField(coordinates [][]float64) *LineStringField {
	return &LineStringField{
		Type:        FeatureTypeLineString,
		Coordinates: coordinates,
	}
}

// Validate: Returns an error if the deepest level of LineStringField coordinates has more then 2 values
func (f *LineStringField) Validate() error {
	if len(f.Coordinates) > 2 {
		return ErrCoordinateTooManyValues
	}
	for _, elem := range f.Coordinates {
		if len(elem) > 2 {
			return ErrCoordinateTooManyValues
		}
	}
	return nil
}

// PolygonField: Holds the value of a Polygon
type PolygonField struct {
	Type        FeatureType   `json:"type" bson:"type"`
	Coordinates [][][]float64 `json:"coordinates" bson:"coordinates"`
}

// NewPolygonField: Returns a new PolygonField with the provided coordinates
func NewPolygonField(coordinates [][][]float64) *PolygonField {
	return &PolygonField{
		Type:        FeatureTypePolygon,
		Coordinates: coordinates,
	}
}

// Validate: Returns an error if the deepest level of PolygonField coordinates has more then 2 values
func (f *PolygonField) Validate() error {
	for _, elem := range f.Coordinates {
		for _, subElem := range elem {
			if len(subElem) > 2 {
				return ErrCoordinateTooManyValues
			}
		}
	}
	return nil
}

// MultiPointField: Holds the value of a MultiPoint
type MultiPointField struct {
	Type        FeatureType `json:"type" bson:"type"`
	Coordinates [][]float64 `json:"coordinates" bson:"coordinates"`
}

// NewMultiPointField: Returns a new MultiPointField with the provided coordinates
func NewMultiPointField(coordinates [][]float64) *MultiPointField {
	return &MultiPointField{
		Type:        FeatureTypeMultiPoint,
		Coordinates: coordinates,
	}
}

// Validate: Returns an error if the deepest level of MultiPointField coordinates has more then 2 values
func (f *MultiPointField) Validate() error {
	for _, elem := range f.Coordinates {
		if len(elem) > 2 {
			return ErrCoordinateTooManyValues
		}
	}
	return nil
}

// MultiLineStringField: Holds the value of a MultiLineString
type MultiLineStringField struct {
	Type        FeatureType   `json:"type" bson:"type"`
	Coordinates [][][]float64 `json:"coordinates" bson:"coordinates"`
}

// NewMultiLineStringField: Returns a new MultiLineStringField with the provided coordinates
func NewMultiLineStringField(coordinates [][][]float64) *MultiLineStringField {
	return &MultiLineStringField{
		Type:        FeatureTypeMultiLineString,
		Coordinates: coordinates,
	}
}

// Validate: Returns an error if the deepest level of MultiLineStringField coordinates has more then 2 values
func (f *MultiLineStringField) Validate() error {
	for _, elem := range f.Coordinates {
		for _, subElem := range elem {
			if len(subElem) > 2 {
				return ErrCoordinateTooManyValues
			}
		}
	}
	return nil
}

// MultiPolygonField: Holds the value of a MultiPolygon
type MultiPolygonField struct {
	Type        FeatureType   `json:"type" bson:"type"`
	Coordinates [][][]float64 `json:"coordinates" bson:"coordinates"`
}

// NewMultiPolygonField: Returns a new MultiPolygonField with the provided coordinates
func NewMultiPolygonField(coordinates [][][]float64) *MultiPolygonField {
	return &MultiPolygonField{
		Type:        FeatureTypeMultiPolygon,
		Coordinates: coordinates,
	}
}

// Validate: Returns an error if the deepest level of MultiPolygonField coordinates has more then 2 values
func (f *MultiPolygonField) Validate() error {
	for _, elem := range f.Coordinates {
		for _, subElem := range elem {
			if len(subElem) > 2 {
				return ErrCoordinateTooManyValues
			}
		}
	}
	return nil
}
