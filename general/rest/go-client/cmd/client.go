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
	"github.com/apache/dubbo-samples/golang/general/rest/go-client/pkg"
	"time"
)

import (
	"github.com/dubbogo/gost/log"
)

import (
	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	"github.com/apache/dubbo-go/config"
	_ "github.com/apache/dubbo-go/protocol/rest"
	_ "github.com/apache/dubbo-go/registry/protocol"

	_ "github.com/apache/dubbo-go/filter/filter_impl"

	_ "github.com/apache/dubbo-go/cluster/cluster_impl"
	_ "github.com/apache/dubbo-go/cluster/loadbalance"
	_ "github.com/apache/dubbo-go/registry/zookeeper"
)

var (
	userProvider = new(pkg.UserProvider)
)

// need to setup environment variable "CONF_CONSUMER_FILE_PATH" to "conf/client.yml" before run
func main() {
	config.SetConsumerService(userProvider)
	config.Load()
	time.Sleep(3 * time.Second)

	gxlog.CInfo("\n\ntest")
	test()
}

func test() {
	gxlog.CInfo("\n\n\nstart to test rest")
	user := &pkg.User{}
	err := userProvider.GetUser(context.TODO(), []interface{}{"A003"}, user)
	if err != nil {
		panic(err)
	}
	gxlog.CInfo("response result: %v", user)

	gxlog.CInfo("\n\n\nstart to test rest - GetUser0")
	ret, err := userProvider.GetUser0("A003", "Moorse中文", 30)
	if err != nil {
		panic(err)
	}
	gxlog.CInfo("response result: %v", ret)

	gxlog.CInfo("\n\n\nstart to test rest - GetUsers")
	ret1, err := userProvider.GetUsers([]interface{}{&pkg.User{ID: "A002"}})
	if err != nil {
		panic(err)
	}
	gxlog.CInfo("response result: %v", ret1)

	gxlog.CInfo("\n\n\nstart to test rest - GetUser3")
	err = userProvider.GetUser3()
	if err != nil {
		panic(err)
	}
	gxlog.CInfo("succ!")

	gxlog.CInfo("\n\n\nstart to test rest illegal method")
	err = userProvider.GetUser1(context.TODO(), []interface{}{"A003"}, user)
	if err == nil {
		panic("err is nil")
	}
	gxlog.CInfo("error: %v", err)
}
