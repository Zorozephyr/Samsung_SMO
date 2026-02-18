package controller

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	databasev1 "github.com/example/postgres-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *DatabaseReconciler) buildStatefulSet(db *databasev1.Database) *appsv1.StatefulSet {
	replicas := int32(1)
	if db.Spec.Replicas != nil {
		replicas = *db.Spec.Replicas
	}

	image := db.Spec.Image
	if image == "" {
		image = "postgres:14"
	}

	secretName := r.secretName(db)

	return &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      db.Name,
			Namespace: db.Namespace,
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":      "database",
					"database": db.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":      "database",
						"database": db.Name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "postgres",
							Image: image,
							Env: []corev1.EnvVar{
								{
									Name:  "POSTGRES_DB",
									Value: db.Spec.DatabaseName,
								},
								{
									Name: "POSTGRES_USER",
									ValueFrom: &corev1.EnvVarSource{
										SecretKeyRef: &corev1.SecretKeySelector{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: secretName,
											},
											Key: "username",
										},
									},
								},
								{
									Name: "POSTGRES_PASSWORD",
									ValueFrom: &corev1.EnvVarSource{
										SecretKeyRef: &corev1.SecretKeySelector{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: secretName,
											},
											Key: "password",
										},
									},
								},
								{
									Name:  "PGDATA",
									Value: "/var/lib/postgresql/data/pgdata",
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "data",
									MountPath: "/var/lib/postgresql/data",
								},
							},
						},
					},
				},
			},
			VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "data",
					},
					Spec: corev1.PersistentVolumeClaimSpec{
						AccessModes: []corev1.PersistentVolumeAccessMode{
							corev1.ReadWriteOnce,
						},
						Resources: corev1.VolumeResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceStorage: resource.MustParse(db.Spec.Storage.Size),
							},
						},
					},
				},
			},
		},
	}
}

func (r *DatabaseReconciler) secretName(db *databasev1.Database) string {
	return fmt.Sprintf("%s-credentials", db.Name)
}

func generatePassword(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func (r *DatabaseReconciler) buildService(db *databasev1.Database) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      db.Name,
			Namespace: db.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app":      "database",
				"database": db.Name,
			},
			Ports: []corev1.ServicePort{
				{
					Port: 5432,
					Name: "postgres",
				},
			},
		},
	}
}

func (r *DatabaseReconciler) listDatabasesInNameSpace(ctx context.Context, namespace string) (*databasev1.DatabaseList, error) {
	list := &databasev1.DatabaseList{}
	err := r.List(ctx, list, client.InNamespace(namespace))
	return list, err
}

func (r *DatabaseReconciler) listDatabasesByLabel(ctx context.Context, labels map[string]string) (*databasev1.DatabaseList, error) {
	list := &databasev1.DatabaseList{}
	err := r.List(ctx, list, client.MatchingLabels(labels))
	return list, err
}

func (r *DatabaseReconciler) patchStatefulSetReplicas(ctx context.Context, statefulSet *appsv1.StatefulSet, replicas int32) error {
	patch := client.MergeFrom(statefulSet.DeepCopy())
	statefulSet.Spec.Replicas = &replicas
	return r.Patch(ctx, statefulSet, patch)
}
