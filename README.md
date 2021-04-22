# Marrying Terraform and Serverless


## Introduction

In this article, we're going to explore how to marry two popular tools in the DevOps space: [Terraform](https://terraform.io) and the [Serverless Framework](https://serverless.com). The article assumes a basic knowledge of both tools, but before we get started we'll review each in a little more detail. As we'll see, each is useful for slightly different cases, and by combining them together we can get the best features of both.

Terraform is used primarily as an Infrastructure as Code framework, which makes it simple to manage infrastructure across multiple platforms using a single declarative style language ([Hashicorp Configuration Languiage](https://www.terraform.io/docs/language/syntax/configuration.html)). Through the use of a rich library of [providers](https://www.terraform.io/docs/language/providers/index.html), Terraform can be used to manage everything from cloud infrastructure on [AWS](https://registry.terraform.io/providers/hashicorp/aws/latest), [Azure](https://registry.terraform.io/providers/hashicorp/azurerm/latest), and [Google Cloud](https://registry.terraform.io/providers/hashicorp/google/latest), to other platforms like [Kubernetes](https://registry.terraform.io/providers/hashicorp/kubernetes/latest), [Kafka](https://registry.terraform.io/providers/Mongey/kafka/latest), and other Hashicorp tools such as [Vault](https://registry.terraform.io/providers/hashicorp/vault/latest), [Consul](https://registry.terraform.io/providers/hashicorp/consul/latest), or [Nomad](https://registry.terraform.io/providers/hashicorp/nomad/latest).

Serverless is a framework used for deploying serverkess functions in various platforms, such as [AWS Lambda](https://www.serverless.com/framework/docs/providers/aws/), [Azure Functions](https://www.serverless.com/framework/docs/providers/azure/), and even non-cloud provider specific ones such as [KNative](https://www.serverless.com/framework/docs/providers/knative/). It provides excellent [templated examples](https://www.serverless.com/examples/) to get you started using various platforms and either synchronous or asynchronous message based mechanisms (such as AWS SQS, Azure Event Hub, etc).

Both are valuable and widely used tools, however their strengths lie in different areas. Terraform is excellent for managing large infrastructure deployments, especially when principles like [module composition](https://www.terraform.io/docs/language/modules/develop/composition.html) are followed. There is also a large [registry](https://registry.terraform.io/) of providers and community built modules, which makes it easy to get started deploying related sets of infrastructure. But from my experience, doing things like deploying a Lambda can get un-necessarily complex, because you have to consider things like creating a bucket for the code to be uploaded to, packaging the code, and managing versions. 

Serverless makes it very easy to deploy code to various functions-as-a-service platforms, and has an excellent [library of examples](https://www.serverless.com/examples/) that make it very easy to begin programming against the same platforms. For example, with one command, you can have [a boilerplate Lambda function](https://www.serverless.com/examples/aws-golang-simple-http-endpoint/) that accepts input from an HTTP endpoint. And while you _can_ get into the nitty gritty details of things like the S3 bucket storing your code (and you should as you scale), you need not worry about such things when you're doing an initial deployment. However, while you _can_ provision infrastructure outside of the function and direct supports using Serverless's [resource](https://www.serverless.com/framework/docs/providers/aws/guide/resources/) functionality, this comes with a number of disadvantages over using Terraform for these things. Chiefly, the combination of a lack of true [Terraform plan](https://www.terraform.io/docs/cli/commands/plan.html) like functionality and that every time you want to change your function code you might potentially effect a change on your other infrastructure should raise the anxiety level for any competant infrastructure engineer.

The good news is that by combining the two, we can get the best of both worlds.

## Principal of Division of Labor

The key to marrying these two tools is to use each only for what they are best at. 
- Terraform - good at stateful stuff, see changes, version things, etc
- Serverless - great for rapidly deploying applications
- Why not do everything in Serverless?
- Why not do everything in Terraform?
- Split out for risk reduction / blast radius

## Let's code!

- We need a way to get values from things that are created by Terraform, securely, such that Serverless (or the stuff deployed by Serverless) can read them
- Why SSM?
- Do's and don'ts
  - Don't let Lambda read everything!
  - Don't use the built-in [SSM variable referencing in Serverless](https://www.serverless.com/framework/docs/providers/aws/guide/variables#reference-variables-using-the-ssm-parameter-store)!
  - Do let your code read the SSM values directly
    - Link to various libraries for reading configs from SSM
    - Advantage: every time your Lambda is invoked, it'll get the right setting
- Walk through
  - What does the Terraform deployment do?
  - What does the Serverless deployment do?
  - How are they tied together?
  - Deploy and walk through full example


