// file to specify the tags for v1 api
// tags - are the way to control the behaviour of code generator
// global tags - doc.go. 
// example reference from upstream: https://github.com/kubernetes/code-generator

// the below ones are global tags
// +k8s:deepcopy-gen=package
// +k8s:defaulter-gen=TypeMeta
// +groupName=radha

//code-generator command not working for me
// ./generate-groups.sh all github.com/KR411-prog/kcluster/pkg/client github.com/KR411-prog/kcluster/pkg/apis radha:v1 --go-header-file ./hack/boilerplate.go.txt  
package v1

