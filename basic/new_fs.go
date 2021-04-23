package basic

import (
	"os"

	fs "github.com/aos-dev/go-service-fs/v2"
	"github.com/aos-dev/go-storage/v3/pairs"
	"github.com/aos-dev/go-storage/v3/types"
)

func NewFs() (types.Storager, error) {
	return fs.NewStorager(
		// WorkDir: https://aos.dev/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_FS_WORKDIR")),
	)
}
