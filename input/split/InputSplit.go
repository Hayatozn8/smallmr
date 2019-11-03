package split

type InputSplit interface {
	GetLength() int64
	//getLocations() ([]string, error)
	//getLocationInfo([]SplitLocationInfo, error)
}
