package reporter

type TraceReporter interface {
	ReportTrace(traceBean TraceBean)
}
