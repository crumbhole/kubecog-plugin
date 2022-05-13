# crumblecog-plugin

This is an init plugin for [argocd-lovely-plugin](https://github.com/crumbhole/argocd-lovely-plugin) to allow for templating in the crumblecog project.

## Usage

Install into argocd alongside lovely and and add to LOVELY_PREPROCESSORS.

It uses go's templating language, the same as helm does. It doesn't have any of the special features that helm adds to the language. Values will appear directly as {{ .foo.bar }} with no .Values prefix, unlike helm.

You probably want to set COG_VALUES_PATH is where to find the values file. This can be a local path.
