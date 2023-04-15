package grpc

import (
	"context"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"go-pkg/logger"
)

// FieldSpec provides a mapping between a metadata context field and a logging field.
type FieldSpec struct {
	metadataKey string
	tagKey      string
}

// ExtractMetadataField creates a specification for converting gRPC metadata fields
// to log tags.
func ExtractMetadataField(metadataKey, tagKey string) FieldSpec {
	return FieldSpec{metadataKey, tagKey}
}

type extractMetadata struct {
	fields []FieldSpec
}

func (r *extractMetadata) ServerReporter(ctx context.Context, meta interceptors.CallMeta) (interceptors.Reporter, context.Context) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fields := []any{}
		logContext := logger.With()
		for _, field := range r.fields {
			value, ok := md[field.metadataKey]
			if ok {
				joinedValue := strings.Join(value, ",")
				fields = append(fields, field.tagKey)
				fields = append(fields, joinedValue)
				logContext = logContext.Str(field.tagKey, joinedValue)
			}
		}
		logContext = logContext.Str("grpc.method", meta.Method).Str("grpc.service", meta.Service)
		ctx = logging.InjectFields(ctx, fields)
		loggerForContext := logContext.Logger()
		ctx = loggerForContext.WithContext(ctx)
	}

	return interceptors.NoopReporter{}, ctx
}

// UnaryServerInterceptor creates an interceptor for extracting fields from requests
// and setting them as log tags.
func UnaryServerInterceptor(fields ...FieldSpec) grpc.UnaryServerInterceptor {
	return interceptors.UnaryServerInterceptor(&extractMetadata{fields})
}

// StreamServerInterceptor creates an interceptor for extracting fields from requests
// and setting them as log tags.
func StreamServerInterceptor(fields ...FieldSpec) grpc.StreamServerInterceptor {
	return interceptors.StreamServerInterceptor(&extractMetadata{fields})
}
