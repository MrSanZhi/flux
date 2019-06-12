package testdata_test

import "testing"

option now = () => (2030-01-01T00:00:00Z)

inData = "
#datatype,string,long,dateTime:RFC3339,double,string,string,string,string,string,string
#group,false,false,false,false,true,true,true,true,true,true
#default,_result,,,,,,,,,
,result,table,_time,_value,_field,_measurement,device,fstype,host,path
,,0,2018-05-22T00:00:00Z,30,used_percent,disk,disk1s1,apfs,host.local,/
,,0,2018-05-22T00:00:10Z,30,used_percent,disk,disk1s1,apfs,host.local,/
,,0,2018-05-22T00:00:20Z,30,used_percent,disk,disk1s1,apfs,host.local,/
,,0,2018-05-22T00:00:30Z,40,used_percent,disk,disk1s1,apfs,host.local,/
,,0,2018-05-22T00:00:40Z,40,used_percent,disk,disk1s1,apfs,host.local,/
,,0,2018-05-22T00:00:50Z,40,used_percent,disk,disk1s1,apfs,host.local,/
"

outData = "
#group,false,false,false,false,true,true,true,true,true,true
#datatype,string,long,dateTime:RFC3339,double,string,string,string,string,string,string
#default,_result,,,,,,,,,
,result,table,_time,_value,_field,_measurement,device,fstype,host,path
,,0,2018-05-22T00:00:20Z,30,used_percent,disk,disk1s1,apfs,host.local,/
,,0,2018-05-22T00:00:30Z,33.333333333333336,used_percent,disk,disk1s1,apfs,host.local,/
,,0,2018-05-22T00:00:40Z,36.666666666666664,used_percent,disk,disk1s1,apfs,host.local,/
,,0,2018-05-22T00:00:50Z,40,used_percent,disk,disk1s1,apfs,host.local,/
"
moving_average = (table=<-) =>
	(table
		|> movingAverage(n:3))

test _moving_average = () =>
	({input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: moving_average})
