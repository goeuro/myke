## myke

> A higher order task aggregator with cascading configuration, suitable as a wrapper over existing task runners.

### What does it do?

Running `myke` on this folder prints:

```
project    tags       tasks
---------  ---------  ------------------
child                 build test
depends               before after before_after itself
env                   env
example               build
extends               task1 task2 task3
params                render
myke                  test
tags1      tagA tagB  tag
tags2      tagB tagC  tag
```

myke allows you to define tasks in `.yml` files and aggregates all of them. This helps you to manage multiple components in multiple projects in multiple repositories.

### vs

* `maven` tries to do too many things and ends up being good at none (scm/release/lifecycle/build/etc). myke has only one purpose of task aggregation
* `bazel` `buck` `pants` `gradle` `...` These tools *replace* your current buildchain by giving you a totally new DSL to compile your programs (`java_binary`, etc). myke simply acts as a yml-based interface to your existing tools and workflows, thereby not needing to change your project and IDE setup
* `grunt` `rake` `gulp` `pyinvoke` `robo` `...` myke allows aggregation of tasks through hierarchies, templates and tags, which these tools dont. [robo](https://github.com/tj/robo) is a close relative without these features. myke is also language agnostic - you don't need to know python to use myke because you only deal with simple yml files
* `make` `scons` `ninja` `...` All these are very low-level tools with crux of local file change tracking (run compile only if src folder is changed, etc). Most buildchains today (maven/docker/etc) are already intelligent enough to process only changed files, so myke completely bypasses file tracking. Secondly, these tools struggle dealing with remote artifacts (docker/nexus/gcloud/etc)
* `capistrano` `fabric` `...` myke is not a deployment tool for remote machines, and does not do anything over SSH
* `ansible` `salt` `...` myke is not a configuration management tool, its a task runner. Its like `ansible-local` but with a clear way to aggregate and run individual tasks

### Usage

For the alpha release, setup myke as follows:

```
git clone <this repo>
pip install --user -r requirements.txt`
export PATH=$(pwd)/bin:$PATH`
```

Create a `myke.yml` file in your project. Feel free to skip fields that are not required:

```yml
---
project: project_name
desc: project description
tags:
  - tagA
  - ...
env:
  key: value
  ...
env_files:
  - # by default includes <this-yml-file-name>.env
  - path/to/other/env/files
  - ...
includes:
  - path/to/folder/or/yml
  - ...
extends:
  - path/to/folder/or/yml
  - ...
tasks:
  <task-name-1>:
    desc: task description
    cmd: command with {{jinja}} templating
    before:
      - other-tasks-to-be-run-before-this
      - ...
    after:
      - other-tasks-to-be-run-after-this
      - ...
  <task-name-2>:
    cmd: |-
      multi-line
      commands
  ...
```

And start using myke:

```
# list all tasks
myke
# run task in project
myke project-name/task-name
# run task by tags
myke tag-name/task-name
```

### Features

* Define tasks in language-agnostic `.yml` files
* Run tasks with project/tag filtering
  * `myke build` runs build in all projects
  * `myke <project>/build` runs build in that specific `<project>`
  * `myke <tag>/build` runs build in all projects tagged `<tag>`
  * `myke <tagA>/<tagB>/.../build` can match tasks by many tags (AND)
* Many ways to configure
  * Environment files
    * Projects can define environment variables using multiple methods (see section below)
  * Runtime parameters
    * If your build task command is: `echo {{key1|required}} {{key2|required}}`
    * You can run it as: `myke build --key1=value1 --key2=value2`
    * Use runtime parameters to pass values that are dynamic each time you run `myke <project/task>`, otherwise prefer environment variables
* Template inheritance
  * Projects can extend other template(s) using `extends` keyword
  * Allows reuse of shared tasks, but still remain different using environment variables or parameters

### Examples

* Run `./myke` to list all the tasks
* Run `./myke test` to use myke to test itself
* Explore the self documenting `examples` folder

### Environment variables

`myke` should execute a given task with the same environment, irrespective of whether you invoke from a child folder or a parent folder. For this reason, parent project's environment variables are **not cascaded down** to child projects. Rather, a child project must **explicitly reference** shared environment variables using `env_files` or `extends` in the yml to avoid ambiguous behavior.

* Many ways to load environment variables:
  * Use `env` property in the project's yaml
  * Use `env_files` property to load custom .env files
  * myke will by default load `<yml-file-name>.env` if it exists
  * for each `.env` file, myke will also try to load a corresponding `.env.local` if exists. Users can use these `.env.local` files to override the default files, and gitignore them
* Additional project-specific environment variables that are set:
  * `$MyKE_PROJECT`: Name of project being executed
  * `$MyKE_TASK`: Name of task
  * `$MyKE_CWD`: Full path to directory where the task is defined
  * `$MyKE_CWD/bin` is prepended to path
* Same goes for every yml file that your project `extends`
  * So all the env variables naturally loaded by each `extends` file (`env`, `env_files`, `[extends-file].env`, `[extends-file].env.local`, `$PATH=EXTENDS_FILE_CWD/bin`, etc) are also made available to the child project
* Any environment variables set in command line override the above

### How do I share common logic in tasks?

Firstly, use single-purpose shared scripts, like the Unix philosophy. If the scripts are complex, make them as standalone scripts in language of your choice (python/etc). Put these scripts under `bin` folder inside your project. From the above Environment Variables section, you can find that `PROJECT_CWD/bin` is always added to the `PATH`, so you can start using these scripts straight away in your tasks.

If multiple projects need to share the same scripts, then another way is to leverage the behavior of `extends` templates. Remember that when your project `extends` another yml file, it **also** extends the environment variables and `bin` folder of that extended project. So you can model a template as a folder:

* `java-template`
  * `template.yml` - project template with tasks
  * `template.env` - environment vars, can be overridden by extending projects
  * `bin` - gets added to the PATH of extending projects
    * any shared scripts that you want
* `kubernetes-template`
  * ...
  * ...

### Roadmap

* New name
* Small enhancements (parameters in dependencies, warn duplicate project names, dry run, etc)
* CI (we already have many test cases, add linting)
* Experiment with golang? Users only use yml files, so we can change anything here without breaking the yml file contract. This is purely optional