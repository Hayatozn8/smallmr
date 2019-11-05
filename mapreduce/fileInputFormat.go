package mapreduce

//package inputformat

import (
	"os"
	//"fmt"
	"github.com/Hayatozn8/smallmr/config"
	intpuSplit "github.com/Hayatozn8/smallmr/split"
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

func NewFileInputFormat() InputFormat{
	return &FileInputFormat{}
}

// implements
// not have PathFilter
func (fif *FileInputFormat) GetSplits(job JobContext) ([]intpuSplit.InputSplit, error) {
	paths := job.GetInputPaths()
	minSize := util.MaxInt64(formatMinSplitSize, fif.getMinSplitSize(job))
	maxSize := fif.getMaxSplitSize(job)

	var splits = make([]intpuSplit.InputSplit, 0, default_split_count)
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
				// fmt.Println(splitSize)

				// how to compute :
				// bytesRemaining = fileLength - n * splitSize
				// start = fileLength - bytesRemaining = n * splitSize
				// readLength = splitSize or = fileLength%splitSize
				bytesRemaining := fileLength

				splitSizeFloat := float64(splitSize)
				for float64(bytesRemaining)/splitSizeFloat > split_slop {
					splits = append(splits, intpuSplit.NewFileSplit(path, fileLength-bytesRemaining, splitSize))
					bytesRemaining -= splitSize
				}

				if bytesRemaining != 0 {
					splits = append(splits, intpuSplit.NewFileSplit(path, fileLength-bytesRemaining, bytesRemaining))
				}
			} else {
				// TODO splits.append
				splits = append(splits, intpuSplit.NewFileSplit(path, 0, fileLength))
			}
		}
		// TODO: create empty array for zroe length files
		// else {

		// }
	}
	return splits, nil
}

// implements
func (fif *FileInputFormat) CreateRecordReader(split intpuSplit.InputSplit, context TaskContext) RecordReader {
	delimiter := context.GetConfiguration().GetString(config.FILE_DELIMITER)
	if delimiter != "" {
		return NewLineRecordReader([]byte(delimiter))
	} else {
		return NewLineRecordReader(nil)
	}
}

func (fif *FileInputFormat) getMaxSplitSize(context JobContext) int64 {
	return context.GetConfiguration().GetInt64(config.SPLIT_MAXSIZE)
}

func (fif *FileInputFormat) getMinSplitSize(context JobContext) int64 {
	return context.GetConfiguration().GetInt64(config.SPLIT_MINSIZE)
}

func (fif *FileInputFormat) isSplitable(path string) bool {
	return true
}

func (fif *FileInputFormat) computeSplitSize(blockSize int64, minSize int64, maxSize int64) int64 {
	return util.MaxInt64(minSize, util.MinInt64(blockSize, maxSize))
}
