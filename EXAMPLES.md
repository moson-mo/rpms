# rpms
## Renoir power metrics server
</br>

## Examples

Below you'll find some examples how to use the metric service.

### Terminal / curl

You can query metric data from the terminal, e.g. with `curl`
By default the server listens only on the loopback interface and port 8090.

To get a list with all metrics in plain text use:

`curl http://127.0.0.1:8090/pmtab?format=plain`

The output will be a list in the format `metric-name: value`.


```
STAPM LIMIT: 20.000000
STAPM VALUE: 2.218381
PPT LIMIT FAST: 20.000000
PPT VALUE FAST: 2.195250
PPT LIMIT SLOW: 15.000001
PPT VALUE SLOW: 2.274918
PPT LIMIT APU: 15.000001
PPT VALUE APU: 0.000000
...
```

### Conky script

If we want to show for example the current power consumption of the CPU, you can use curl in your conky script:

`${exec curl -s http://127.0.0.1:8090/pmval?metric=SOCKET%20POWER | cut -c1-4} W`

With curl we query the value for metric `SOCKET POWER`.</br>
Then we use cut to trim the value since we only want to see the first 4 characters.</br>
(Note: `%20` is the urlencoded value for a whitespace).

![conky example](https://github.com/moson-mo/rpms/raw/main/assets/rpms_conky.png?inline=true)