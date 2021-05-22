## Purpose

Command line utility to check for available covid vaccines based on district, age and dosage preferences.

> **Disclaimer**: This is created just to ease the search for vaccine availability, this is not supported or endorsed by any official governing authorities nor by the author of this utility. So use it with caution and at your own risk.       

## Prerequisites

You would need `golang` and `curl` installed on your machine. 

## Usage 

```
go build
./jabfinder -h
```

To check for availability use the command `./jabfinder check -h`

```
Checks the availability of vaccine

Usage:
  jabfinder check [flags]

Flags:
  -a, --age int               18 or 45 - Age group to find slots for (default 18)
  -d, --districtCode string   Numeric district code
  -e, --dose int              1 or 2 - Dose to filter by (default 1)
  -h, --help                  help for check

Global Flags:
      --config string   config file (default is $HOME/.jabfinder.yaml)
```

**Examples:**

Refer the following examples to know how to use this utility 
```
## For 18 years and dosage 1
./jabfinder -d <districtCode> -a 18 -e 1


## For 45 years and dosage 2
./jabfinder -d <districtCode> -a 45 -e 2

## Replace districtCode with the district you are searching for 
```

#### Sample Response
```
+------------+------------+--------+--------+-------------------------------+--------------------------------+
|    DATE    |  VACCINE   | DOSE 1 | DOSE 2 |            CENTER             |            ADDRESS             |
+------------+------------+--------+--------+-------------------------------+--------------------------------+
| 23-05-2021 | COVISHIELD | 17     | 25     | Srirampura UPHC               | Near Sai Baba Nagar, 560021    |
+------------+------------+--------+--------+-------------------------------+--------------------------------+
| 23-05-2021 | COVISHIELD | 0      | 25     | Rajajinagar UPHC              | 57th Cross5th Block            |
|            |            |        |        |                               | Rajajinagar, 560010            |
+------------+------------+--------+--------+-------------------------------+--------------------------------+
| 23-05-2021 | COVISHIELD | 0      | 25     | Manjunathanagar UPHC          | 1st Main Road Manjunathanagar  |
|            |            |        |        |                               | Banglore-560010, 560010        |
+------------+------------+--------+--------+-------------------------------+--------------------------------+
```

### Finding district code

You will have to first know the state code you are searching for. You can find the state code by running the following command.

```
./jabfinder states
```

Use the state code you have identified like this to find the district code you are searching for by running the following command.
```
./jabfinder districts -s <STATECODE>

e.g. for Karnataka

./jabfinder districts -s 16
```

Find the code to the corresponding district you want to check availability for and use it in the `jabfinder check` command.


