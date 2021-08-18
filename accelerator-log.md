# Accelerator Log

## Options
```json
{
  "branchName" : "main",
  "gitRepoName" : "garreeoke",
  "imageRepoName" : "garreeoke",
  "projectDescription" : "",
  "projectName" : "go-kontinue"
}
```
## Log
```
┏ engine (Chain)
┃  Info Running Chain(Chain, UniquePath)
┃ ┏ engine.transformations[0] (Chain)
┃ ┃  Info Running Chain(Include, ReplaceText)
┃ ┃ ┏ engine.transformations[0].transformations[0] (Include)
┃ ┃ ┃  Info Will include [**.go, **.md, k8s/**.yaml]
┃ ┃ ┃ Debug router.go matched [**.go, **.md, k8s/**.yaml] -> included
┃ ┃ ┃ Debug types.go matched [**.go, **.md, k8s/**.yaml] -> included
┃ ┃ ┃ Debug handlers.go matched [**.go, **.md, k8s/**.yaml] -> included
┃ ┃ ┃ Debug main.go matched [**.go, **.md, k8s/**.yaml] -> included
┃ ┃ ┃ Debug k8s/workload.yaml matched [**.go, **.md, k8s/**.yaml] -> included
┃ ┃ ┃ Debug README.md matched [**.go, **.md, k8s/**.yaml] -> included
┃ ┃ ┗ Debug routes.go matched [**.go, **.md, k8s/**.yaml] -> included
┃ ┃ ┏ engine.transformations[0].transformations[1] (ReplaceText)
┃ ┗ ┗  Info Will replace [git-repo-name->garreeoke, image-repo-name->garreeoke, sample-go-api->go-kontinue, branch-name->main]
┗ ╺ engine.transformations[1] (UniquePath)
```
