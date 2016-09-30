# tt
timetool

Simple command line tool for converting a time to a variety of formats.

##Installation

    go get github.com/telecoda/tt

##Parameters
    Usage of ./tt:

    Simple time conversion tool.
    Provide a time input in a variety of formats
    If no time provided, defaults to current time

    -n int
            time in nanoseconds
    -s string
            time in format 2006-01-02T15:04:05Z07:00
    -u int
            time in unix
            
##Examples

Use current time
    
    $ ./tt

    Format      :                     Value
    =======================================
    Nano        :       1475222813988046767
    Unix        :                1475222813
    RFC3339     : 2016-09-30T09:06:53+01:00
    RFC3339Nano : 2016-09-30T09:06:53.988046767+01:00


Use unix time

    $ ./tt -u 1475222813

    Format      :                     Value
    =======================================
    Nano        :       1475222813000000000
    Unix        :                1475222813
    RFC3339     : 2016-09-30T09:06:53+01:00
    RFC3339Nano : 2016-09-30T09:06:53+01:00


Use nano seconds

    $ ./tt -n 1475222813000000000
    Format      :                     Value
    =======================================
    Nano        :       1475222813000000000
    Unix        :                1475222813
    RFC3339     : 2016-09-30T09:06:53+01:00
    RFC3339Nano : 2016-09-30T09:06:53+01:00
    
Use a text formatted time

    $ ./tt -s 2016-09-30T09:06:53+01:00
    Format      :                     Value
    =======================================
    Nano        :       1475222813000000000
    Unix        :                1475222813
    RFC3339     : 2016-09-30T09:06:53+01:00
    RFC3339Nano : 2016-09-30T09:06:53+01:00
    
or

    $ ./tt -s 2016-09-30T09:06:53Z
    Format      :                     Value
    =======================================
    Nano        :       1475226413000000000
    Unix        :                1475226413
    RFC3339     : 2016-09-30T09:06:53Z
    RFC3339Nano : 2016-09-30T09:06:53Z