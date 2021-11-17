# sample-apibuilder-demo

https://github.com/kubernetes-sigs/apiserver-builder-alpha/blob/master/docs/tools_user_guide.md

```
$ apiserver-boot init repo --domain ui.bytebuilder.dev

$ apiserver-boot create group version resource --group kubedb --version v1alpha1 --kind MongoDBInfoView

$ apiserver-boot create group version resource --group kubedb --version v1alpha1 --kind PostgresOverview

$ make generate manifests fmt

$ apiserver-boot run local
```
