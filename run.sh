#! /bin/bash
# Copyright 2021 CloudWeGo Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# specify test directories in tests/
tests+=("crosstest/dubbo2kitex")
tests+=("crosstest/kitex2dubbo")
tests+=("crosstest/kitex2kitex")
tests+=("registry_test/registry")
tests+=("registry_test/resolver")

integrate_test() {
  sub_dir=$1
  full_path=$(pwd)/tests/$sub_dir
  go test --count=1 "$full_path"
  return $?
}

# change codec-dubbo version for CI
LOCAL_REPO=$1
cd ./code/kitex || exit
if [[ -n $LOCAL_REPO ]]; then
  go mod edit -replace github.com/kitex-contrib/codec-dubbo="${LOCAL_REPO}"
fi
go mod tidy
cd - || exit
if [[ -n $LOCAL_REPO ]]; then
  go mod edit -replace github.com/kitex-contrib/codec-dubbo="${LOCAL_REPO}"
fi
go mod tidy

# run zookeeper
make run

# run tests
for ((i = 0; i < ${#tests[*]}; i++)); do
  integrate_test "${tests[i]}"
  res=$?
  if [ $res -ne 0 ]; then
    make stop
    exit $res
  fi
done

make stop
