// Code generated by gowrap. DO NOT EDIT.
// template: https://raw.githubusercontent.com/hexdigest/gowrap/6c8f05695fec23df85903a8da0af66ac414e2a63/templates/opentelemetry
// gowrap: http://github.com/hexdigest/gowrap

package traced

import (
	"context"

	"github.com/dell/goobjectscale/pkg/client/api"
	"github.com/dell/goobjectscale/pkg/client/model"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// TenantsInterfaceWithTracing implements api.TenantsInterface interface instrumented with opentracing spans
type TenantsInterfaceWithTracing struct {
	api.TenantsInterface
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewTenantsInterfaceWithTracing returns TenantsInterfaceWithTracing
func NewTenantsInterfaceWithTracing(base api.TenantsInterface, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) TenantsInterfaceWithTracing {
	d := TenantsInterfaceWithTracing{
		TenantsInterface: base,
		_instance:        instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// Create implements api.TenantsInterface
func (_d TenantsInterfaceWithTracing) Create(ctx context.Context, payload model.TenantCreate) (tp1 *model.Tenant, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "api.TenantsInterface.Create")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":     ctx,
				"payload": payload}, map[string]interface{}{
				"tp1": tp1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.TenantsInterface.Create(ctx, payload)
}

// Delete implements api.TenantsInterface
func (_d TenantsInterfaceWithTracing) Delete(ctx context.Context, name string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "api.TenantsInterface.Delete")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":  ctx,
				"name": name}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.TenantsInterface.Delete(ctx, name)
}

// DeleteQuota implements api.TenantsInterface
func (_d TenantsInterfaceWithTracing) DeleteQuota(ctx context.Context, name string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "api.TenantsInterface.DeleteQuota")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":  ctx,
				"name": name}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.TenantsInterface.DeleteQuota(ctx, name)
}

// Get implements api.TenantsInterface
func (_d TenantsInterfaceWithTracing) Get(ctx context.Context, name string, params map[string]string) (tp1 *model.Tenant, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "api.TenantsInterface.Get")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"name":   name,
				"params": params}, map[string]interface{}{
				"tp1": tp1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.TenantsInterface.Get(ctx, name, params)
}

// GetQuota implements api.TenantsInterface
func (_d TenantsInterfaceWithTracing) GetQuota(ctx context.Context, name string, params map[string]string) (tp1 *model.TenantQuota, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "api.TenantsInterface.GetQuota")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"name":   name,
				"params": params}, map[string]interface{}{
				"tp1": tp1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.TenantsInterface.GetQuota(ctx, name, params)
}

// List implements api.TenantsInterface
func (_d TenantsInterfaceWithTracing) List(ctx context.Context, params map[string]string) (tp1 *model.TenantList, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "api.TenantsInterface.List")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"params": params}, map[string]interface{}{
				"tp1": tp1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.TenantsInterface.List(ctx, params)
}

// SetQuota implements api.TenantsInterface
func (_d TenantsInterfaceWithTracing) SetQuota(ctx context.Context, name string, payload model.TenantQuotaSet) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "api.TenantsInterface.SetQuota")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":     ctx,
				"name":    name,
				"payload": payload}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.TenantsInterface.SetQuota(ctx, name, payload)
}

// Update implements api.TenantsInterface
func (_d TenantsInterfaceWithTracing) Update(ctx context.Context, payload model.TenantUpdate, name string) (err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "api.TenantsInterface.Update")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":     ctx,
				"payload": payload,
				"name":    name}, map[string]interface{}{
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.TenantsInterface.Update(ctx, payload, name)
}
