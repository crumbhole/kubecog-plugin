# Kubecog-plugin

This is an init plugin for [argocd-lovely-plugin](https://github.com/crumbhole/argocd-lovely-plugin) to allow for templating in the kubecog project.

## Usage

Use the docker image from this repo, which includes lovely, and is used as a sidecar for argocd.

You must place a .kubecog.yaml in the directory you want processing with this plugin, otherwise it will not do anything.

It uses gomplate (go's templating language), the same as helm does. Left and right delimiters default to [[ and ]] to avoid the clash with the default helm delimiters. They can be overridden in .kubecog.yaml for your specific templates if that is helpful.

Environment:

- KUBECOG_URL_PREFIX: *Required* Set this to a gomplate URL for where to find your values files for context in gomplate.
- KUBECOG_GOMPLATE_PATH: Set this to where to find gomplate binary if it is not on your path.

## .kubecog.yaml
kubecog:
  contextname: context.yml
delimiters: (all optional)
  left: <leftdelim>
  right: <rightdelim>
