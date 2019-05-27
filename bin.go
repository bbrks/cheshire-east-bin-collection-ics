package main

const (
	Unknown = Bin(iota)
	GeneralWaste
	MixedRecycling
	GardenWaste
)

type Bin uint8

// String returns the name for the given Bin
func (b Bin) String() string {
	switch b {
	case GeneralWaste:
		return "Black Bin (General Waste)"
	case MixedRecycling:
		return "Grey Bin (Mixed Recycling)"
	case GardenWaste:
		return "Brown Bin (Garden Waste)"
	}
	return "Unknown Bin"
}

// Color returns the CSS3 color name for the given Bin
func (b Bin) Color() string {
	switch b {
	case GeneralWaste:
		return "black"
	case MixedRecycling:
		return "silver"
	case GardenWaste:
		return "saddlebrown"
	}
	return ""
}
