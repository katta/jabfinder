## Purpose

Command line utility to check for available covid vaccines based on district, age and dosage preferences.

> **Disclaimer**: This is created just to ease the search for vaccine availability, this is not supported or endorsed by any official governing authorities. So use it with caution.       

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
  -a, --age int               Age group to find slots for (default 18)
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

### Finding district code

You will have to first know the state code you are searching for. You can find the state code from this file [states](./pkg/cowin/states.json)

```
e.g.: If you are looking for a district in Karnataka, search for "karnataka" sates file mentioned above. The number 16 (state_id) in this example is the state code for Karnataka.
    
{
  "state_id": 16,
  "state_name": "Karnataka"
}

```

Use the state code you have identified like this to find the district code you are searching for by running the following comamnd.
```
## Replace STATE_CODE in the following command with the state code you have identified

curl --location --request GET 'https://cdn-api.co-vin.in/api/v2/admin/location/districts/<STATE_CODE>' \
--header 'User-Agent: Mozilla/5.0'

e.g. For Karnataka , where the state code is 16 use this follwoing command

curl --location --request GET 'https://cdn-api.co-vin.in/api/v2/admin/location/districts/16' \
--header 'User-Agent: Mozilla/5.0'
```
You will get a JSON response which will have all the districts along with its id. Find the district you are looking for and then pick its id which should be used as `districtCode` parameter while using this utility. You can find the sample districts file [here](./pkg/cowin/districts.json)