# DOCODE

Projects onboarding tool using [docker](http://www.docker.io)

## HOW IT WORKS

Instead of polluting your machine with different dependencies for different projects, `docode` allows each project to have their dependencies in isolation.

Imagine checking out a project you never worked before and being able to run it with only one command, without worring about conflicts with our local environment.

In this scenario, each project would have their own `Docodefile` which contains instructions for setting up the development enviroment. For example:

``` yaml
image_name: 'my_project_docker_image'
image_tag: 'latest'
ports:
  80: 8080
  443: 4443
run_list:
  - start_background_services.sh
  - tmux
```

Running `docode` in this project folder will:
- load up that file
- fire docker with that image
- map the ports
- execute the run list, in the end opening tmux for the developer.

It also mounts the current folder as the  workdir inside the running container.

## TODO

* Import SSH key
* Custom ENV sets
* Custom extra mount points
* Forwarding Host ENV option
* Command line options:
  * set target docodefile
