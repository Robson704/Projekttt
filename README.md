# Todo List

![architecture](arch.png)

The Todo List web app contains the following:
* Lambda
* DynamoDB
* S3 Bucket
* API Gateway
* Code Pipeline

The **Lambda** function will run our code in response to events and automatically manages the underlying compute resources. It will handle every request coming from the API gateway and will return a response from the API gateway. The `TodoListFunction` lambda function has an access to the DynamoDB of `users` and `tasks` table for reading and writing permission.

**DynamoDB** will serve as our database to store our data for users and tasks. Both `users` and `tasks` table has a partition key called `id`. The partition key is our primary key for the said table.

The **S3 Bucket** is for the frontend where it is going to be deployed as a static web application. It is a container for our static files and is the source code for our frontend.

An **API Gateway** will route all requests to the lambda function and is configured with CORS. It has an API endpoint for `tasks` and `users`. The `tasks` endpoint has a method of *GET* and *POST*. It also has a resource where it accepts a `task_id` with the method of *POST* and *DELETE*. Whilst the `users` endpoint has the same structure as `tasks` but it accepts the `user_id` path parameter.

Lastly, **Code Pipeline** is used for the automation of the software deployment process CI/CD. It automatically builds, tests, and launches the application each time there is a change in our code. Code Pipeline is integrated with a third-party service called Github. When the developer commit changes to the repository, Code Pipeline automatically detects the changes. Those changes are built, and if there are tests that are configured, they will run. After the tests are complete, the code is built and deployed to the staging. The pipeline setup we have has the developer stage and the production stage, where it needs to be manually approved to be deployed.

## Project Directory
* `cmd/todoList/main.go` is our lambda handler where our function code processes the API events.
* `lib/stacks/todo-list-stack.ts` is where your CDK application’s main stack is defined.
* `lib/stage/todo-list-stage.ts` is where you instantiate your resource stack.
* `lib/pipeline-stack.ts` is the initial structure of your pipeline and instantiating the `developer` and `production` stage.
* `bin/todo-list.ts` is the entrypoint of the CDK application. It will load the stack defined in `lib/todo-list-stack.ts`.
* `web_app` contains our front-end design for Todo List App that will be placed in an S3 bucket called `todo-list-app-dev`. It has a functionality of create, read, update, and delete for `users` and `tasks`.
* `package.json` is your npm module manifest. It includes information like the name of your app, version, dependencies and build scripts like “watch” and “build” (package-lock.json is maintained by npm)
* `cdk.json` tells the toolkit how to run your app. In our case it will be "npx ts-node bin/todo-list.ts"
* `tsconfig.json` your project’s typescript configuration
* `.gitignore` and `.npmignore` tell git and npm which files to include/exclude from source control and when publishing this module to the package manager.
* `node_modules` is maintained by npm and includes all your project’s dependencies.

---
**NOTE**

For the **`web_app`** front-end design to work, you'll need to deploy your application first to **CloudFormation** using AWS CDK. After the stack is completely deployed, go to API Gateway console -> `todo-list-api` -> Stages. Select on the stage menu you have (e.g. `prod`) then it will show you the list of APIs with its method and invoke URL. Get the invoke URL for `tasks` and `users`, then update the `web_app/dashboard.html` and `web_app/index.html` URL.

---

## AWS Configure
Configure your workstation with your credentials and an AWS region.
```bash
dev@dev:~$ aws configure
```

To create multiple accounts for AWS CLI:
```bash
dev@dev:~$ aws configure --profile profile_name
```

Provide your AWS access key ID, secret access key and default region when prompted. You can switch between the accounts by passing the profile on the command.

```bash
dev@dev:~$ aws s3 ls --profile profile_name
```

When no `--profile` parameter in the command, `default` profile will be used.

## Install aws-todo-list-app
Run `npm install` in the root of your project and this will install all the dependencies.

```bash
dev@dev:~:aws-todo-list-app$ npm install
npm WARN deprecated source-map-url@0.4.1: See https://github.com/lydell/source-map-url#deprecated
npm WARN deprecated urix@0.1.0: Please see https://github.com/lydell/urix#deprecated
npm WARN deprecated resolve-url@0.2.1: https://github.com/lydell/resolve-url#deprecated
npm WARN deprecated source-map-resolve@0.5.3: See https://github.com/lydell/source-map-resolve#deprecated
npm WARN deprecated sane@4.1.0: some dependency vulnerabilities fixed, support for node < 10 dropped, and newer ECMAScript syntax/features added

added 518 packages, and audited 537 packages in 3s

27 packages are looking for funding
  run `npm fund` for details

found 0 vulnerabilities
```

## Synthesis

```bash
dev@dev:~:aws-todo-list-app$ cdk synth
Resources:
  TodoListcustomRole78A60DC9:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Statement:
          - Action: sts:AssumeRole
            Effect: Allow
            Principal:
              Service: lambda.amazonaws.com
        Version: "2012-10-17"
      ManagedPolicyArns:
        - Fn::Join:
            - ""
            - - "arn:"
              - Ref: AWS::Partition
              - :iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
      RoleName: TodoList_customRole
    Metadata:
      aws:cdk:path: TodoListAppStack/TodoList_customRole/Resource
.....
Parameters:
  BootstrapVersion:
    Type: AWS::SSM::Parameter::Value<String>
    Default: /cdk-bootstrap/zep458faq/version
    Description: Version of the CDK Bootstrap resources in this environment, automatically retrieved from SSM Parameter Store. [cdk:skip]
Rules:
  CheckBootstrapVersion:
    Assertions:
      - Assert:
          Fn::Not:
            - Fn::Contains:
                - - "1"
                  - "2"
                  - "3"
                  - "4"
                  - "5"
                - Ref: BootstrapVersion
        AssertDescription: CDK bootstrap stack version 6 required. Please run 'cdk bootstrap' with a recent version of the CDK CLI.
```

## Bootstrapping

```bash
dev@dev:~:aws-todo-list-app$ cdk bootstrap --profile profile_name
  ⏳  Bootstrapping environment aws://xxxxxxxxxx/xx-xxxx-x...
Trusted accounts for deployment: (none)
Trusted accounts for lookup: (none)
Using default execution policy of 'arn:aws:iam::aws:policy/AdministratorAccess'. Pass '--cloudformation-execution-policies' to customize.
CDKToolkit: creating CloudFormation changeset...

 ✅  Environment aws://xxxxxxxxxx/xx-xxxx-x bootstrapped.
****************************************************
*** Newer version of CDK is available [2.31.1]   ***
*** Upgrade recommended (npm install -g aws-cdk) ***
****************************************************
```

### Updgrade CDK
```bash
dev@dev:~:aws-todo-list-app$ npm install -g aws-cdk
added 1 package, and audited 2 packages in 1s

found 0 vulnerabilities
npm notice 
npm notice New minor version of npm available! 8.5.0 -> 8.13.2
npm notice Changelog: https://github.com/npm/cli/releases/tag/v8.13.2
npm notice Run npm install -g npm@8.13.2 to update!
npm notice 
```

## Deploy
This is to deploy your CDK app. You should see a warning asking if `Do you wish to deploy these changes (y/n)?`. This is warning you that deploying the app contains security-sensitive changes. Since we need to allow/deny the resources, enter `y` to deploy the stack and create the resources.

```bash
dev@dev:~$ cdk deploy --profile profile_name
✨  Synthesis time: 3.89s

This deployment will make potentially sensitive changes according to your current security approval level (--require-approval broadening).
Please confirm you intend to make the following modifications:

IAM Statement Changes
┌───┬──────────────────────────────────────────────┬────────┬──────────────────────────────────────────────┬──────────────────────────────────────────────┬────────────────────────────────────────────────┐
│   │ Resource                                     │ Effect │ Action                                       │ Principal                                    │ Condition                                      │
├───┼──────────────────────────────────────────────┼────────┼──────────────────────────────────────────────┼──────────────────────────────────────────────┼────────────────────────────────────────────────┤
│ + │ ${Custom::CDKBucketDeployment9351AE39958944B │ Allow  │ sts:AssumeRole                               │ Service:lambda.amazonaws.com                 │                                                │
│   │ 45CED0CD9AC6950C/ServiceRole.Arn}            │        │                                              │                                              │                                                │
├───┼──────────────────────────────────────────────┼────────┼──────────────────────────────────────────────┼──────────────────────────────────────────────┼────────────────────────────────────────────────┤
│ + │ ${Custom::S3AutoDeleteObjectsCustomResourceP │ Allow  │ sts:AssumeRole                               │ Service:lambda.amazonaws.com                 │                                                │
│   │ rovider/Role.Arn}                            │        │                                              │                                              │                                                │
├───┼──────────────────────────────────────────────┼────────┼──────────────────────────────────────────────┼──────────────────────────────────────────────┼────────────────────────────────────────────────┤

...........

(NOTE: There may be security-related changes not in this list. See https://github.com/aws/aws-cdk/issues/1299)

Do you wish to deploy these changes (y/n)? 
```

The output shoukd look like the this:
```bash
TodoListAppStack: deploying...
[0%] start: Publishing xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx:current_account-current_region
[100%] success: Published xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx:current_account-current_region
TodoListAppStack: creating CloudFormation changeset...

 ✅  TodoListAppStack

✨  Deployment time: 106.99s

Stack ARN:
arn:aws:cloudformation:xx-xxxx-x:xxxxxxxxxx:stack/TodoListAppStack/xxxxxxxxxxxxxxxxxxxxxxxxx

✨  Total time: 110.66s
```

## Using `Makefile` to install, bootstrap, and deploy the project

1. Install all the dependencies and bootstrap your project
    ```bash
    dev@dev:~:aws-todo-list-app$ make init
    ```

    To initialize the project with specific AWS profile, you can pass a parameter called `profile`.
    ```bash
    dev@dev:~:aws-todo-list-app$ make init profile=profile_name
    ```

2. Deploy the project.
    ```bash
    dev@dev:~:aws-todo-list-app$ make deploy
    # Deploying with specific AWS profile
    dev@dev:~:aws-todo-list-app$ make deploy profile=profile_name
    ```