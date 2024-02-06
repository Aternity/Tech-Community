# 401 - APM Analysis Server on AWS EC2

This cookbook deploys the ALLUVIO Aternity [APM](https://www.riverbed.com/products/application-performance-monitoring/) Analysis Server in few clicks into your existing VPC/Subnet in your own AWS account.

<div align="center">
<img src="images/alluvio-apm-as-on-aws-ec2_login.png" alt="APM AS on AWS EC2" width="70%" height="auto">
</div>

## Prerequisites

1. an AWS Account
2. a temporary URL of the APM Analysis Server installer package, for example https://yourstorage.com/folder/installer.tar?token=123&validity=12
3. the checksum of the installer package, for example "123412341234..."

> [!TIP]
> Check the [Cookbook FAQ](#FAQ) and refer to the [Riverbed Support website](https://support.riverbed.com/content/support/software/aternity-dem/aternity-apm.html)

## Quick Start

In the table bewlo, hit the **Launch Stack** button of the region where you want to deploy the APM Analysis Server.

Then follow the wizard in your AWS CloudFormation console. There, just enter the **temporary URL** and **checksum** of the installer, select a **Subnet** and a **Security Group**, scroll down and finally hit the **Create Stack** button to deploy the APM Analysis Server.

<div align="center">
  
|   | AWS Region | Quick Start |
| --- | --- | --- | 
| us-west-1 | US West (N. California) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=us-west-1#/stacks/create/review?templateURL=TEMPLATE1&stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2) |
| us-west-2 | US West (Oregon) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=us-west-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| us-east-1 | US East (N. Virginia) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| us-east-2 | US East (Ohio) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ca-west-1 | Canada (Calgary) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ca-west-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ca-central-1 | Canada (Central) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ca-central-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| af-south-1 | Africa (Cape Town) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=af-south-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| eu-west-1 | Europe (Ireland) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| eu-west-2 | Europe (London) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| eu-west-3 | Europe (Paris) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-3#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| eu-north-1 | Europe (Stockholm) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-north-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| eu-south-1 | Europe (Milan) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-south-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| eu-south-2 | Europe (Spain) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-south-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| eu-central-1 | Europe (Frankfurt) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-central-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| eu-central-2 | Europe (Zurich) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-central-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| il-central-1 | Israel (Tel Aviv) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=il-central-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| me-central-1 | Middle East (UAE) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=me-central-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| me-south-1 | Middle East (Bahrain) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=me-south-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-northeast-1 | Asia Pacific (Tokyo) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-northeast-2 | Asia Pacific (Seoul) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-northeast-3 | Asia Pacific (Osaka) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-3#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-south-1 | Asia Pacific (Mumbai) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-south-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-south-2 | Asia Pacific (Hyderabad) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-south-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-east-1 | Asia Pacific (Hong Kong) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-east-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-southeast-1 | Asia Pacific (Singapore) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-southeast-2 | Asia Pacific (Sydney) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-southeast-3 | Asia Pacific (Jakarta) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-3#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-southeast-4 | Asia Pacific (Melbourne) | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-4#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |

</div>

It is ready after just few minutes (usually less than 5 minutes) and you can connect and log into the webconsole.


## FAQ

### How to get a temporary URL for the installer package and the checksum?

From the [Riverbed Support page](https://support.riverbed.com/content/support/software/aternity-dem/aternity-apm.html) check for *Aternity APM Analysis Server (Linux Installer)* and download the latest version of the installer package. Then you can upload it in a storage of your choice, for example in a AWS S3 Bucket from which you can generate a temporary URL. Whether private or public the must be accessible to the EC2 instance.

On the same [Riverbed Support page](https://support.riverbed.com/content/support/software/aternity-dem/aternity-apm.html), you can download the checksum file. It is a text file that contains the required checksum.

### How to connect to the APM Analysis Server console?

When the stack is created you should be able to connect to web console using the Public URL or Private URL:

- **Public URL**: [https://ec2-your-instance-ip.your_region.compute.amazonaws.com](https://ec2-your-instance.your_region.compute.amazonaws.com)
- **Private URL**: [https://ec2-your-instance-ip.your_region.compute.internal](https://ec2-your-instance.your_region.compute.internal)

> [!TIP]
> - To find the actual **Public URL** or **Private URL**, check the Outputs tab in the AWS CloudFormation stack. Or in the EC2 service, you shoud have a instance named **ALLUVIO Aternity APM Analysis Server**
> - If you cannot reach the page, check the connectivty. Possibly the Security Group associated to the EC2 might be blocking port 443
> - For the login / password, refer to the User Guide on [Riverbed Support](https://support.riverbed.com/content/support/software/aternity-dem/aternity-apm.html)

#### License

Copyright (c) 2024 Riverbed

The contents provided here are licensed under the terms and conditions of the MIT License accompanying the software ("License"). The scripts are distributed "AS IS" as set forth in the License. The script also include certain third party code. All such third party code is also distributed "AS IS" and is licensed by the respective copyright holders under the applicable terms and conditions (including, without limitation, warranty and liability disclaimers) identified in the license notices accompanying the software.
