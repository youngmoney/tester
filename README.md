# Tester

Context dependent test wrapper.

## Usage

``` bash
tester test [args ...]
```

For tests that have an "all" command

``` bash
tester -a test
```

### Looping

Test, Read Logs, Loop

``` bash
tester -l ...
```

## Config

``` yaml
tester:
  log_reader:
    command: vim $@
  tests:
  - match_path_regex: .*
    pre_test_command: echo run before every test
    test_all_command: echo run all tests
    test_command: |
      echo complicated command
      for t in "$@"; do
        [ "$t" == "fail" ] && exit 1
        echo $t
      done
    failed_log_list_command: ls /path/to/logs | grep 'failure'
```

See `examples/config.yaml`
