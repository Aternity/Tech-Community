# 401 - Cookbook ALLUVIO APM Analysis Server on AWS EC2

This cookbook deploys the ALLUVIO APM Analysis Server in your own AWS account in few clicks.

![ALLUVIO APM AS on AWS EC2](images/alluvio-apm-as-on-aws-ec2_login.png)

## Prerequisites

1. AWS Account
2. A temporary URL of the APM AS installer package, for example https://yourstorage.com/folder/installer.tar?token=123&validity=12
3. The checksum of the installer package, for example "123412341234..."

> [!TIP]
> To get the installer package and more details, check on the [Riverbed Support](https://support.riverbed.com/content/support/software/aternity-dem/aternity-apm.html) for *Aternity APM Analysis Server (Linux Installer) version 2023.11.0* 

## Quick Start

In the following table, hit the **Launch Stack** button of the region where you want to deploy the APM Analysis Server.

It will open the wizard of the **Quick Start** deployment in the AWS CloudFormation console. There, just enter the **temporary URL** and **checksum** of the installer and hit the **Create Stack** button to deploy the APM Analysis Server in your default VPC and Subnet. For more customization, use the Launch Stack button in the [Custom Deployment](#custom-deployment) in the section below.

<div align="center">
  
| AWS Region | Quick Start |
| --- | --- | 
| us-west-1 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=us-west-1#/stacks/create/review?templateURL=TEMPLATE1&stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2) |
| us-west-2 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=us-west-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| us-east-1 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| us-east-2 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| eu-west-1 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| eu-west-2 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| eu-west-3 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-3#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| eu-central-1 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-central-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-northeast-1 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-northeast-2 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-northeast-3 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-3#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-southeast-1 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |
| ap-southeast-2 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2&templateURL=TEMPLATE1) |

</div>

## Custom deployment

The following table contains the "Launch Stack" buttons that launch custom deployment.

As a **Quick Start** deployment, the **Custom** deployment will install the ALLUVIO APM Analysis Server on an EC2 instance. While the **Quick Start** uses default network configuration, the **Custom** has more parameters that allow to customize the configuration in AWS, for example selecting the VPC/Subnet and the Security Group.

<div align="center">

| AWS Region | Custom |
| --- | --- | 
| us-west-1 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=us-west-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2-Custom&templateURL=TEMPLATE2) |
| us-west-2 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=us-west-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2-Custom&templateURL=TEMPLATE2) |
| us-east-1 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2-Custom&templateURL=TEMPLATE2) |
| us-east-2 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2-Custom&templateURL=TEMPLATE2) |
| eu-west-1 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2-Custom&templateURL=TEMPLATE2) |
| eu-west-2 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2-Custom&templateURL=TEMPLATE2) |
| eu-west-3 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-3#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2-Custom&templateURL=TEMPLATE2) |
| eu-central-1 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-central-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2-Custom&templateURL=TEMPLATE2) |
| ap-northeast-1 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2-Custom&templateURL=TEMPLATE2) |
| ap-northeast-2 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2-Custom&templateURL=TEMPLATE2) |
| ap-northeast-3 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-3#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2-Custom&templateURL=TEMPLATE2) |
| ap-southeast-1 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-1#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2-Custom&templateURL=TEMPLATE2) |
| ap-southeast-2 | [![Deploy to Azure](https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-2#/stacks/new?stackName=Riverbed-Community-Cookbook-ALLUVIO-APM-AS-on-EC2-Custom&templateURL=TEMPLATE2) |

</div>

## FAQ

### How to get a temporary URL for the installer package?

From the [Riverbed Support page](https://support.riverbed.com/content/support/software/aternity-dem/aternity-apm.html) check for *Aternity APM Analysis Server (Linux Installer) version 2023.11.0* and download the installer package. Then you can upload it in a storage of your choice from which you can generate temporary and secure URL. For example, in a AWS S3 bucket.

### In the CloudFormation console, how to create the stack?

After filling the parameters that are required for the deployment, for example the **URL** and **checksum** in the section "ALLUVIO APM Analysis Installer", and having configured other section, for example in the section "EC2 Instance Configuration" you can select different sizing for the EC2 instance, just scroll down and hit the **Create Stack** button.

> [!NOTE]
> - The Create Stack button starts the deployment the ALLUVIO APM Analysis Server on an EC2 instance in the selected region

## How to connect to the APM AS console?

When the stack is created, usually in less than 5 minutes, in a Web Browser you should be able to connect to the Public URL or Private URL of the APM Analysis Server:

- **Public URL**: [https://ec2-your-instance-ip.your_region.compute.amazonaws.com](https://ec2-your-instance.your_region.compute.amazonaws.com)
- **Private URL**: [https://ec2-your-instance-ip.your_region.compute.internal](https://ec2-your-instance.your_region.compute.internal)

> [!TIP]
> - To find the actual **Public URL** or **Private URL**, just go to the Outputs tabs of the AWS CloudFormation stack, or else go the EC2 service and find the instance named **ALLUVIO APM Analysis Server**
> - If you cannot reach the page, check the port 443 is open and the connectivity. Typically the Security Group associated to the EC2 might be blocking port 443
> - For the login / password, check the guide on [Riverbed Support](https://support.riverbed.com/content/support/software/aternity-dem/aternity-apm.html)

#### License

Copyright (c) 2024 Riverbed

The contents provided here are licensed under the terms and conditions of the MIT License accompanying the software ("License"). The scripts are distributed "AS IS" as set forth in the License. The script also include certain third party code. All such third party code is also distributed "AS IS" and is licensed by the respective copyright holders under the applicable terms and conditions (including, without limitation, warranty and liability disclaimers) identified in the license notices accompanying the software.
