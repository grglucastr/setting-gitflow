# Bnunes.dev GitFlow Foundation
Here is the foundation of my GitFlow process. I'll be using and improving this process through all my new projects and this repository will be my testing one for future improvements.

# Branches
The main braches are:
- main
- develop

Both branches are locked to commit directly and the only way to publish to these branches are through Pull Requests.

The pipeline will create the Pull Request **automatically** if you follow the branch naming convention as described below.


## To add new features
To add new features you have to create a **feature** branch from the **develop** branch as shown in the example below.

```
git checkout develop
git checkout -b feature/new-feature
```

## To add a bugfix
To add a bugfix you have to also create a **bugfix** branch from the **develop** branch as shown in the example below.

```
git checkout develop
git checkout -b bugfix/new-bugfix
```

# Workflows
Above is described the main stages of the GitFlow process

## Create PR automatically to develop