package mapreduce

//package inputformat

import (
	"os"

	"github.com/Hayatozn8/smallmr/config"
	"github.com/Hayatozn8/smallmr/input/recordReader"
	"github.com/Hayatozn8/smallmr/input/split"
	intpuSplit "github.com/Hayatozn8/smallmr/input/split"
	"github.com/Hayatozn8/smallmr/mapreduce"
	"github.com/Hayatozn8/smallmr/util"
)

const (
	split_slop          float64 = 1.1
	formatMinSplitSize  int64   = 1
	default_split_count         = 10
)

//single file
type FileInputFormat struct {
	//TODO
}

// implements
// not have PathFilter
func (fif *FileInputFormat) GetSplits(job mapreduce.JobContext) ([]intpuSplit.InputSplit, error) {
	paths := job.GetInputPaths()
	minSize := util.MaxInt64(formatMinSplitSize, fif.getMinSplitSize(job))
	maxSize := fif.getMaxSplitSize(job)

	var splits = make([]intpuSplit.InputSplit, default_split_count)
	for _, path := range paths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return nil, err
		}

		fileLength := fileInfo.Size()
		if fileLength != 0 {
			if fif.isSplitable(path) {
				//long blockSize = file.getBlockSize();
				splitSize := fif.computeSplitSize(util.BlockSize, minSize, maxSize)

				// bytesRemaining = fileLength - n * splitSize
				// start = fileLength - bytesRemaining = n * splitSize
				// readLength = splitSize or = fileLength%splitSize
				bytesRemaining := fileLength
				for float64(bytesRemaining/splitSize) > split_slop {
					splits = append(splits, split.NewFileSplit(path, fileLength-bytesRemaining, splitSize))
					bytesRemaining -= splitSize
				}

				if bytesRemaining != 0 {
					splits = append(splits, split.NewFileSplit(path, fileLength-bytesRemaining, bytesRemaining))
				}
			} else {
				// TODO splits.append
				splits = append(splits, split.NewFileSplit(path, 0, fileLength))
			}
		}
		// TODO: create empty array for zroe length files
		// else {

		// }
	}
	return splits, nil
}

// implements
func (fif *FileInputFormat) createRecordReader(split intpuSplit.InputSplit, context mapreduce.TaskContext) recordReader.RecordReader {
	return recordReader.NewLineRecordReader(, )
}

func (fif *FileInputFormat) getMaxSplitSize(context mapreduce.JobContext) int64 {
	return context.GetConfiguration().GetInt64(config.SPLIT_MAXSIZE)
}

func (fif *FileInputFormat) getMinSplitSize(context mapreduce.JobContext) int64 {
	return context.GetConfiguration().GetInt64(config.SPLIT_MINSIZE)
}

func (fif *FileInputFormat) isSplitable(path string) bool {
	return false
}

func (fif *FileInputFormat) computeSplitSize(blockSize int64, minSize int64, maxSize int64) int64 {
	return util.MaxInt64(minSize, util.MinInt64(blockSize, maxSize))
}
