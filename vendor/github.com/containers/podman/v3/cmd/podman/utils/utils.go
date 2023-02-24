package utils

import (
	"fmt"
	"os"

	"github.com/containers/podman/v3/pkg/domain/entities"
	"github.com/containers/podman/v3/pkg/domain/entities/reports"
)

// IsDir returns true if the specified path refers to a directory.
func IsDir(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return file.IsDir()
}

// FileExists returns true if path refers to an existing file.
func FileExists(path string) bool {
	file, err := os.Stat(path)
	// All errors return file == nil
	if err != nil {
		return false
	}
	return !file.IsDir()
}

func PrintPodPruneResults(podPruneReports []*entities.PodPruneReport, heading bool) error {
	var errs OutputErrors
	if heading && len(podPruneReports) > 0 {
		fmt.Println("Deleted Pods")
	}
	for _, r := range podPruneReports {
		if r.Err == nil {
			fmt.Println(r.Id)
		} else {
			errs = append(errs, r.Err)
		}
	}
	return errs.PrintErrors()
}

func PrintContainerPruneResults(containerPruneReports []*reports.PruneReport, heading bool) error {
	var errs OutputErrors
	if heading && (len(containerPruneReports) > 0) {
		fmt.Println("Deleted Containers")
	}
	for _, v := range containerPruneReports {
		fmt.Println(v.Id)
		if v.Err != nil {
			errs = append(errs, v.Err)
		}
	}
	return errs.PrintErrors()
}

func PrintVolumePruneResults(volumePruneReport []*reports.PruneReport, heading bool) error {
	var errs OutputErrors
	if heading && len(volumePruneReport) > 0 {
		fmt.Println("Deleted Volumes")
	}
	for _, r := range volumePruneReport {
		if r.Err == nil {
			fmt.Println(r.Id)
		} else {
			errs = append(errs, r.Err)
		}
	}
	return errs.PrintErrors()
}

func PrintImagePruneResults(imagePruneReports []*reports.PruneReport, heading bool) error {
	if heading {
		fmt.Println("Deleted Images")
	}
	for _, r := range imagePruneReports {
		fmt.Println(r.Id)
		if r.Err != nil {
			fmt.Fprint(os.Stderr, r.Err.Error()+"\n")
		}
	}

	return nil
}
