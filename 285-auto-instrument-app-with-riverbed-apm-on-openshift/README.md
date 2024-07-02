# 285-auto-instrument-app-with-riverbed-apm-on-openshift

Enable Riverbed APM on an app running in a Red Hat OpenShift cluster, using the [Riverbed Operator](https://github.com/riverbed/riverbed-operator) and configuring the automatic instrumentation.

## Prerequisites

1. a SaaS account for [Riverbed APM](https://www.riverbed.com/products/application-performance-monitoring)

2. a Red Hat OpenShift environment: Red Hat console access, OpenShift client CLI (`oc` command), a cluster with Linux machine pools and apps

> [!NOTE]
> This cookbook has been tested with Red Hat OpenShift version 4

## Step 1. Get the details for Riverbed APM

In the Riverbed APM web console, from the Home page, hit "Deploy Collectors" and "Install now" button (or else navigate via the traditional menu: CONFIGURE > AGENTS > Install Agents).

Then in the Linux agent panel, switch to the "Standard Agent Install" to find and grab the values of your:

1. **Customer Id**, for example *12341234-12341234-13241234*

2. **SaaS Analysis Server Host**, for example *agents.apm.your_environment.aternity.com*

## Step 2. Connect to the cluster with `oc`

Start a shell and connect to your cluster using `oc login`, the full command would look like this:

```shell
oc login --token=yourtoken --server=yourserver
```


## Step 3. Deploy the Riverbed Operator on the Cluster

### 3.1 Permissions

Run the following command to configure the required permissions for the Riverbed Operator and APM Agent:

```shell
# Set the permissions
oc apply -f https://raw.githubusercontent.com/Aternity/Tech-Community/main/285-instrument-app-with-riverbed-apm-on-openshift/riverbed-operator-permissions-openshift.yaml
```

### 3.2 Riverbed Operator

Execute the scripts below to deploy the Riverbed Operator (using the `oc` command)

```shell
# Install Cert-manager prerequisite
oc apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.14.5/cert-manager.yaml

# Wait at least 2 min for cert-manager to get ready
sleep 120

# Install the Riverbed Operator
oc apply -f https://raw.githubusercontent.com/riverbed/riverbed-operator/1.0.0/riverbed-operator.yaml
```

### 3.3 Riverbed Operator configuration

3. Run the following command, that will open the configuration in your editor before applying it to the cluster. There, you can configure the lines `customerId: ""` and `analysisServerHost: "agents.apm.YOUR-ENV.aternity.com"` adding the values of the **Customer Id** and **SaaS Analysis Server Host** obtained in Step 1.. When done, just save and close the file.

```shell
# Configure the Riverbed Operator with your Customer Id and SaaS Analysis Server Host
oc create -f https://raw.githubusercontent.com/riverbed/riverbed-operator/1.0.0/riverbed_configuration_v1.0.0.yaml --namespace=riverbed-operator --edit
```

For example the related lines in the configuration will look like this:

```
...
  analysisServerHost: agents.apm.your_environment.aternity.com
  customerId: 12341234-12341234-13241234
...
```

> [!NOTE]
> Please refer to the [Riverbed Operator](https://github.com/riverbed/riverbed-operator) to learn the details on how to deploy the Riverbed Operator on a Kubernetes cluster.

## Step 4. Check the setup is ready

You can check the Pods in the `riverbed-operator` namespace are ready.

```shell
oc get pod -n riverbed-operator
```

The output will look like this in a cluster having few nodes

```console
NAME                                                    READY   STATUS    RESTARTS   AGE
riverbed-apm-agent-6l4k8                                1/1     Running   0          43s
riverbed-apm-agent-g5q4m                                1/1     Running   0          43s
...
riverbed-operator-controller-manager-56dd4ddf78-n2t66   2/2     Running   0          4m8s
```


## Step 5. *optional* Instrument a demo app

1. Run the following command to deploy the demo application `yourapp` in the namespace `cookbook-app`. The app uses the docker image of a simple java webapp.

```shell
oc apply -f https://raw.githubusercontent.com/Aternity/Tech-Community/main/285-instrument-app-with-riverbed-apm-on-openshift/app/yourapp.yaml
```

2. In a separate shell, run this command to open a local port that will give access the app. For example on port 8888

```shell
oc port-forward -n cookbook-app service/yourapp --address 127.0.0.1 8888:80
```

Then you should be able to access from your browser to [http://127.0.0.1:8888](http://127.0.0.1:8888) or from the CLI `curl http://127.0.0.1:8888`

<details>
<summary>Clean up the demo app</summary>

Simply run the following 

```shell
oc delete  -f https://raw.githubusercontent.com/Aternity/Tech-Community/main/285-instrument-app-with-riverbed-apm-on-openshift/app/yourapp.yaml
```

</details>

## Step 5. Instrument your app

Simply enable the automatic instrumentation adding the annotation.

> [!NOTE]
> For more details about automatic instrumentation, please refer to the readme page of the [Riverbed Operator](https://github.com/riverbed/riverbed-operator).


For example, with the demo application `yourapp`, which is a java app, the command below patches the app deployment to add the annotation `"instrument.apm.riverbed/inject-java":"true"`:

```shell
oc patch deployment -n cookbook-app cookbook-app -p '{"spec": {"template":{"metadata":{"annotations":{"instrument.apm.riverbed/inject-java":"true"}}}} }'
```

> [!TIP]
> The APM instrumentation annotation can also be added to the manifest, for example [app/yourapp-with-apm.yaml](app/yourapp-with-apm.yaml) is annotated, and based on the original manifest [app/yourapp.yaml](app/yourapp.yaml).


## Step 6. Navigate the app and monitor in Riverbed APM web console 

Go to the APM web console to observe every transaction of the application. Every Pod is automatically instrumented.

## Notes

### How to clean-up the demo app

Simply run the following 

```shell
oc delete  -f https://raw.githubusercontent.com/Aternity/Tech-Community/main/285-instrument-app-with-riverbed-apm-on-openshift/app/yourapp.yaml
```

### How to clean-up the operator


```shell
oc delete -f https://raw.githubusercontent.com/riverbed/riverbed-operator/1.0.0/riverbed-operator.yaml
```

> [!NOTE]
> Please refer to the [Riverbed Operator](https://github.com/riverbed/riverbed-operator) to learn the details on how to deploy the Riverbed Operator on a Kubernetes cluster.

#### License

Copyright (c) 2024 Riverbed

The contents provided here are licensed under the terms and conditions of the MIT License accompanying the software ("License"). The scripts are distributed "AS IS" as set forth in the License. The script also include certain third party code. All such third party code is also distributed "AS IS" and is licensed by the respective copyright holders under the applicable terms and conditions (including, without limitation, warranty and liability disclaimers) identified in the license notices accompanying the software.
