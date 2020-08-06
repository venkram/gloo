---
title: AWS Lambda + EKS ServiceAccounts
weight: 101
description: Using EKS ServiceAccounts with Gloo for AWS Lambda
---

# How to use EKS ServiceAccounts to authenticate AWS Lambda requests with Gloo

Recently, AWS added the ability to associate Kubernetes ServiceAccounts with IAM roles.
This [blog post](https://aws.amazon.com/blogs/opensource/introducing-fine-grained-iam-roles-service-accounts/) 
explains the feature in more detail.

Gloo Api Gateway now supports discovering, and authenticating AWS Lambdas in kubernetes using 
these projected ServiceAccounts

## Configuring EKS cluster to use IAM ServiceAccount roles

The first step to enabling this IAM ServiceAccount roles with Gloo is creating/configuring an EKS
cluster to use this feature.

A full tutorial can be found [in AWS' docs](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html).
Once the cluster exists and is configured properly, return here for the rest of the tutorial.

Note: The aws role needs to be associated with a policy which has access to the following 4
Actions for this tutorial to function properly.

    * lambda:ListFunctions
    * lambda:InvokeFunction
    * lambda:GetFunction
    * lambda:InvokeAsync

## Deploying the EKS Pod Identity webhook

In order for the Gloo workloads to recieve the projected AWS ServiceAccount tokens, special environment
variables and volumes need to be added to the Pod spec. Luckily AWS has a [mutating webhook](https://github.com/aws/amazon-eks-pod-identity-webhook) 
available which can do this for us. Installation instructions can be found [here](https://github.com/aws/amazon-eks-pod-identity-webhook).
If you run into issues pulling the images feel free to use the solo hosted images instead, those can be found
at `quay.io/solo-io/aws-identity-webhook`.

The injection process can be accomplished manually by adding these values to the Gloo deployments, but
we highly recommend using the webhook. The rest of this tutorial assumes the webhook is being used.

## Deploying Gloo

As this feature is brand new, it is currently only available on a beta branch of gloo. The following 
are the version requirements for closed source and open source Gloo.

    Closed Source: v1.5.0-beta7
    
    Open Source: v1.5.0-beta16

For the purpose of this tutorial we will be installing open source Gloo, but closed source Gloo 
should work exactly the same

```shell script
helm install gloo https://storage.googleapis.com/solo-public-helm/charts/gloo-1.5.0-beta16.tgz \
 --namespace gloo-system --create-namespace --values - <<EOF
settings:
  aws:
    enableServiceAccountCredentials: true
gateway:
  proxyServiceAccount:
    extraAnnotations:
      eks.amazonaws.com/role-arn: $AWS_ROLE_ARN
discovery:
  serviceAccount:
    extraAnnotations:
      eks.amazonaws.com/role-arn: $AWS_ROLE_ARN
EOF
```

Once helm has finished installing, which we can check by running the following, we're ready to move on.
```shell script
kubectl rollout status deployment -n gloo-system gateway-proxy
kubectl rollout status deployment -n gloo-system gloo
kubectl rollout status deployment -n gloo-system gateway
kubectl rollout status deployment -n gloo-system discovery
```


## Routing to our Lambda

Now that Gloo is running with our credentials set up, we can go ahead and create our Gloo config to 
enable routing to our AWS Lambda

First we need to create our Upstream. The following names and qualifier should be replaced with the lambda
you wish to route to.
```yaml
apiVersion: gloo.solo.io/v1
kind: Upstream
metadata:
  annotations:
  name: lambda
  namespace: gloo-system
spec:
  aws:
    lambdaFunctions:
    - lambdaFunctionName: uppercase
      logicalName: uppercase1
      qualifier: "1"
    region: us-east-1
```

Once the Upstream has been accepted we can go ahead and create our Virtual Service
```yaml

---
apiVersion: gateway.solo.io/v1
kind: VirtualService
metadata:
  name: default
  namespace: gloo-system
spec:
  virtualHost:
    domains:
    - '*'
    routes:
    - matchers:
      - prefix: /lambda
      routeAction:
        single:
          destinationSpec:
            aws:
              logicalName: uppercase1
          upstream:
            name: lambda
            namespace: gloo-system
```

Now we can go ahead and try our route!
```shell script
curl -v $(glooctl proxy url)/lambda --data '"abc"' --request POST -H"content-type: application/json"
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying 3.129.77.154...
* TCP_NODELAY set
* Connected to a665a63a0d73311eaa9e206ac0b0c480-1372496097.us-east-2.elb.amazonaws.com (3.129.77.154) port 80 (#0)
> POST /lambda HTTP/1.1
> Host: a665a63a0d73311eaa9e206ac0b0c480-1372496097.us-east-2.elb.amazonaws.com
> User-Agent: curl/7.64.1
> Accept: */*
> content-type: application/json
> Content-Length: 5
>
* upload completely sent off: 5 out of 5 bytes
< HTTP/1.1 200 OK
< date: Wed, 05 Aug 2020 17:59:58 GMT
< content-type: application/json
< content-length: 5
< x-amzn-requestid: e5cc4545-2989-4105-a4b2-49707d654bce
< x-amzn-remapped-content-length: 0
< x-amz-executed-version: 1
< x-amzn-trace-id: root=1-5f2af39e-5b3e38488ffeb5ec541107d4;sampled=0
< x-envoy-upstream-service-time: 53
< server: envoy
<
* Connection #0 to host a665a63a0d73311eaa9e206ac0b0c480-1372496097.us-east-2.elb.amazonaws.com left intact
"ABC"* Closing connection 0
```

We can also optionally override the role ARN used to authenticate our lambda requests, by adding it into our Upstream
like so:
```yaml
apiVersion: gloo.solo.io/v1
kind: Upstream
metadata:
  annotations:
  name: lambda
  namespace: gloo-system
spec:
  aws:
    lambdaFunctions:
    - lambdaFunctionName: uppercase
      logicalName: uppercase1
      qualifier: "1"
    region: us-east-1
    roleArn: $SECONDARY_AWS_ROLE_ARN
```

