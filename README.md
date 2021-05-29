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
  -n, --notify                Notifies via email when the slots are available
  -v, --vaccine string        COVISHIELD or COVAXIN - vaccine types available (default "COVISHIELD")

Global Flags:
      --config string   config file (default is $HOME/.jabfinder.yaml)
      --generateDoc     Set to true to generate Documents (Must be run from SCM cloned location
```

For detailed usage of all commands in `jabfinder` checkout [docs](./docs)

**Examples:**

Refer the following examples to know how to use this utility 
```
## Replace districtCode with the district you are searching for. Check this doc below on how to find the district code.

## For 18 years and dosage 1
./jabfinder check -d <districtCode> -a 18 -e 1

## For 45 years and dosage 2
./jabfinder check -d <districtCode> -a 45 -e 2

## For checking continuousy use the "--notify" flag. The  command will check for availability every 10 seconds by default. 
./jabfinder check -d <districtCode> -a 45 -e 2 --notify

## You can change the duration between checks by providing the interval in the command as shown below.
JABF_NOTIFY_INTERVALINSECONDS=20 ./jabfinder check -d <districtCode> -a 45 -e 2 --notify 
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

### Sending email notifications

You could use this tool to send you a email notification as and when the new slots are available. Use the following commands to set it up.

1. Uses Google's smtp server by default `smtp.google.com`
2. You should have your email (gmail) and password from which you would want to trigger an email from
3. You should provide the email to which you want to trigger notification to
4. Setup these details as environment variables as shown below

```
export JABF_SMTP_EMAIL=jabfinderindia@gmail.com  # Replace the email with your email
export JABF_SMTP_PASSWORD=xxx                    # Replace xxx with the password
export JABF_NOTIFY_TOEMAIL=abc@xyz.com           # Replace abc@xyz.com to the email to which you want to send notifications to

# Runt he following command in the same shell where you have set these environment variables

./jabfinder check --districtCode 294 --age 45 --notify
```
> You may face errors sending email based on your smtp email account settings. Please check the Google's [documentation](https://support.google.com/mail/answer/7126229) to enable the use of SMTP on your email based on your preferences.

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

### Using Docker

You would have to pass environment variables for SMTP config.

Create a `.env-file` locally with the following and replace the values accordingly
   
```
JABF_NOTIFY_INTERVALINSECONDS=300
JABF_SMTP_EMAIL=abc@xyz.com
JABF_SMTP_PASSWORD=xxx
JABF_NOTIFY_TOEMAIL=abc@xyz.com
```  

#### 1. Using published image

Pulls the image from [dockerhub](https://hub.docker.com/repository/docker/vatsakatta/jabfinder)

```
docker run -it --name jabfinder --env-file=.env-file vatsakatta/jabfinder:1.0.0 ./jabfinder check -d 294 -a 18 -e 1 -n
```

#### 2. Building Locally

You can build the docker image using the following command

```
docker build . -t jabfinder
```

#### 3. Running Local Build

Run the docker image built using above command with the necessary parameters 

1. Pass in the command arguments along with `docker run`
   ```
   docker run -d --name jabfinder --env-file=.env-file jabfinder:latest ./jabfinder check -d 294 -a 18 -e 1 -n
  
   ## The above command will check for open slots every 5 minutes for the age limit > 18 and dose 1 and sends a mail to abc@xyz.com
   ```
   
2. Verify if the jabfinder is running in docker using the logs

    ```
    docker logs -f jabfinder    
    ``` 

 


