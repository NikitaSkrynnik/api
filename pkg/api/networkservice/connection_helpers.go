// Copyright (c) 2018-2021 VMware, Inc.
//
// Copyright (c) 2021 Doc.ai and/or its affiliates.
//
// Copyright (c) 2022 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package networkservice

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Equals returns if connection equals given connection
func (x *Connection) Equals(connection protoreflect.ProtoMessage) bool {
	// use as proto.Message
	return proto.Equal(x, connection)
}

// Clone clones connection
func (x *Connection) Clone() *Connection {
	// use as proto.Message
	return proto.Clone(x).(*Connection)
}

// GetCurrentPathSegment - Get the current path segment of the connection
func (x *Connection) GetCurrentPathSegment() *PathSegment {
	if x == nil {
		return nil
	}
	if len(x.GetPath().GetPathSegments()) == 0 {
		return nil
	}
	if int(x.GetPath().GetIndex()) > len(x.GetPath().GetPathSegments())-1 {
		return nil
	}
	return x.GetPath().GetPathSegments()[x.GetPath().GetIndex()]
}

// GetPrevPathSegment - Get the previous path segment of the connection
func (x *Connection) GetPrevPathSegment() *PathSegment {
	if x == nil {
		return nil
	}
	if len(x.GetPath().GetPathSegments()) == 0 {
		return nil
	}
	if int(x.GetPath().GetIndex()) == 0 {
		return nil
	}
	if int(x.GetPath().GetIndex())-1 > len(x.GetPath().GetPathSegments()) {
		return nil
	}
	return x.GetPath().GetPathSegments()[x.GetPath().GetIndex()-1]
}

// GetNextPathSegment - Get the next path segment of the connection
func (x *Connection) GetNextPathSegment() *PathSegment {
	if x == nil {
		return nil
	}
	if len(x.GetPath().GetPathSegments()) == 0 {
		return nil
	}
	if len(x.GetPath().GetPathSegments())-1 < int(x.GetPath().GetIndex())+1 {
		return nil
	}
	return x.GetPath().GetPathSegments()[x.GetPath().GetIndex()+1]
}
