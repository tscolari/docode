# DOCODE

Project dependency management tool

## HOW IT WORKS

You shouldn't need to worry with libraries and versions every time you start
working on a new project. Keeping them from conflicting from the ones from
another project you are working on, or the one you are going to start next week.

`Docode` is a simple solution for project dependency management, it wraps docker
and keep the project dependencies in a docker image.

It uses a `DocodeFile` to describe the project world, that gets automatically
loaded when you run `docode` in a folder.

For example:

``` yaml
image_name: my_project_docker_image
image_tag: latest
ssh_key: /home/me/.ssh/my_key
ports:
  80: 8080
  443: 4443
run_list:
  - start_background_services.sh
  - tmux
```

Running `docode` in this project folder will:
- start docker with the `my_project_docker_image:latest` image, which should ideally have the project dependencies already installed.
- map the ports from the container to the host, so you can access through a browser for example.
- load the ssh_key inside the docker container.
- execute the run list, in this case the developer.

It also mounts the current folder (the project folder) as the /workdir inside the running container,
allowing external and internal tools to have access to the folder.

## COMMAND LINE OPTIONS

* `-k <ssh-key path>` will override/use the given ssh-key.
* `-n` will skip the image pull step (use the local one)
* `-t <image tag>` will override the image tag given
* `-i <image name>` will override the image given

## License

Docode is released under the MIT License.
