package nodelocal

import "github.com/gocrane-io/crane/pkg/utils"

type nodeLocalCollector interface {
	name() string
	collect() (map[string]utils.TimeSeries, error)
}
