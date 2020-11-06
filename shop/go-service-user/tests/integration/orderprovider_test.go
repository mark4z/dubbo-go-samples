// +build integration

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

package integration

import (
	"context"
	"github.com/apache/dubbo-samples/golang/shop/go-service-user/pkg"
	"testing"
)
import (
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	order := &pkg.Order{}
	err := orderProvider.GetOrder(context.TODO(), []interface{}{"A001"}, order)
	assert.Nil(t, err)
	assert.Equal(t, "A001", order.Name)
	assert.Equal(t, "A001", order.Id)
	assert.Equal(t, 200, order.Price)
	assert.Equal(t, "A001", order.Product.Id)
	assert.Equal(t, 100, order.Product.Quantity)
	assert.Equal(t, 100, order.Product.Price)
}
