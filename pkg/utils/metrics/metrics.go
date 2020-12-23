package metrics

import (
	"github.com/emicklei/go-restful"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	apimachineryversion "k8s.io/apimachinery/pkg/version"
	compbasemetrics "k8s.io/component-base/metrics"
	ksVersion "kubesphere.io/kubesphere/pkg/version"
	"net/http"
)

var (
	Defaults        DefaultMetrics
	defaultRegistry compbasemetrics.KubeRegistry
	// MustRegister registers registerable metrics but uses the defaultRegistry, panic upon the first registration that causes an error
	MustRegister func(...compbasemetrics.Registerable)
	// Register registers a collectable metric but uses the defaultRegistry
	Register func(compbasemetrics.Registerable) error

	RawMustRegister func(...prometheus.Collector)

)

func init() {
	defaultRegistry = compbasemetrics.NewKubeRegistry()
	MustRegister = defaultRegistry.MustRegister
	Register = defaultRegistry.Register
	RawMustRegister = defaultRegistry.RawMustRegister

	RawMustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	RawMustRegister(prometheus.NewGoCollector())
}

// DefaultMetrics installs the default prometheus metrics handler
type DefaultMetrics struct{}

// Install adds the DefaultMetrics handler
func (m DefaultMetrics) Install(c *restful.Container) {
	c.Handle("/kapis/metrics", Handler())
}

//Overwrite version.Get
func versionGet() apimachineryversion.Info {
	info := ksVersion.Get()
	return apimachineryversion.Info{
		Major:        info.GitMajor,
		Minor:        info.GitMinor,
		GitVersion:   info.GitVersion,
		GitCommit:    info.GitCommit,
		GitTreeState: info.GitTreeState,
		BuildDate:    info.BuildDate,
		GoVersion:    info.GoVersion,
		Compiler:     info.Compiler,
		Platform:     info.Platform,
	}
}

// Handler returns an HTTP handler for the DefaultGatherer. It is
// already instrumented with InstrumentHandler (using "prometheus" as handler
// name).
func Handler() http.Handler {
	return promhttp.InstrumentMetricHandler(prometheus.DefaultRegisterer, promhttp.HandlerFor(defaultRegistry, promhttp.HandlerOpts{}))
}
