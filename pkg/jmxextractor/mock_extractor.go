// Copyright © 2021 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jmxextractor

import "github.com/banzaicloud/koperator/api/v1beta1"

type mockJmxExtractor struct{}

func (exp *mockJmxExtractor) ExtractDockerImageAndVersion(brokerId int32, brokerConfig *v1beta1.BrokerConfig,
	clusterImage string, headlessServiceEnabled bool) (*v1beta1.KafkaVersion, error) {
	return &v1beta1.KafkaVersion{Image: clusterImage, Version: "2.7.0"}, nil
}
