/*
   Copyright 2019 Tuxera Oy. All Rights Reserved.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package driver

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func createSession() *session.Session {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"),
	}))
	return sess
}

func TestSubnet(t *testing.T) {

	svc := ec2.New(createSession())

	_, err := getDefaultSubnet(svc)
	if err != nil {
		t.Error("Obtained error: ", err)
	}
}
