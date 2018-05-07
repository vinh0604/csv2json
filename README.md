# csv2json

The command line program helps converting CSV data to JSON

### Usage

```
csv2json -input sample.csv > out.json
```

```
cat sample.csv | csv2json > out.json
```

By default, the program uses the first line in the CSV input as header, in which each column is mapped to the corresponding JSON key, and value is always treated as string.

`-header` option allows users to specify the arbitrary header as well as the data type and default value

Format :
```
-header <keyName>:<dataType>:<defaultValue>,<keyName>:<dataType>:<defaultValue>
```

Supported data types are:
```
string
number
boolean
```

Example:
```
csv2json -input sample.csv -header name:string:,age:number:0
```


