// to register the go struct type to Kubernetes objects using schemes
// register a custom type in kubernetes
// ? what is he doing in 18:42 in https://www.youtube.com/watch?v=89PdRvRUcPU&list=PLh4KH3LtJvRTtFWz1WGlyDa7cKjj2Sns0&index=2

// example reference from upstream: https://github.com/kubernetes/code-generator/blob/master/examples/apiserver/apis/example/v1/register.go#L49

package v1

import (
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var  SchemeGroupVersion = schema.GroupVersion{
	Group: "radha", 
	Version: "v1",
}

var (
	
	SchemeBuilder      runtime.SchemeBuilder
)

// ? can you explain the purpose and usage of this init
func init() {
	//call a function to register the type 
	SchemeBuilder.Register(addKnownTypes)
}

func addKnownTypes(scheme *runtime.Scheme) error{
  scheme.AddKnownTypes(SchemeGroupVersion, &Kcluster{}, &KclusterList{})
  metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
  return nil
}

