package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"testing"
)

type binaryTest struct {
	name        string
	installName string
	checksum    string
}

// nolint: gochecknoglobals
var binaries = []binaryTest{
	{"packr", "gobuffalo/packr@1.21.9", "194f4db8bfcdbc3dbf83aaa4866692276622feea2d7aa4c80cd126d50518cdac"},
	{"dep", "golang/dep@0.5.0", "1a7bdb0d6c31ecba8b3fd213a1170adf707657123e89dff234871af9e0498be2"},
	{"golangci-lint", "golangci/golangci-lint@1.12.5", "d7edca5704c823c46554bfb8a5966d78b3d953eee1f7f1feea0aebc31431c215"},
	{"protoc", "google/protobuf@3.6.1", "216476214cd0ffea236df1e53891abd6b9d92518f179a9c3a33dbe010133c062"},
	{"licensei", "goph/licensei@0.0.7", "ebb553863821f8c149cf1ecce55d8c1340bf74f523bc0841879fdc25ef2616dc"},
	{"goreleaser", "goreleaser/goreleaser@0.95.2", "a43df5bb9d0ffd7d1e74f17d39611c6786a8d32bc245680b7d1fa75dde581a40"},
	{"gotestsum", "gotestyourself/gotestsum@0.3.2", "8e1707dbdc3675988ffa929bd8f3468ffcdd58c680c9ba21aa57c9cfb155937f"},
	{"migrate", "golang-migrate/migrate@4.2.1", "a699e4188e2013b06e139171a8fe3750818a470a9de57306007f5f85d657e125"},
	{"gobin", "myitcv/gobin@0.0.4", "7b4cbc1a14bde2aacac2d5c019797b03fdd22490c9fd859bd73574c2607928ba"},
	{"protolock", "nilslice/protolock@0.10.0", "a75cd17e0ace23269121567ee46c4e77cbc55b0b671389026f79a715c6f164f1"},
	{"jq", "stedolan/jq@1.6", "5c0a0a3ea600f302ee458b30317425dd9632d1ad8882259fcaf4e9b868b2b1ef"},
	{"jq", "stedolan/jq@1.4", "335a99a68eb9a1ecacfc947550003f103cfed627d3116c8bcae9ac11dd26d337"},
	{"prototool", "uber/prototool@1.3.0", "23c858991b8f54b5e17eee4427b447bbb3ff3352a8c50a4dce8b1f464e27704b"},
}

// nolint: gochecknoglobals
var vanityBinaries = []binaryTest{
	{"dep", "dep@0.5.0", "1a7bdb0d6c31ecba8b3fd213a1170adf707657123e89dff234871af9e0498be2"},
	{"golangci-lint", "golangci-lint@1.12.5", "d7edca5704c823c46554bfb8a5966d78b3d953eee1f7f1feea0aebc31431c215"},
	{"protoc", "protoc@3.6.1", "216476214cd0ffea236df1e53891abd6b9d92518f179a9c3a33dbe010133c062"},
	{"protoc", "protobuf@3.6.1", "216476214cd0ffea236df1e53891abd6b9d92518f179a9c3a33dbe010133c062"},
	{"goreleaser", "goreleaser@0.95.2", "a43df5bb9d0ffd7d1e74f17d39611c6786a8d32bc245680b7d1fa75dde581a40"},
	{"gotestsum", "gotestsum@0.3.2", "8e1707dbdc3675988ffa929bd8f3468ffcdd58c680c9ba21aa57c9cfb155937f"},
	{"gobin", "gobin@0.0.4", "7b4cbc1a14bde2aacac2d5c019797b03fdd22490c9fd859bd73574c2607928ba"},
	{"protolock", "protolock@0.10.0", "a75cd17e0ace23269121567ee46c4e77cbc55b0b671389026f79a715c6f164f1"},
	{"jq", "jq@1.6", "5c0a0a3ea600f302ee458b30317425dd9632d1ad8882259fcaf4e9b868b2b1ef"},
	{"prototool", "prototool@1.3.0", "23c858991b8f54b5e17eee4427b447bbb3ff3352a8c50a4dce8b1f464e27704b"},
}

func TestAcceptance(t *testing.T) {
	if m := flag.Lookup("test.run").Value.String(); m == "" || !regexp.MustCompile(m).MatchString(t.Name()) {
		t.Skip("skipping as execution was not requested explicitly using go test -run")
	}

	t.Parallel()

	if _, err := os.Stat("build/binbrew"); os.IsNotExist(err) {
		t.Fatal("Build binbrew binary first!")
	}

	t.Run("test", test)
	t.Run("testMulti", testMulti)
	t.Run("testVanity", testVanity)
}

func test(t *testing.T) {
	t.Parallel()

	cleanDir(t, "test")

	for _, bin := range binaries {
		bin := bin

		t.Run(bin.installName, func(t *testing.T) {
			t.Parallel()

			dir := "tests/test"

			cmd := exec.Command("../../build/binbrew", "install", bin.installName)
			cmd.Dir = dir

			err := cmd.Run()
			if err != nil {
				t.Fatal(err)
			}

			f, err := os.Open(dir + "/bin/" + bin.name)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			h := sha256.New()
			if _, err := io.Copy(h, f); err != nil {
				t.Fatal(err)
			}

			checksum := fmt.Sprintf("%x", h.Sum(nil))
			if checksum != bin.checksum {
				t.Errorf("checksum mismatch\ngot:  %s\nwant: %s", checksum, bin.checksum)
			}
		})
	}
}

func testMulti(t *testing.T) {
	t.Parallel()

	cleanDir(t, "multi_test")

	var bins []binaryTest
	var installArgs = []string{"install"}

	for _, bin := range binaries[:3] {
		bins = append(bins, bin)
		installArgs = append(installArgs, bin.installName)
	}

	dir := "tests/multi_test"

	cmd := exec.Command("../../build/binbrew", installArgs...)
	cmd.Dir = dir

	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}

	for _, bin := range bins {
		f, err := os.Open(dir + "/bin/" + bin.name)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			t.Fatal(err)
		}

		checksum := fmt.Sprintf("%x", h.Sum(nil))
		if checksum != bin.checksum {
			t.Errorf("checksum mismatch\ngot:  %s\nwant: %s", checksum, bin.checksum)
		}
	}
}

func testVanity(t *testing.T) {
	t.Parallel()

	cleanDir(t, "vanity_test")

	for _, bin := range vanityBinaries {
		bin := bin

		t.Run(bin.installName, func(t *testing.T) {
			t.Parallel()

			dir := "tests/vanity_test"

			cmd := exec.Command("../../build/binbrew", "install", bin.installName)
			cmd.Dir = dir

			err := cmd.Run()
			if err != nil {
				t.Fatal(err)
			}

			f, err := os.Open(dir + "/bin/" + bin.name)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			h := sha256.New()
			if _, err := io.Copy(h, f); err != nil {
				t.Fatal(err)
			}

			checksum := fmt.Sprintf("%x", h.Sum(nil))
			if checksum != bin.checksum {
				t.Errorf("checksum mismatch\ngot:  %s\nwant: %s", checksum, bin.checksum)
			}
		})
	}
}

func cleanDir(t *testing.T, suite string) {
	t.Helper()

	if err := os.RemoveAll("tests/" + suite); err != nil {
		t.Fatal(err)
	}

	if err := os.MkdirAll("tests/"+suite, 0755); err != nil {
		t.Fatal(err)
	}
}
