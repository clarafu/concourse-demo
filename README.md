# concourse-demo

A dummy repository that is used for a basic Concourse pipeline demo.

This repository contains a go app that reads from a file named "commit.txt" and prints out a message that wraps the contents of the commit file. There is also a version file that is used by semver resource from the demo pipeline.

### Demo steps

The pipeline exists at https://ci.concourse-ci.org/teams/main/pipelines/demo-pipeline, where we can first explain what a Concourse pipeline looks like. Topics I would suggest to go over is:

- Jobs (shown from the 4 green boxes: bump-version, test, build and deploy)
- Resources (black boxes: app, version and app-release)
  - When explaining resources, we can click into the resources page to show the versions. The one I would use is the "app" resource, and show that the versions it grabs is the commit version because it is a git resource. Then I would click on the version to show the metadata. I would explain how resources are just abstract external locations where your pipeline will monitor for changes, fetch and push bits to. Resources use resource types, like the git resource type that the "app" resource uses, there are Concourse provided resources and also community made ones, you can also easily make your own.
- The lines between the jobs to show the passed constraints, where each version that gets grabbed is passed through each job one by one.

Then I would click into the "test" job to show the different steps within a job. The first step in the "test" job is the get step that grabs the latest version of the "app" resource (which is a git resource). The second step is a task step that uses the "app" resource that was fetched by the get step as an input to the task. I would explain how a task step works, with it  being something that runs any script that you give it and the step succeeds by it successfully running the script. I would also explain how each step is run within it's own container.

I would then go to the "build" job to show how it packages up the app code by building the go app (you can show how it builds the app by clicking on the build-app step) and then uses a put step to push the packaged go app to an external storage (in this case is a dev-release on github). To pass things between jobs within Concourse requires you to push them to external service, in order to maintain statelessness.

The last job is the "deploy" job, that takes the packaged go app from the dev-release and releases it as a github release with the "version" resource as the app version.

One option to make it more interactive is that before you start the explaination of the pipeline, you can run the script within `scripts/add-commit` to create a new version of the "app" resource in order to trigger off the pipeline. And as the pipeline runs with that new version, you can explain all the topics above as it progresses. In order to run the script, you need to create a file under `scripts/app_private_key` which the private key can be found in the Concourse lastpass under `OSCON demo keys`. 

After you are finished explaining these basic topics of the demo, you can also show the pipeline config which is a yaml file that lives in https://github.com/clarafu/concourse-demo/blob/master/ci/demo-pipeline.yml. There are also the tasks that the pipeline uses that lives within the `ci/tasks` folder. You can also change some things within the pipeline config and then use the fly CLI (make sure to have the cli step up beforehand) to show how to set the pipeline and see the changes in the UI. In order to set the pipeline, you need to have the `app-repo-key` and `dummy-user-access-token` vars set. These keys are also found in the Concourse lastpass `OSCON demo keys`. 
