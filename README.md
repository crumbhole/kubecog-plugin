# crumblecog-plugin

This is an init plugin for [argocd-lovely-plugin](https://github.com/crumbhole/argocd-lovely-plugin) to allow for templating in the crumblecog project.

## Usage

Install into argocd alongside lovely and and add to ARGOCD_ENV_LOVELY_PREPROCESSORS.

You must place a .kubecog.yaml in the directory you want processing with this plugin, otherwise it will not do anything.

It uses gomplate (go's templating language), the same as helm does. Left and right delimiters default to [[ and ]] to avoid the clash with the default helm delimiters. They can be overridden in .kubecog.yaml for your specific templates if that is helpful.

Environment:

- ARGOCD_ENV_KUBECOG_URL_PREFIX: *Required* Set this to a gomplate URL for where to find your values files for context in gomplate.
- ARGOCD_ENV_KUBECOG_GOMPLATE_PATH: Set this to where to find gomplate binary if it is not on your path.

## .kubecog.yaml
kubecog:
  contextname: context.yml
delimiters: (all optional)
  left: <leftdelim>
  right: <rightdelim>
