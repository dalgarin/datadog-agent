package testutil

import (
	"math/rand"
	"time"

	"github.com/DataDog/datadog-agent/pkg/trace/pb"
)

type SpanConfig struct {
	// MinTags specifies the minimum number of tags this span should have.
	MinTags int
	// MaxTags specifies the maximum number of tags this span should have.
	MaxTags int
}

type TraceConfig struct {
	// MinSpans specifies the minimum number of spans per trace.
	MinSpans int
	// MaxSpans specifies the maximum number of spans per trace.
	MaxSpans int
}

// GeneratePayload generates a new payload.
func GeneratePayload(n int, tc TraceConfig, sc SpanConfig) pb.Traces {
	if n == 0 {
		return pb.Traces{}
	}
	if tc.MinSpans == 0 {
		tc.MinSpans = 1
	}
	if tc.MaxSpans < tc.MinSpans {
		tc.MaxSpans = tc.MinSpans
	}
	out := make(pb.Traces, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, GenerateTrace(tc, sc))
	}
	return out
}

// GenerateTrace generates a valid trace using the given config.
func GenerateTrace(tc TraceConfig, sc SpanConfig) pb.Trace {
	n := tc.MinSpans
	if tc.MaxSpans > tc.MinSpans {
		n += rand.Intn(tc.MaxSpans - tc.MinSpans)
	}
	t := make(pb.Trace, 0, n)
	var (
		maxd int64
		root *pb.Span
	)
	for i := 0; i < n; i++ {
		s := GenerateSpan(sc)
		if s.Duration > maxd {
			root = s
			maxd = s.Duration
		}
		t = append(t, s)
	}
	for _, span := range t {
		if span == root {
			continue
		}
		span.TraceID = root.TraceID
		span.ParentID = root.SpanID
		span.Start = root.Start + rand.Int63n(root.Duration-span.Duration)
	}
	return t
}

// GenerateSpan generates a random root span with all fields filled in.
func GenerateSpan(c SpanConfig) *pb.Span {
	pickString := func(all []string) string { return all[rand.Intn(len(all))] }
	id := uint64(rand.Int63())
	duration := 1 + rand.Int63n(1_000_000_000) // between 1ns and 1s
	s := &pb.Span{
		Service:  pickString(services),
		Name:     pickString(names),
		Resource: pickString(resources),
		TraceID:  id,
		SpanID:   id,
		ParentID: 0,
		Start:    time.Now().UnixNano() - duration,
		Duration: duration,
		Error:    int32(rand.Intn(2)),
		Meta:     make(map[string]string),
		Metrics:  make(map[string]float64),
		Type:     pickString(types),
	}
	if c.MaxTags == 0 {
		return s
	}
	ntags := c.MinTags + rand.Intn(c.MaxTags-c.MinTags)
	nmetrics := 0
	if ntags > 4 {
		// make 25% of tags Metrics when we have more than 4
		nmetrics = ntags / 4
	}
	for i := 0; i < ntags-nmetrics; i++ {
		for k := range metas {
			s.Meta[k] = pickString(metas[k])
			break
		}
	}
	for i := 0; i < nmetrics; i++ {
		s.Metrics[pickString(metrics)] = rand.Float64()
	}
	return s
}
