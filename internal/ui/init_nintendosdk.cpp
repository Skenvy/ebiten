// Copyright 2023 The Ebitengine Authors
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

//go:build nintendosdk

// The actual implementation will be provided by github.com/hajimehoshi/uwagaki.

#include "init_nintendosdk.h"

extern "C" NativeWindowType ebitengine_Initialize() { return 0; }

extern "C" void ebitengine_InitializeProfiler() {}

extern "C" void ebitengine_RecordProfilerHeartbeat() {}
