// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package context

import "math/rand/v2"

type Context struct {
	Rand     *rand.Rand
	Current  int
	Total    int
	Progress float64
}
