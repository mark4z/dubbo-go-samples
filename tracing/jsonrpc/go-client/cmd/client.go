/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"time"
)

import (
	"github.com/apache/dubbo-samples/golang/tracing/jsonrpc/go-client/pkg"
	gxlog "github.com/dubbogo/gost/log"
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"

	_ "github.com/apache/dubbo-go/cluster/cluster_impl"
	_ "github.com/apache/dubbo-go/cluster/loadbalance"
	"github.com/apache/dubbo-go/common/logger"
	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	"github.com/apache/dubbo-go/config"
	_ "github.com/apache/dubbo-go/filter/filter_impl"
	_ "github.com/apache/dubbo-go/protocol/jsonrpc"
	_ "github.com/apache/dubbo-go/registry/protocol"
	_ "github.com/apache/dubbo-go/registry/zookeeper"
)

// need to setup environment variable "CONF_CONSUMER_FILE_PATH" to "conf/client.yml" before run
func main() {
	userProvider := new(pkg.UserProvider)
	config.SetConsumerService(userProvider)
	config.Load()
	initZipkin()
	time.Sleep(3 * time.Second)

	gxlog.CInfo("\n\n\nstart to test jsonrpc")
	user := &pkg.JsonRPCUser{}

	// FIXME: client side doesn't work
	ctx := context.WithValue(context.Background(), "MyTracing", "Tracing123")
	span, spanCtx := opentracing.StartSpanFromContext(ctx, "User-Test-Service")
	defer span.Finish()

	err := userProvider.GetUser(spanCtx, []interface{}{"A003"}, user)
	if err != nil {
		panic(err)
	}
	gxlog.CInfo("response result: %v", user)
}

func initZipkin() {
	// set up a span reporter
	reporter := zipkinhttp.NewReporter("http://localhost:9411/api/v2/spans")

	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint("client-service", "myservice.mydomain.com:80")
	if err != nil {
		logger.Errorf("unable to create local endpoint: %+v\n", err)
	}

	// initialize our tracer
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		logger.Errorf("unable to create tracer: %+v\n", err)
	}

	// use zipkin-go-opentracing to wrap our tracer
	tracer := zipkinot.Wrap(nativeTracer)

	// optionally set as Global OpenTracing tracer instance
	opentracing.SetGlobalTracer(tracer)
}
