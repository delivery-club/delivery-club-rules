package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	projectName = "delivery-club-rules"
	binary      = "dcRules"
	releaseDir  = "bin"
)

type platformInfo struct {
	goos   string
	goarch string
}

func (p platformInfo) String() string { return p.goos + "-" + p.goarch }

func main() {
	log.SetFlags(0)

	version := flag.String("version", "", "dcRules release version")
	flag.Parse()

	if *version == "" {
		log.Fatal("version argument is not set")
	}

	platforms := []platformInfo{
		{"linux", "amd64"},
		{"linux", "arm64"},
		{"darwin", "amd64"},
		{"darwin", "arm64"},
		{"windows", "amd64"},
		{"windows", "arm64"},
	}

	err := os.Mkdir(releaseDir, 0755)
	if err != nil && err != os.ErrExist {
		log.Printf("on release dir: %s", err)
		return
	}

	checksums, err := os.Create(filepath.Join(releaseDir, projectName+"-"+*version+"-checksums.txt"))
	if err != nil {
		log.Printf("on create checksums: %s", err)
		return
	}
	defer checksums.Close()

	for _, platform := range platforms {
		if err := prepareArchive(checksums, platform, *version); err != nil {
			log.Printf("error: build %s: %v", platform, err)
			return
		}
	}
}

func prepareArchive(checksums io.Writer, platform platformInfo, version string) error {
	log.Printf("building %s", platform)

	buildCmd := exec.Command("make", "build-release")
	buildCmd.Env = append([]string{}, os.Environ()...) // Copy env slice
	buildCmd.Env = append(buildCmd.Env, "GOOS="+platform.goos)
	buildCmd.Env = append(buildCmd.Env, "GOARCH="+platform.goarch)
	out, err := buildCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("run %s: %v: %s", buildCmd, err, out)
	}

	filename := binary
	if platform.goos == "windows" {
		filename = binary + ".exe"
		err = os.Rename(filepath.Join(releaseDir, binary), filepath.Join(releaseDir, filename))
		if err != nil {
			return fmt.Errorf("on file rename: from %s to %s", binary, filename)
		}
	}

	archiveName := projectName + "-" + version + "-" + platform.String() + ".zip"
	zipCmd := exec.Command("zip", archiveName, filename)
	zipCmd.Dir = releaseDir
	log.Printf("creating %s archive", archiveName)
	if out, err := zipCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("make archive: %v: %s", err, out)
	}

	shaCmd := exec.Command("shasum", "-a", "256", archiveName)
	shaCmd.Dir = releaseDir

	out, err = shaCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("on create shasum: %s", err)
	}
	fmt.Fprint(checksums, string(out))

	return nil
}
