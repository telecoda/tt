# tt
timetool

Simple command line tool for converting a time to a variety of formats.

##Installation

    go get github.com/telecoda/tt

##Parameters
    Usage of ./tt:

    Simple time conversion tool.
    Provide a time input in a variety of formats
    It takes one parameter which is a timestamp. This can be a second, 
    nanosecond or millisecond unix timestamp. The tool makes some 
    assumptions about the time being relatively recent. It uses the number of 
    digits to decide which format your timestamp is in.  Try not to use times 
    prior to 1973-03-03 09:46:40 +0000 UTC and expect the nanosecond or 
    millisecond time to work. You can also provide a string field as long as 
    the string is in format RFC3339.
    If no time provided, defaults to current time
            
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

    $ ./tt 1475222813

    Format      :                     Value
    =======================================
    Nano        :       1475222813000000000
    Unix        :                1475222813
    RFC3339     : 2016-09-30T09:06:53+01:00
    RFC3339Nano : 2016-09-30T09:06:53+01:00




    $ ./tt 1475222813000000000
    Format      :                     Value
    =======================================
    Nano        :       1475222813000000000
    Unix        :                1475222813
    RFC3339     : 2016-09-30T09:06:53+01:00
    RFC3339Nano : 2016-09-30T09:06:53+01:00
    
    $ ./tt 147522281300000
    Format      :                     Value
    =======================================
    Nano        :       1475222813000000000
    Unix        :                1475222813
    RFC3339     : 2016-09-30T09:06:53+01:00
    RFC3339Nano : 2016-09-30T09:06:53+01:00

Use a text formatted time

    $ ./tt 2016-09-30T09:06:53+01:00
    Format      :                     Value
    =======================================
    Nano        :       1475222813000000000
    Unix        :                1475222813
    RFC3339     : 2016-09-30T09:06:53+01:00
    RFC3339Nano : 2016-09-30T09:06:53+01:00
    
or

    $ ./tt 2016-09-30T09:06:53Z
    Format      :                     Value
    =======================================
    Nano        :       1475226413000000000
    Unix        :                1475226413
    RFC3339     : 2016-09-30T09:06:53Z
    RFC3339Nano : 2016-09-30T09:06:53Z