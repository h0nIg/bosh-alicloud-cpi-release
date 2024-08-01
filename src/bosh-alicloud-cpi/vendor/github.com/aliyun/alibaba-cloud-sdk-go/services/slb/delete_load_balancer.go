package slb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteLoadBalancer invokes the slb.DeleteLoadBalancer API synchronously
func (client *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
	response = CreateDeleteLoadBalancerResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLoadBalancerWithChan invokes the slb.DeleteLoadBalancer API asynchronously
func (client *Client) DeleteLoadBalancerWithChan(request *DeleteLoadBalancerRequest) (<-chan *DeleteLoadBalancerResponse, <-chan error) {
	responseChan := make(chan *DeleteLoadBalancerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLoadBalancer(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteLoadBalancerWithCallback invokes the slb.DeleteLoadBalancer API asynchronously
func (client *Client) DeleteLoadBalancerWithCallback(request *DeleteLoadBalancerRequest, callback func(response *DeleteLoadBalancerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLoadBalancerResponse
		var err error
		defer close(result)
		response, err = client.DeleteLoadBalancer(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DeleteLoadBalancerRequest is the request struct for api DeleteLoadBalancer
type DeleteLoadBalancerRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EnableEipReserve     string           `position:"Query" name:"EnableEipReserve"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

// DeleteLoadBalancerResponse is the response struct for api DeleteLoadBalancer
type DeleteLoadBalancerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLoadBalancerRequest creates a request to invoke DeleteLoadBalancer API
func CreateDeleteLoadBalancerRequest() (request *DeleteLoadBalancerRequest) {
	request = &DeleteLoadBalancerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DeleteLoadBalancer", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLoadBalancerResponse creates a response to parse from DeleteLoadBalancer response
func CreateDeleteLoadBalancerResponse() (response *DeleteLoadBalancerResponse) {
	response = &DeleteLoadBalancerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
