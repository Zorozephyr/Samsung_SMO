/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	databasev1 "github.com/example/postgres-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	errors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// DatabaseReconciler reconciles a Database object
type DatabaseReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=database.example.com,resources=databases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=database.example.com,resources=databases/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=database.example.com,resources=databases/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete

func (r *DatabaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := logf.FromContext(ctx)

	db := &databasev1.Database{}
	if err := r.Get(ctx, req.NamespacedName, db); err != nil {
		if errors.IsNotFound(err) {
			logger.Info("Database not found")
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	logger.Info("Reconciling Database", "name", db.Name)

	// Reconcile Secret (must be done before StatefulSet)
	if err := r.reconcileSecret(ctx, db); err != nil {
		return ctrl.Result{}, err
	}

	// Reconcile StatefulSet
	if err := r.reconcileStatefulSet(ctx, db); err != nil {
		return ctrl.Result{}, err
	}

	// Reconcile Service
	if err := r.reconcileService(ctx, db); err != nil {
		return ctrl.Result{}, err
	}

	// Update status
	if err := r.updateStatus(ctx, db); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *DatabaseReconciler) reconcileSecret(ctx context.Context, db *databasev1.Database) error {
	logger := logf.FromContext(ctx)

	secretName := r.secretName(db)
	secret := &corev1.Secret{}

	err := r.Get(ctx, client.ObjectKey{
		Name:      secretName,
		Namespace: db.Namespace,
	}, secret)

	if errors.IsNotFound(err) {
		password, err := generatePassword(16)
		if err != nil {
			return fmt.Errorf("failed to generate password: %w", err)
		}

		secret = &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      secretName,
				Namespace: db.Namespace,
				Labels:    db.Labels,
			},
			Type: corev1.SecretTypeOpaque,
			Data: map[string][]byte{
				"username": []byte(db.Spec.Username),
				"password": []byte(password),
				"database": []byte(db.Spec.DatabaseName),
			},
		}

		if err := ctrl.SetControllerReference(db, secret, r.Scheme); err != nil {
			return err
		}

		logger.Info("Creating secret", "name", secret.Name)
		return r.Create(ctx, secret)
	}

	return nil
}

func (r *DatabaseReconciler) reconcileStatefulSet(ctx context.Context, db *databasev1.Database) error {
	logger := logf.FromContext(ctx)

	statefulSet := &appsv1.StatefulSet{}
	err := r.Get(ctx, client.ObjectKey{
		Name:      db.Name,
		Namespace: db.Namespace,
	}, statefulSet)

	desiredStatefulSet := r.buildStatefulSet(db)

	if errors.IsNotFound(err) {
		// Set owner reference
		if err := ctrl.SetControllerReference(db, desiredStatefulSet, r.Scheme); err != nil {
			return err
		}
		logger.Info("Creating StatefulSet", "name", desiredStatefulSet.Name)
		return r.Create(ctx, desiredStatefulSet)
	} else if err != nil {
		return err
	}

	if *statefulSet.Spec.Replicas != *desiredStatefulSet.Spec.Replicas {
		logger.Info("Updating replicas via patch")
		return r.patchStatefulSetReplicas(ctx, statefulSet, *desiredStatefulSet.Spec.Replicas)
	}

	// Update if needed
	if statefulSet.Spec.Replicas != desiredStatefulSet.Spec.Replicas ||
		statefulSet.Spec.Template.Spec.Containers[0].Image != desiredStatefulSet.Spec.Template.Spec.Containers[0].Image {
		statefulSet.Spec = desiredStatefulSet.Spec
		logger.Info("Updating StatefulSet", "name", statefulSet.Name)
		if err := r.updateWithRetry(ctx, statefulSet, 3); err != nil {
			return err
		}
	}

	return nil
}

func (r *DatabaseReconciler) reconcileService(ctx context.Context, db *databasev1.Database) error {
	service := &corev1.Service{}
	err := r.Get(ctx, client.ObjectKey{
		Name:      db.Name,
		Namespace: db.Namespace,
	}, service)

	desiredService := r.buildService(db)

	if errors.IsNotFound(err) {
		if err := ctrl.SetControllerReference(db, desiredService, r.Scheme); err != nil {
			return err
		}
		return r.Create(ctx, desiredService)
	} else if err != nil {
		return err
	}

	// Service updates are less common, but handle if needed
	return nil
}

func (r *DatabaseReconciler) updateStatus(ctx context.Context, db *databasev1.Database) error {
	// Set the secret name in status
	db.Status.SecretName = r.secretName(db)

	// Check StatefulSet status
	statefulSet := &appsv1.StatefulSet{}
	err := r.Get(ctx, client.ObjectKey{
		Name:      db.Name,
		Namespace: db.Namespace,
	}, statefulSet)

	if err != nil {
		db.Status.Phase = "Pending"
		db.Status.Ready = false
	} else {
		if statefulSet.Status.ReadyReplicas == *statefulSet.Spec.Replicas {
			db.Status.Phase = "Ready"
			db.Status.Ready = true
			db.Status.Endpoint = fmt.Sprintf("%s.%s.svc.cluster.local:5432", db.Name, db.Namespace)
		} else {
			db.Status.Phase = "Creating"
			db.Status.Ready = false
		}
	}

	return r.Status().Update(ctx, db)
}

// SetupWithManager sets up the controller with the Manager.
func (r *DatabaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&databasev1.Database{}).
		Owns(&appsv1.StatefulSet{}).
		Owns(&corev1.Service{}).
		Owns(&corev1.Secret{}).
		Named("database").
		Complete(r)
}
