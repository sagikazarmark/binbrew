package command

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/hashicorp/go-getter" // nolint: goimports
	"github.com/sagikazarmark/binbrew/internal/provider"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

type installOptions struct {
	binaries   []string
	noProgress bool
}

// NewInstallCommand returns a cobra command for `binbrew install`.
func NewInstallCommand() *cobra.Command {
	options := installOptions{}

	cmd := &cobra.Command{
		Use:   "install [flags] NAME[@VERSION]...",
		Short: "Install one or more binaries",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.binaries = args

			return runInstall(options)
		},
		DisableFlagsInUseLine: true,
	}

	cmd.Flags().BoolVar(&options.noProgress, "no-progress", false, "Do not show download progress")

	return cmd
}

func runInstall(options installOptions) error {
	p := provider.NewGithubProvider()

	for _, binaryReference := range options.binaries {
		binaryReferenceParts := strings.SplitN(binaryReference, "@", 2)

		name := binaryReferenceParts[0]
		version := "latest"

		if len(binaryReferenceParts) > 1 {
			version = binaryReferenceParts[1]
		}

		binary, err := p.Resolve(name, version)
		if err != nil {
			return err
		}

		tmp := filepath.Join("bin/tmp", binary.Name)

		err = os.MkdirAll(tmp, 0744)
		if err != nil {
			return err
		}

		pb := newProgressBar()
		var getterOptions = make([]getter.ClientOption, 0)
		if !options.noProgress {
			getterOptions = append(getterOptions, getter.WithProgress(pb))
		}

		err = getter.GetAny(tmp, binary.URL, getterOptions...)
		pb.progress.Wait()
		if err != nil {
			return err
		}

		input, err := ioutil.ReadFile(filepath.Join(tmp, binary.File))
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(filepath.Join("bin", binary.Name), input, 0744)
		if err != nil {
			return err
		}
	}

	return nil
}

type progressBar struct {
	// lock everything below
	lock sync.Mutex

	progress *mpb.Progress
}

func newProgressBar() *progressBar {
	return &progressBar{
		progress: mpb.New(),
	}
}

// TrackProgress instantiates a new progress bar that will
// display the progress of stream until closed.
// total can be 0.
func (cpb *progressBar) TrackProgress(src string, currentSize, totalSize int64, stream io.ReadCloser) io.ReadCloser {
	cpb.lock.Lock()
	defer cpb.lock.Unlock()

	if cpb.progress == nil {
		cpb.progress = mpb.New()
	}
	bar := cpb.progress.AddBar(
		totalSize,
		mpb.PrependDecorators(
			decor.CountersKibiByte("% 6.1f / % 6.1f"),
		),
		mpb.AppendDecorators(
			decor.EwmaETA(decor.ET_STYLE_MMSS, float64(totalSize)/2048),
			decor.Name(" ] "),
			decor.AverageSpeed(decor.UnitKiB, "% .2f"),
		),
	)

	reader := bar.ProxyReader(stream)

	return &readCloser{
		Reader: reader,
		close: func() error {
			cpb.lock.Lock()
			defer cpb.lock.Unlock()

			return nil
		},
	}
}

type readCloser struct {
	io.Reader
	close func() error
}

func (c *readCloser) Close() error { return c.close() }
