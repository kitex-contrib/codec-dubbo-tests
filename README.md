# codec-dubbo-tests

This repository contains integrated tests for [codec-dubbo](https://github.com/kitex-contrib/codec-dubbo).

## Integrated Tests Categories

All the tests are stored in **tests/** directory.

1. crosstest/dubbo2kitex
2. crosstest/kitex2dubbo
3. crosstest/kitex2kitex
4. registry_test/registry
5. registry_test/resolver

These tests categories correspond to the test subdirectories in **run.sh**.

### Add Test Cases in Existing Integrated Tests Categories

Please refer to the concrete category for detail.

### Add New Test Category

1. Create sub-dir in **tests/** directory. **Make sure that this sub-dir could be tested
by ```go test``` directly.**
2. Add this sub-dir to **tests** variable in **run.sh**

## Local Test

### Modify Client/Server Code

If you need to modify Client/Server Code, please refer to code/ directory. For the kitex part, 
please refer to [codec-dubbo](https://github.com/kitex-contrib/codec-dubbo) for detailed usage.

### Test What You Want

**run.sh** would test all integrated test categories stored in **tests** variable by default. 
To save time, you could remain the test category you want and comment others.

### Test

```shell
./run.sh
```

## For codec-dubbo CI

**run.sh** would accept a parameter to replace codec-dubbo version for CI.