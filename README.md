# Toastmasters International Club DCP Status Checker
Check the Club DCP status page to generate the current DCP report in Yaml and image formats.

## Install tools
Install go lang and nodejs

`make install-tools`

## Format the Code
Format the code

`make format`

## Generate Report
Generate Club DCP report in Yaml format

`make generate-report`

![dcp_report.yaml](./reports/dcp_report.yaml)

```
DCP Report 2024-2025:
    Administration:
        Club officer list on time:
            achieved: "1"
            status: ""
            target: "Y"
        Membership-renewal dues on time:
            achieved: "1"
            status: Achieved
            target: "Y"
    DCP Status:
        Membership:
            Base: "22"
            Required: "20"
            To Date: "26"
        Overall:
            Current: "4"
            Distinguished: "No"
            Target: "10"
    Education:
        Level 1 awards:
            achieved: "2"
            status: 2 Level 1s needed
            target: "4"
        Level 2 awards:
            achieved: "0"
            status: 2 Level 2s needed
            target: "2"
        Level 3 awards:
            achieved: "0"
            status: 2 Level 3s needed
            target: "2"
        Level 4, Level 5, or DTM award:
            achieved: "1"
            status: Achieved
            target: "1"
        More Level 2 awards:
            achieved: "0"
            status: 2 Level 2s needed
            target: "2"
        One more Level 4, Level 5, or DTM award:
            achieved: "1"
            status: Achieved
            target: "1"
    Membership:
        More new members:
            achieved: "0"
            status: 4 New Members needed
            target: "4"
        New members:
            achieved: "4"
            status: Achieved
            target: "4"
    Training:
        Club officers trained June-August:
            achieved: "4"
            status: First Training Period Achieved
            target: "4"
        Club officers trained November-February:
            achieved: "0"
            status: Second Training Period 4 needed
            target: "4"
```


Capture the Club DCP report screenthot png

`make generate-screenshot`

![Club DCP Report](./reports/dcp_report.png)

Generate Club DCP report in Yaml and capture the Club DCP report screenthot

`make generate-all`

Use the generated Yaml report and summarize the achievements using Chat GPT API. Send an announcement message to the MS Teams channel with the summary and image captured.

`make send-notification`
