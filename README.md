# cb-perf-tester

Command line tool that allows conducting some performance tests against Couchbase related to [Sub-Document API](https://docs.couchbase.com/c-sdk/current/concept-docs/subdocument-operations.html)

# configuration file

The configuration file contains information necessary to connect to a Couchbase bucket. This bucket is utilized for storing and conducting tests. 
By default, the configuration file is located at: `$HOME/.cb-perf-tester.yaml`. However, it can be overridden using the runtime parameter `--config`. 
A sample configuration is provided [within the repository](./configuration/local.yaml).

# runtime parameters

There are two runtime commands available: `regular` and `subdoc`. 

The `regular` command is designed to establish the performance baseline for standard key-value operations that fetch the entire document.
On the other hand, the `subdoc` command is more sophisticated, allowing you to evaluate the performance of the Sub-Document API by adjusting the difficulty of finding sub-paths within scanned documents.

| Parameter       | Applicable for | Default Value                 | Description                                                                                 |
|-----------------|----------------|-------------------------------|---------------------------------------------------------------------------------------------|
| `--config`      | regular, subdoc| `${HOME}/.cb-perf-tester.yaml`| Location of the configuration file.                                                         |
| `--repeat`      | regular, subdoc| 1000                          | Number of operations to perform.                                                            |
| `--parallel`    | regular, subdoc| 1                             | Number of goroutines that will be used to concurrently perform operations.                  |
| `--keys`        | regular, subdoc| 10000                         | Number of sub-elements that constitute the sample document; it affects the document's size. |
| `--search-keys` | subdoc         | 4                             | Number of searched sub-elements while performing a single operation.                        |
| `--difficulty`  | subdoc         | easy                          | The level of difficulty in searching for a single sub-element within the sample document.   |


The structure of sample document is as follows:

```
{
    "key": "test-subdoc",
    "data": {
        "subkey-000000": "value-000000",
        "subkey-000001": "value-000001",
        . . .
        "subkey-0….N": "value-0…..N",
    }
}
```

## understanding the difficulty parameter

The difficulty parameter can take one of four values:

- **easy**: Searched elements are located at the beginning of the sample document.
- **medium**: Searched elements are located in the middle of the sample document.
- **hard**: Searched elements are located at the end of the sample document.
- **impossible**: There are no searched elements in the sample document.

# example environment

## spawn couchbase instance

```
docker run couchbase:community-6.6.0
```

## build cb-perf-tester

In the main directory of repo:

```
make build
```

## run sample performance test

```
cd bin
./cb-perf-tester subdoc --parallel 200 --repeat 50 --search-keys 10 --difficulty hard --config ../configuration/local.yaml
Using config file: ../configuration/local.yaml
benchmark params: keys=10000, level=Hard, search-keys=10, repeats=50, parallel=200
Generated doc with subkeys: 10000, byte size is: 310043

search for subkeys [level=Hard]: [data.subkey-009999 data.subkey-009998 data.subkey-009997 data.subkey-009996 data.subkey-009995 data.subkey-009994 data.subkey-009993 data.subkey-009992 data.subkey-009991 data.subkey-009990]

subdoc report: successes: 10000, errors: 0, duration: 1m9.89228332s, rps: 143.077312, success rps: 143.077312

```

## interpreting the output

The last line of the output displays the number of successful and failed operations, the duration of the entire test, and the computed rate of operations (total) and successful.
Typically, there should be no errors; if there are, the benchmark results are not considered relevant.

