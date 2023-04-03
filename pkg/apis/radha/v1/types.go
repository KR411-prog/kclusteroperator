//https://www.youtube.com/watch?v=89PdRvRUcPU&list=PLh4KH3LtJvRTtFWz1WGlyDa7cKjj2Sns0&index=2

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/api/meta/v1" // ? how is mentioned as v1. when I see only https://pkg.go.dev/k8s.io/apimachinery@v0.26.3/pkg/api/meta
	//then I found this k8s.io/apimachinery/pkg/apis/meta/v1
)
// 1. typeMeta 2.ObjectMeta 3.Spec are present in a Kubernetes resource
type Kcluster struct {
	metav1.TypeMeta //seems to be struct of structs
	metav1.ObjectMeta

	Spec KclusterSpec
}