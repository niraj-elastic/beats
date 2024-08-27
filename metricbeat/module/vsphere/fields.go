// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package vsphere

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "vsphere", asset.ModuleFieldsPri, AssetVsphere); err != nil {
		panic(err)
	}
}

// AssetVsphere returns asset data.
// This is the base64 encoded zlib format compressed contents of module/vsphere.
func AssetVsphere() string {
	return "eJzsl8tu2zwQhfd+ikH28QN48QNBfqRZNGmBNN0aNDWy2FAcgTOy4T59QUlWrYvtOKbaTbkwEIo65wsvZ6hbeMPdAjZcZOhxBiBGLC7gZvNS9dzMABJk7U0hhtwC/psBADRPIaektOE1jxYV4wLWagaQGrQJL6qht+BUjocWocmuCIM9lUXTM+LSFToU07ZkQd/2jwmG1mKtUNRB/6hZ3YbSfYhDkESJYiGP8/A3d8bsod5wtyWf9J6dQAjts2EBSkFZC5Ih/L83qowZFDNpowQT2BrJqjEN+vwMqabSySipJbe+DPO5zFfoA+hv/SP2vNTkUrOeqyQ3zIbcXJMTT3aOTq0s9meoZloRWVTuMqwHq9YgmRJIUNDnxiHDNkPJ0AOLN1qgxYAGAwxDQ3L2f5iImMC4xGgl2NKSB0fSnrjHO0hRSenxLG5GLBMtdpDmE66RD4PtHYZHYvnoOXAoW/JvE03MXv2k9cRR8VzbXDJBg5A4m6wnmMZEuvVh304la/iNN0fd+Bw1TDnox7N8MBZ5x4I5DITbOqYKpY3s5kKi7Hy1kyMb4/I9+S0oQqUYNkdY9rGFCS0lnytZwNB+wJl6xKiYDx4xOmXJmESlfGVMpqEs9HgIsVYWk2VqSfUHvIe1QK/RyXtpm+Ht00EwhFS/JhN67//tOGgLyPgKFWW9OHn2M+YGuv/6CsbBU0/10LbOgXi+dQq8w7g62PF8q2N9xjbHnPxUh/WpEg/2Y9LnT2kDN1UuR8KbJo6vhWvuOcvI95zOvWaQUBvjpVQ2Vzoz7qr7y1Gly1Orugyb8c+EjwdXT7BjFn4mysp9Mell+VQZ/b1eBniq1+F4XFPEPdY3/VKgV2LcGl7qq9y/ojFd0bjbKGPDB+0llWNdIstk9YNS+BQMrk7qirU6oROi1ukQpeTFn9ZO4Ys2r9UWjA97WAevZ9UlC+XLulKMEtLqBw6+BOrO5RVxdl8Zw4jxH63VvwIAAP//7RO+5g=="
}
